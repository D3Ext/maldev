# Execution examples

## Commands

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/exec"
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

  "github.com/D3Ext/maldev/exec"
)

func main(){
  out := exec.ExecutePowershell("Get-Process -Name notepad") // Using powershell
  fmt.Println(out)
}
```



