# Regex Examples

## Find substring between

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/src/regex"
)

func main(){
  matches := regex.FindSubstringBetween("Look: ABC some example test ABC", "A", "C")
  fmt.Println(matches) // Output: [B B]
}
```

