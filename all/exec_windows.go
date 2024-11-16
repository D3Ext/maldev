package all

import (
	"github.com/D3Ext/maldev/src/exec"
	"golang.org/x/sys/windows"
)

func ExecuteCommand(comm string) string {
	return exec.ExecuteCommand(comm)
}

func ExecutePowershell(comm string) string {
	return exec.ExecutePowershell(comm)
}

func ExecuteWithToken(comm string, token windows.Token) string {
	return exec.ExecuteWithToken(comm, token)
}
