package network

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	DefaultTimeout = 5000
)

type RequestInfo struct {
	Response      string
	StatusCode    int
	ContentLength int
	Protocol      string
}

func PostHttpReq(url string, post_data url.Values, timeout int) (RequestInfo, error) {
	client := &http.Client{Timeout: time.Duration(timeout) * time.Millisecond}

	req, err := http.NewRequest("POST", url, strings.NewReader(post_data.Encode()))
	if err != nil {
		return RequestInfo{}, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return RequestInfo{}, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return RequestInfo{}, err
	}

	return RequestInfo{
		Response:      string(bodyBytes),
		StatusCode:    resp.StatusCode,
		ContentLength: int(resp.ContentLength),
		Protocol:      resp.Proto,
	}, err
}
