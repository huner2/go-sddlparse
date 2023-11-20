package sddlparse_test

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/huner2/go-sddlparse"
)

func TestB64SDDLLarge(t *testing.T) {
	b64, err := os.ReadFile("./testdata/large_test_data.b64")
	if err != nil {
		t.Fatal(err)
	}
	sddl, err := sddlparse.SDDLFromBase64Encoded(b64)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(sddl)
	if err != nil {
		t.Fatal(err)
	}
	compareJs, err := os.ReadFile("./testdata/large_test.json")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(js, compareJs) {
		t.Fatal("json output is not equal")
	}
}

func TestB64SDDLLargeStr(t *testing.T) {
	b64, err := os.ReadFile("./testdata/large_test_data.b64")
	if err != nil {
		t.Fatal(err)
	}
	sddl, err := sddlparse.SDDLFromBase64Encoded(b64)
	if err != nil {
		t.Fatal(err)
	}
	str, err := sddl.MustString()
	if err != nil {
		t.Fatal(err)
	}
	compareStr, err := os.ReadFile("./testdata/large_test.sddl")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal([]byte(str), compareStr) {
		t.Fatal("string output is not equal")
	}
}

func TestB64SDDLSmall(t *testing.T) {
	b64, err := os.ReadFile("./testdata/small_test_data.b64")
	if err != nil {
		t.Fatal(err)
	}
	sddl, err := sddlparse.SDDLFromBase64Encoded(b64)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(sddl)
	if err != nil {
		t.Fatal(err)
	}
	compareJs, err := os.ReadFile("./testdata/small_test.json")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(js, compareJs) {
		t.Fatal("json output is not equal")
	}
}

func TestB64SDDLSmallStr(t *testing.T) {
	b64, err := os.ReadFile("./testdata/small_test_data.b64")
	if err != nil {
		t.Fatal(err)
	}
	sddl, err := sddlparse.SDDLFromBase64Encoded(b64)
	if err != nil {
		t.Fatal(err)
	}
	str, err := sddl.MustString()
	if err != nil {
		t.Fatal(err)
	}
	compareStr, err := os.ReadFile("./testdata/small_test.sddl")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal([]byte(str), compareStr) {
		t.Fatal("string output is not equal")
	}
}

func TestParseGUID(t *testing.T) {
	guidStrRandom := "f8a0b131-5f68-486c-8040-7e8fc3c85e4f"

	guid, err := sddlparse.GuidFromString(guidStrRandom)
	if err != nil {
		t.Fatal(err)
	}
	if guid.String() != guidStrRandom {
		t.Fatal("guid string is not equal")
	}
}

func TestParseSDDLSmallStr(t *testing.T) {
	sddl, err := os.ReadFile("./testdata/small_test.sddl")
	if err != nil {
		t.Fatal(err)
	}
	sddlStr := string(sddl)
	sddlParsed, err := sddlparse.SDDLFromString(sddlStr, "0", "0")
	if err != nil {
		t.Fatal(err)
	}
	sddlParsedStr, err := sddlParsed.MustString()
	if err != nil {
		t.Fatal(err)
	}
	if sddlParsedStr != sddlStr {
		t.Fatal("sddl string is not equal")
	}
}

func TestParseSDDLLargeStr(t *testing.T) {
	sddl, err := os.ReadFile("./testdata/large_test.sddl")
	if err != nil {
		t.Fatal(err)
	}
	sddlStr := string(sddl)
	sddlParsed, err := sddlparse.SDDLFromString(sddlStr, "0", "0")
	if err != nil {
		t.Fatal(err)
	}
	sddlParsedStr, err := sddlParsed.MustString()
	if err != nil {
		t.Fatal(err)
	}
	if sddlParsedStr != sddlStr {
		t.Fatal("sddl string is not equal")
	}
}

func TestBinaryOutputSmall(t *testing.T) {
	b64, err := os.ReadFile("./testdata/small_test_data.b64")
	if err != nil {
		t.Fatal(err)
	}
	sddl, err := sddlparse.SDDLFromBase64Encoded(b64)
	if err != nil {
		t.Fatal(err)
	}
	bin, err := sddl.ToBase64Encoded()
	if err != nil {
		t.Fatal(err)
	}
	sddl, err = sddlparse.SDDLFromBase64Encoded(bin)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(sddl)
	if err != nil {
		t.Fatal(err)
	}
	compareJs, err := os.ReadFile("./testdata/small_test.json")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(js, compareJs) {
		t.Fatal("json output is not equal")
	}
}

func TestBinaryOutputLarge(t *testing.T) {
	b64, err := os.ReadFile("./testdata/large_test_data.b64")
	if err != nil {
		t.Fatal(err)
	}
	sddl, err := sddlparse.SDDLFromBase64Encoded(b64)
	if err != nil {
		t.Fatal(err)
	}
	bin, err := sddl.ToBase64Encoded()
	if err != nil {
		t.Fatal(err)
	}
	sddl, err = sddlparse.SDDLFromBase64Encoded(bin)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(sddl)
	if err != nil {
		t.Fatal(err)
	}
	compareJs, err := os.ReadFile("./testdata/large_test.json")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(js, compareJs) {
		t.Fatal("json output is not equal")
	}
}

func TestModifySmallStr(t *testing.T) {
	sddl, err := os.ReadFile("./testdata/small_test.sddl")
	if err != nil {
		t.Fatal(err)
	}
	sddlStr := string(sddl)
	sddlParsed, err := sddlparse.SDDLFromString(sddlStr, "0", "0")
	if err != nil {
		t.Fatal(err)
	}
	sddlParsed.SACL = append(sddlParsed.SACL, &sddlparse.ACE{
		Type:       sddlparse.ACETYPE_ACCESS_ALLOWED,
		AccessMask: sddlparse.ACCESS_MASK_GENERIC_ALL,
		Flags:      sddlparse.ACEFLAG_OBJECT_INHERIT,
		SID:        "S-1-2-3-4",
	})
	sddlParsedStr, err := sddlParsed.MustString()
	if err != nil {
		t.Fatal(err)
	}
	if sddlParsedStr != sddlStr+"(A;OI;GA;;;S-1-2-3-4)" {
		t.Fatal("sddl string isn't equal")
	}
}
