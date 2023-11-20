package sddlparse

import (
	"fmt"
	"strings"
)

// Error messages
const errInvalidSDDL = "invalid SDDL"
const errInvalidSID = "invalid SID"
const errInvalidACE = "invalid ACE"
const errInvalidACL = "invalid ACL"
const errInvalidGUID = "invalid GUID"

// SDDL Control Flags
const (
	SDDL_OWNERDEFAULTED = 1 << iota
	SDDL_GROUPDEFAULTED
	SDDL_DACLPRESENT
	SDDL_DACLDEFAULTED
	SDDL_SACLPRESENT
	SDDL_SACLDEFAULTED
	SDDL_SERVERSECURITY
	SDDL_DACLTRUSTED
	SDDL_DACLINHERITANCEREQUIRED
	SDDL_INHERITEANCEREQUIRED
	SDDL_DACLAUTOINHERITED
	SDDL_SACLAUTOINHERITED
	SDDL_DACLPROTECTED
	SDDL_SACLPROTECTED
	SDDL_CONTROLVALID
	SDDL_SELFRELATIVE
)

// Access Mask Flags
type AccessMask uint32

const (
	ACCESS_MASK_GENERIC_READ           AccessMask = 0x80000000
	ACCESS_MASK_GENERIC_WRITE          AccessMask = 0x40000000
	ACCESS_MASK_GENERIC_EXECUTE        AccessMask = 0x20000000
	ACCESS_MASK_GENERIC_ALL            AccessMask = 0x10000000
	ACCESS_MASK_MAXIMUM_ALLOWED        AccessMask = 0x02000000
	ACCESS_MASK_ACCESS_SYSTEM_SECURITY AccessMask = 0x01000000
	ACCESS_MASK_SYNCHRONIZE            AccessMask = 0x00100000
	ACCESS_MASK_WRITE_OWNER            AccessMask = 0x00080000
	ACCESS_MASK_WRITE_DACL             AccessMask = 0x00040000
	ACCESS_MASK_READ_CONTROL           AccessMask = 0x00020000
	ACCESS_MASK_DELETE                 AccessMask = 0x00010000

	ACCESS_MASK_ADS_RIGHT_DS_CREATE_CHILD   AccessMask = 0x00000001
	ACCESS_MASK_ADS_RIGHT_DS_DELETE_CHILD   AccessMask = 0x00000002
	ACCESS_MASK_ADS_RIGHT_DS_LIST_CONTENTS  AccessMask = 0x00000004
	ACCESS_MASK_ADS_RIGHT_DS_SELF           AccessMask = 0x00000008
	ACCESS_MASK_ADS_RIGHT_DS_READ_PROP      AccessMask = 0x00000010
	ACCESS_MASK_ADS_RIGHT_DS_WRITE_PROP     AccessMask = 0x00000020
	ACCESS_MASK_ADS_RIGHT_DS_DELETE_TREE    AccessMask = 0x00000040
	ACCESS_MASK_ADS_RIGHT_DS_LIST_OBJECT    AccessMask = 0x00000080
	ACCESS_MASK_ADS_RIGHT_DS_CONTROL_ACCESS AccessMask = 0x00000100
)

// String returns the access mask as a string. If the access mask is not known,
// it returns the mask as a hex string.
func (mask *AccessMask) String() string {
	if str, err := mask.MustString(); err == nil {
		return str
	} else {
		return fmt.Sprintf("%08x", *mask)
	}
}

// MustString returns the access mask as a string. If the access mask is not known
// or it doesn't have a string representation, it returns an error.
func (mask *AccessMask) MustString() (string, error) {
	maskString := ""
	if *mask&ACCESS_MASK_GENERIC_READ != 0 {
		maskString += "GR"
	}
	if *mask&ACCESS_MASK_GENERIC_WRITE != 0 {
		maskString += "GW"
	}
	if *mask&ACCESS_MASK_GENERIC_EXECUTE != 0 {
		maskString += "GX"
	}
	if *mask&ACCESS_MASK_GENERIC_ALL != 0 {
		maskString += "GA"
	}
	if *mask&ACCESS_MASK_MAXIMUM_ALLOWED != 0 {
		maskString += "MA"
	}
	if *mask&ACCESS_MASK_ACCESS_SYSTEM_SECURITY != 0 {
		maskString += "AS"
	}
	if *mask&ACCESS_MASK_SYNCHRONIZE != 0 {
		maskString += "SY"
	}
	if *mask&ACCESS_MASK_WRITE_OWNER != 0 {
		maskString += "WO"
	}
	if *mask&ACCESS_MASK_WRITE_DACL != 0 {
		maskString += "WD"
	}
	if *mask&ACCESS_MASK_READ_CONTROL != 0 {
		maskString += "RC"
	}
	if *mask&ACCESS_MASK_DELETE != 0 {
		maskString += "DE"
	}
	if *mask&ACCESS_MASK_ADS_RIGHT_DS_CREATE_CHILD != 0 {
		maskString += "CC"
	}
	if *mask&ACCESS_MASK_ADS_RIGHT_DS_DELETE_CHILD != 0 {
		maskString += "DC"
	}
	if *mask&ACCESS_MASK_ADS_RIGHT_DS_SELF != 0 {
		maskString += "SW"
	}
	if *mask&ACCESS_MASK_ADS_RIGHT_DS_READ_PROP != 0 {
		maskString += "RP"
	}
	if *mask&ACCESS_MASK_ADS_RIGHT_DS_WRITE_PROP != 0 {
		maskString += "WP"
	}
	if *mask&ACCESS_MASK_ADS_RIGHT_DS_CONTROL_ACCESS != 0 {
		maskString += "CR"
	}
	if *mask&ACCESS_MASK_ADS_RIGHT_DS_LIST_CONTENTS != 0 {
		maskString += "LC"
	}
	if *mask&ACCESS_MASK_ADS_RIGHT_DS_DELETE_TREE != 0 {
		maskString += "DT"
	}
	if *mask&ACCESS_MASK_ADS_RIGHT_DS_LIST_OBJECT != 0 {
		maskString += "LO"
	}
	if maskString == "" && *mask != 0 {
		return "", fmt.Errorf("unknown access mask: %08x", *mask)
	}
	return maskString, nil
}

// Ace Types
type AceType byte

/*
Not all conditional ACE types are supported in the SDDL. The conditional ACE
types ACCESS_ALLOWED_CALLBACK_ACE and ACCESS_DENIED_CALLBACK_ACE are not supported in
Windows Vista and earlier client releases or Windows Server 2008 and earlier server releases. The
conditional ACE types ACCESS_ALLOWED_CALLBACK_OBJECT_ACE and
SYSTEM_AUDIT_CALLBACK_ACE are not supported in Windows 7 or Windows Server 2008 R2.
*/
const (
	ACETYPE_ACCESS_ALLOWED AceType = iota
	ACETYPE_ACCESS_DENIED
	ACETYPE_SYSTEM_AUDIT
	ACETYPE_SYSTEM_ALARM
	ACETYPE_ACCESS_ALLOWED_COMPOUND
	ACETYPE_ACCESS_ALLOWED_OBJECT
	ACETYPE_ACCESS_DENIED_OBJECT
	ACETYPE_SYSTEM_AUDIT_OBJECT
	ACETYPE_SYSTEM_ALARM_OBJECT
	ACETYPE_ACCESS_ALLOWED_CALLBACK
	ACETYPE_ACCESS_DENIED_CALLBACK
	ACETYPE_ACCESS_ALLOWED_CALLBACK_OBJECT
	ACETYPE_ACCESS_DENIED_CALLBACK_OBJECT
	ACETYPE_SYSTEM_AUDIT_CALLBACK
	ACETYPE_SYSTEM_ALARM_CALLBACK
	ACETYPE_SYSTEM_AUDIT_CALLBACK_OBJECT
	ACETYPE_SYSTEM_ALARM_CALLBACK_OBJECT
	ACETYPE_SYSTEM_MANDATORY_LABEL
	ACETYPE_SYSTEM_RESOURCE_ATTRIBUTE
	ACETYPE_SYSTEM_SCOPED_POLICY_ID
)

// String returns the ACE type as a string. If the ACE type is not known, it
// returns the type as a hex string.
//
// Some of the ACE types don't have a string representation that I can find.
//
// https://learn.microsoft.com/en-us/windows/win32/secauthz/ace-strings
func (aceType *AceType) String() string {
	if str, err := aceType.MustString(); err == nil {
		return str
	} else {
		return fmt.Sprintf("%02x", *aceType)
	}
}

// MustString returns the ACE type as a string. If the ACE type is not known
// or it doesn't have a string representation, it returns an error.
func (aceType *AceType) MustString() (string, error) {
	if *aceType == 0 {
		return "A", nil
	} else if *aceType == ACETYPE_ACCESS_DENIED {
		return "D", nil
	} else if *aceType == ACETYPE_ACCESS_ALLOWED_OBJECT {
		return "OA", nil
	} else if *aceType == ACETYPE_ACCESS_DENIED_OBJECT {
		return "OD", nil
	} else if *aceType == ACETYPE_SYSTEM_AUDIT {
		return "AU", nil
	} else if *aceType == ACETYPE_SYSTEM_ALARM {
		return "AL", nil
	} else if *aceType == ACETYPE_ACCESS_ALLOWED_COMPOUND {
		return "AC", nil
	} else if *aceType == ACETYPE_SYSTEM_AUDIT_OBJECT {
		return "OU", nil
	} else if *aceType == ACETYPE_SYSTEM_ALARM_OBJECT {
		return "OL", nil
	} else if *aceType == ACETYPE_SYSTEM_MANDATORY_LABEL {
		return "ML", nil
	} else if *aceType == ACETYPE_ACCESS_ALLOWED_CALLBACK {
		return "XA", nil
	} else if *aceType == ACETYPE_ACCESS_DENIED_CALLBACK {
		return "XD", nil
	} else if *aceType == ACETYPE_SYSTEM_RESOURCE_ATTRIBUTE {
		return "RA", nil
	} else if *aceType == ACETYPE_SYSTEM_SCOPED_POLICY_ID {
		return "SP", nil
	} else if *aceType == ACETYPE_SYSTEM_AUDIT_CALLBACK {
		return "XU", nil
	} else if *aceType == ACETYPE_ACCESS_ALLOWED_CALLBACK_OBJECT {
		return "ZA", nil
	} else {
		return "", fmt.Errorf("unknown ACE type: %02x", *aceType)
	}
}

// Ace Flags
type AceFlag byte

const (
	ACEFLAG_OBJECT_INHERIT       AceFlag = 0x01
	ACEFLAG_CONTAINER_INHERIT    AceFlag = 0x02
	ACEFLAG_NO_PROPAGATE_INHERIT AceFlag = 0x04
	ACEFLAG_INHERIT_ONLY         AceFlag = 0x08
	ACEFLAG_INHERITED            AceFlag = 0x10
	ACEFLAG_SUCCESSFUL_ACCESS    AceFlag = 0x40
	ACEFLAG_FAILED_ACCESS        AceFlag = 0x80
)

// String returns the ACE flag as a string. If the ACE flag is not known, it
// returns the flag as a hex string.
func (aceFlag *AceFlag) String() string {
	if str, err := aceFlag.MustString(); err == nil {
		return str
	} else {
		return fmt.Sprintf("%02x", *aceFlag)
	}
}

// MustString returns the ACE flag as a string. If the ACE flag is not known
// or it doesn't have a string representation, it returns an error.
func (aceFlag *AceFlag) MustString() (string, error) {
	flagString := ""
	if *aceFlag&ACEFLAG_OBJECT_INHERIT != 0 {
		flagString += "OI"
	}
	if *aceFlag&ACEFLAG_CONTAINER_INHERIT != 0 {
		flagString += "CI"
	}
	if *aceFlag&ACEFLAG_NO_PROPAGATE_INHERIT != 0 {
		flagString += "NP"
	}
	if *aceFlag&ACEFLAG_INHERIT_ONLY != 0 {
		flagString += "IO"
	}
	if *aceFlag&ACEFLAG_INHERITED != 0 {
		flagString += "ID"
	}
	if *aceFlag&ACEFLAG_SUCCESSFUL_ACCESS != 0 {
		flagString += "SA"
	}
	if *aceFlag&ACEFLAG_FAILED_ACCESS != 0 {
		flagString += "FA"
	}
	if flagString == "" && *aceFlag != 0 {
		return "", fmt.Errorf("unknown ACE flag: %02x", *aceFlag)
	}
	return flagString, nil
}

type ObjectFlag byte

const (
	ACE_OBJECT_INVALID                ObjectFlag = 0x00
	ACE_OBJECT_TYPE_PRESENT           ObjectFlag = 0x01
	ACE_INHERITED_OBJECT_TYPE_PRESENT ObjectFlag = 0x02
)

func sidToSDDLAlias(sid string) string {
	if strings.HasPrefix(sid, "S-1-5-21-") && strings.HasSuffix(sid, "512") {
		return "DA"
	} else if strings.HasPrefix(sid, "S-1-5-21-") && strings.HasSuffix(sid, "514") {
		return "DG"
	} else if sid == "S-1-5-9" {
		return "ED"
	} else if strings.HasPrefix(sid, "S-1-5-21-") && strings.HasSuffix(sid, "516") {
		return "DD"
	} else if strings.HasPrefix(sid, "S-1-5-21-") && strings.HasSuffix(sid, "515") {
		return "DC"
	} else if sid == "S-1-5-32-544" {
		return "BA"
	} else if sid == "S-1-5-32-546" {
		return "BG"
	} else if sid == "S-1-5-32-545" {
		return "BU"
	} else if strings.HasPrefix(sid, "S-1-5-21-") && strings.HasSuffix(sid, "500") {
		// Technically, the domain admin SID is S-1-5-21-<domain>-500, but
		// we don't know the domain, so we'll just have to use the sid
		// as we can't distinguish between a local admin and a domain admin.
		//
		// Returning the SID gives the most information, so we'll do that.
		return sid
	} else if strings.HasPrefix(sid, "S-1-5-21-") && strings.HasSuffix(sid, "501") {
		return "LG"
	} else if sid == "S-1-5-32-548" {
		return "AO"
	} else if sid == "S-1-5-32-551" {
		return "BO"
	} else if sid == "S-1-5-32-550" {
		return "PO"
	} else if sid == "S-1-5-32-549" {
		return "SO"
	} else if sid == "S-1-5-11" {
		return "AU"
	} else if sid == "S-1-5-10" {
		return "PS"
	} else if sid == "S-1-3-0" {
		return "CO"
	} else if sid == "S-1-3-1" {
		return "CG"
	} else if sid == "S-1-5-18" {
		return "SY"
	} else if sid == "S-1-5-32-547" {
		return "PU"
	} else if sid == "S-1-1-0" {
		return "WD"
	} else if sid == "S-1-5-32-552" {
		return "RE"
	} else if sid == "S-1-5-4" {
		return "IU"
	} else if sid == "S-1-5-2" {
		return "NU"
	} else if sid == "S-1-5-6" {
		return "SU"
	} else if sid == "S-1-5-12" {
		return "RC"
	} else if sid == "S-1-5-33" {
		return "WR"
	} else if sid == "S-1-5-7" {
		return "AN"
	} else if strings.HasPrefix(sid, "S-1-5-21-") && strings.HasSuffix(sid, "518") {
		return "SA"
	} else if strings.HasPrefix(sid, "S-1-5-21-") && strings.HasSuffix(sid, "517") {
		return "CA"
	} else if strings.HasPrefix(sid, "S-1-5-21-") && strings.HasSuffix(sid, "553") {
		return "RS"
	} else if strings.HasPrefix(sid, "S-1-5-21-") && strings.HasSuffix(sid, "519") {
		return "EA"
	} else if strings.HasPrefix(sid, "S-1-5-21-") && strings.HasSuffix(sid, "520") {
		return "PA"
	} else if sid == "S-1-5-32-554" {
		return "RU"
	} else if sid == "S-1-5-19" {
		return "LS"
	} else if sid == "S-1-5-20" {
		return "NS"
	} else if sid == "S-1-5-32-555" {
		return "RD"
	} else if sid == "S-1-5-32-556" {
		return "NO"
	} else if sid == "S-1-5-32-558" {
		return "MU"
	} else if sid == "S-1-5-32-559" {
		return "LU"
	} else if sid == "S-1-5-32-568" {
		return "IS"
	} else if sid == "S-1-5-32-569" {
		return "CY"
	} else if sid == "S-1-3-4" {
		return "OW"
	} else if sid == "S-1-5-32-573" {
		return "ER"
	} else if strings.HasPrefix(sid, "S-1-5-21-") && strings.HasSuffix(sid, "498") {
		return "RO"
	} else if sid == "S-1-5-32-569" {
		return "CD"
	} else if sid == "S-1-15-2-1" {
		return "AC"
	} else if sid == "S-1-5-32-575" {
		return "RA"
	} else if sid == "S-1-5-32-576" {
		return "ES"
	} else if sid == "S-1-5-32-577" {
		return "MS"
	} else if sid == "S-1-5-84-0-0-0-0-0" {
		return "UD"
	} else if sid == "S-1-5-32-578" {
		return "HA"
	} else if strings.HasPrefix(sid, "S-1-5-21-") && strings.HasSuffix(sid, "522") {
		return "CN"
	} else if sid == "S-1-5-32-579" {
		return "AA"
	} else if sid == "S-1-5-32-580" {
		return "RM"
	} else if sid == "S-1-16-4096" {
		return "LW"
	} else if sid == "S-1-16-8192" {
		return "ME"
	} else if sid == "S-1-16-8448" {
		return "MP"
	} else if sid == "S-1-16-12288" {
		return "HI"
	} else if sid == "S-1-16-16384" {
		return "SI"
	} else {
		return sid
	}
}

func sddlAliasToSID(alias, domainIdentity, machineIdentity string) (string, error) {
	if alias == "DA" {
		return "S-1-5-21-" + domainIdentity + "-512", nil
	} else if alias == "DG" {
		return "S-1-5-21-" + domainIdentity + "-514", nil
	} else if alias == "ED" {
		return "S-1-5-9", nil
	} else if alias == "DD" {
		return "S-1-5-21-" + domainIdentity + "-516", nil
	} else if alias == "DC" {
		return "S-1-5-21-" + domainIdentity + "-515", nil
	} else if alias == "BA" {
		return "S-1-5-32-544", nil
	} else if alias == "BG" {
		return "S-1-5-32-546", nil
	} else if alias == "BU" {
		return "S-1-5-32-545", nil
	} else if alias == "LA" {
		return "S-1-5-21-" + machineIdentity + "-500", nil
	} else if alias == "LG" {
		return "S-1-5-21-" + machineIdentity + "-501", nil
	} else if alias == "AO" {
		return "S-1-5-32-548", nil
	} else if alias == "BO" {
		return "S-1-5-32-551", nil
	} else if alias == "PO" {
		return "S-1-5-32-550", nil
	} else if alias == "SO" {
		return "S-1-5-32-549", nil
	} else if alias == "AU" {
		return "S-1-5-11", nil
	} else if alias == "PS" {
		return "S-1-5-10", nil
	} else if alias == "CO" {
		return "S-1-3-0", nil
	} else if alias == "CG" {
		return "S-1-3-1", nil
	} else if alias == "SY" {
		return "S-1-5-18", nil
	} else if alias == "PU" {
		return "S-1-5-32-547", nil
	} else if alias == "WD" {
		return "S-1-1-0", nil
	} else if alias == "RE" {
		return "S-1-5-32-552", nil
	} else if alias == "IU" {
		return "S-1-5-4", nil
	} else if alias == "NU" {
		return "S-1-5-2", nil
	} else if alias == "SU" {
		return "S-1-5-6", nil
	} else if alias == "RC" {
		return "S-1-5-12", nil
	} else if alias == "WR" {
		return "S-1-5-33", nil
	} else if alias == "AN" {
		return "S-1-5-7", nil
	} else if alias == "SA" {
		return "S-1-5-21-" + domainIdentity + "-518", nil
	} else if alias == "CA" {
		return "S-1-5-21-" + domainIdentity + "-517", nil
	} else if alias == "RS" {
		return "S-1-5-21-" + domainIdentity + "-553", nil
	} else if alias == "EA" {
		return "S-1-5-21-" + domainIdentity + "-519", nil
	} else if alias == "PA" {
		return "S-1-5-21-" + domainIdentity + "-520", nil
	} else if alias == "RU" {
		return "S-1-5-32-554", nil
	} else if alias == "LS" {
		return "S-1-5-19", nil
	} else if alias == "NS" {
		return "S-1-5-20", nil
	} else if alias == "RD" {
		return "S-1-5-32-555", nil
	} else if alias == "NO" {
		return "S-1-5-32-556", nil
	} else if alias == "MU" {
		return "S-1-5-32-558", nil
	} else if alias == "LU" {
		return "S-1-5-32-559", nil
	} else if alias == "IS" {
		return "S-1-5-32-568", nil
	} else if alias == "CY" {
		return "S-1-5-32-569", nil
	} else if alias == "OW" {
		return "S-1-3-4", nil
	} else if alias == "ER" {
		return "S-1-5-32-573", nil
	} else if alias == "RO" {
		return "S-1-5-21-" + domainIdentity + "-498", nil
	} else if alias == "CD" {
		return "S-1-5-32-569", nil
	} else if alias == "AC" {
		return "S-1-15-2-1", nil
	} else if alias == "RA" {
		return "S-1-5-32-575", nil
	} else if alias == "ES" {
		return "S-1-5-32-576", nil
	} else if alias == "MS" {
		return "S-1-5-32-577", nil
	} else if alias == "UD" {
		return "S-1-5-84-0-0-0-0-0", nil
	} else if alias == "HA" {
		return "S-1-5-32-578", nil
	} else if alias == "CN" {
		return "S-1-5-21-" + domainIdentity + "-522", nil
	} else if alias == "AA" {
		return "S-1-5-32-579", nil
	} else if alias == "RM" {
		return "S-1-5-32-580", nil
	} else if alias == "LW" {
		return "S-1-16-4096", nil
	} else if alias == "ME" {
		return "S-1-16-8192", nil
	} else if alias == "MP" {
		return "S-1-16-8448", nil
	} else if alias == "HI" {
		return "S-1-16-12288", nil
	} else if alias == "SI" {
		return "S-1-16-16384", nil
	} else {
		return "", fmt.Errorf("unknown alias: %s", alias)
	}
}
