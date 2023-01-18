# Scanning Examples

## Check if URL is up

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/scanning"
)

func main(){
  timeout := 1000 // Url timeout must be especified in milliseconds, however you can use "scanning.UrlTimeout" to use default timeout
  check := scanning.CheckUrl("https://google.com", timeout) // func CheckUrl(url_to_check string, timeout int) (bool)
  fmt.Println(check)
}
```

## Get full URL of domain

```go
package main

import (
  "fmt"
  "log"

  "github.com/D3Ext/maldev/scanning"
)

func main(){
  url, err := scanning.GetHttpFromDomain("google.com", 1000) // GetHttpFromDomain(domain string, timeout int) (string, error)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(url)
}
```

## Hostscan

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/scanning"
)

func main(){
  ips, err := scanning.Hostscan("192.168.1.0/24", 300) // Hostscan(ip_range string, ping_timeout int) ([]string, error)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(ips)
}
```

## Check if IP is up

```go
package main

import (
  "fmt"
  "log"

  "github.com/D3Ext/maldev/scanning"
)

func main(){
  timeout := 1200 // Timeout in milliseconds, you can use "scanning.PingTimeout" as a default timeout
  check, err := scanning.CheckIfUp("192.168.1.6", timeout)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(check)
}
```

## Scan ports of ip

```go
package main

import (
  "fmt"
  "log"

  "github.com/D3Ext/maldev/scanning"
)

func main(){
  all_ports, err := scanning.PortscanAll("192.168.1.6") // Scan all ports of an ip (65535)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(all_ports)

  common_ports, err := scanning.PortscanCommon("192.168.1.6") // Scan most common ports of an ip
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(common_ports)
}
```

## Check if TCP/UDP port is open

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/scanning"
)

func main(){
  check := scanning.CheckOpenTcpPort("192.168.1.6", "8080") // func CheckOpenTcpPort(ip string, port string) (bool)
  fmt.Println(check)

  check2 := scanning.CheckOpenUdpPort("192.168.1.6", "53") // func CheckOpenUdpPort(ip string, port string) (bool)
  fmt.Println(check2)
}
```

## Get all subdomains of a domain

- It uses some third-party APIs like crt.sh, AlienVault, HackerTarget (don't use domains from CTFs because they don't exist)

```go
package main

import (
  "fmt"
  "log"

  "github.com/D3Ext/maldev/scanning"
)

func main(){
  subdoms, err := scanning.GetAllSubdomains("hackthebox.com") // func GetAllSubdomains(dom string) ([]string, error)
  if err != nil {
    log.Fatal(err)
  }
}
```

## Whois (domain)

```go
package main

import (
  "fmt"
  "log"

  "github.com/D3Ext/maldev/scanning"
)

func main(){
  data, err := scanning.WhoisDomain("google.com")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(data.Domain)
}
```



