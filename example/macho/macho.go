package main

import (
	"fmt"
	"github.com/Laky-64/gologging"
	"github.com/Laky-64/swift/demangling/utils"
	"github.com/Laky-64/swift/proxy"
	"github.com/blacktop/go-macho"
	"log"
)

func main() {
	open, err := macho.Open("FileName")
	if err != nil {
		log.Fatal(err)
	}
	proxyMacho, err := proxy.New(open)
	if err != nil {
		log.Fatal(err)
	}
	fields, _ := proxyMacho.GetSwiftFields()
	for _, t := range fields {
		if t.Type == "YOUR_FIELD_TYPE" {
			node, err := proxyMacho.Demangle(t.Records[0].MangledTypeNameOffset.GetAddress())
			if err != nil {
				gologging.Fatal(err)
			}
			fmt.Println(utils.ToString(node, 0))
		}
	}
	defer func() {
		_ = proxyMacho.Close()
	}()
}
