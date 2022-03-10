package whois

import (
	"time"

	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

func new() *whois.Client {
	c := whois.NewClient()
	c.SetTimeout(5 * time.Second)
	return c
}

func query(domain string) (whoisparser.WhoisInfo, error) {
	var whoisInfo whoisparser.WhoisInfo

	c := new()
	response, err := c.Whois(domain)
	if err != nil {
		return whoisInfo, err
	}

	whoisInfo, err = whoisparser.Parse(response)
	if err != nil {
		return whoisInfo, err
	}

	return whoisInfo, nil
}

func ExpirationDate(domain string) (string, error) {
	whoisInfo, err := query(domain)
	if err != nil {
		return "", err
	}

	return whoisInfo.Domain.ExpirationDate, nil
}
