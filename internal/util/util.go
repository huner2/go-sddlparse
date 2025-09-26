package util

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const errInvalidSID = "invalid SID"
const errInvalidGUID = "invalid GUID"

func SidFromLEBytes(data []byte) (string, uint32, error) {
	if len(data) < 8 {

		return "", 0, errors.New(errInvalidSID)
	}
	revision := data[0]
	if revision != 0x01 {

		return "", 0, errors.New(errInvalidSID)
	}
	subAuthorityCount := data[1]
	if subAuthorityCount < 1 || subAuthorityCount > 15 || len(data) < 8+int(subAuthorityCount)*4 {
		return "", 0, errors.New(errInvalidSID)
	}
	identifierAuthority := uint64(data[2])<<40 | uint64(data[3])<<32 | uint64(data[4])<<24 | uint64(data[5])<<16 | uint64(data[6])<<8 | uint64(data[7])
	subAuthorities := make([]string, subAuthorityCount)
	for i := 0; i < int(subAuthorityCount); i++ {
		subAuthorities[i] = fmt.Sprint(binary.LittleEndian.Uint32(data[8+i*4 : 12+i*4]))
	}

	return "S-1-" + fmt.Sprint(identifierAuthority) + "-" + strings.Join(subAuthorities, "-"), uint32(8 + subAuthorityCount*4), nil
}

func SidToLEBytes(sid string) ([]byte, error) {
	if len(sid) == 0 {
		data := make([]byte, 8)
		data[0] = 0x01
		return data, nil
	}
	sid = strings.TrimPrefix(sid, "S-1-")
	sections := strings.Split(sid, "-")
	if len(sections) < 2 || len(sections) > 15 {
		return nil, errors.New(errInvalidSID)
	}
	identifierAuthority, err := strconv.ParseUint(sections[0], 10, 48)
	if err != nil {
		return nil, errors.New(errInvalidSID)
	}
	subAuthorityCount := len(sections) - 1
	if subAuthorityCount < 1 || subAuthorityCount > 15 {
		return nil, errors.New(errInvalidSID)
	}
	data := make([]byte, 8+subAuthorityCount*4)
	data[0] = 0x01
	data[1] = byte(subAuthorityCount)
	data[2] = byte(identifierAuthority >> 40)
	data[3] = byte(identifierAuthority >> 32)
	data[4] = byte(identifierAuthority >> 24)
	data[5] = byte(identifierAuthority >> 16)
	data[6] = byte(identifierAuthority >> 8)
	data[7] = byte(identifierAuthority)
	for i := 0; i < subAuthorityCount; i++ {
		subAuthority, err := strconv.ParseUint(sections[i+1], 10, 32)
		if err != nil {
			return nil, errors.New(errInvalidSID)
		}
		binary.LittleEndian.PutUint32(data[8+i*4:12+i*4], uint32(subAuthority))
	}
	if len(data)%4 != 0 {
		data = append(data, make([]byte, len(data)%4)...)
	}

	return data, nil
}

// GUID represents a Windows GUID.
//
// Redefining this here to avoid importing golang.org/x/sys/windows.
type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

func GuidFromBytes(data []byte) (GUID, error) {
	if len(data) < 16 {
		return GUID{}, errors.New(errInvalidGUID)
	}
	return GUID{
		Data1: binary.LittleEndian.Uint32(data[0:4]),
		Data2: binary.LittleEndian.Uint16(data[4:6]),
		Data3: binary.LittleEndian.Uint16(data[6:8]),
		Data4: [8]byte{data[8], data[9], data[10], data[11], data[12], data[13], data[14], data[15]},
	}, nil
}

func GuidToBytes(guid GUID) []byte {
	data := make([]byte, 16)
	binary.LittleEndian.PutUint32(data[0:4], guid.Data1)
	binary.LittleEndian.PutUint16(data[4:6], guid.Data2)
	binary.LittleEndian.PutUint16(data[6:8], guid.Data3)
	data[8] = guid.Data4[0]
	data[9] = guid.Data4[1]
	data[10] = guid.Data4[2]
	data[11] = guid.Data4[3]
	data[12] = guid.Data4[4]
	data[13] = guid.Data4[5]
	data[14] = guid.Data4[6]
	data[15] = guid.Data4[7]
	return data
}

// IsNull returns true if the GUID is null.
func (g *GUID) IsNull() bool {
	return g.Data1 == 0 && g.Data2 == 0 && g.Data3 == 0 && g.Data4 == [8]byte{0, 0, 0, 0, 0, 0, 0, 0}
}

// String returns the GUID as a string.
//
// Format: {xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}
func (guid *GUID) String() string {
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x", guid.Data1, guid.Data2, guid.Data3, guid.Data4[0:2], guid.Data4[2:8])
}

// GuidFromString returns a GUID from a string.
//
// Format: {xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx} or xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
//
// Returns an error if the string is not a valid GUID.
func GuidFromString(gstring string) (*GUID, error) {
	var guid GUID
	gstring = strings.Trim(gstring, "{}")
	if len(gstring) != 36 {
		return nil, errors.New("invalid length guid string")
	}
	sections := strings.Split(gstring, "-")
	if len(sections) != 5 {
		return nil, errors.New("invalid guid string")
	}
	data1, err := strconv.ParseUint(sections[0], 16, 32)
	if err != nil {
		return nil, err
	}
	data2, err := strconv.ParseUint(sections[1], 16, 16)
	if err != nil {
		return nil, err
	}
	data3, err := strconv.ParseUint(sections[2], 16, 16)
	if err != nil {
		return nil, err
	}
	data4, err := strconv.ParseUint(sections[3]+sections[4], 16, 64)
	if err != nil {
		return nil, err
	}
	guid.Data1 = uint32(data1)
	guid.Data2 = uint16(data2)
	guid.Data3 = uint16(data3)
	guid.Data4 = [8]byte{byte(data4 >> 56), byte(data4 >> 48), byte(data4 >> 40), byte(data4 >> 32), byte(data4 >> 24), byte(data4 >> 16), byte(data4 >> 8), byte(data4)}
	return &guid, nil
}
