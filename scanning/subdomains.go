package scanning

/*

References:
https://github.com/famasoon/crtsh
https://github.com/projectdiscovery/subfinder

*/

import (
	"encoding/json"
	"fmt"
	"github.com/D3Ext/maldev/slices"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type CTLogs []*CTLog

type CTLog struct { // Custom struct for crt.sh api
	IssuerCaID        int    `json:"issuer_ca_id"`
	IssuerName        string `json:"issuer_name"`
	NameValue         string `json:"name_value"`
	MinCertID         int    `json:"min_cert_id"`
	MinEntryTimestamp string `json:"min_entry_timestamp"`
	NotBefore         string `json:"not_before"`
	NotAfter          string `json:"not_after"`
}

type AlienvaultResponse struct {
	PassiveDNS []struct {
		Hostname string `json:"hostname"`
	} `json:"passive_dns"`
}

// Define API urls
const crtsh_url string = "https://crt.sh/?output=json&CN="
const hackertarget_url string = "https://api.hackertarget.com/hostsearch/?q="
const alienvault_url string = "https://otx.alienvault.com/api/v1/indicators/domain/"

func GetCrt(dom string) ([]string, error) {
	var crt_subdomains []string

	// Use crt.sh API
	req := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := req.Get(crtsh_url + dom)
	if err != nil {
		return []string{""}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = fmt.Errorf("crt.sh is down or seems to be slow")
		return []string{""}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []string{""}, err
	}

	var logs CTLogs
	err = json.Unmarshal(body, &logs)
	if err != nil {
		return []string{""}, err
	}

	for _, c := range logs {
		crt_subdomains = append(crt_subdomains, c.NameValue)
	}

	return crt_subdomains, nil
}

func GetHackerTarget(dom string) ([]string, error) {
	var hackertarget_subdomains []string

	// Use HackerTarget API
	req := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := req.Get(hackertarget_url + dom)
	if err != nil {
		return []string{""}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = fmt.Errorf("HackerTarget is down or seems to be slow")
		return []string{""}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []string{""}, err
	}

	for _, entry := range strings.Split(string(body), "\n") {
		hackertarget_subdomains = append(hackertarget_subdomains, strings.Split(entry, ",")[0])
	}

	return hackertarget_subdomains, nil
}

func GetAlienVault(dom string) ([]string, error) {
	var alienvault_subdomains []string

	// Use HackerTarget API
	req := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := req.Get(alienvault_url + dom + "/passive_dns")
	if err != nil {
		return []string{""}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = fmt.Errorf("AlienVault is down or seems to be slow")
		return []string{""}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []string{""}, err
	}

	var alien AlienvaultResponse
	err = json.Unmarshal(body, &alien)
	if err != nil {
		return []string{""}, err
	}

	for _, entry := range alien.PassiveDNS {
		alienvault_subdomains = append(alienvault_subdomains, entry.Hostname)
	}

	return alienvault_subdomains, nil
}

func GetAllSubdomains(dom string) ([]string, error) {
	var final_subdomains []string

	// Get all subdomains
	subdomains1, err := GetCrt(dom)
	if err != nil {
		return []string{""}, err
	}

	subdomains2, err := GetHackerTarget(dom)
	if err != nil {
		return []string{""}, err
	}

	subdomains3, err := GetAlienVault(dom)
	if err != nil {
		return []string{""}, err
	}

	for _, entry := range subdomains1 { // Add all subdomains from Crt.sh
		final_subdomains = append(final_subdomains, entry)
	}

	for _, entry := range subdomains2 { // Add all subdomains from HackerTarget
		final_subdomains = append(final_subdomains, entry)
	}

	for _, entry := range subdomains3 { // Add subdomains from AlienVault
		final_subdomains = append(final_subdomains, entry)
	}

	// Finally remove duplicates and return subdomains
	final_subdomains = slices.RemoveDuplicatesStr(final_subdomains)
	return final_subdomains, nil
}
