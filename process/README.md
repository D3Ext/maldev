# Process Examples

- Every function is different for Linux or Windows

## List all processes

```go
package main

import (
  "fmt"
  "log"

  "github.com/D3Ext/maldev/process"
)

func main(){
  // Linux
  p, err := processes.GetLinuxProcesses()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(p)

  // Windows
  p2, err := process.GetWinProcesses()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(p2)
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
  // Linux
  process_name, err := processes.GetNameByPidLin(1234)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(process_name)

  // Windows
  process_name2, err := process.GetNameByPidWin(1234)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(process_name2)
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
  // Linux
  pids, err := process.GetPidByNameLin("firefox")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(pids)

  // Windows
  pids2, err := process.GetPidByNameWin("firefox.exe")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(pids2)
}
```


