package main

import (
	"flag"

	"github.com/huner2/go-sddlparse/v2"
)

func main() {
	sddl := flag.String("sddl", "", "SDDL string to parse")
	flag.Parse()
	if *sddl == "" {
		flag.Usage()
		return
	}

	parsed, err := sddlparse.SDDLFromBase64Encoded([]byte(*sddl))
	if err != nil {
		panic(err)
	}
	str, err := parsed.MustString()
	if err != nil {
		panic(err)
	}
	println(str)
}
