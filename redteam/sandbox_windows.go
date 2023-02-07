package redteam

/*

This file contains some useful functions to detect sandboxes and avoid being detected

References:
https://www.picussecurity.com/resource/virtualization/sandbox-evasion-how-attackers-avoid-malware-analysis

*/

import (
  "errors"
  "unsafe"
  "runtime"
  "time"

  proc "github.com/D3Ext/maldev/process"

  "golang.org/x/sys/windows"
)

type SandboxConfig struct {
  MinSpace    int
  Cores       int
  Memory      int
  Path        string
  Time        int
  Internet    bool
  Processes   bool
  Network     bool
}

const (
  defaultDiskUsage  = 68719476736 // less than 64GB is rare in real systems
  defaultCpuCores   = 2           // default number of CPU cores (if system has 2 or less, it's probably a sandbox)
  defaultMemory     = 4174967296  // less than 4GB of memory may be a sign of sandbox
  defaultPath       = "C:\\"      // root path to get its total usage
  defaultTime       = 10000       // duration in milliseconds (to detect Sleep Skipping technique)
  defaultSoftware   = true        // look for useful software like Chrome, Firefox and other application
  defaultInternet   = true        // check if system has internet connection
  defaultProcesses  = true        // enable process analysis to check sandboxing
  defaultNetwork    = true        // enable network analysis to check sandboxing
)

var ( // Define default sandbox enumeration config (recommended)
  DefaultConfig = SandboxConfig{
    MinSpace: defaultDiskUsage,
    Memory: defaultMemory,
    Cores: defaultCpuCores,
    Path: defaultPath,
    Time: defaultTime,
    Internet: defaultInternet,
    Processes: defaultProcesses,
    Network: defaultNetwork,
  }
)

var drivers []string = []string{ // Sandbox drivers taken from https://github.com/LordNoteworthy/al-khaser
  "C:\\Windows\\System32\\drivers\\VBoxMouse.sys",
  "C:\\Windows\\System32\\drivers\\VBoxGuest.sys",
  "C:\\Windows\\System32\\drivers\\VBoxSF.sys",
  "C:\\Windows\\System32\\drivers\\VBoxVideo.sys",
  "C:\\Windows\\System32\\vboxdisp.dll",
  "C:\\Windows\\System32\\vboxhook.dll",
  "C:\\Windows\\System32\\vboxmrxnp.dll",
  "C:\\Windows\\System32\\vboxogl.dll",
  "C:\\Windows\\System32\\vboxoglarrayspu.dll",
  "C:\\Windows\\System32\\vboxservice.exe",
  "C:\\Windows\\System32\\vboxtray.exe",
  "C:\\Windows\\System32\\VBoxControl.exe",
  "C:\\Windows\\System32\\drivers\\vmmouse.sys",
  "C:\\Windows\\System32\\drivers\\vmhgfs.sys",
  "C:\\Windows\\System32\\drivers\\vmci.sys",
  "C:\\Windows\\System32\\drivers\\vmmemctl.sys",
  "C:\\Windows\\System32\\drivers\\vmmouse.sys",
  "C:\\Windows\\System32\\drivers\\vmrawdsk.sys",
  "C:\\Windows\\System32\\drivers\\vmusbmouse.sys",
}

var processes []string = []string{ // Sandbox processes taken from https://github.com/LordNoteworthy/al-khaser
  "vboxservice.exe",
  "vboxtray.exe",
  "vmtoolsd.exe",
  "vmwaretray.exe",
  "vmwareuser",
  "VGAuthService.exe",
  "vmacthlp.exe",
  "vmsrvc.exe",
  "vmusrvc.exe",
  "xenservice.exe",
  "qemu-ga.exe",
  // Some extra processes (added by me)
  "wireshark.exe",
  "Procmon.exe",
  "Procmon64.exe",
}

type memStatusEx struct { // Auxiliary struct to retrieve total memory
  dwLength     uint32
  dwMemoryLoad uint32
  ullTotalPhys uint64
  ullAvailPhys uint64
  unused       [5]uint64
}

var ( // Load necessary DLLs and functions
  kernel32              = windows.NewLazySystemDLL("kernel32.dll")
  getDiskFreeSpaceExW   = kernel32.NewProc("GetDiskFreeSpaceExW")
  globalMemoryStatusEx  = kernel32.NewProc("GlobalMemoryStatusEx")
  getTickCount64        = kernel32.NewProc("GetTickCount64")
)

func MemoryCheck() (int, error) { // Get total memory space
  msx := &memStatusEx{
    dwLength: 64,
  }
  
  r1, _, _ := globalMemoryStatusEx.Call(uintptr(unsafe.Pointer(msx)))
  if r1 == 0 {
    return 0, errors.New("An error has ocurred while executing GlobalMemoryStatusEx")
  }

  return int(msx.ullTotalPhys), nil
}

func CpuCheck() (int) { // Get number of CPUs in system, if system doesn't have more than 2 cores it's probably a sandbox/VM
  return runtime.NumCPU()
}

func DiskCheck(path string) (int, error) { // Get disk space of given path (by default "C:\\")
  lpTotalNumberOfBytes := int64(0)

  diskret, _, err := getDiskFreeSpaceExW.Call(
    uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(path))),
    uintptr(0),
    uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
    uintptr(0),
  )
  
  if diskret == 0 {
    return 0, err
  }

  return int(lpTotalNumberOfBytes), nil
}

func TimeCheck(timeout int) (int) {
  // Get current time
  r1, _, _ := getTickCount64.Call(0, 0, 0, 0)

  // Sleep for given timeout (by default 10000ms = 10s)
  time.Sleep(time.Duration(timeout) * time.Millisecond)

  r2, _, _ := getTickCount64.Call(0, 0, 0, 0)
  elapsedTime := int(r2) - int(r1)

  return elapsedTime
}

func ProcessesCheck() (bool, error) { // List all processes and check if process name is a known sandboxing process
  all, err := proc.GetProcesses()
  if err != nil {
    return false, err
  }

  for _, p := range all {
    for _, p_name := range processes {
      if p.Exe == p_name {
        return true, nil
      }
    }
  }

  return false, nil
}



