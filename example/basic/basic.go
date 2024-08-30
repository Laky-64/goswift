package main

import (
	"github.com/Laky-64/swift/demangling"
	"github.com/Laky-64/swift/demangling/utils"
	"log"
)

func main() {
	demangler, err := demangling.New([]byte("_$ss5Int64V"))
	if err != nil {
		log.Fatal(demangler)
	}
	node, err := demangler.Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(utils.ToString(node, 0))
}
