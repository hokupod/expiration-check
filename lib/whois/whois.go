package whois

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	whoisparser "github.com/icamys/whois-parser"
	"github.com/itchyny/timefmt-go"
	"github.com/likexian/whois"
)

func new() *whois.Client {
	c := whois.NewClient()
	c.SetTimeout(10 * time.Second)
	return c
}

func query(domain string) (string, error) {
	c := new()
	if strings.HasSuffix(domain, ".jp") {
		domain = domain + "/e"
	}
	response, err := c.Whois(domain)

	if err != nil {
		return "", err
	}

	return response, nil
}

func parse(domain string, whoisRaw string) (*whoisparser.Record, error) {
	whoisRecord := whoisparser.Parse(domain, whoisRaw)
	if whoisRecord.ErrCode != 0 {
		return whoisRecord, fmt.Errorf("ParseError")
	}

	return whoisRecord, nil
}

func ExpirationDate(domain string, isDuration bool) (string, error) {
	whoisRaw, err := query(domain)
	if err != nil {
		return "", err
	}

	whoisRecord, err := parse(domain, whoisRaw)
	if err != nil {
		return "", err
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

	if isDuration {
		expDate, err := timefmt.ParseInLocation(dateStr, "%Y-%m-%d", time.UTC)
		if err != nil {
			return "", err
		}

		nowDateStr := timefmt.Format(time.Now().In(time.UTC), "%Y-%m-%d")
		nowDate, err := timefmt.ParseInLocation(nowDateStr, "%Y-%m-%d", time.UTC)
		if err != nil {
			return "", err
		}

		duration := expDate.Sub(nowDate)
		durationDay := math.Floor(duration.Hours() / 24.0)
		durationStr := strconv.Itoa(int(durationDay))

		dateStr = durationStr
	}

	return dateStr, nil
}
