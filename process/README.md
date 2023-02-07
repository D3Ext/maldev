# Process Examples

- This package provides some easy-to-use functions to work with processes. Supported for Linux and Windows

## List all processes

```go
package main

import (
  "fmt"
  "log"

  "github.com/D3Ext/maldev/process"
)

func main(){
  p, err := processes.GetProcesses() // func GetProcesses() ([]ps.Process, error)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(p)
}
```

## Find name by pid

```go
package main

import (
  "fmt"
  "log"

  "github.com/D3Ext/maldev/process"
)

func main(){
  process_name, err := processes.FindNameByPid(1234) // func FindNameByPid(pid int) (string, error)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(process_name)
}

```

## Find pid(s) by name

```go
package main

import (
  "fmt"
  "log"

  "github.com/D3Ext/maldev/process"
)

func main(){
  pids, err := process.FindPidByName("firefox") // func FindPidByName(name string) ([]int, error)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(pids)
}
```


