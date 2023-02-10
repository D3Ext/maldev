# System Examples

- This package provides some functions which can be really useful during system reconnaissance or whatever you want

## General info

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/system"
)

func main(){
  home, err := system.Home() // Get home folder
  // handle error
  fmt.Println(home)

  fmt.Println(system.Env()) // Acts as a wrapper of os.Environ()

  cwd, err := system.Pwd() // Get current working directory
  // handle error
  fmt.Println(cwd)

  user, err := system.Whoami() // Get current system username
  // handle error
  fmt.Println(user)

  groups, err := system.GetGroups() // List user groups names
  // handle error
  fmt.Println(groups)

  pipes, err := system.GetPipes() // Returns a []string (only for Windows)
  // handle error
  fmt.Println(pipes)

  software, err := system.GetUsefulSoftware() // Returns a []string with useful installed software (e.g. python.exe)
  // handle error
  fmt.Println(software)

  uptime := system.Uptime() // Only for windows
  fmt.Println(uptime)

  sid, err := system.GetSid() // Returns user SID
  // handle error
  fmt.Println(sid)

  rid, err := system.GetRid() // Returns RID
  // handle error
  fmt.Println(rid)
}
```

## Find installed AVs/EDRs

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/system"
)

func main(){
  all_info := system.GetEdrInfo() // Returns a custom struct see edr_windows.go for help

  fmt.Println(all_info.Format()) // Returns a formatted string with info
}
```

## Get full user information

This functions attemps to be a native Golang "whoami /all"

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/system"
)

func main(){
  user_info, err := system.WhoamiAll()
  // handle error

  fmt.Println(user_info.Username)
  fmt.Println(user_info.IntegrityLevel)
  fmt.Println(user_info.Privileges)
}
```


