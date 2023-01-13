package system

import (
  "os"
)

func Env() ([]string) {
  return os.Environ()
}
