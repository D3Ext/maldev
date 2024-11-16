package all

import (
	"net/url"

	"github.com/D3Ext/maldev/src/network"
)

func GetAllInterfaces() ([]string, error) {
	return network.GetAllInterfaces()
}

func DownloadFile(file_url string) error {
	return network.DownloadFile(file_url)
}

func PostHttpReq(url string, post_data url.Values, timeout int) (network.RequestInfo, error) {
	return network.PostHttpReq(url, post_data, timeout)
}

func GetInterfaceInfo(interface_name string) (network.CustomInterface, error) {
	return network.GetInterfaceInfo(interface_name)
}

func CheckInternet() bool {
	return network.CheckInternet()
}

func Netstat() ([]*network.PortsInfo, error) {
	return network.Netstat()
}

func FormatedNetstat() (string, error) {
	return network.FormatedNetstat()
}

func GetPublicIp() string {
	return network.GetPublicIp()
}

func GetStatusCode(url_to_check string) (int, error) {
	return network.GetStatusCode(url_to_check)
}
