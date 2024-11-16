# Slices Examples

## Remove duplicates

- From a []string

```go
package main

import (
  "fmt"
  
  "github.com/D3Ext/maldev/src/slices"
)

func main(){
  unique := slices.RemoveDuplicatesStr([]string{"first", "second", "third", "first", "third"})
  fmt.Println(unique)
}
```

- From a []int

```go
package main

import (
  "fmt"
  
  "github.com/D3Ext/maldev/src/slices"
)

func main(){
  unique := slices.RemoveDuplicatesInt([]int{1, 2, 3, 1, 3})
  fmt.Println(unique)
}
```

## Check if slice contains

- Strings

```go
package main

import (
  "fmt"
  
  "github.com/D3Ext/maldev/src/slices"
)

func main(){
  check := slices.StringSliceContains([]string{"first", "second", "third", "first", "third"}, "second")
  fmt.Println(check)
}
```

- Ints

```go
package main

import (
  "fmt"
  
  "github.com/D3Ext/maldev/src/slices"
)

func main(){
  check := slices.IntSliceContains([]int{1, 2, 3, 1, 3}, 2)
  fmt.Println(check)
}
```

- Strings (insensitive)

```go
package main

import (
  "fmt"
  
  "github.com/D3Ext/maldev/src/slices"
)

func main(){
  check := slices.StringSliceContainsInsensitive([]string{"fiRSt", "seCoNd", "tHirD", "FirSt", "ThiRd"}, "second")
  fmt.Println(check)
}
```



