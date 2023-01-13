package system

import (
  "os"
)

func Pwd() (string, error) { // Get current working directory
  path, err := os.Getwd()
  if err != nil {
    return "", err
  }

  return path, nil
}
