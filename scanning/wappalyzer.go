package scanning

import (
  "io/ioutil"
  "net/http"

  wappalyzer "github.com/projectdiscovery/wappalyzergo"
)

func Wappalyzer(url string) ([]string, error) {
  var technologies []string

  resp, err := http.DefaultClient.Get(url)
  if err != nil {
    return technologies, err
  }
  
  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return technologies, err
  }

  wappalyzerClient, err := wappalyzer.New()
  if err != nil {
    return technologies, err
  }

  fingerprints := wappalyzerClient.Fingerprint(resp.Header, data)
  
  for entry, _ := range fingerprints {
    technologies = append(technologies, entry)
  }

  return technologies, nil
}
