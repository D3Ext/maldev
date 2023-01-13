package crypto

import (
  "github.com/alwindoss/morse"
  "strings"
)

func Morse(input string) (string, error) {
  h := morse.NewHacker()
  morseCode, err := h.Encode(strings.NewReader(input))
  if err != nil {
    return "", err
  }

  return string(morseCode), nil
}
