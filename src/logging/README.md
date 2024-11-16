# Logging Examples

## Status logging functions

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/src/logging"
)

func main(){
  test = "Example text!"

  logging.Success(test)
  logging.Error(test)
  logging.Info(test)
  logging.Warning(test)
  logging.Debug(test)
}
```

## Use colors

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/src/logging"
)

func main(){
  test = "Example text!"

  logging.Green(test)
  logging.Red(test)
  logging.Purple(test)
  logging.Blue(test)
  logging.Cyan(test)

  green_text := logging.SGreen(test)
  red_text := logging.SRed(test)
  purple_text := logging.SPurple(test)
  blue_text := logging.SBlue(test)
  cyan_text := logging.SCyan(test)

  fmt.Println(green_text)
  fmt.Println(red_text)
  fmt.Println(purple_text)
  fmt.Println(blue_text)
  fmt.Println(cyan_text)
}
```

## Create banners

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/src/logging"
)

func main(){
  logging.PrintBanner("Maldev")

  fmt.Println("==================")

  ascii := logging.GetBanner("Maldev")
  fmt.Println(ascii)
}
```



