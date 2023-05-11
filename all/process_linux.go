package all

import (
  ps "github.com/mitchellh/go-ps"
  "github.com/D3Ext/maldev/process"
)

func GetProcesses() ([]ps.Process, error) {
  return process.GetProcesses()
}

func FindPidByName(name string) ([]int, error) {
  return process.FindPidByName(name)
}

func FindNameByPid(pid int) (string, error) {
  return process.FindNameByPid(pid)
}


