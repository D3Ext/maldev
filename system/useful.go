package system

import (
	"os/exec"
)

func GetUsefulSoftware() ([]string, error) {
	var binaries []string = []string{"docker", "nc", "netcat", "python", "python3", "php", "perl", "ruby", "gcc", "g++", "ping", "base64", "socat", "curl", "wget", "certutil", "xterm", "gpg", "mysql", "ssh"}
	var discovered_software []string

	for _, b := range binaries {
		path, _ := exec.LookPath(b)
		discovered_software = append(discovered_software, path)
	}

	return discovered_software, nil
}
