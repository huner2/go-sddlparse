package conditionalparser

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf16"

	"github.com/antlr4-go/antlr/v4"
	"github.com/huner2/go-sddlparse/internal/util"
)

/*
	Conditional order of operations:
	1. Exists, Not Exists
	2. Member_of, Not_Member_of, Device_Member_of, Not_Device_Member_of,
		Member_of_Any, Not_Member_of_Any, Device_Member_of_Any,
		Not_Device_Member_of_Any, Contains, Not_Contains, Any_of,
		Not_Any_of, ==, !=, <, <=, >, >=
	3. !
	4. &&
	5. ||
*/

// ConditionalTokenType represents the type of a token in a conditional ACE.
type ConditionalTokenType int

const (
	CONDITIONAL_TOKEN_TYPE_PADDING ConditionalTokenType = 0x00

	// Numeric types are stored in two's complement format.
	CONDITIONAL_TOKEN_TYPE_SIGNED_BYTE     ConditionalTokenType = 0x01
	CONDITIONAL_TOKEN_TYPE_SIGNED_SHORT    ConditionalTokenType = 0x02
	CONDITIONAL_TOKEN_TYPE_SIGNED_LONG     ConditionalTokenType = 0x03
	CONDITIONAL_TOKEN_TYPE_SIGNED_LONGLONG ConditionalTokenType = 0x04

	CONDITIONAL_TOKEN_TYPE_UNICODE_STRING ConditionalTokenType = 0x10
	CONDITIONAL_TOKEN_TYPE_OCTET_STRING   ConditionalTokenType = 0x18
	CONDITIONAL_TOKEN_TYPE_COMPOSITE      ConditionalTokenType = 0x50
	CONDITIONAL_TOKEN_TYPE_SID            ConditionalTokenType = 0x51

	// Relational Operators
	CONDITIONAL_TOKEN_TYPE_MEMBER_OF                ConditionalTokenType = 0x89
	CONDITIONAL_TOKEN_TYPE_DEVICE_MEMBER_OF         ConditionalTokenType = 0x8A
	CONDITIONAL_TOKEN_TYPE_MEMBER_OF_ANY            ConditionalTokenType = 0x8B
	CONDITIONAL_TOKEN_TYPE_DEVICE_MEMBER_OF_ANY     ConditionalTokenType = 0x8C
	CONDITIONAL_TOKEN_TYPE_NOT_MEMBER_OF            ConditionalTokenType = 0x90
	CONDITIONAL_TOKEN_TYPE_NOT_DEVICE_MEMBER_OF     ConditionalTokenType = 0x91
	CONDITIONAL_TOKEN_TYPE_NOT_MEMBER_OF_ANY        ConditionalTokenType = 0x92
	CONDITIONAL_TOKEN_TYPE_NOT_DEVICE_MEMBER_OF_ANY ConditionalTokenType = 0x93

	// Binary Operators
	CONDITIONAL_TOKEN_TYPE_EQUALS                 ConditionalTokenType = 0x80
	CONDITIONAL_TOKEN_TYPE_NOT_EQUALS             ConditionalTokenType = 0x81
	CONDITIONAL_TOKEN_TYPE_LESS_THAN              ConditionalTokenType = 0x82
	CONDITIONAL_TOKEN_TYPE_LESS_THAN_OR_EQUALS    ConditionalTokenType = 0x83
	CONDITIONAL_TOKEN_TYPE_GREATER_THAN           ConditionalTokenType = 0x84
	CONDITIONAL_TOKEN_TYPE_GREATER_THAN_OR_EQUALS ConditionalTokenType = 0x85
	CONDITIONAL_TOKEN_TYPE_CONTAINS               ConditionalTokenType = 0x86
	CONDITIONAL_TOKEN_TYPE_ANY_OF                 ConditionalTokenType = 0x88
	CONDITIONAL_TOKEN_TYPE_NOT_CONTAINS           ConditionalTokenType = 0x8E
	CONDITIONAL_TOKEN_TYPE_NOT_ANY_OF             ConditionalTokenType = 0x8F

	// Logical Operators
	CONDITIONAL_TOKEN_TYPE_EXISTS      ConditionalTokenType = 0x87
	CONDITIONAL_TOKEN_TYPE_NOT_EXISTS  ConditionalTokenType = 0x8D
	CONDITIONAL_TOKEN_TYPE_LOGICAL_NOT ConditionalTokenType = 0xA2
	CONDITIONAL_TOKEN_TYPE_LOGICAL_AND ConditionalTokenType = 0xA0
	CONDITIONAL_TOKEN_TYPE_LOGICAL_OR  ConditionalTokenType = 0xA1

	// Attributes
	CONDITIONAL_TOKEN_TYPE_LOCAL_ATTRIBUTE    ConditionalTokenType = 0xF8
	CONDITIONAL_TOKEN_TYPE_USER_ATTRIBUTE     ConditionalTokenType = 0xF9
	CONDITIONAL_TOKEN_TYPE_RESOURCE_ATTRIBUTE ConditionalTokenType = 0xFA
	CONDITIONAL_TOKEN_TYPE_DEVICE_ATTRIBUTE   ConditionalTokenType = 0xFB
)

// ConditionalBase represents the base of a conditional token.
// Only applies to numeric types.
type ConditionalBase int

const (
	CONDITIONAL_BASE_OCTAL       ConditionalBase = 0x01
	CONDITIONAL_BASE_DECIMAL     ConditionalBase = 0x02
	CONDITIONAL_BASE_HEXADECIMAL ConditionalBase = 0x03
)

// ConditionalSign represents the sign of a conditional token.
// Only applies to numeric types.
type ConditionalSign int

const (
	CONDITIONAL_SIGN_POSITIVE ConditionalSign = 0x01
	CONDITIONAL_SIGN_NEGATIVE ConditionalSign = 0x02
	CONDITIONAL_SIGN_NONE     ConditionalSign = 0x03
)

type Conditional struct {
	TokenType ConditionalTokenType
	Base      ConditionalBase
	Sign      ConditionalSign
	Length    uint32
	Value     interface{}
}

type ConditionalExpression struct {
	Conditions []Conditional
}

const errInvalidApplicationData = "invalid application data"

func (ce *ConditionalExpression) ToBytes() ([]byte, error) {
	if len(ce.Conditions) == 0 {
		return nil, nil
	}

	var buf []byte

	// Add the magic header "artx"
	buf = append(buf, 0x61, 0x72, 0x74, 0x78)

	for _, cond := range ce.Conditions {
		// Write token type
		buf = append(buf, byte(cond.TokenType))

		switch cond.TokenType {
		case CONDITIONAL_TOKEN_TYPE_LOCAL_ATTRIBUTE,
			CONDITIONAL_TOKEN_TYPE_USER_ATTRIBUTE,
			CONDITIONAL_TOKEN_TYPE_RESOURCE_ATTRIBUTE,
			CONDITIONAL_TOKEN_TYPE_DEVICE_ATTRIBUTE,
			CONDITIONAL_TOKEN_TYPE_UNICODE_STRING:

			str := cond.Value.(string)
			// Convert to UTF-16LE
			utf16Data := utf16.Encode([]rune(str))
			// Add null terminator
			utf16Data = append(utf16Data, 0)

			// Convert to bytes
			var strBytes []byte
			for _, r := range utf16Data {
				strBytes = append(strBytes, byte(r), byte(r>>8))
			}

			// Write length (DWORD)
			lengthBytes := make([]byte, 4)
			binary.LittleEndian.PutUint32(lengthBytes, uint32(len(strBytes)))
			buf = append(buf, lengthBytes...)

			// Write string data
			buf = append(buf, strBytes...)

		case CONDITIONAL_TOKEN_TYPE_SIGNED_BYTE,
			CONDITIONAL_TOKEN_TYPE_SIGNED_SHORT,
			CONDITIONAL_TOKEN_TYPE_SIGNED_LONG,
			CONDITIONAL_TOKEN_TYPE_SIGNED_LONGLONG:

			val := cond.Value.(int64)

			// Store the value as-is in two's complement format
			storeVal := uint64(val)

			// Write 8-byte value (QWORD) in little endian
			valueBytes := make([]byte, 8)
			binary.LittleEndian.PutUint64(valueBytes, storeVal)
			buf = append(buf, valueBytes...)

			// Write sign byte
			buf = append(buf, byte(cond.Sign))

			// Write base byte
			buf = append(buf, byte(cond.Base))

		case CONDITIONAL_TOKEN_TYPE_OCTET_STRING:
			value := cond.Value.([]byte)

			// Write length (DWORD)
			lengthBytes := make([]byte, 4)
			binary.LittleEndian.PutUint32(lengthBytes, uint32(len(value)))
			buf = append(buf, lengthBytes...)

			// Write data
			buf = append(buf, value...)

		case CONDITIONAL_TOKEN_TYPE_COMPOSITE:
			conditions := cond.Value.([]Conditional)

			// Create a nested ConditionalExpression and serialize it
			nestedExpr := &ConditionalExpression{Conditions: conditions}
			nestedBytes, err := nestedExpr.ToBytes()
			if err != nil {
				return nil, fmt.Errorf("error serializing composite: %v", err)
			}

			// Remove the "artx" header from nested data
			if len(nestedBytes) >= 4 {
				nestedBytes = nestedBytes[4:]
			}

			// Write length (DWORD)
			lengthBytes := make([]byte, 4)
			binary.LittleEndian.PutUint32(lengthBytes, uint32(len(nestedBytes)))
			buf = append(buf, lengthBytes...)

			// Write data
			buf = append(buf, nestedBytes...)

		case CONDITIONAL_TOKEN_TYPE_SID:
			sidStr := cond.Value.(string)
			sidBytes, err := util.SidToLEBytes(sidStr)
			if err != nil {
				return nil, fmt.Errorf("error converting SID to bytes: %v", err)
			}

			// Write length (DWORD)
			lengthBytes := make([]byte, 4)
			binary.LittleEndian.PutUint32(lengthBytes, uint32(len(sidBytes)))
			buf = append(buf, lengthBytes...)

			// Write SID data
			buf = append(buf, sidBytes...)

		case CONDITIONAL_TOKEN_TYPE_EXISTS,
			CONDITIONAL_TOKEN_TYPE_NOT_EXISTS,
			CONDITIONAL_TOKEN_TYPE_LOGICAL_NOT,
			CONDITIONAL_TOKEN_TYPE_LOGICAL_AND,
			CONDITIONAL_TOKEN_TYPE_LOGICAL_OR,
			CONDITIONAL_TOKEN_TYPE_MEMBER_OF,
			CONDITIONAL_TOKEN_TYPE_DEVICE_MEMBER_OF,
			CONDITIONAL_TOKEN_TYPE_MEMBER_OF_ANY,
			CONDITIONAL_TOKEN_TYPE_DEVICE_MEMBER_OF_ANY,
			CONDITIONAL_TOKEN_TYPE_NOT_MEMBER_OF,
			CONDITIONAL_TOKEN_TYPE_NOT_DEVICE_MEMBER_OF,
			CONDITIONAL_TOKEN_TYPE_NOT_MEMBER_OF_ANY,
			CONDITIONAL_TOKEN_TYPE_NOT_DEVICE_MEMBER_OF_ANY,
			CONDITIONAL_TOKEN_TYPE_EQUALS,
			CONDITIONAL_TOKEN_TYPE_NOT_EQUALS,
			CONDITIONAL_TOKEN_TYPE_LESS_THAN,
			CONDITIONAL_TOKEN_TYPE_LESS_THAN_OR_EQUALS,
			CONDITIONAL_TOKEN_TYPE_GREATER_THAN,
			CONDITIONAL_TOKEN_TYPE_GREATER_THAN_OR_EQUALS,
			CONDITIONAL_TOKEN_TYPE_CONTAINS,
			CONDITIONAL_TOKEN_TYPE_ANY_OF,
			CONDITIONAL_TOKEN_TYPE_NOT_CONTAINS,
			CONDITIONAL_TOKEN_TYPE_NOT_ANY_OF:
			// Operators have no additional data

		default:
			return nil, fmt.Errorf("unsupported token type: 0x%x", cond.TokenType)
		}
	}

	length := len(buf)
	// Align to DWORD boundary
	if length%4 != 0 {
		padding := 4 - (length % 4)
		for i := 0; i < padding; i++ {
			buf = append(buf, 0x00)
		}
	}

	return buf, nil
}

func (c *Conditional) String() string {
	switch c.TokenType {
	case CONDITIONAL_TOKEN_TYPE_LOCAL_ATTRIBUTE:
		return c.Value.(string)
	case CONDITIONAL_TOKEN_TYPE_USER_ATTRIBUTE:
		return fmt.Sprintf("@user.%s", c.Value.(string))
	case CONDITIONAL_TOKEN_TYPE_RESOURCE_ATTRIBUTE:
		return fmt.Sprintf("@resource.%s", c.Value.(string))
	case CONDITIONAL_TOKEN_TYPE_DEVICE_ATTRIBUTE:
		return fmt.Sprintf("@device.%s", c.Value.(string))
	case CONDITIONAL_TOKEN_TYPE_UNICODE_STRING:
		return fmt.Sprintf("\"%s\"", c.Value.(string))
	case CONDITIONAL_TOKEN_TYPE_SIGNED_BYTE, CONDITIONAL_TOKEN_TYPE_SIGNED_SHORT,
		CONDITIONAL_TOKEN_TYPE_SIGNED_LONG, CONDITIONAL_TOKEN_TYPE_SIGNED_LONGLONG:
		val := c.Value.(int64)
		switch c.Base {
		case CONDITIONAL_BASE_OCTAL:
			return fmt.Sprintf("0%o", val)
		case CONDITIONAL_BASE_HEXADECIMAL:
			return fmt.Sprintf("0x%x", val)
		default:
			return fmt.Sprintf("%d", val)
		}
	case CONDITIONAL_TOKEN_TYPE_SID:
		return fmt.Sprintf("SID(%s)", c.Value.(string))
	case CONDITIONAL_TOKEN_TYPE_OCTET_STRING:
		// BLOB format: # followed by hexadecimal numbers
		return fmt.Sprintf("#%x", c.Value.([]byte))
	case CONDITIONAL_TOKEN_TYPE_COMPOSITE:
		conds := c.Value.([]Conditional)
		str := "{"
		for i, cond := range conds {
			if i > 0 {
				str += ","
			}
			str += cond.String()
		}
		str += "}"
		return str
	default:
		return "UNKNOWN"
	}
}

func (ce *ConditionalExpression) ToInfixString() (string, error) {
	if len(ce.Conditions) == 0 {
		return "", nil
	}

	var stack []string

	for _, cond := range ce.Conditions {
		switch cond.TokenType {
		case CONDITIONAL_TOKEN_TYPE_LOCAL_ATTRIBUTE,
			CONDITIONAL_TOKEN_TYPE_USER_ATTRIBUTE,
			CONDITIONAL_TOKEN_TYPE_RESOURCE_ATTRIBUTE,
			CONDITIONAL_TOKEN_TYPE_DEVICE_ATTRIBUTE,
			CONDITIONAL_TOKEN_TYPE_UNICODE_STRING,
			CONDITIONAL_TOKEN_TYPE_SIGNED_BYTE,
			CONDITIONAL_TOKEN_TYPE_SIGNED_SHORT,
			CONDITIONAL_TOKEN_TYPE_SIGNED_LONG,
			CONDITIONAL_TOKEN_TYPE_SIGNED_LONGLONG,
			CONDITIONAL_TOKEN_TYPE_OCTET_STRING,
			CONDITIONAL_TOKEN_TYPE_COMPOSITE,
			CONDITIONAL_TOKEN_TYPE_SID:
			stack = append(stack, cond.String())

		case CONDITIONAL_TOKEN_TYPE_EXISTS:
			if len(stack) < 1 {
				return "", errors.New("insufficient operands for EXISTS")
			}
			operand := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, fmt.Sprintf("Exists %s", operand))

		case CONDITIONAL_TOKEN_TYPE_NOT_EXISTS:
			if len(stack) < 1 {
				return "", errors.New("insufficient operands for NOT_EXISTS")
			}
			operand := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, fmt.Sprintf("Not_exists %s", operand))

		case CONDITIONAL_TOKEN_TYPE_LOGICAL_NOT:
			if len(stack) < 1 {
				return "", errors.New("insufficient operands for NOT")
			}
			operand := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, fmt.Sprintf("!(%s)", operand))

		case CONDITIONAL_TOKEN_TYPE_EQUALS:
			if len(stack) < 2 {
				return "", errors.New("insufficient operands for EQUALS")
			}
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, fmt.Sprintf("%s==%s", left, right))

		case CONDITIONAL_TOKEN_TYPE_NOT_EQUALS:
			if len(stack) < 2 {
				return "", errors.New("insufficient operands for NOT_EQUALS")
			}
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, fmt.Sprintf("%s!=%s", left, right))

		case CONDITIONAL_TOKEN_TYPE_LESS_THAN:
			if len(stack) < 2 {
				return "", errors.New("insufficient operands for LESS_THAN")
			}
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, fmt.Sprintf("%s<%s", left, right))

		case CONDITIONAL_TOKEN_TYPE_LESS_THAN_OR_EQUALS:
			if len(stack) < 2 {
				return "", errors.New("insufficient operands for LESS_THAN_OR_EQUALS")
			}
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, fmt.Sprintf("%s<=%s", left, right))

		case CONDITIONAL_TOKEN_TYPE_GREATER_THAN:
			if len(stack) < 2 {
				return "", errors.New("insufficient operands for GREATER_THAN")
			}
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, fmt.Sprintf("%s>%s", left, right))

		case CONDITIONAL_TOKEN_TYPE_GREATER_THAN_OR_EQUALS:
			if len(stack) < 2 {
				return "", errors.New("insufficient operands for GREATER_THAN_OR_EQUALS")
			}
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, fmt.Sprintf("%s>=%s", left, right))

		case CONDITIONAL_TOKEN_TYPE_CONTAINS:
			if len(stack) < 2 {
				return "", errors.New("insufficient operands for CONTAINS")
			}
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, fmt.Sprintf("%s Contains %s", left, right))

		case CONDITIONAL_TOKEN_TYPE_NOT_CONTAINS:
			if len(stack) < 2 {
				return "", errors.New("insufficient operands for NOT_CONTAINS")
			}
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, fmt.Sprintf("%s Not_Contains %s", left, right))

		case CONDITIONAL_TOKEN_TYPE_ANY_OF:
			if len(stack) < 2 {
				return "", errors.New("insufficient operands for ANY_OF")
			}
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, fmt.Sprintf("%s Any_of %s", left, right))

		case CONDITIONAL_TOKEN_TYPE_NOT_ANY_OF:
			if len(stack) < 2 {
				return "", errors.New("insufficient operands for NOT_ANY_OF")
			}
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, fmt.Sprintf("%s Not_Any_of %s", left, right))

		case CONDITIONAL_TOKEN_TYPE_MEMBER_OF:
			if len(stack) < 2 {
				return "", errors.New("insufficient operands for MEMBER_OF")
			}
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, fmt.Sprintf("Member_of %s", right))

		case CONDITIONAL_TOKEN_TYPE_NOT_MEMBER_OF:
			if len(stack) < 2 {
				return "", errors.New("insufficient operands for NOT_MEMBER_OF")
			}
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, fmt.Sprintf("Not_Member_of %s", right))

		case CONDITIONAL_TOKEN_TYPE_DEVICE_MEMBER_OF:
			if len(stack) < 2 {
				return "", errors.New("insufficient operands for DEVICE_MEMBER_OF")
			}
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, fmt.Sprintf("Device_Member_of %s", right))

		case CONDITIONAL_TOKEN_TYPE_NOT_DEVICE_MEMBER_OF:
			if len(stack) < 2 {
				return "", errors.New("insufficient operands for NOT_DEVICE_MEMBER_OF")
			}
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, fmt.Sprintf("Not_Device_Member_of %s", right))

		case CONDITIONAL_TOKEN_TYPE_MEMBER_OF_ANY:
			if len(stack) < 2 {
				return "", errors.New("insufficient operands for MEMBER_OF_ANY")
			}
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, fmt.Sprintf("Member_of_Any %s", right))

		case CONDITIONAL_TOKEN_TYPE_NOT_MEMBER_OF_ANY:
			if len(stack) < 2 {
				return "", errors.New("insufficient operands for NOT_MEMBER_OF_ANY")
			}
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, fmt.Sprintf("Not_Member_of_Any %s", right))

		case CONDITIONAL_TOKEN_TYPE_DEVICE_MEMBER_OF_ANY:
			if len(stack) < 2 {
				return "", errors.New("insufficient operands for DEVICE_MEMBER_OF_ANY")
			}
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, fmt.Sprintf("Device_Member_of_Any %s", right))

		case CONDITIONAL_TOKEN_TYPE_NOT_DEVICE_MEMBER_OF_ANY:
			if len(stack) < 2 {
				return "", errors.New("insufficient operands for NOT_DEVICE_MEMBER_OF_ANY")
			}
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, fmt.Sprintf("Not_Device_Member_of_Any %s", right))

		case CONDITIONAL_TOKEN_TYPE_LOGICAL_AND:
			if len(stack) < 2 {
				return "", errors.New("insufficient operands for AND")
			}
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			stack = append(stack, fmt.Sprintf("(%s)&&(%s)", left, right))

		case CONDITIONAL_TOKEN_TYPE_LOGICAL_OR:
			if len(stack) < 2 {
				return "", errors.New("insufficient operands for OR")
			}
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, fmt.Sprintf("(%s)||(%s)", left, right))

		default:
			return "", fmt.Errorf("unknown token type: 0x%x", cond.TokenType)
		}
	}

	if len(stack) != 1 {
		return "", errors.New("invalid expression: stack should contain exactly one element")
	}

	return stack[0], nil
}

// ParseApplicationData parses the binary conditional data and returns a ConditionalExpression
func ParseApplicationData(data []byte) (*ConditionalExpression, error) {
	if len(data) == 0 {
		return &ConditionalExpression{}, nil
	}
	if len(data) < 4 {
		return nil, errors.New(errInvalidApplicationData)
	}
	if data[0] != 0x61 || data[1] != 0x72 || data[2] != 0x74 || data[3] != 0x78 {
		return nil, errors.New(errInvalidApplicationData)
	}

	var conditions []Conditional
	i := 4

	for i < len(data) {
		// If current byte is padding, check that the rest are also padding
		if data[i] == 0x00 {
			for j := i; j < len(data); j++ {
				if data[j] != 0x00 {
					return nil, errors.New(errInvalidApplicationData)
				}
			}
			break
		}

		tokenType := ConditionalTokenType(data[i])
		if len(data)-i < 1 {
			return nil, errors.New(errInvalidApplicationData)
		}
		i++

		cond := Conditional{
			TokenType: tokenType,
		}

		switch tokenType {
		case CONDITIONAL_TOKEN_TYPE_LOCAL_ATTRIBUTE,
			CONDITIONAL_TOKEN_TYPE_USER_ATTRIBUTE,
			CONDITIONAL_TOKEN_TYPE_RESOURCE_ATTRIBUTE,
			CONDITIONAL_TOKEN_TYPE_DEVICE_ATTRIBUTE,
			CONDITIONAL_TOKEN_TYPE_UNICODE_STRING:
			// Next 4 bytes are a length
			if len(data)-i < 4 {
				return nil, errors.New(errInvalidApplicationData)
			}
			length := binary.LittleEndian.Uint32(data[i : i+4])
			i += 4
			if len(data)-i < int(length) {
				return nil, errors.New(errInvalidApplicationData)
			}
			value := data[i : i+int(length)]
			i += int(length)

			// Convert UTF-16LE to string for Unicode strings and attributes
			if tokenType == CONDITIONAL_TOKEN_TYPE_UNICODE_STRING ||
				tokenType == CONDITIONAL_TOKEN_TYPE_LOCAL_ATTRIBUTE ||
				tokenType == CONDITIONAL_TOKEN_TYPE_USER_ATTRIBUTE ||
				tokenType == CONDITIONAL_TOKEN_TYPE_RESOURCE_ATTRIBUTE ||
				tokenType == CONDITIONAL_TOKEN_TYPE_DEVICE_ATTRIBUTE {

				if len(value)%2 != 0 {
					return nil, errors.New("invalid UTF-16 string length")
				}

				utf16Data := make([]uint16, len(value)/2)
				for j := 0; j < len(value); j += 2 {
					utf16Data[j/2] = binary.LittleEndian.Uint16(value[j : j+2])
				}

				// Remove null terminator if present
				if len(utf16Data) > 0 && utf16Data[len(utf16Data)-1] == 0 {
					utf16Data = utf16Data[:len(utf16Data)-1]
				}

				cond.Value = string(utf16.Decode(utf16Data))
			} else {
				cond.Value = value
			}
			cond.Length = length

		case CONDITIONAL_TOKEN_TYPE_SIGNED_BYTE,
			CONDITIONAL_TOKEN_TYPE_SIGNED_SHORT,
			CONDITIONAL_TOKEN_TYPE_SIGNED_LONG,
			CONDITIONAL_TOKEN_TYPE_SIGNED_LONGLONG:
			// 1 QWORD, least significant byte first, 2's complement.
			// Next byte is sign
			// Final byte is base
			if len(data)-i < 10 {
				return nil, errors.New(errInvalidApplicationData)
			}
			value := data[i : i+8]
			i += 8
			sign := ConditionalSign(data[i])
			i++
			base := ConditionalBase(data[i])
			i++

			// Convert bytes to int64
			val := int64(binary.LittleEndian.Uint64(value))

			// Apply sign
			if sign == CONDITIONAL_SIGN_NEGATIVE {
				val = -val
			}

			cond.Value = val
			cond.Sign = sign
			cond.Base = base

		case CONDITIONAL_TOKEN_TYPE_OCTET_STRING:
			// DWORD length followed by that many bytes
			if len(data)-i < 4 {
				return nil, errors.New(errInvalidApplicationData)
			}
			length := binary.LittleEndian.Uint32(data[i : i+4])
			i += 4
			if len(data)-i < int(length) {
				return nil, errors.New(errInvalidApplicationData)
			}
			value := data[i : i+int(length)]
			i += int(length)
			cond.Value = value
			cond.Length = length

		case CONDITIONAL_TOKEN_TYPE_COMPOSITE:
			// DWORD length followed by that many bytes
			if len(data)-i < 4 {
				return nil, errors.New(errInvalidApplicationData)
			}
			length := binary.LittleEndian.Uint32(data[i : i+4])
			i += 4
			if len(data)-i < int(length) {
				return nil, errors.New(errInvalidApplicationData)
			}
			value := data[i : i+int(length)]
			i += int(length)
			cond.Length = length
			nestedExpr, err := ParseApplicationData(append([]byte{0x61, 0x72, 0x74, 0x78}, value...))
			if err != nil {
				return nil, err
			}
			cond.Value = nestedExpr.Conditions

		case CONDITIONAL_TOKEN_TYPE_SID:
			// DWORD length followed by that many bytes
			if len(data)-i < 4 {
				return nil, errors.New(errInvalidApplicationData)
			}
			length := binary.LittleEndian.Uint32(data[i : i+4])
			i += 4
			if len(data)-i < int(length) {
				return nil, errors.New(errInvalidApplicationData)
			}
			value := data[i : i+int(length)]
			i += int(length)
			var err error
			cond.Value, _, err = util.SidFromLEBytes(value)
			if err != nil {
				return nil, err
			}
			cond.Length = length

		case CONDITIONAL_TOKEN_TYPE_EXISTS,
			CONDITIONAL_TOKEN_TYPE_NOT_EXISTS,
			CONDITIONAL_TOKEN_TYPE_LOGICAL_NOT,
			CONDITIONAL_TOKEN_TYPE_LOGICAL_AND,
			CONDITIONAL_TOKEN_TYPE_LOGICAL_OR,
			CONDITIONAL_TOKEN_TYPE_MEMBER_OF,
			CONDITIONAL_TOKEN_TYPE_DEVICE_MEMBER_OF,
			CONDITIONAL_TOKEN_TYPE_MEMBER_OF_ANY,
			CONDITIONAL_TOKEN_TYPE_DEVICE_MEMBER_OF_ANY,
			CONDITIONAL_TOKEN_TYPE_NOT_MEMBER_OF,
			CONDITIONAL_TOKEN_TYPE_NOT_DEVICE_MEMBER_OF,
			CONDITIONAL_TOKEN_TYPE_NOT_MEMBER_OF_ANY,
			CONDITIONAL_TOKEN_TYPE_NOT_DEVICE_MEMBER_OF_ANY,
			CONDITIONAL_TOKEN_TYPE_EQUALS,
			CONDITIONAL_TOKEN_TYPE_NOT_EQUALS,
			CONDITIONAL_TOKEN_TYPE_LESS_THAN,
			CONDITIONAL_TOKEN_TYPE_LESS_THAN_OR_EQUALS,
			CONDITIONAL_TOKEN_TYPE_GREATER_THAN,
			CONDITIONAL_TOKEN_TYPE_GREATER_THAN_OR_EQUALS,
			CONDITIONAL_TOKEN_TYPE_CONTAINS,
			CONDITIONAL_TOKEN_TYPE_ANY_OF,
			CONDITIONAL_TOKEN_TYPE_NOT_CONTAINS,
			CONDITIONAL_TOKEN_TYPE_NOT_ANY_OF:
			// No additional data for operators

		default:
			return nil, errors.New(errInvalidApplicationData)
		}

		conditions = append(conditions, cond)
	}

	return &ConditionalExpression{Conditions: conditions}, nil
}

// ConditionalExpressionBuilder is a listener that builds our ConditionalExpression from the parse tree
type ConditionalExpressionBuilder struct {
	*BaseconditionalexpressionListener
	conditions []Conditional
	stack      []interface{} // Stack to build the expression in postfix order
	hasError   bool
	errorMsg   string
}

// NewConditionalExpressionBuilder creates a new builder
func NewConditionalExpressionBuilder() *ConditionalExpressionBuilder {
	return &ConditionalExpressionBuilder{
		BaseconditionalexpressionListener: &BaseconditionalexpressionListener{},
		conditions:                        make([]Conditional, 0),
		stack:                             make([]interface{}, 0),
		hasError:                          false,
		errorMsg:                          "",
	}
}

// GetConditionalExpression returns the built ConditionalExpression
func (b *ConditionalExpressionBuilder) GetConditionalExpression() (*ConditionalExpression, error) {
	if b.hasError {
		return nil, fmt.Errorf("parse error: %s", b.errorMsg)
	}
	return &ConditionalExpression{Conditions: b.conditions}, nil
}

// Helper method to add a conditional to our list
func (b *ConditionalExpressionBuilder) addConditional(cond Conditional) {
	b.conditions = append(b.conditions, cond)
}

// ExitSimpleAttrName - Handle simple attribute names (local attributes)
func (b *ConditionalExpressionBuilder) ExitSimpleAttrName(ctx *SimpleAttrNameContext) {
	attrName := ctx.GetText()
	cond := Conditional{
		TokenType: CONDITIONAL_TOKEN_TYPE_LOCAL_ATTRIBUTE,
		Value:     attrName,
	}
	b.addConditional(cond)
}

// ExitPrefixedAttrName - Handle prefixed attribute names (@user., @device., @resource.)
func (b *ConditionalExpressionBuilder) ExitPrefixedAttrName(ctx *PrefixedAttrNameContext) {
	attrText := ctx.GetText()

	var tokenType ConditionalTokenType
	var attrName string

	if strings.HasPrefix(attrText, "@user.") {
		tokenType = CONDITIONAL_TOKEN_TYPE_USER_ATTRIBUTE
		attrName = strings.TrimPrefix(attrText, "@user.")
	} else if strings.HasPrefix(attrText, "@device.") {
		tokenType = CONDITIONAL_TOKEN_TYPE_DEVICE_ATTRIBUTE
		attrName = strings.TrimPrefix(attrText, "@device.")
	} else if strings.HasPrefix(attrText, "@resource.") {
		tokenType = CONDITIONAL_TOKEN_TYPE_RESOURCE_ATTRIBUTE
		attrName = strings.TrimPrefix(attrText, "@resource.")
	} else {
		b.hasError = true
		b.errorMsg = fmt.Sprintf("invalid attribute prefix: %s", attrText)
		return
	}

	cond := Conditional{
		TokenType: tokenType,
		Value:     attrName,
	}
	b.addConditional(cond)
}

// ExitIntegerValue - Handle integer values
func (b *ConditionalExpressionBuilder) ExitIntegerValue(ctx *IntegerValueContext) {
	intText := ctx.GetText()

	var value int64
	var base ConditionalBase
	var sign ConditionalSign = CONDITIONAL_SIGN_POSITIVE
	var err error

	if strings.HasPrefix(intText, "+") {
		sign = CONDITIONAL_SIGN_POSITIVE
		intText = intText[1:]
	} else if strings.HasPrefix(intText, "-") {
		sign = CONDITIONAL_SIGN_NEGATIVE
		intText = intText[1:]
	}

	if ctx.HEX_INTEGER() != nil {
		base = CONDITIONAL_BASE_HEXADECIMAL
		hexText := ctx.HEX_INTEGER().GetText()
		value, err = strconv.ParseInt(hexText[2:], 16, 64) // Remove "0x"
	} else if ctx.OCTAL_INTEGER() != nil {
		base = CONDITIONAL_BASE_OCTAL
		octalText := ctx.OCTAL_INTEGER().GetText()
		value, err = strconv.ParseInt(octalText[1:], 8, 64) // Remove leading "0"
	} else if ctx.DECIMAL_INTEGER() != nil {
		base = CONDITIONAL_BASE_DECIMAL
		value, err = strconv.ParseInt(ctx.DECIMAL_INTEGER().GetText(), 10, 64)
	} else {
		b.hasError = true
		b.errorMsg = fmt.Sprintf("unknown integer format: %s", intText)
		return
	}

	if err != nil {
		b.hasError = true
		b.errorMsg = fmt.Sprintf("invalid integer: %s", ctx.GetText())
		return
	}

	// Apply sign
	if sign == CONDITIONAL_SIGN_NEGATIVE {
		value = -value
	}

	// Determine the appropriate token type based on value range
	var tokenType ConditionalTokenType
	if value >= -128 && value <= 127 {
		tokenType = CONDITIONAL_TOKEN_TYPE_SIGNED_BYTE
	} else if value >= -32768 && value <= 32767 {
		tokenType = CONDITIONAL_TOKEN_TYPE_SIGNED_SHORT
	} else if value >= -2147483648 && value <= 2147483647 {
		tokenType = CONDITIONAL_TOKEN_TYPE_SIGNED_LONG
	} else {
		tokenType = CONDITIONAL_TOKEN_TYPE_SIGNED_LONGLONG
	}

	cond := Conditional{
		TokenType: tokenType,
		Base:      base,
		Sign:      sign,
		Value:     value,
	}
	b.addConditional(cond)
}

// ExitValue - Handle string literals from STRING_LITERAL token
func (b *ConditionalExpressionBuilder) ExitValue(ctx *ValueContext) {
	if ctx.STRING_LITERAL() != nil {
		text := ctx.STRING_LITERAL().GetText()
		// Remove surrounding quotes
		if len(text) >= 2 && text[0] == '"' && text[len(text)-1] == '"' {
			text = text[1 : len(text)-1]
		}

		cond := Conditional{
			TokenType: CONDITIONAL_TOKEN_TYPE_UNICODE_STRING,
			Value:     text,
		}
		b.addConditional(cond)
	} else if ctx.OCTET_STRING() != nil {
		text := ctx.OCTET_STRING().GetText()
		// Remove the leading '#'
		text = strings.TrimPrefix(text, "#")

		// Decode hex string
		data, err := hex.DecodeString(text)
		if err != nil {
			b.hasError = true
			b.errorMsg = fmt.Sprintf("invalid octet string: %s", ctx.GetText())
			return
		}

		cond := Conditional{
			TokenType: CONDITIONAL_TOKEN_TYPE_OCTET_STRING,
			Value:     data,
		}
		b.addConditional(cond)
	}
}

// ExitLiteralSID - Handle SID literals
func (b *ConditionalExpressionBuilder) ExitLiteralSID(ctx *LiteralSIDContext) {
	text := ctx.GetText()
	// Remove SID() wrapper
	if strings.HasPrefix(text, "SID(") && strings.HasSuffix(text, ")") {
		text = text[4 : len(text)-1]
	}

	cond := Conditional{
		TokenType: CONDITIONAL_TOKEN_TYPE_SID,
		Value:     text,
	}
	b.addConditional(cond)
}

// ExitSidString - Handle SID string content (either SID_TOKEN or SID_FORMAT)
func (b *ConditionalExpressionBuilder) ExitSidString(ctx *SidStringContext) {
	text := ctx.GetText()

	cond := Conditional{
		TokenType: CONDITIONAL_TOKEN_TYPE_SID,
		Value:     text,
	}
	b.addConditional(cond)
}

// ExitValueArray - Handle value arrays (composite values)
// Value arrays are only composite if they contain more than one value I think.
func (b *ConditionalExpressionBuilder) ExitValueArray(ctx *ValueArrayContext) {
	// For arrays, we need to collect all the values that were just processed
	// and create a composite conditional
	values := ctx.AllValue()
	if len(values) > 1 {
		// This is an array, wrap the last N conditions in a composite
		numValues := len(values)
		if len(b.conditions) >= numValues {
			// Extract the last N conditions
			arrayConditions := b.conditions[len(b.conditions)-numValues:]
			// Remove them from the main list
			b.conditions = b.conditions[:len(b.conditions)-numValues]

			// Prevents issues with circular references
			copied := make([]Conditional, len(arrayConditions))
			copy(copied, arrayConditions)

			// Create composite conditional
			cond := Conditional{
				TokenType: CONDITIONAL_TOKEN_TYPE_COMPOSITE,
				Value:     copied,
			}
			b.addConditional(cond)
		}
	}
	// If it's a single value, it's already been handled by the value's exit method
}

// ExitSidArray - Handle SID arrays
// Sid arrays always seem to be composite, even if there's only one SID
func (b *ConditionalExpressionBuilder) ExitSidArray(ctx *SidArrayContext) {
	sids := ctx.AllLiteralSID()
	if len(sids) > 0 {
		// This is an array, wrap the last N SID conditions in a composite
		numSids := len(sids)
		if len(b.conditions) >= numSids {
			// Extract the last N conditions
			sidConditions := b.conditions[len(b.conditions)-numSids:]
			// Remove them from the main list
			b.conditions = b.conditions[:len(b.conditions)-numSids]

			// Prevents issues with circular references
			copied := make([]Conditional, len(sidConditions))
			copy(copied, sidConditions)

			// Create composite conditional
			cond := Conditional{
				TokenType: CONDITIONAL_TOKEN_TYPE_COMPOSITE,
				Value:     copied,
			}
			b.addConditional(cond)
		}
	}
}

// Operator handlers - these add the appropriate operator tokens

// ExitExistsOp - Handle Exists and Not_exists operations
func (b *ConditionalExpressionBuilder) ExitExistsOp(ctx *ExistsOpContext) {
	var tokenType ConditionalTokenType

	if ctx.EXISTS() != nil {
		tokenType = CONDITIONAL_TOKEN_TYPE_EXISTS
	} else if ctx.NOT_EXISTS() != nil {
		tokenType = CONDITIONAL_TOKEN_TYPE_NOT_EXISTS
	} else {
		b.hasError = true
		b.errorMsg = fmt.Sprintf("unknown exists operation: %s", ctx.GetText())
		return
	}

	cond := Conditional{
		TokenType: tokenType,
	}
	b.addConditional(cond)
}

// ExitMemberofOp - Handle membership operations
func (b *ConditionalExpressionBuilder) ExitMemberofOp(ctx *MemberofOpContext) {
	var tokenType ConditionalTokenType

	if ctx.MEMBER_OF() != nil {
		tokenType = CONDITIONAL_TOKEN_TYPE_MEMBER_OF
	} else if ctx.NOT_MEMBER_OF() != nil {
		tokenType = CONDITIONAL_TOKEN_TYPE_NOT_MEMBER_OF
	} else if ctx.MEMBER_OF_ANY() != nil {
		tokenType = CONDITIONAL_TOKEN_TYPE_MEMBER_OF_ANY
	} else if ctx.NOT_MEMBER_OF_ANY() != nil {
		tokenType = CONDITIONAL_TOKEN_TYPE_NOT_MEMBER_OF_ANY
	} else if ctx.DEVICE_MEMBER_OF() != nil {
		tokenType = CONDITIONAL_TOKEN_TYPE_DEVICE_MEMBER_OF
	} else if ctx.DEVICE_MEMBER_OF_ANY() != nil {
		tokenType = CONDITIONAL_TOKEN_TYPE_DEVICE_MEMBER_OF_ANY
	} else if ctx.NOT_DEVICE_MEMBER_OF() != nil {
		tokenType = CONDITIONAL_TOKEN_TYPE_NOT_DEVICE_MEMBER_OF
	} else if ctx.NOT_DEVICE_MEMBER_OF_ANY() != nil {
		tokenType = CONDITIONAL_TOKEN_TYPE_NOT_DEVICE_MEMBER_OF_ANY
	} else {
		b.hasError = true
		b.errorMsg = fmt.Sprintf("unknown membership operation: %s", ctx.GetText())
		return
	}

	cond := Conditional{
		TokenType: tokenType,
	}
	b.addConditional(cond)
}

// ExitRelOp - Handle relational operations (<, <=, >, >=)
func (b *ConditionalExpressionBuilder) ExitRelOp(ctx *RelOpContext) {
	var tokenType ConditionalTokenType

	if ctx.LT() != nil {
		tokenType = CONDITIONAL_TOKEN_TYPE_LESS_THAN
	} else if ctx.LTE() != nil {
		tokenType = CONDITIONAL_TOKEN_TYPE_LESS_THAN_OR_EQUALS
	} else if ctx.GT() != nil {
		tokenType = CONDITIONAL_TOKEN_TYPE_GREATER_THAN
	} else if ctx.GTE() != nil {
		tokenType = CONDITIONAL_TOKEN_TYPE_GREATER_THAN_OR_EQUALS
	} else {
		b.hasError = true
		b.errorMsg = fmt.Sprintf("unknown relational operation: %s", ctx.GetText())
		return
	}

	cond := Conditional{
		TokenType: tokenType,
	}
	b.addConditional(cond)
}

// ExitRelOp2 - Handle equality operations (==, !=)
func (b *ConditionalExpressionBuilder) ExitRelOp2(ctx *RelOp2Context) {
	var tokenType ConditionalTokenType

	if ctx.EQ() != nil {
		tokenType = CONDITIONAL_TOKEN_TYPE_EQUALS
	} else if ctx.NEQ() != nil {
		tokenType = CONDITIONAL_TOKEN_TYPE_NOT_EQUALS
	} else {
		b.hasError = true
		b.errorMsg = fmt.Sprintf("unknown equality operation: %s", ctx.GetText())
		return
	}

	cond := Conditional{
		TokenType: tokenType,
	}
	b.addConditional(cond)
}

// ExitContainsOp - Handle Contains and Not_Contains operations
func (b *ConditionalExpressionBuilder) ExitContainsOp(ctx *ContainsOpContext) {
	var tokenType ConditionalTokenType

	if ctx.CONTAINS() != nil {
		tokenType = CONDITIONAL_TOKEN_TYPE_CONTAINS
	} else if ctx.NOT_CONTAINS() != nil {
		tokenType = CONDITIONAL_TOKEN_TYPE_NOT_CONTAINS
	} else {
		b.hasError = true
		b.errorMsg = fmt.Sprintf("unknown contains operation: %s", ctx.GetText())
		return
	}

	cond := Conditional{
		TokenType: tokenType,
	}
	b.addConditional(cond)
}

// ExitAnyofOp - Handle Any_of and Not_Any_of operations
func (b *ConditionalExpressionBuilder) ExitAnyofOp(ctx *AnyofOpContext) {
	var tokenType ConditionalTokenType

	if ctx.ANY_OF() != nil {
		tokenType = CONDITIONAL_TOKEN_TYPE_ANY_OF
	} else if ctx.NOT_ANY_OF() != nil {
		tokenType = CONDITIONAL_TOKEN_TYPE_NOT_ANY_OF
	} else {
		b.hasError = true
		b.errorMsg = fmt.Sprintf("unknown any_of operation: %s", ctx.GetText())
		return
	}

	cond := Conditional{
		TokenType: tokenType,
	}
	b.addConditional(cond)
}

// ExitFactor - Handle logical NOT
func (b *ConditionalExpressionBuilder) ExitFactor(ctx *FactorContext) {
	// Check if this is a logical NOT operation
	if ctx.NOT() != nil {
		cond := Conditional{
			TokenType: CONDITIONAL_TOKEN_TYPE_LOGICAL_NOT,
		}
		b.addConditional(cond)
	}
	// Parentheses don't need special handling in postfix notation
}

// ExitSuperTerm - Handle logical AND operations
func (b *ConditionalExpressionBuilder) ExitSuperTerm(ctx *SuperTermContext) {
	factors := ctx.AllFactor()
	// Add AND operators for each additional factor
	for i := 1; i < len(factors); i++ {
		cond := Conditional{
			TokenType: CONDITIONAL_TOKEN_TYPE_LOGICAL_AND,
		}
		b.addConditional(cond)
	}
}

// ExitExpr - Handle logical OR operations
func (b *ConditionalExpressionBuilder) ExitExpr(ctx *ExprContext) {
	superTerms := ctx.AllSuperTerm()
	// Add OR operators for each additional super term
	for i := 1; i < len(superTerms); i++ {
		cond := Conditional{
			TokenType: CONDITIONAL_TOKEN_TYPE_LOGICAL_OR,
		}
		b.addConditional(cond)
	}
}

// CustomErrorListener implements ANTLR's ErrorListener interface
type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	hasError bool
	errorMsg string
}

func NewCustomErrorListener() *CustomErrorListener {
	return &CustomErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
		hasError:             false,
		errorMsg:             "",
	}
}

func (e *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, ex antlr.RecognitionException) {
	e.hasError = true
	e.errorMsg = fmt.Sprintf("line %d:%d - %s", line, column, msg)
}

// ParseConditionalExpression parses a conditional expression string into a ConditionalExpression
func ParseConditionalExpression(input string) (*ConditionalExpression, error) {
	// Create input stream
	inputStream := antlr.NewInputStream(input)

	// Create lexer
	lexer := NewconditionalexpressionLexer(inputStream)

	// Create error listener
	errorListener := NewCustomErrorListener()
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)

	// Create token stream
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)

	// Create parser
	parser := NewconditionalexpressionParser(tokenStream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(errorListener)

	// Parse starting from the root rule
	tree := parser.CondExpr()

	// Check for lexer/parser errors
	if errorListener.hasError {
		return nil, fmt.Errorf("syntax error: %s", errorListener.errorMsg)
	}

	// Create and use our custom listener
	builder := NewConditionalExpressionBuilder()
	antlr.ParseTreeWalkerDefault.Walk(builder, tree)

	// Return the built expression
	return builder.GetConditionalExpression()
}
