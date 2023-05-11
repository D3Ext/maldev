package all

import (
  sc "github.com/D3Ext/maldev/shellcode"
)

func ConvertDllToShellcode(dll_file string, func_name string, data string) ([]byte, error) {
  return sc.ConvertDllToShellcode(dll_file, func_name, data)
}

func CreateRemoteThread(shellcode []byte, pid int) (error) {
  return sc.CreateRemoteThread(shellcode, pid)
}

func CreateProcess(shellcode []byte, pid int) (error) {
  return sc.CreateProcess(shellcode, pid)
}

func Fibers(shellcode []byte) (error) {
  return sc.Fibers(shellcode)
}

func QueueUserApc(shellcode []byte) (error) {
  return sc.QueueUserApc(shellcode)
}

func UuidFromString(shellcode []byte) (error) {
  return sc.UuidFromString(shellcode)
}

func RtlCreateUserThread(shellcode []byte, pid int) (error) {
  return sc.RtlCreateUserThread(shellcode, pid)
}

func GetShellcodeFromUrl(sc_url string) ([]byte, error) {
  return sc.GetShellcodeFromUrl(sc_url)
}

func GetShellcodeFromFile(file string) ([]byte, error) {
  return sc.GetShellcodeFromFile(file)
}

func WriteShellcodeToFile(filename string, shellcode []byte) (error) {
  return sc.WriteShellcodeToFile(filename, shellcode)
}



