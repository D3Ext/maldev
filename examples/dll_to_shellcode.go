package main

import (
  "time"
  "flag"
  "os"

  l "github.com/D3Ext/maldev/logging"
  "github.com/D3Ext/maldev/shellcode"
)

func ParseFlags() (string, string, string) {
  var pe_filename string
  var output_filename string
  var function_name string

  flag.StringVar(&pe_filename, "p", "", "Portable Executable to convert to shellcode")
  flag.StringVar(&output_filename, "o", "", "raw shellcode output file")
  flag.StringVar(&function_name, "f", "", "function name to execute (if using msfvenom DLL use random name)")
  flag.Parse()

  return pe_filename, output_filename, function_name // Return all param values
}

func main(){
  pe, output, function := ParseFlags()

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


