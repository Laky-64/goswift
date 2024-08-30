package main

import (
	"log"
	"swift/demangling"
	"swift/demangling/utils"
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
