# Basic SDDL parsing functionality for AD SDDLs in Go

This library provides functions to parse and modify the [SDDL](https://msdn.microsoft.com/en-us/library/cc230366.aspx) format, centered around Active Directory Access Control Lists (ACLs).

## Features

* Parse SDDL strings into a structured format
* Modify the structured format
* Convert the structured format back to SDDL strings
* Parse SDDL binary data into a structured format
* Convert the structured format back to SDDL binary data

## Usage

Given a base64 output from ldapsearch, or a binary format from
a library such as [Go LDAP](github.com/go-ldap/ldap), you can
parse the binary data into a structured format:

```go
package main

import (
    parser "github.com/huner2/go-sddlparse"
)

func main() {
    b64 = "base64 encoded SDDL"
    sddl, err := parser.SDDLFromBase64Encoded(b64)
    if err != nil {
        panic(err)
    }
    // Do something with the SDDL, such as add a new ACE
    sddl.DACL = append(sddl.DACL, &parser.ACE{
        Type: parser.ACETYPE_ACCESS_ALLOWED,
        Flags: 0,
        Mask: parser.ACEMASK_GENERIC_ALL,
        SID: "S-1-5-32-544",
    })

    // Convert the SDDL to binary
    bin, err = sddl.ToBinary()
    if err != nil {
        panic(err)
    }
}
```

The same can be done with a SDDL string, using `SDDLFromString`.

An example using the `go-ldap` library:

```go
package main

import (
    "github.com/go-ldap/ldap/v3"
    parser "github.com/huner2/go-sddlparse"
    "log"
)

func main() {
    conn, err := ldap.DialURL("ldap://10.137.137.2")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    err = conn.Bind("testuser@test.test", "testpassword")
    if err != nil {
        log.Fatal(err)
    }
    searchRequest := ldap.NewSearchRequest(
        "DC=test,DC=test",
        ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
        "(objectClass=*)",
        []string{"nTSecurityDescriptor"},
        nil,
    )
    sr, err := conn.Search(searchRequest)
    if err != nil {
        log.Fatal(err)
    }
    descriptor := sr.Entries[0].GetAttributeValu    ("nTSecurityDescriptor")
    sddl, err := sddlparse.SDDLFromBinary([]byte(descriptor))
    if err != nil {
        log.Fatal(err)
    }
}

```

## License

This library is licensed under the MIT license. See the LICENSE file for more details.
