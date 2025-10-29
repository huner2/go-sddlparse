package sddlparse

import (
	"encoding/base64"
	"encoding/binary"
	"errors"

	conditionalparser "github.com/huner2/go-sddlparse/v2/internal/conditionalParser"
	"github.com/huner2/go-sddlparse/v2/internal/util"
)

// ACE is an Access Control Entry.
// This is a simplified version of the Windows ACE structure.
//
// This struct contains all fields for all the covered ACE types.
// Only the fields for the resolved ACE type will be populated.
type ACE struct {
	// Common fields
	Type       AceType
	Flags      AceFlag
	AccessMask AccessMask
	SID        string

	// Object-specific fields
	ObjectType          util.GUID
	InheritedObjectType util.GUID
	ObjectFlags         ObjectFlag

	// Callback-specific fields
	ApplicationData *conditionalparser.ConditionalExpression

	// System-resource-specific fields
	AttributeData []byte
}

// String returns the ACE as a string.
// Shortcut for MustString, returning an empty string on error.
func (ace *ACE) String() string {
	if str, err := ace.MustString(); err == nil {
		return str
	} else {
		return ""
	}
}

// MustString returns the ACE as a string.
// Returns an error if the ACE has unhandled elements.
func (ace *ACE) MustString() (string, error) {
	str, err := ace.Type.MustString()
	if err != nil {
		return "", err
	}
	str += ";"
	flags, err := ace.Flags.MustString()
	if err != nil {
		return "", err
	}
	str += flags
	str += ";"
	mask, err := ace.AccessMask.MustString()
	if err != nil {
		return "", err
	}
	str += mask
	str += ";"
	if !ace.ObjectType.IsNull() {
		str += ace.ObjectType.String()
	}
	str += ";"
	if !ace.InheritedObjectType.IsNull() {
		str += ace.InheritedObjectType.String()
	}
	str += ";"
	str += sidToSDDLAlias(ace.SID)
	if ace.Type == ACETYPE_ACCESS_ALLOWED_CALLBACK ||
		ace.Type == ACETYPE_ACCESS_DENIED_CALLBACK ||
		ace.Type == ACETYPE_ACCESS_ALLOWED_CALLBACK_OBJECT ||
		ace.Type == ACETYPE_ACCESS_DENIED_CALLBACK_OBJECT ||
		ace.Type == ACETYPE_SYSTEM_AUDIT_CALLBACK ||
		ace.Type == ACETYPE_SYSTEM_AUDIT_CALLBACK_OBJECT {
		str += ";"

		if ace.ApplicationData != nil {
			appStr, err := ace.ApplicationData.ToInfixString()
			if err != nil {
				return "", err
			}
			str += "(" + appStr + ")"
		}
	}
	return "(" + str + ")", nil
}

// ACL represents a Windows ACL.
// Simplified to a slice of ACEs.
type ACL []*ACE

// SDDL represents a Windows SDDL string.
// This is a subset of the SDDL format, only containing the fields we care about.
//
// See https://docs.microsoft.com/en-us/windows/win32/secauthz/security-descriptor-string-format for more information.
type SDDL struct {
	Version      byte
	ControlFlags uint16
	// Owner and Group are SIDs as strings.
	Owner string
	Group string
	SACL  ACL
	DACL  ACL
}

// String returns the SDDL as a string.
// Shortcut for MustString, returning an empty string on error.
func (sddl *SDDL) String() string {
	if str, err := sddl.MustString(); err == nil {
		return str
	} else {
		return ""
	}
}

// MustString returns the SDDL as a string.
// Returns an error if the SDDL has unhandled elements.
func (sddl *SDDL) MustString() (string, error) {
	str := ""
	if sddl.Owner != "" {
		str += "O:" + sidToSDDLAlias(sddl.Owner)
	}
	if sddl.Group != "" {
		str += "G:" + sidToSDDLAlias(sddl.Group)
	}
	/*
		The security descriptor string format does not support NULL ACLs.
		To denote an empty ACL, the security descriptor string includes the D: or S: token with no additional string information.
	*/
	if sddl.ControlFlags&SDDL_DACLPRESENT != 0 {
		str += "D:"
	}
	if sddl.DACL != nil {
		daclString := ""
		if sddl.ControlFlags&SDDL_DACLPROTECTED != 0 {
			daclString += "P"
		}
		if sddl.ControlFlags&SDDL_DACLINHERITANCEREQUIRED != 0 {
			daclString += "AR"
		}
		if sddl.ControlFlags&SDDL_DACLAUTOINHERITED != 0 {
			daclString += "AI"
		}
		for _, ace := range sddl.DACL {
			aceString, err := ace.MustString()
			if err != nil {
				return "", err
			}
			daclString += aceString
		}

		str += daclString
	}
	if sddl.ControlFlags&SDDL_SACLPRESENT != 0 {
		str += "S:"
	}
	if sddl.SACL != nil {
		saclString := ""
		if sddl.ControlFlags&SDDL_SACLPROTECTED != 0 {
			saclString += "P"
		}
		if sddl.ControlFlags&SDDL_INHERITEANCEREQUIRED != 0 {
			saclString += "AR"
		}
		if sddl.ControlFlags&SDDL_SACLAUTOINHERITED != 0 {
			saclString += "AI"
		}
		for _, ace := range sddl.SACL {
			aceString, err := ace.MustString()
			if err != nil {
				return "", err
			}
			saclString += aceString
		}

		str += saclString
	}
	return str, nil
}

func aclFromBytes(data []byte) (ACL, error) {
	if len(data) < 8 {
		return nil, errors.New(errInvalidACL)
	}

	revision := data[0]
	if revision != 0x02 && revision != 0x04 {
		return nil, errors.New(errInvalidACL)
	}
	if data[1] != 0x00 {
		return nil, errors.New(errInvalidACL)
	}
	aclSize := binary.LittleEndian.Uint16(data[2:4])
	aceCount := binary.LittleEndian.Uint16(data[4:6])
	if binary.LittleEndian.Uint16(data[6:8]) != 0x00 {
		return nil, errors.New(errInvalidACL)
	}
	if len(data) < int(aclSize) {
		return nil, errors.New(errInvalidACL)
	}
	acl := make(ACL, aceCount)
	tmp := data[8:]
	for i := 0; i < int(aceCount); i++ {
		ace, aceSize, err := aceFromBytes(tmp)
		if err != nil {

			return nil, err
		}
		acl[i] = ace
		tmp = tmp[aceSize:]
	}
	return acl, nil
}

func aclToBytes(acl ACL) ([]byte, error) {
	data := make([]byte, 8)
	data[0] = 0x04
	binary.LittleEndian.PutUint16(data[4:6], uint16(len(acl)))
	for _, ace := range acl {
		aceData, _, err := aceToBytes(ace)
		if err != nil {
			return nil, err
		}
		data = append(data, aceData...)
	}
	binary.LittleEndian.PutUint16(data[2:4], uint16(len(data)))
	return data, nil
}

// SDDLFromBinary decodes a binary SDDL string.
// This expects the generic binary format Windows uses.
//
// Returns an error if the SDDL is invalid.
// The SDDL will be deemed invalid for any of the following reasons:
//
//   - The SDDL is too short
//   - The SDDL does not start with 0x01 0x00
//   - The owner or group SID is invalid (if present)
//   - The SACL or DACL is invalid (if present)
//   - An unknown or unhandled ACE type is encountered
func SDDLFromBinary(data []byte) (*SDDL, error) {
	if len(data) < 16 {
		return nil, errors.New(errInvalidSDDL)
	}
	if data[0] != 0x01 {
		return nil, errors.New(errInvalidSDDL)
	}
	if data[1] != 0x00 {
		return nil, errors.New(errInvalidSDDL)
	}
	sddl := &SDDL{
		Version:      0x1,
		ControlFlags: binary.LittleEndian.Uint16(data[2:4]),
	}

	ownerOffset := binary.LittleEndian.Uint32(data[4:8])
	groupOffset := binary.LittleEndian.Uint32(data[8:12])
	saclOffset := binary.LittleEndian.Uint32(data[12:16])
	daclOffset := binary.LittleEndian.Uint32(data[16:20])

	if ownerOffset != 0 {
		if ownerOffset >= uint32(len(data)) {
			return nil, errors.New(errInvalidSDDL)
		}
		owner, _, err := util.SidFromLEBytes(data[ownerOffset:])
		if err != nil {

			return nil, err
		}
		sddl.Owner = owner
	}

	if groupOffset != 0 {
		if groupOffset >= uint32(len(data)) {
			return nil, errors.New(errInvalidSDDL)
		}
		group, _, err := util.SidFromLEBytes(data[groupOffset:])
		if err != nil {

			return nil, err
		}
		sddl.Group = group
	}

	if saclOffset != 0 {
		if saclOffset >= uint32(len(data)) {
			return nil, errors.New(errInvalidSDDL)
		}
		sacl, err := aclFromBytes(data[saclOffset:])
		if err != nil {

			return nil, err
		}
		sddl.SACL = sacl
	}

	if daclOffset != 0 {
		if daclOffset >= uint32(len(data)) {
			return nil, errors.New(errInvalidSDDL)
		}
		dacl, err := aclFromBytes(data[daclOffset:])
		if err != nil {

			return nil, err
		}
		sddl.DACL = dacl
	}

	return sddl, nil
}

// SDDLFromBase64Encoded decodes a base64 encoded SDDL string.
// This is a shortcut for base64.StdEncoding.DecodeString and SDDLFromBinary.
//
// Returns an error if the base64 decoding fails or if the SDDL is invalid.
func SDDLFromBase64Encoded(data []byte) (*SDDL, error) {
	b64, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return nil, err
	}
	return SDDLFromBinary(b64)
}

// SDDLFromString decodes the standard SDDL string format.
//
// Domain and machine identity are used to resolve SID aliases to SIDs.
// If not provided, then SIDs will look like S-1-5--500.
//
// Machine identity is likely unnecessary for any active directory SDDLs,
// whereas domain identity is likely unnecessary for any local machine results.
func SDDLFromString(sddl, domainIdentity, machineIdentity string) (*SDDL, error) {
	// Have to explicitly set version to 1 here, as it's not included in the string.
	obj := &SDDL{
		Version: 0x1,
	}
	datalen := len(sddl)

	if datalen < 2 {
		// Technically, I think this is valid, but it's not useful.
		return nil, errors.New(errInvalidSDDL)
	}
	if sddl[0] == 'O' {
		if datalen < 4 {
			return nil, errors.New(errInvalidSDDL)
		}
		if sddl[1] != ':' {
			return nil, errors.New(errInvalidSDDL)
		}
		if sddl[2:4] == "S-" {
			sid := "S-"
			i := 4
			for i < datalen && ((sddl[i] >= '0' && sddl[i] <= '9') || sddl[i] == '-') {
				sid += string(sddl[i])
				i++
			}
			obj.Owner = sid
			if i < datalen {
				sddl = sddl[i:]
			} else {
				return obj, nil
			}
		} else {
			sid, err := sddlAliasToSID(sddl[2:4], domainIdentity, machineIdentity)
			if err != nil {
				return nil, err
			}
			obj.Owner = sid
			if datalen > 4 {
				sddl = sddl[4:]
			} else {
				return obj, nil
			}
		}
	}

	datalen = len(sddl)
	if datalen < 4 {
		return nil, errors.New(errInvalidSDDL)
	}
	if sddl[0] == 'G' {
		if sddl[1] != ':' {
			return nil, errors.New(errInvalidSDDL)
		}
		if sddl[2:4] == "S-" {
			sid := "S-"
			i := 4
			for i < datalen && ((sddl[i] >= '0' && sddl[i] <= '9') || sddl[i] == '-') {
				sid += string(sddl[i])
				i++
			}
			obj.Group = sid
			if i < datalen {
				sddl = sddl[i:]
			} else {
				return obj, nil
			}
		} else {
			sid, err := sddlAliasToSID(sddl[2:4], domainIdentity, machineIdentity)
			if err != nil {
				return nil, err
			}
			obj.Group = sid
			if datalen > 4 {
				sddl = sddl[4:]
			} else {
				return obj, nil
			}
		}
	}

	datalen = len(sddl)
	if datalen < 4 {
		return nil, errors.New(errInvalidSDDL)
	}

	if sddl[0] == 'D' {
		if sddl[1] != ':' {
			return nil, errors.New(errInvalidSDDL)
		}
		obj.ControlFlags |= SDDL_DACLPRESENT
		flags := 0
		i := 2
		for i < datalen && (sddl[i] == 'P' || sddl[i] == 'A' || sddl[i] == 'R' || sddl[i] == 'I') {
			switch sddl[i] {
			case 'P':
				flags |= SDDL_DACLPROTECTED
			case 'A':
				if i+1 < datalen {
					switch sddl[i+1] {
					case 'R':
						flags |= SDDL_DACLINHERITANCEREQUIRED
					case 'I':
						flags |= SDDL_DACLAUTOINHERITED
					}
				}
			}
			i++
		}
		obj.ControlFlags |= uint16(flags)
		if i < datalen {
			sddl = sddl[i:]
			datalen = len(sddl)
		} else {
			return obj, nil
		}

		dacl, aclLen, err := aclFromString(sddl, domainIdentity, machineIdentity)
		if err != nil {
			return nil, err
		}
		obj.DACL = dacl
		if datalen > aclLen {
			sddl = sddl[aclLen:]
		} else {
			return obj, nil
		}
	}

	datalen = len(sddl)
	if datalen < 4 {
		return nil, errors.New(errInvalidSDDL)
	}

	if sddl[0] == 'S' {
		if sddl[1] != ':' {
			return nil, errors.New(errInvalidSDDL)
		}
		obj.ControlFlags |= SDDL_SACLPRESENT
		flags := 0
		i := 2
		for i < datalen && (sddl[i] == 'P' || sddl[i] == 'A' || sddl[i] == 'R' || sddl[i] == 'I') {
			switch sddl[i] {
			case 'P':
				flags |= SDDL_SACLPROTECTED
			case 'A':
				if i+1 < datalen {
					switch sddl[i+1] {
					case 'R':
						flags |= SDDL_INHERITEANCEREQUIRED
					case 'I':
						flags |= SDDL_SACLAUTOINHERITED
					}
				}
			}
			i++
		}
		obj.ControlFlags |= uint16(flags)
		if i < datalen {
			sddl = sddl[i:]
		} else {
			return obj, nil
		}

		sacl, aclLen, err := aclFromString(sddl, domainIdentity, machineIdentity)
		if err != nil {
			return nil, err
		}
		obj.SACL = sacl
		if datalen > aclLen {
			sddl = sddl[aclLen:]
		} else {
			return obj, nil
		}
	}

	if len(sddl) > 0 {
		return nil, errors.New(errInvalidSDDL)
	}

	return obj, nil
}

func aclFromString(sddl, domainIdentity, machineIdentity string) (ACL, int, error) {
	var acl ACL

	datalen := len(sddl)
	if datalen < 8 {
		return nil, 0, errors.New(errInvalidACL)
	}
	if sddl[0] != '(' {
		return nil, 0, errors.New(errInvalidACL)
	}
	inParen := true
	i := 1
	count := 0
	for inParen {
		count++
		parsedLen := 0
		if i+1 >= datalen {
			return nil, 0, errors.New(errInvalidACL)
		}
		curAce := &ACE{}
		typeString := sddl[i : i+2]
		extra := 1
		switch typeString {
		case "A;":
			curAce.Type = ACETYPE_ACCESS_ALLOWED
			extra = 0
		case "D;":
			curAce.Type = ACETYPE_ACCESS_DENIED
			extra = 0
		case "OA":
			curAce.Type = ACETYPE_ACCESS_ALLOWED_OBJECT
		case "OD":
			curAce.Type = ACETYPE_ACCESS_DENIED_OBJECT
		case "AU":
			curAce.Type = ACETYPE_SYSTEM_AUDIT
		case "AL":
			curAce.Type = ACETYPE_SYSTEM_ALARM
		case "AC":
			curAce.Type = ACETYPE_ACCESS_ALLOWED_COMPOUND
		case "OU":
			curAce.Type = ACETYPE_SYSTEM_AUDIT_OBJECT
		case "OL":
			curAce.Type = ACETYPE_SYSTEM_ALARM_OBJECT
		case "ML":
			curAce.Type = ACETYPE_SYSTEM_MANDATORY_LABEL
		case "XA":
			curAce.Type = ACETYPE_ACCESS_ALLOWED_CALLBACK
		case "XD":
			curAce.Type = ACETYPE_ACCESS_DENIED_CALLBACK
		case "RA":
			curAce.Type = ACETYPE_SYSTEM_RESOURCE_ATTRIBUTE
		case "SP":
			curAce.Type = ACETYPE_SYSTEM_SCOPED_POLICY_ID
		case "XU":
			curAce.Type = ACETYPE_SYSTEM_AUDIT_CALLBACK
		case "ZA":
			curAce.Type = ACETYPE_ACCESS_ALLOWED_CALLBACK_OBJECT
		default:
			return nil, 0, errors.New(errInvalidACE)
		}
		parsedLen += 2 + extra
		if i+parsedLen >= datalen {
			return nil, 0, errors.New(errInvalidACL)
		}
		if sddl[i+parsedLen-1] != ';' {
			return nil, 0, errors.New(errInvalidACL)
		}
		flags := AceFlag(0)
		for i+parsedLen+1 < datalen && sddl[i+parsedLen] != ';' {
			flagString := sddl[i+parsedLen : i+parsedLen+2]
			switch flagString {
			case "OI":
				flags |= ACEFLAG_OBJECT_INHERIT
			case "CI":
				flags |= ACEFLAG_CONTAINER_INHERIT
			case "NP":
				flags |= ACEFLAG_NO_PROPAGATE_INHERIT
			case "IO":
				flags |= ACEFLAG_INHERIT_ONLY
			case "ID":
				flags |= ACEFLAG_INHERITED
			case "SA":
				flags |= ACEFLAG_SUCCESSFUL_ACCESS
			case "FA":
				flags |= ACEFLAG_FAILED_ACCESS
			default:
				return nil, 0, errors.New(errInvalidACE)
			}
			parsedLen += 2
		}
		curAce.Flags = flags
		if i+parsedLen >= datalen {
			return nil, 0, errors.New(errInvalidACL)
		}
		if sddl[i+parsedLen] != ';' {
			return nil, 0, errors.New(errInvalidACL)
		}
		parsedLen++
		if i+parsedLen >= datalen {
			return nil, 0, errors.New(errInvalidACL)
		}
		accessMask := AccessMask(0)
		for i+parsedLen+1 < datalen && sddl[i+parsedLen] != ';' {
			maskString := sddl[i+parsedLen : i+parsedLen+2]
			switch maskString {
			case "GR":
				accessMask |= ACCESS_MASK_GENERIC_READ
			case "GW":
				accessMask |= ACCESS_MASK_GENERIC_WRITE
			case "GX":
				accessMask |= ACCESS_MASK_GENERIC_EXECUTE
			case "GA":
				accessMask |= ACCESS_MASK_GENERIC_ALL
			case "MA":
				accessMask |= ACCESS_MASK_MAXIMUM_ALLOWED
			case "AS":
				accessMask |= ACCESS_MASK_ACCESS_SYSTEM_SECURITY
			case "SY":
				accessMask |= ACCESS_MASK_SYNCHRONIZE
			case "WO":
				accessMask |= ACCESS_MASK_WRITE_OWNER
			case "WD":
				accessMask |= ACCESS_MASK_WRITE_DACL
			case "RC":
				accessMask |= ACCESS_MASK_READ_CONTROL
			case "SD":
				accessMask |= ACCESS_MASK_DELETE
			case "CC":
				accessMask |= ACCESS_MASK_ADS_RIGHT_DS_CREATE_CHILD
			case "DC":
				accessMask |= ACCESS_MASK_ADS_RIGHT_DS_DELETE_CHILD
			case "SW":
				accessMask |= ACCESS_MASK_ADS_RIGHT_DS_SELF
			case "RP":
				accessMask |= ACCESS_MASK_ADS_RIGHT_DS_READ_PROP
			case "WP":
				accessMask |= ACCESS_MASK_ADS_RIGHT_DS_WRITE_PROP
			case "CR":
				accessMask |= ACCESS_MASK_ADS_RIGHT_DS_CONTROL_ACCESS
			case "LC":
				accessMask |= ACCESS_MASK_ADS_RIGHT_DS_LIST_CONTENTS
			case "DT":
				accessMask |= ACCESS_MASK_ADS_RIGHT_DS_DELETE_TREE
			case "LO":
				accessMask |= ACCESS_MASK_ADS_RIGHT_DS_LIST_OBJECT
			default:
				return nil, 0, errors.New(errInvalidACE)
			}
			parsedLen += 2
		}
		curAce.AccessMask = accessMask
		if i+parsedLen >= datalen {
			return nil, 0, errors.New(errInvalidACL)
		}
		if sddl[i+parsedLen] != ';' {
			return nil, 0, errors.New(errInvalidACL)
		}
		parsedLen++
		if i+parsedLen >= datalen {
			return nil, 0, errors.New(errInvalidACL)
		}
		if sddl[i+parsedLen] != ';' {
			if i+parsedLen+36 >= datalen {
				return nil, 0, errors.New(errInvalidACL)
			}
			objectGuid := sddl[i+parsedLen : i+parsedLen+36]
			guid, err := util.GuidFromString(objectGuid)
			if err != nil {
				return nil, 0, err
			}
			curAce.ObjectType = *guid
			parsedLen += 36
			if i+parsedLen >= datalen {
				return nil, 0, errors.New(errInvalidACL)
			}
			if sddl[i+parsedLen] != ';' {
				return nil, 0, errors.New(errInvalidACL)
			}
		}
		parsedLen++
		if i+parsedLen >= datalen {
			return nil, 0, errors.New(errInvalidACL)
		}
		if sddl[i+parsedLen] != ';' {
			if i+parsedLen+36 >= datalen {
				return nil, 0, errors.New(errInvalidACL)
			}
			objectGuid := sddl[i+parsedLen : i+parsedLen+36]
			guid, err := util.GuidFromString(objectGuid)
			if err != nil {
				return nil, 0, err
			}
			curAce.InheritedObjectType = *guid
			parsedLen += 36
			if i+parsedLen >= datalen {
				return nil, 0, errors.New(errInvalidACL)
			}
			if sddl[i+parsedLen] != ';' {
				return nil, 0, errors.New(errInvalidACL)
			}
		}
		parsedLen++
		if i+parsedLen >= datalen {
			return nil, 0, errors.New(errInvalidACL)
		}
		if sddl[i+parsedLen] != ';' {
			if i+parsedLen+2 >= datalen {
				return nil, 0, errors.New(errInvalidACL)
			}
			sidString := sddl[i+parsedLen : i+parsedLen+2]
			if sidString == "S-" {
				sid := "S-"
				j := i + parsedLen + 2
				for j < datalen && ((sddl[j] >= '0' && sddl[j] <= '9') || sddl[j] == '-') {
					sid += string(sddl[j])
					j++
				}
				curAce.SID = sid
				parsedLen += j - (i + parsedLen)
			} else {
				sid, err := sddlAliasToSID(sidString, domainIdentity, machineIdentity)
				if err != nil {

					return nil, 0, err
				}
				curAce.SID = sid
				parsedLen += 2
			}
		}
		if i+parsedLen >= datalen {
			return nil, 0, errors.New(errInvalidACL)
		}
		if curAce.Type == ACETYPE_ACCESS_ALLOWED_CALLBACK ||
			curAce.Type == ACETYPE_ACCESS_DENIED_CALLBACK ||
			curAce.Type == ACETYPE_ACCESS_ALLOWED_CALLBACK_OBJECT ||
			curAce.Type == ACETYPE_ACCESS_DENIED_CALLBACK_OBJECT ||
			curAce.Type == ACETYPE_SYSTEM_AUDIT_CALLBACK ||
			curAce.Type == ACETYPE_SYSTEM_AUDIT_CALLBACK_OBJECT {
			if i+parsedLen+2 > datalen {
				return nil, 0, errors.New(errInvalidACL)
			}
			openCount := 0
			closedCount := 0
			conditionalString := ""
			if sddl[i+parsedLen] != ';' || sddl[i+parsedLen+1] != '(' {
				return nil, 0, errors.New(errInvalidACL)
			}
			parsedLen += 1
			for j := i + parsedLen; j < datalen; j++ {
				if sddl[j] == '(' {
					openCount++
				} else if sddl[j] == ')' {
					closedCount++
				}
				conditionalString += string(sddl[j])
				if openCount > 0 && openCount == closedCount {
					parsedLen += j - (i + parsedLen) + 1
					break
				}
			}
			if openCount != closedCount {
				return nil, 0, errors.New(errInvalidACL)
			}
			condExpr, err := conditionalparser.ParseConditionalExpression(conditionalString)
			if err != nil {
				return nil, 0, err
			}
			curAce.ApplicationData = condExpr
		}

		if sddl[i+parsedLen] != ')' {
			return nil, 0, errors.New(errInvalidACL)
		}
		parsedLen++
		if i+parsedLen >= datalen {
			inParen = false
			i += parsedLen
		} else {
			if sddl[i+parsedLen] != '(' {
				inParen = false
			} else {
				parsedLen++
			}
			i += parsedLen
		}
		acl = append(acl, curAce)
	}

	return acl, i, nil
}

// ToBinary encodes the SDDL as a binary string.
// This is the generic binary format Windows uses.
//
// Returns an error if the SDDL is invalid.
func (sddl *SDDL) ToBinary() ([]byte, error) {
	var data []byte

	if sddl.Version != 0x1 {
		return nil, errors.New(errInvalidSDDL)
	}
	data = append(data, 0x01)
	data = append(data, 0x00)
	data = append(data, byte(sddl.ControlFlags))
	data = append(data, byte(sddl.ControlFlags>>8))

	data = append(data, 0x00, 0x00, 0x00, 0x00) // Owner offset
	data = append(data, 0x00, 0x00, 0x00, 0x00) // Group offset
	data = append(data, 0x00, 0x00, 0x00, 0x00) // SACL offset
	data = append(data, 0x00, 0x00, 0x00, 0x00) // DACL offset

	if sddl.Owner != "" {
		sid, err := util.SidToLEBytes(sddl.Owner)
		if err != nil {
			return nil, err
		}
		data = append(data, sid...)

		binary.LittleEndian.PutUint32(data[4:8], uint32(len(data)-len(sid)))
	}

	if sddl.Group != "" {
		sid, err := util.SidToLEBytes(sddl.Group)
		if err != nil {
			return nil, err
		}
		data = append(data, sid...)

		binary.LittleEndian.PutUint32(data[8:12], uint32(len(data)-len(sid)))
	}

	if sddl.SACL != nil {
		sacl, err := aclToBytes(sddl.SACL)
		if err != nil {
			return nil, err
		}
		data = append(data, sacl...)

		binary.LittleEndian.PutUint32(data[12:16], uint32(len(data)-len(sacl)))
	}

	if sddl.DACL != nil {
		dacl, err := aclToBytes(sddl.DACL)
		if err != nil {
			return nil, err
		}
		data = append(data, dacl...)

		binary.LittleEndian.PutUint32(data[16:20], uint32(len(data)-len(dacl)))
	}

	return data, nil
}

// ToBase64Encoded encodes the SDDL as a base64 byte array.
// This is a shortcut for base64.StdEncoding.Encode and SDDLToBinary.
//
// Returns an error if the SDDL is invalid.
func (sddl *SDDL) ToBase64Encoded() ([]byte, error) {
	bin, err := sddl.ToBinary()
	if err != nil {
		return nil, err
	}
	b64 := make([]byte, base64.StdEncoding.EncodedLen(len(bin)))
	base64.StdEncoding.Encode(b64, bin)
	return b64, nil
}

// CreateConditionalExpression creates a conditional expression from a string.
// This can be used to create (or modify) the ApplicationData field of a conditional ACE.
func CreateConditionalExpression(condStr string) (*conditionalparser.ConditionalExpression, error) {
	return conditionalparser.ParseConditionalExpression(condStr)
}
