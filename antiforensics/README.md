# Antiforensics Examples

## Wipe a file or dir (recursive)

```go
package main

import (
  "log"

  "github.com/D3Ext/maldev/antiforensics"
)

func main(){
  err = antiforensics.Wipe("file.txt")
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
  err := antiforensics.Timestomp("./file.txt", 5) // func TimestompFile(filename string, count int) (error)
  if err != nil {
    log.Fatal(err)
  }
}
```

