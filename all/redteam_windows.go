package all

import (
	"github.com/D3Ext/maldev/src/redteam"
	"golang.org/x/sys/windows"
)

func ConvertDllToShellcode(dll_file string, func_name string, data string) ([]byte, error) {
	return redteam.ConvertDllToShellcode(dll_file, func_name, data)
}

func AutoDumpHashes() ([]redteam.Hash, error) {
	return redteam.AutoDumpHashes()
}

func DumpSamHashes(system_file string, sam_file string) ([]redteam.Hash, error) {
	return redteam.DumpSamHashes(system_file, sam_file)
}

func DumpDitHashes(system_file string, ntds_file string) ([]redteam.Hash, error) {
	return redteam.DumpDitHashes(system_file, ntds_file)
}

func GetToken() (windows.Token, error) {
	return redteam.GetToken()
}

func Impersonate(pid int) (windows.Token, error) {
	return redteam.Impersonate(pid)
}

func AllSandboxChecks(cfg redteam.SandboxConfig) (redteam.SandboxResults, error) {
	return redteam.AllSandboxChecks(cfg)
}

func MemoryCheck() (int, error) {
	return redteam.MemoryCheck()
}

func CpuCheck() int {
	return redteam.CpuCheck()
}

func UptimeCheck() int {
	return redteam.UptimeCheck()
}

func DiskCheck(path string) (int, error) {
	return redteam.DiskCheck(path)
}

func TimeCheck(timeout int) int {
	return redteam.TimeCheck(timeout)
}

func InternetCheck() bool {
	return redteam.InternetCheck()
}

func DriversCheck() (bool, string) {
	return redteam.DriversCheck()
}

func ProcessesCheck() (bool, string, error) {
	return redteam.ProcessesCheck()
}

