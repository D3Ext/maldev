package misc

/*

References:
https://gist.github.com/D3Ext/845bdc6a22bbdd50fe409d78b7d59b96
https://rosettacode.org/wiki/De_Bruijn_sequences#Go

*/

import (
  "bytes"
  "strings"
)

const characters string = "abcdefABCDEF123"

func deBruijn(n int) (string) {
  var k int = 15
  alphabet := characters[0:k]

  a := make([]byte, k*n)
  var seq []byte
  var db func(int, int)

  db = func(t, p int) {
    if t > n {
      if (n % p == 0) {
        seq = append(seq, a[1:p+1]...)
      }
    } else {
      a[t] = a[t-p]
      db(t+1, p)
      for j := int(a[t-p] + 1); j < k; j++ {
        a[t] = byte(j)
        db(t+1, t)
      }
    }
  }
  
  db(1, 1)
  var buf bytes.Buffer

  for _, i := range seq {
    buf.WriteByte(alphabet[i])
  }
  b := buf.String()

  return b + b[0:n-1]
}

func GeneratePattern(length int) (string) {
  return deBruijn(4)[0:length]
}

func GetPatternOffset(pattern_str string) (int) {
  original_pattern := deBruijn(4)
  return len(strings.Split(original_pattern, pattern_str)[0]) // Split original cyclic pattern so 0 position is the number of characters to return
}
