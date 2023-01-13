package system

import (
  "os/user"
)

func Whoami() (string, error) {
  u, err := user.Current()
  if err != nil {
    return "", err
  }
  return u.Username, nil
}
