package scanning

import (
  "net/http"
  "time"
  "crypto/tls"
)

func CheckUrl(url_to_check string, timeout int) (bool) { // This function check if provided URL is up
  req := http.Client{
    Timeout: time.Duration(timeout) * time.Millisecond,
  }

  response, err := req.Get(url_to_check)
  if err != nil {
    return false
  }

  if response.StatusCode != 0 {
    return true
  } else {
    return false
  }
}

func GetHttpFromDomain(domain string, timeout int) (string, error) {
  http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

  req := http.Client{
    Timeout: time.Duration(timeout) * time.Millisecond,
  }

  response, err := req.Get("http://" + domain) // Send request against http:// so if the web is https:// it will redirect
  if err != nil {
    return "", err
  }

  final_url := response.Request.URL.String() // Finally return final url
  return final_url, nil
}
