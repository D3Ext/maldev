# Network Examples

## Download a file from url

```go
package main

import (
  "fmt"
  "log"

  "github.com/D3Ext/maldev/network"
)

func main(){
  err := network.DownloadFile("https://example.com/file.txt") // File saved as filename, in this case file.txt
  if err != nil {
    log.Fatal(err)
  }
}
```

## Get status code of url

```go
package main

import (
  "fmt"
  "log"

  "github.com/D3Ext/maldev/network"
)

func main(){
  code, err := network.GetStatusCode("https://google.com")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(code)
}
```

## List all network interfaces

```go
package main

import (
  "fmt"
  "log"

  "github.com/D3Ext/maldev/network"
)

func main(){
  interfaces, err := network.ListAllInterfaces() // func ListAllInterfaces() ([]string, error)
  if err != nil {
    log.Fatal(err)
  }
}
```

## Get all information about an interface

```go
package main

import (
  "fmt"
  "log"

  "github.com/D3Ext/maldev/network"
)

func main(){
  info, err := network.GetInterfaceInfo("wlan0") // func GetInterfaceInfo(interface_name string) (net.Interface, error)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(info)
  format := network.FormatInfo(info)
  fmt.Println(format)
}
```

## Check internet connection

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/network"
)

func main(){
  check := network.CheckInternet()
  fmt.Println(check) // true or false
}
```

## List active ports

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/network"
)

func main(){
  raw_netstat := network.Netstat() // func Netstat() ([]*PortsInfo, error)
  fmt.Println(raw_netstat)

  formated_netstat := network.FormatedNetstat() // func FormatedNetstat() (string, error)
  fmt.Println(formated_netstat) // Represented in columns
}
```

## Get public ip

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/network"
)

func main(){
  ip := network.GetPublicIp()
  fmt.Println(ip)
}
```




