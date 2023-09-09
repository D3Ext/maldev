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
	l "github.com/D3Ext/maldev/logging"
	"github.com/D3Ext/maldev/misc"
)

func main() {
	var pattern_length int
	var pattern_to_search string

	flag.IntVar(&pattern_length, "l", 0, "length of the pattern to create")
	flag.StringVar(&pattern_to_search, "p", "", "pattern to search its offset")
	flag.Parse()

	if pattern_length != 0 && pattern_to_search != "" { // Check if two flags were especified to avoid error (just use one)
		l.Badln("Bad parameters!")
		fmt.Println("Usage: ./cyclic_pattern -l 256")
		fmt.Println("       ./cyclic_pattern -p <pattern to search>")
		os.Exit(0)
	} else if pattern_length == 0 && pattern_to_search == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if pattern_length != 0 {
		l.Infoln("Pattern length: " + strconv.Itoa(pattern_length))
		cyclic_pattern := misc.GeneratePattern(pattern_length) // Create cyclic pattern using De Brujin algorithm
		l.Goodln("Cyclic pattern: " + cyclic_pattern)
		os.Exit(0) // Exit program
	} else if pattern_to_search != "" {
		cyclic_pattern := misc.GeneratePattern(50000)
		if strings.Contains(cyclic_pattern, pattern_to_search) {
			l.Infoln("Pattern to search: " + l.SCyan(pattern_to_search))
			offset := misc.GetPatternOffset(pattern_to_search)
			l.Goodln("Offset found at: " + l.SCyan(strconv.Itoa(offset)))
		} else {
			l.Badln("Especified pattern offset not found!")
			os.Exit(0)
		}
	}
}
