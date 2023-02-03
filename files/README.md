# File Examples

## Get file content

```go
package main

import (
  "fmt"
  "log"

  "github.com/D3Ext/maldev/files"
)

func main(){
  content, err := files.GetFileContent("/path/to/file.txt")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(content)
}
```

## Check if file exists

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/file"
)

func main(){
  check := files.Exists("/path/to/file")
  fmt.Println(check) // true or false
}
```

## Check if path is a file

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/files"
)

func main(){
  check := files.IsFile("/path/to/file")
  fmt.Println(check)
}
```

## Check if path is a directory

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/files"
)

func main(){
  check := files.IsDir("/path/to/dir")
  fmt.Println(check)
}
```

## Copy a file or dir recursively

```go
package main

import (
  "log"

  "github.com/D3Ext/maldev/files"
)

func main(){
  err := files.Copy("/path/to/file.txt", "/path/to/destination")
  if err != nil {
    log.Fatal(err)
  }
}
```

## Move a file or dir

```go
package main

import (
  "log"

  "github.com/D3Ext/maldev/files"
)

func main(){
  err := files.Move("/path/to/file.txt", "/path/to/destination/file.txt")
  if err != nil {
    log.Fatal(err)
  }
}
```


