package network

/*

References:
https://www.tutorialspoint.com/how-to-get-the-response-status-code-in-golang
https://stackoverflow.com/questions/72607685/getting-the-status-code-from-a-get-request-in-golang

*/

import (
  "net/http"
)

func GetStatusCode(url_to_check string) (int, error) { // Return status code of given url
  req, err := http.NewRequest("GET", url_to_check, nil) // Create request
  if err != nil {
    return 0, err
  }

  req.Header.Set("Accept", "application/x-www-form-urlencoded") // Set common headers
  req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")

  client := &http.Client{}
  resp, err := client.Do(req) // Send request
  if err != nil {
    return 0, err
  }
  defer resp.Body.Close()

  return resp.StatusCode, nil // no errors
}


