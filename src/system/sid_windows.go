package system

import (
	"strings"
	"syscall"
)

func GetSid() (string, error) {
	var token syscall.Token

	pid, err := syscall.GetCurrentProcess() // Get self process
	if err != nil {
		return "", err
	}

	err = syscall.OpenProcessToken(pid, syscall.TOKEN_QUERY, &token) // Open token process
	if err != nil {
		return "", err
	}
	defer syscall.CloseHandle(syscall.Handle(token))

	info, err := token.GetTokenUser() // Get information about token
	if err != nil {
		return "", err
	}

	sid, err := info.User.Sid.String() // Convert Sid to string
	if err != nil {
		return "", err
	}

	return sid, nil // Return sid
}

func GetRid() (string, error) {
	full_sid, err := GetSid()
	if err != nil {
		return "", err
	}

	rid := strings.Split(full_sid, "-")[len(strings.Split(full_sid, "-"))-1] // Get last part of the SID which is the RID
	return rid, nil
}
