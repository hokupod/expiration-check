package whois

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	whoisparser "github.com/icamys/whois-parser"
	"github.com/itchyny/timefmt-go"
	tld "github.com/jpillora/go-tld"
	"github.com/likexian/whois"
)

type IHolder interface {
	Name() string
	ExpirationDate() (*time.Time, error)
}

type Holder struct {
	expirationDate *time.Time
}

func (holder *Holder) getClient() *whois.Client {
	c := whois.NewClient()
	c.SetTimeout(10 * time.Second)
	return c
}

func (holder *Holder) query(domain string) (string, error) {
	c := holder.getClient()
	if strings.HasSuffix(domain, ".jp") {
		domain = domain + "/e"
	}
	response, err := c.Whois(domain)

	if err != nil {
		return "", err
	}

	return response, nil
}

func (holder *Holder) parse(domain string, whoisRaw string) (*whoisparser.Record, error) {
	whoisRecord := whoisparser.Parse(domain, whoisRaw)
	if whoisRecord.ErrCode != 0 {
		return whoisRecord, fmt.Errorf("ParseError")
	}

	return whoisRecord, nil
}

func (holder *Holder) Name() string {
	return "whois"
}

func (holder *Holder) ExpirationDate(domain string) (*time.Time, error) {
	if holder.expirationDate != nil {
		return holder.expirationDate, nil
	}

	prepared, err := prepareDomain(domain)
	if err != nil {
		return nil, err
	}

	whoisRaw, err := holder.query(prepared)
	if err != nil {
		return nil, err
	}

	whoisRecord, err := holder.parse(domain, whoisRaw)
	if err != nil {
		return nil, err
	}

	reDate := regexp.MustCompile(`.+?\((.+?)\)`)
	_date := whoisRecord.Registrar.ExpirationDate
	if strings.Contains(_date, "(") {
		_date = reDate.ReplaceAllString(_date, "$1")
	}

	if strings.Contains(_date, "/") {
		_date = strings.ReplaceAll(_date, "/", "-")
	}

	dateStr := _date[:10]
	expDate, err := timefmt.ParseInLocation(dateStr, "%Y-%m-%d", time.UTC)
	if err != nil {
		return nil, err
	}

	holder.expirationDate = &expDate

	return holder.expirationDate, nil
}

func prepareDomain(domain string) (string, error) {
	url, err := tld.Parse("https://" + domain)
	if err != nil {
		return "", err
	}

	return url.Domain + "." + url.TLD, nil
}
