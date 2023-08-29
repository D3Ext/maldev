package all

import (
	"github.com/D3Ext/maldev/exec"
	"golang.org/x/sys/windows"
)

func ExecuteCommand(comm string) string {
	return exec.ExecuteCommand(comm)
}

func ExecutePowershell(comm string) string {
	return exec.ExecutePowershell(comm)
}

func ExecWithToken(comm string, token windows.Token) string {
	return exec.ExecWithToken(comm, token)
}
