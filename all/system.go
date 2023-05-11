package all

import "github.com/D3Ext/maldev/system"

func GetGroups() ([]string, error) {
  return system.GetGroups()
}

func Env() ([]string) {
  return system.Env()
}

