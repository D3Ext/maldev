package misc

/*

References:
https://gist.github.com/D3Ext/845bdc6a22bbdd50fe409d78b7d59b96
https://rosettacode.org/wiki/De_Bruijn_sequences#Go

*/

import (
  "math/rand"
  "time"
)

func RandomString(length int) (string) { // Return random string passing an integer (length)
  var seededRand *rand.Rand = rand.New(
  rand.NewSource(time.Now().UnixNano()))
  const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}

func RandomInt(max int, min int) (int) { // Return a random number between max and min
  rand.Seed(time.Now().UnixNano())
  rand_int := rand.Intn(max - min + 1) + min
  return rand_int
}


