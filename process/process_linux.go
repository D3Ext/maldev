package process

/*

References:
https://github.com/mitchellh/go-ps/

type Process interface {
  Pid()         int
  PPid()        int
	Executable()  string
}

*/

import (
	"errors"
	ps "github.com/mitchellh/go-ps"
)

func GetProcesses() ([]ps.Process, error) {
	processes, err := ps.Processes()
	if err != nil {
		return processes, err
	}

	return processes, nil
}

func FindPidByName(name string) ([]int, error) {
	var pids []int
	processes, err := ps.Processes() // Get all processes
	if err != nil {
		return pids, err
	}

	for _, proc := range processes { // Iterate over them
		if proc.Executable() == name { // Check given binary name
			pids = append(pids, proc.Pid()) // Append the pid
		}
	}

	return pids, nil
}

func FindNameByPid(pid int) (string, error) {
	processes, err := ps.Processes()
	if err != nil {
		return "", err
	}

	for _, proc := range processes {
		if proc.Pid() == pid {
			return proc.Executable(), nil
		}
	}

	return "", errors.New("Pid not found!")
}
