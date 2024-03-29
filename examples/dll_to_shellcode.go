package main

import (
	"flag"
	"os"
	"time"

	l "github.com/D3Ext/maldev/logging"
	"github.com/D3Ext/maldev/shellcode"
)

func main() {
	var pe string
	var output string
	var function string

	flag.StringVar(&pe, "p", "", "Portable Executable to convert to shellcode")
	flag.StringVar(&output, "o", "", "raw shellcode output file")
	flag.StringVar(&function, "f", "", "function name to execute (if using msfvenom DLL use random name)")
	flag.Parse()

	if (pe == "") || (output == "") || (function == "") {
		l.Badln("Parameters missing! Example: ", os.Args[0], " -p evil.dll -f xyz -o shellcode.bin")
		flag.PrintDefaults()
		os.Exit(0)
	}

	l.Infoln("Converting " + pe + " to independent position shellcode...")
	sc, err := shellcode.ConvertDllToShellcode(pe, function, "")
	if err != nil {
		l.Fatal(err)
	}
	time.Sleep(200 * time.Millisecond)
	l.Infoln("Writing raw shellcode to " + output)

	err = shellcode.WriteShellcodeToFile(output, sc)
	if err != nil {
		l.Fatal(err)
	}

	l.Goodln("Process completed!")
}
