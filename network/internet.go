package network

/*

References:
https://stackoverflow.com/questions/50056144/check-for-internet-connection-from-application

*/

import (
	"net/http"
	"time"
)

func CheckInternet() bool {
	client := http.Client{
		Timeout: 3000 * time.Millisecond, // 3s timeout (more than necessary)
	}

	_, err := client.Get("https://google.com")

	if err != nil {
		return false
	}

	return true
}
