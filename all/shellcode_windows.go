package all

import (
	"github.com/D3Ext/maldev/src/redteam"
)

func CreateRemoteThread(shellcode []byte, pid int) error {
	return redteam.CreateRemoteThread(shellcode, pid)
}

func CreateProcess(shellcode []byte, pid int) error {
	return redteam.CreateProcess(shellcode, pid)
}

func Fibers(shellcode []byte) error {
	return redteam.Fibers(shellcode)
}

func QueueUserApc(shellcode []byte) error {
	return redteam.QueueUserApc(shellcode)
}

func UuidFromString(shellcode []byte) error {
	return redteam.UuidFromString(shellcode)
}

func RtlCreateUserThread(shellcode []byte, pid int) error {
	return redteam.RtlCreateUserThread(shellcode, pid)
}

func ProcessHollowing(shellcode []byte, proc string, blockdlls bool) error {
  return redteam.ProcessHollowing(shellcode, proc, blockdlls)
}

