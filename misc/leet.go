package misc

/*

References:
https://simple.wikipedia.org/wiki/Leet

*/

import (
  "strings"
)

func TextToLeet(input string) (string) {
  input = strings.ToLower(input)
  input = strings.ReplaceAll(input, "e", "3")
  input = strings.ReplaceAll(input, "a", "@")
  input = strings.ReplaceAll(input, "t", "7")
  input = strings.ReplaceAll(input, "l", "1")
  input = strings.ReplaceAll(input, "s", "$")
  input = strings.ReplaceAll(input, "o", "0")
  input = strings.ReplaceAll(input, "h", "#")
  input = strings.ReplaceAll(input, "g", "6")
  input = strings.ReplaceAll(input, "i", "¡")

  return input
}
