package network

/*

References:
https://stackoverflow.com/questions/11692860/how-can-i-efficiently-download-a-large-file-using-go

*/

import (
  "net/http"
  "strings"
  "io"
  "os"
)

func DownloadFile(file_url string) (error) { // This function downloads file from url (e.g. https://example.com/myfile.png)
  s := strings.Split(file_url, "/")
  filename := s[len(s)-1]

  f, err := os.Create(filename) // Create filename
  if err != nil {
    return err
  }
  defer f.Close()

  req, err := http.NewRequest("GET", file_url, nil) // Create request
  if err != nil {
    return err
  }

  req.Header.Set("Accept", "application/x-www-form-urlencoded") // Set common headers
  req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")

  client := &http.Client{}
  resp, err := client.Do(req) // Send request
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  // Write the body to file
  _, err = io.Copy(f, resp.Body) // Copy request body to file descriptor
  if err != nil  {
    return err
  }

  return nil // no errors
}
