package main

/*

Rewrite of the metasploit pattern_create.rb in Golang

References:
https://github.com/rapid7/metasploit-framework/blob/master/tools/exploit/pattern_create.rb

*/

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	// Maldev packages
	l "github.com/D3Ext/maldev/src/logging"
	"github.com/D3Ext/maldev/src/misc"
)

func main() {
	var pattern_length int
	var pattern_to_search string

	flag.IntVar(&pattern_length, "l", 0, "length of the pattern to create")
	flag.StringVar(&pattern_to_search, "p", "", "pattern to search its offset")
	flag.Parse()

	if (pattern_length != 0 && pattern_to_search != "") { // Check if two flags were especified to avoid error (just use one)
		l.Error("Bad parameters!")
		fmt.Println("Usage: ./cyclic_pattern -l 256")
		fmt.Println("       ./cyclic_pattern -p <pattern to search>")
		os.Exit(0)
	} else if (pattern_length == 0 && pattern_to_search == "") {
    fmt.Println("cyclic_pattern:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if (pattern_length != 0) {
		l.Info("Pattern length: " + strconv.Itoa(pattern_length))
		cyclic_pattern := misc.GeneratePattern(pattern_length) // Create cyclic pattern using De Brujin algorithm
		l.Success("Cyclic pattern: " + cyclic_pattern)
		os.Exit(0) // Exit program
	} else if (pattern_to_search != "") {
		cyclic_pattern := misc.GeneratePattern(5000)
		if (strings.Contains(cyclic_pattern, pattern_to_search)) {
			l.Info("Pattern to search: " + l.SCyan(pattern_to_search))
			offset := misc.GetPatternOffset(pattern_to_search)
			l.Success("Offset found at: " + l.SCyan(strconv.Itoa(offset)))
		} else {
			l.Error("Especified pattern offset not found!")
			os.Exit(0)
		}
	}
}
