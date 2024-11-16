# Miscelaneous Examples

This directory contains a bunch of miscelaneous functions which can be really useful

## Generate random string of n length

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/src/misc"
)

func main(){
  rand_str := misc.RandomString(10)
  fmt.Println(rand_str)
}
```

## Generate random int of n length

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/src/misc"
)

func main(){
  rand_int := misc.RandomInt(10)
  fmt.Println(rand_int) // 3894276149
}
```

## Create a cyclic pattern of n length (De Bruijn algorithm)

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/src/misc"
)

func main(){
  pattern := misc.GeneratePattern(1000)
  fmt.Println(pattern)

  // Get offset from some 4bits substring of pattern
  offset := misc.GetOffset("Jeek")
}
```

## Epoch date

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/src/misc"
)

func main(){
  // Convert date to epoch format
  epoch := misc.DateToEpoch(2023, 1, 8, 11, 23, 0) // DateToEpoch(year, month, day, hour, minute, second int)
  fmt.Println(epoch) // Output: 1673176980

  // Convert epoch to date
  date := misc.EpochToDate(epoch)
  fmt.Println(date) // Output: 2023-1-8 11:23:0
}
```

## Convert text to Leet

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/src/misc"
)

func main(){
  leet := misc.TextToLeet("This is a test")
  fmt.Println(leet) // Output: 7#¡$ ¡$ @ 73$7
}
```

## Get a random user-agent

- This function picks a random user-agent from a generic list

```go
package main

import (
  "fmt"

  "github.com/D3Ext/maldev/src/misc"
)

func main(){
  user_agent := misc.GetRandomAgent()
  fmt.Println(user_agent) // Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36
}
```

## Convert bytes to GB

```go
package main

import (
  "fmt"
  "github.com/D3Ext/maldev/src/misc"
)

func main(){
  total_gb := misc.BytesToGb(21474836480)
  fmt.Println(total_gb)
}
```

## Convert GB to bytes

```go
package main

import (
  "fmt"
  "github.com/D3Ext/maldev/src/misc"
)

func main(){
  total_bytes := misc.GbToBytes(20)
  fmt.Println(total_bytes)
}
```

