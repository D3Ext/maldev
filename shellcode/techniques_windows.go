package shellcode

import "github.com/D3Ext/Hooka/pkg/hooka"

func CreateRemoteThread(shellcode []byte, pid int) error {
	return hooka.CreateRemoteThread(shellcode, pid)
}

func CreateProcess(shellcode []byte, pid int) error {
	return hooka.CreateProcess(shellcode, pid)
}

func Fibers(shellcode []byte) error {
	return hooka.Fibers(shellcode)
}

func QueueUserApc(shellcode []byte) error {
	return hooka.QueueUserApc(shellcode)
}

func UuidFromString(shellcode []byte) error {
	return hooka.UuidFromString(shellcode)
}

func RtlCreateUserThread(shellcode []byte, pid int) error {
	return hooka.RtlCreateUserThread(shellcode, pid)
}
