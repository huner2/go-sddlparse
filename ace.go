package sddlparse

import (
	"encoding/binary"
	"errors"

	conditionalparser "github.com/huner2/go-sddlparse/internal/conditionalParser"
	"github.com/huner2/go-sddlparse/internal/util"
)

func aceFromBytes(data []byte) (*ACE, uint16, error) {
	if len(data) < 4 {

		return nil, 0, errors.New(errInvalidACE)
	}
	aceType := AceType(data[0])
	aceFlags := AceFlag(data[1])
	aceSize := binary.LittleEndian.Uint16(data[2:4])
	if aceSize < 4 || len(data) < int(aceSize) || aceSize%4 != 0 {

		return nil, 0, errors.New(errInvalidACE)
	}
	var ace *ACE
	var err error
	tmp := data[4:aceSize]
	switch aceType {
	case ACETYPE_ACCESS_ALLOWED:
		fallthrough
	case ACETYPE_ACCESS_DENIED:
		fallthrough
	case ACETYPE_SYSTEM_AUDIT:
		fallthrough
	case ACETYPE_SYSTEM_MANDATORY_LABEL:
		fallthrough
	case ACETYPE_SYSTEM_SCOPED_POLICY_ID:
		ace, err = accessAceFromBytes(tmp)
	case ACETYPE_ACCESS_ALLOWED_OBJECT:
		fallthrough
	case ACETYPE_ACCESS_DENIED_OBJECT:
		ace, err = accessObjectAceFromBytes(tmp)
	case ACETYPE_ACCESS_ALLOWED_CALLBACK:
		fallthrough
	case ACETYPE_ACCESS_DENIED_CALLBACK:
		fallthrough
	case ACETYPE_SYSTEM_AUDIT_CALLBACK:
		ace, err = accessCallbackAceFromBytes(tmp)
	case ACETYPE_ACCESS_ALLOWED_CALLBACK_OBJECT:
		fallthrough
	case ACETYPE_ACCESS_DENIED_CALLBACK_OBJECT:
		fallthrough
	case ACETYPE_SYSTEM_AUDIT_OBJECT:
		fallthrough
	case ACETYPE_SYSTEM_AUDIT_CALLBACK_OBJECT:
		ace, err = accessCallbackObjectAceFromBytes(tmp)
	case ACETYPE_SYSTEM_RESOURCE_ATTRIBUTE:
		ace, err = systemResourceAttributeAceFromBytes(tmp)
	default:

		return nil, 0, errors.New(errInvalidACE)
	}
	if err != nil {
		return nil, 0, err
	}
	ace.Type = aceType
	ace.Flags = aceFlags

	return ace, aceSize, nil
}

func aceToBytes(ace *ACE) ([]byte, uint16, error) {
	var err error
	var data []byte
	switch ace.Type {
	case ACETYPE_ACCESS_ALLOWED:
		fallthrough
	case ACETYPE_ACCESS_DENIED:
		fallthrough
	case ACETYPE_SYSTEM_AUDIT:
		fallthrough
	case ACETYPE_SYSTEM_MANDATORY_LABEL:
		fallthrough
	case ACETYPE_SYSTEM_SCOPED_POLICY_ID:
		data, err = accessAceToBytes(ace)
	case ACETYPE_ACCESS_ALLOWED_OBJECT:
		fallthrough
	case ACETYPE_ACCESS_DENIED_OBJECT:
		data, err = accessObjectAceToBytes(ace)
	case ACETYPE_ACCESS_ALLOWED_CALLBACK:
		fallthrough
	case ACETYPE_ACCESS_DENIED_CALLBACK:
		fallthrough
	case ACETYPE_SYSTEM_AUDIT_CALLBACK:
		data, err = accessCallbackAceToBytes(ace)
	case ACETYPE_ACCESS_ALLOWED_CALLBACK_OBJECT:
		fallthrough
	case ACETYPE_ACCESS_DENIED_CALLBACK_OBJECT:
		fallthrough
	case ACETYPE_SYSTEM_AUDIT_OBJECT:
		fallthrough
	case ACETYPE_SYSTEM_AUDIT_CALLBACK_OBJECT:
		data, err = accessCallbackObjectAceToBytes(ace)
	case ACETYPE_SYSTEM_RESOURCE_ATTRIBUTE:
		data, err = systemResourceAttributeAceToBytes(ace)
	default:

		return nil, 0, errors.New(errInvalidACE)
	}
	if err != nil {

		return nil, 0, err
	}
	aceSize := uint16(len(data) + 4)
	data = append([]byte{byte(ace.Type), byte(ace.Flags), 0, 0}, data...)
	data = append(data, make([]byte, aceSize%4)...)
	if aceSize%4 != 0 {
		aceSize += aceSize % 4
	}
	binary.LittleEndian.PutUint16(data[2:4], aceSize)

	return data, aceSize, nil
}

// AccessMask - 4 bytes
// SID - variable
func accessAceFromBytes(data []byte) (*ACE, error) {
	if len(data) < 4 {
		return nil, errors.New(errInvalidACE)
	}
	accessMask := binary.LittleEndian.Uint32(data[0:4])
	sid, _, err := util.SidFromLEBytes(data[4:])
	if err != nil {
		return nil, err
	}
	return &ACE{
		AccessMask: AccessMask(accessMask),
		SID:        sid,
	}, nil
}

func accessAceToBytes(ace *ACE) ([]byte, error) {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data[0:4], uint32(ace.AccessMask))
	sid, err := util.SidToLEBytes(ace.SID)
	if err != nil {
		return nil, err
	}
	data = append(data, sid...)

	return data, nil
}

// AccessMask - 4 bytes
// Flags - 4 bytes
// ObjectType - 16 bytes
// InheritedObjectType - 16 bytes
// SID - variable
func accessObjectAceFromBytes(data []byte) (*ACE, error) {
	var err error
	if len(data) < 4 {
		return nil, errors.New(errInvalidACE)
	}
	accessMask := binary.LittleEndian.Uint32(data[0:4])
	uniqueFlags := ObjectFlag(binary.LittleEndian.Uint32(data[4:8]))

	var objectType util.GUID
	var inheritedObjectType util.GUID

	start := 8

	if uniqueFlags&ACE_OBJECT_TYPE_PRESENT != 0 {
		objectType, err = util.GuidFromBytes(data[8:24])
		if err != nil {
			return nil, err
		}
		start += 16
	}
	if uniqueFlags&ACE_INHERITED_OBJECT_TYPE_PRESENT != 0 {
		inheritedObjectType, err = util.GuidFromBytes(data[start : start+16])
		if err != nil {
			return nil, err
		}
		start += 16
	}
	sid, _, err := util.SidFromLEBytes(data[start:])
	if err != nil {
		return nil, err
	}
	return &ACE{
		AccessMask:          AccessMask(accessMask),
		SID:                 sid,
		ObjectType:          objectType,
		InheritedObjectType: inheritedObjectType,
		ObjectFlags:         uniqueFlags,
	}, nil
}

func accessObjectAceToBytes(ace *ACE) ([]byte, error) {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data[0:4], uint32(ace.AccessMask))
	data = append(data, make([]byte, 4)...)
	binary.LittleEndian.PutUint32(data[4:8], uint32(ace.ObjectFlags))
	if ace.ObjectFlags&ACE_OBJECT_TYPE_PRESENT != 0 {
		objectType := util.GuidToBytes(ace.ObjectType)
		data = append(data, objectType...)
	}
	if ace.ObjectFlags&ACE_INHERITED_OBJECT_TYPE_PRESENT != 0 {
		inheritedObjectType := util.GuidToBytes(ace.InheritedObjectType)
		data = append(data, inheritedObjectType...)
	}
	sid, err := util.SidToLEBytes(ace.SID)
	if err != nil {
		return nil, err
	}
	data = append(data, sid...)

	return data, nil
}

// AccessMask - 4 bytes
// SID - variable
// ApplicationData - variable
func accessCallbackAceFromBytes(data []byte) (*ACE, error) {
	if len(data) < 4 {
		return nil, errors.New(errInvalidACE)
	}
	accessMask := binary.LittleEndian.Uint32(data[0:4])
	sid, length, err := util.SidFromLEBytes(data[4:])
	if err != nil {
		return nil, err
	}
	var ApplicationData []byte
	if len(data) > 4+int(length) {
		ApplicationData = data[4+length:]
	}
	var conditional *conditionalparser.ConditionalExpression
	if len(ApplicationData) > 0 {
		conditional, err = conditionalparser.ParseApplicationData(ApplicationData)
		if err != nil {
			return nil, err
		}
	}
	return &ACE{
		AccessMask:      AccessMask(accessMask),
		SID:             sid,
		ApplicationData: conditional,
	}, nil
}

func accessCallbackAceToBytes(ace *ACE) ([]byte, error) {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data[0:4], uint32(ace.AccessMask))
	sid, err := util.SidToLEBytes(ace.SID)
	if err != nil {
		return nil, err
	}
	data = append(data, sid...)
	var appDataBytes []byte
	if ace.ApplicationData != nil {
		appDataBytes, err = ace.ApplicationData.ToBytes()
		if err != nil {
			return nil, err
		}
	}
	data = append(data, appDataBytes...)

	return data, nil
}

// AccessMask - 4 bytes
// Flags - 4 bytes
// ObjectType - 16 bytes
// InheritedObjectType - 16 bytes
// SID - variable
// ApplicationData - variable
func accessCallbackObjectAceFromBytes(data []byte) (*ACE, error) {
	if len(data) < 4 {
		return nil, errors.New(errInvalidACE)
	}
	accessMask := binary.LittleEndian.Uint32(data[0:4])
	uniqueFlags := ObjectFlag(binary.LittleEndian.Uint32(data[4:8]))

	var err error
	var objectType util.GUID
	var inheritedObjectType util.GUID

	start := 8

	if uniqueFlags&ACE_OBJECT_TYPE_PRESENT != 0 {
		objectType, err = util.GuidFromBytes(data[8:24])
		if err != nil {
			return nil, err
		}
		start += 16
	}
	if uniqueFlags&ACE_INHERITED_OBJECT_TYPE_PRESENT != 0 {
		inheritedObjectType, err = util.GuidFromBytes(data[start : start+16])
		if err != nil {
			return nil, err
		}
		start += 16
	}
	sid, length, err := util.SidFromLEBytes(data[start:])
	if err != nil {
		return nil, err
	}
	var ApplicationData []byte
	if len(data) > start+int(length) {
		ApplicationData = data[length:]
	}

	var conditional *conditionalparser.ConditionalExpression
	if len(ApplicationData) > 0 {
		conditional, err = conditionalparser.ParseApplicationData(ApplicationData)
		if err != nil {
			return nil, err
		}
	}

	return &ACE{
		AccessMask:          AccessMask(accessMask),
		SID:                 sid,
		ObjectType:          objectType,
		InheritedObjectType: inheritedObjectType,
		ObjectFlags:         uniqueFlags,
		ApplicationData:     conditional,
	}, nil
}

func accessCallbackObjectAceToBytes(ace *ACE) ([]byte, error) {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data[0:4], uint32(ace.AccessMask))
	data = append(data, make([]byte, 4)...)
	binary.LittleEndian.PutUint32(data[4:8], uint32(ace.ObjectFlags))
	if ace.ObjectFlags&ACE_OBJECT_TYPE_PRESENT != 0 {

		objectType := util.GuidToBytes(ace.ObjectType)
		data = append(data, objectType...)
	}
	if ace.ObjectFlags&ACE_INHERITED_OBJECT_TYPE_PRESENT != 0 {

		inheritedObjectType := util.GuidToBytes(ace.InheritedObjectType)
		data = append(data, inheritedObjectType...)
	}
	sid, err := util.SidToLEBytes(ace.SID)
	if err != nil {
		return nil, err
	}
	data = append(data, sid...)

	var appDataBytes []byte
	if ace.ApplicationData != nil {
		appDataBytes, err = ace.ApplicationData.ToBytes()
		if err != nil {
			return nil, err
		}
	}
	data = append(data, appDataBytes...)

	return data, nil
}

// Access Mask - 4 bytes
// SID - variable
// Attribute Data - variable
func systemResourceAttributeAceFromBytes(data []byte) (*ACE, error) {
	if len(data) < 4 {
		return nil, errors.New(errInvalidACE)
	}
	accessMask := binary.LittleEndian.Uint32(data[0:4])
	sid, length, err := util.SidFromLEBytes(data[4:])
	if err != nil {
		return nil, err
	}
	var AttributeData []byte
	if len(data) > 4+int(length) {
		AttributeData = data[4+length:]
	}
	return &ACE{
		AccessMask:    AccessMask(accessMask),
		SID:           sid,
		AttributeData: AttributeData,
	}, nil
}

func systemResourceAttributeAceToBytes(ace *ACE) ([]byte, error) {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data[0:4], uint32(ace.AccessMask))
	sid, err := util.SidToLEBytes(ace.SID)
	if err != nil {
		return nil, err
	}
	data = append(data, sid...)
	data = append(data, ace.AttributeData...)

	return data, nil
}
