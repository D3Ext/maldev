# Execution examples

## Commands

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/src/exec"
)

func main(){
  out := exec.ExecuteCommand("uname -a") // With bash or cmd
  fmt.Println(out)
}
```

## Powershell Commands

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/src/exec"
)

func main(){
  out := exec.ExecutePowershell("Get-Process -Name notepad") // Using powershell
  fmt.Println(out)
}
```

## Execute Commands with Token

```go
package main

import (
  "fmt"
  "log"

  "github.com/D3Ext/maldev/src/exec"
  "github.com/D3Ext/maldev/src/redteam"
)

func main(){
  // get a valid Windows token (e.g. impersonating a PID)
  token, err := redteam.Impersonate(1312)
  if err != nil {
    log.Fatal(err)
  }

  out := exec.ExecuteWithToken("whoami", token)
  fmt.Println(out)
}
```


