package process

/*

References:
https://github.com/mitchellh/go-ps
https://stackoverflow.com/questions/36333896/how-to-get-process-id-by-process-name-in-windows-environment
https://www.socketloop.com/tutorials/golang-get-executable-name-behind-process-id-example

*/

import (
  "errors"
  "golang.org/x/sys/windows"
  "unsafe"
  "syscall"
)

const TH32CS_SNAPPROCESS = 0x00000002

type WindowsProcess struct { // Windows process structure
  ProcessID       int     // PID
  ParentProcessID int
  Exe             string  // Cmdline executable (e.g. explorer.exe)
}

func GetPidByName(name string) ([]int, error) { // Return all PIDs of a binary
  var pids_to_return []int
  all_processes, err := GetProcesses()
  if err != nil {
    return pids_to_return, err
  }

  for _, process := range all_processes { // Iterate over all processes
    if process.Exe == name {
      pids_to_return = append(pids_to_return, process.ProcessID) // Append every match
    }
  }
  return pids_to_return, nil
}

func GetNameByPid(pid int) (string, error) { // Return process name (.exe) of a specific PID
  all_processes, err := GetProcesses()
  if err != nil {
    return "", err
  }

  for _, process := range all_processes {
    if process.ProcessID == pid {
      return process.Exe, nil
    }
  }
  return "", errors.New("PID not found")
}

func GetProcesses() ([]WindowsProcess, error) { // Get all processes using native windows API
  handle, err := windows.CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS, 0)
  if err != nil {
    return nil, err
  }
  defer windows.CloseHandle(handle)
  var entry windows.ProcessEntry32
  entry.Size = uint32(unsafe.Sizeof(entry))

  err = windows.Process32First(handle, &entry)
  if err != nil {
    return nil, err
  }

  results := make([]WindowsProcess, 0, 50)
  for {
    results = append(results, NewWindowsProcess(&entry))

    err = windows.Process32Next(handle, &entry)
    if err != nil {
      if err == syscall.ERROR_NO_MORE_FILES {
        return results, nil
      }
      return nil, err
    }
  }
}

func NewWindowsProcess(e *windows.ProcessEntry32) WindowsProcess { // Auxiliar function (don't use it)
  end := 0
  for {
    if e.ExeFile[end] == 0 {
      break
    }
    end++
  }

  return WindowsProcess{
    ProcessID:       int(e.ProcessID),
    ParentProcessID: int(e.ParentProcessID),
    Exe:             syscall.UTF16ToString(e.ExeFile[:end]),
  }
}



