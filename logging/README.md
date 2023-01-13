# Logging Examples

## Status logs 


```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/logging"
)

func main(){
  test = "Example text!"

  logging.Goodln(test)
  logging.Badln(test) // Note the "ln" on the function as fmt functions
  logging.Infoln(test)

  fmt.Println("===============")

  logging.Good(test)
  logging.Bad(test)
  logging.Info(test)
}
```

## Use colors

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/logging"
)

func main(){
  test = "Example text!"

  logging.Greenln(test)
  logging.Redln(test)
  logging.Purpleln(test)
  logging.Blueln(test)
  logging.Cyanln(test)

  fmt.Println("============")

  logging.Green(test)
  logging.Red(test)
  logging.Purple(test)
  logging.Blue(test)
  logging.Cyan(test)

  fmt.Println("\n============")

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

  "github.com/D3Ext/maldev/logging"
)

func main(){
  logging.PrintBanner("Maldev")

  ascii := logging.GetBanner("Maldev")
  fmt.Println(ascii)
}
```

## Print text with time

```go
package main

import (
  "github.com/D3Ext/maldev/logging"
)

func main(){
  logging.TimePrintln("This is an example!")
}
```


