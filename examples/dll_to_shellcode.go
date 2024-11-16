package main

import (
	"flag"
  "log"
  "fmt"
	"os"
	"time"

	l "github.com/D3Ext/maldev/src/logging"
	"github.com/D3Ext/maldev/src/redteam"
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
    fmt.Println("dll_to_shellcode:")
		flag.PrintDefaults()
    fmt.Println()
    l.Error("Parameters missing! Example: dll_to_shellcode -p evil.dll -f xyz -o shellcode.bin")
		os.Exit(0)
	}

	l.Info("Converting " + pe + " to independent position shellcode...")
	sc, err := redteam.ConvertDllToShellcode(pe, function, "")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(200 * time.Millisecond)
	l.Info("Writing raw shellcode to " + output)

	err = redteam.WriteShellcodeToFile(output, sc)
	if err != nil {
		log.Fatal(err)
	}

	l.Success("Process completed!")
}
