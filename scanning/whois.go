package scanning

/*

References:
https://pkg.go.dev/github.com/likexian/whois
https://pkg.go.dev/github.com/likexian/whois-parser

type WhoisInfo struct {
  Domain         *Domain  `json:"domain,omitempty"`
  Registrar      *Contact `json:"registrar,omitempty"`
  Registrant     *Contact `json:"registrant,omitempty"`
  Administrative *Contact `json:"administrative,omitempty"`
  Technical      *Contact `json:"technical,omitempty"`
  Billing        *Contact `json:"billing,omitempty"`
}

type Domain struct {
  ID                   string     `json:"id,omitempty"`
  Domain               string     `json:"domain,omitempty"`
  Punycode             string     `json:"punycode,omitempty"`
  Name                 string     `json:"name,omitempty"`
  Extension            string     `json:"extension,omitempty"`
  WhoisServer          string     `json:"whois_server,omitempty"`
  Status               []string   `json:"status,omitempty"`
  NameServers          []string   `json:"name_servers,omitempty"`
  DNSSec               bool       `json:"dnssec,omitempty"`
  CreatedDate          string     `json:"created_date,omitempty"`
  CreatedDateInTime    *time.Time `json:"created_date_in_time,omitempty"`
  UpdatedDate          string     `json:"updated_date,omitempty"`
  UpdatedDateInTime    *time.Time `json:"updated_date_in_time,omitempty"`
  ExpirationDate       string     `json:"expiration_date,omitempty"`
  ExpirationDateInTime *time.Time `json:"expiration_date_in_time,omitempty"`
}

type Contact struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Organization string `json:"organization,omitempty"`
	Street       string `json:"street,omitempty"`
	City         string `json:"city,omitempty"`
	Province     string `json:"province,omitempty"`
	PostalCode   string `json:"postal_code,omitempty"`
	Country      string `json:"country,omitempty"`
	Phone        string `json:"phone,omitempty"`
	PhoneExt     string `json:"phone_ext,omitempty"`
	Fax          string `json:"fax,omitempty"`
	FaxExt       string `json:"fax_ext,omitempty"`
	Email        string `json:"email,omitempty"`
	ReferralURL  string `json:"referral_url,omitempty"`
}

*/

import (
	"github.com/likexian/whois"
	wp "github.com/likexian/whois-parser"
)

func WhoisDomain(dom string) (wp.WhoisInfo, error) {
	raw_whois, err := whois.Whois(dom)
	if err != nil {
		return wp.WhoisInfo{}, err
	}

	results, err := wp.Parse(raw_whois)
	if err != nil {
		return wp.WhoisInfo{}, err
	}

	return results, nil
}
