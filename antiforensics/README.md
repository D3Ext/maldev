# Antiforensics Examples

## Wipe a file

```go
package main

import (
  "log"

  "github.com/D3Ext/maldev/antiforensics"
)

func main(){
  err = antiforensics.WipeFile("./file.txt")
  if err != nil {
    log.Fatal(err)
  }
}
```

## Timestomping

```go
package main

import (
  "log"

  "github.com/D3Ext/maldev/antiforensics"
)

func main(){
  err := antiforensics.TimestompFile("./file.txt", 5) // func TimestompFile(filename string, count int) (error)
  if err != nil {
    log.Fatal(err)
  }
}
```

