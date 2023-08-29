package network

/*

References:
https://stackoverflow.com/questions/41670155/get-public-ip-in-golang
https://gist.github.com/ankanch/8c8ec5aaf374039504946e7e2b2cdf7f

*/

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type IP struct {
	Query string
}

func GetPublicIp() string { // Send a signle fast http request to retrieve public ip
	req, err := http.Get("http://ip-api.com/json/") // Send GET request
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body) // Read response
	if err != nil {
		return err.Error()
	}

	var ip IP
	json.Unmarshal(body, &ip) // Parse request response with struct
	return ip.Query           // Return ip
}
