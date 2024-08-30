package main

import (
	"flag"
	"fmt"
	"github.com/Laky-64/gologging"
	"github.com/Laky-64/goswift/demangling"
	"github.com/Laky-64/goswift/demangling/utils"
)

func main() {
	var simplified, noSugar, compact, tree, showErrors bool
	flag.Usage = func() {
		fmt.Println("Usage: swift-demangle [options] [mangled-name ...]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}
	flag.BoolVar(&compact, "compact", false, "Compact mode (only emit the demangled names)")
	flag.BoolVar(&simplified, "simplified", false, "Don't display module names or implicit self types")
	flag.BoolVar(&noSugar, "no-sugar", false, "No sugar mode (disable common language idioms such as ? and [] from the output)")
	flag.BoolVar(&tree, "tree", false, "Tree-only mode (do not show the demangled string)")
	flag.BoolVar(&showErrors, "show-errors", false, "Show errors (in demangled output)")
	flag.Parse()
	data := flag.Args()
	options := utils.ModeDefault
	if simplified {
		options |= utils.ModeSimplified
	}
	if noSugar {
		options |= utils.ModeNoSugar
	}
	demangleFunc := func(mangled string) (res string, resErr error) {
		res = mangled
		if context, err := demangling.New([]byte(mangled)); err == nil {
			if result, err := context.Result(); err == nil {
				if tree {
					res = utils.Tree(result)
				} else {
					res = utils.ToString(result, options)
				}
			} else if showErrors {
				resErr = err
			}
		} else if showErrors {
			resErr = err
		}
		return
	}
	writeOutput := func(input string) {
		s, err := demangleFunc(input)
		if err != nil {
			gologging.Error(err)
			return
		}
		if !compact && !tree {
			fmt.Print(fmt.Sprintf("%s ---> ", input))
		}
		fmt.Println(s)
	}
	for _, input := range data {
		writeOutput(input)
	}
	for len(data) == 0 {
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			return
		}
		writeOutput(input)
	}
}
