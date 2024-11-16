package all

import (
	"github.com/D3Ext/maldev/src/process"
	ps "github.com/mitchellh/go-ps"
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
