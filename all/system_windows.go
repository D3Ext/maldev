package all

import "github.com/D3Ext/maldev/src/system"

func GetEdrInfo() (*system.EdrInfo, error) {
	return system.GetEdrInfo()
}

func ListPipes() ([]string, error) {
	return system.ListPipes()
}

func GetWinName() (string, error) {
	return system.GetWinName()
}

func GetUserPrivs() (bool, error) { // Function to check if current user has Administrator privileges
	return system.GetUserPrivs()
}

func Systeminfo() (*system.SystemInfo, error) {
	return system.Systeminfo()
}

func WhoamiAll() (*system.UserInfo, error) {
	return system.WhoamiAll()
}

func GetSid() (string, error) {
	return system.GetSid()
}

func GetRid() (string, error) {
	return system.GetRid()
}

func Uptime() int {
	return system.Uptime()
}
