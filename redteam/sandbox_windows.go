package redteam

/*

This file contains some useful functions to detect sandboxes and avoid being detected

References:
https://www.picussecurity.com/resource/virtualization/sandbox-evasion-how-attackers-avoid-malware-analysis

*/

import (
	"errors"
	"fmt"
	"runtime"
	"time"
	"unsafe"

	filesm "github.com/D3Ext/maldev/files"
	netm "github.com/D3Ext/maldev/network"
	proc "github.com/D3Ext/maldev/process"
	sysm "github.com/D3Ext/maldev/system"

	"golang.org/x/sys/windows"
)

type SandboxConfig struct {
	MinSpace  int
	Cores     int
	Memory    int
	Path      string
	Time      int
	Software  bool
	Drivers   bool
	Internet  bool
	Processes bool
	Network   bool
	Uptime    bool
}

type SandboxResults struct {
	MinSpace  bool
	Cores     bool
	Memory    bool
	Path      string
	Time      bool
	Software  bool
	Drivers   bool
	Internet  bool
	Processes bool
	Network   bool
	Uptime    bool
}

const (
	defaultDiskUsage = 68719476736 // less than 64GB is rare in real systems (use 0 to disable this check)
	defaultCpuCores  = 2           // default number of CPU cores (if system has 2 or less, it's probably a sandbox), use 0 as value to disable check
	defaultMemory    = 4174967296  // less than 4GB of memory may be a sign of sandbox (use 0 to disable check)
	defaultPath      = "C:\\"      // root path to get its total usage
	defaultTime      = 10000       // duration in milliseconds (to detect Sleep Skipping technique), set 0 as value to disable
	defaultSoftware  = true        // look for useful software like Chrome, Firefox and other application
	defaultDrivers   = true        // look for common virtualization drivers
	defaultInternet  = true        // check if system has internet connection
	defaultProcesses = true        // enable process analysis to check sandboxing
	defaultNetwork   = true        // enable network analysis to check sandboxing
	defaultUptime    = true        // get uptime boot to check if date is rare
)

var ( // Define default sandbox enumeration config (recommended)
	DefaultConfig = SandboxConfig{
		MinSpace:  defaultDiskUsage,
		Memory:    defaultMemory,
		Cores:     defaultCpuCores,
		Path:      defaultPath,
		Time:      defaultTime,
		Software:  defaultSoftware,
		Drivers:   defaultDrivers,
		Internet:  defaultInternet,
		Processes: defaultProcesses,
		Network:   defaultNetwork,
		Uptime:    defaultUptime,
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
	"vmware.exe",
	"vmware-vmx.exe",
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
	"volatily.exe",
	"volatily3.exe",
	"DumpIt.exe",
	"dumpit.exe",
}

type memStatusEx struct { // Auxiliary struct to retrieve total memory
	dwLength     uint32
	dwMemoryLoad uint32
	ullTotalPhys uint64
	ullAvailPhys uint64
	unused       [5]uint64
}

var ( // Load necessary DLLs and functions
	kernel32             = windows.NewLazySystemDLL("kernel32.dll")
	getDiskFreeSpaceExW  = kernel32.NewProc("GetDiskFreeSpaceExW")
	globalMemoryStatusEx = kernel32.NewProc("GlobalMemoryStatusEx")
	getTickCount64       = kernel32.NewProc("GetTickCount64")
)

// Main function of this file, receives a config struct (default one recommended)
// Then it returns another custom struct with results, all of them are bools and if it is true it means that may be a sandbox
func AllSandboxChecks(cfg SandboxConfig) (SandboxResults, error) {
	var scanning_results SandboxResults

	if cfg.MinSpace != 0 && cfg.Path != "" { // Check disk usage
		disk_space, err := DiskCheck(cfg.Path)
		if err != nil {
			return scanning_results, err
		}

		if disk_space < cfg.MinSpace { // Check if system space is lower than especified one
			scanning_results.MinSpace = true
		} else {
			scanning_results.MinSpace = false
		}
	}

	if cfg.Cores != 0 { // Check number of CPU cores and compare
		if CpuCheck() < cfg.Cores {
			scanning_results.Cores = true
		} else {
			scanning_results.Cores = false
		}
	}

	if cfg.Memory != 0 { // Get total memory usage
		mem, err := MemoryCheck()
		if err != nil {
			return scanning_results, err
		}
		fmt.Println(mem)

		if mem < cfg.Memory {
			scanning_results.Memory = true
		} else {
			scanning_results.Memory = false
		}
	}

	if cfg.Time != 0 { // Check time (detect Sleep Skipping technique)
		dt := TimeCheck(cfg.Time)
		fmt.Println(dt)
		if dt-50 > cfg.Time/2 {
			scanning_results.Time = false
		} else {
			scanning_results.Time = true
		}
	}

	if cfg.Drivers == true { // Look for common drivers (Virtual Box, VMWare, Qemu...)
		drivers_check, _ := DriversCheck()
		if drivers_check == true {
			scanning_results.Drivers = true
		} else {
			scanning_results.Drivers = false
		}
	}

	if cfg.Internet == true { // Check if there is internet connection
		conn_check := InternetCheck()
		if conn_check == false {
			scanning_results.Internet = true
		} else {
			scanning_results.Internet = false
		}
	}

	if cfg.Processes == true { // Look for potentially analysis apps processes
		proc_check, _, err := ProcessesCheck()
		if err != nil {
			return scanning_results, err
		}

		if proc_check == true {
			scanning_results.Processes = true
		} else {
			scanning_results.Processes = false
		}
	}

	return scanning_results, nil // Return sandbox analysis
}

/*

Here start all the functions which take care one by one of every thing to check

*/

func MemoryCheck() (int, error) { // Get total memory space
	msx := &memStatusEx{
		dwLength: 64,
	}

	r1, _, _ := globalMemoryStatusEx.Call(uintptr(unsafe.Pointer(msx)))
	if r1 == 0 {
		return 0, errors.New("An error has occurred while executing GlobalMemoryStatusEx")
	}

	return int(msx.ullTotalPhys), nil
}

func CpuCheck() int { // Get number of CPUs in system, if system doesn't have more than 2 cores it's probably a sandbox/VM
	return runtime.NumCPU()
}

func UptimeCheck() int {
	return sysm.Uptime()
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

func TimeCheck(timeout int) int {
	// Get current time
	r1, _, _ := getTickCount64.Call(0, 0, 0, 0)

	// Sleep for given timeout (by default 10000ms = 10s)
	time.Sleep(time.Duration(timeout) * time.Millisecond)

	r2, _, _ := getTickCount64.Call(0, 0, 0, 0)
	elapsedTime := int(r2) - int(r1)

	return elapsedTime
}

func InternetCheck() bool {
	return netm.CheckInternet()
}

func DriversCheck() (bool, string) {
	for _, d := range drivers { // Iterate over all drivers to check if they exist
		if filesm.Exists(d) == true {
			return true, d
		}
	}

	return false, ""
}

func ProcessesCheck() (bool, string, error) { // List all processes and check if process name is a known sandboxing process
	all, err := proc.GetProcesses()
	if err != nil {
		return false, "", err
	}

	for _, p := range all {
		for _, p_name := range processes {
			if p.Exe == p_name {
				return true, p.Exe, nil
			}
		}
	}

	return false, "", nil
}
