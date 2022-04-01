package holder

import (
	"math"
	"time"

	"github.com/itchyny/timefmt-go"
)

type IHolder interface {
	Name() string
	ExpirationDate(domain string) (*time.Time, error)
}

type ExpirationChecker struct {
	Domain      string
	Expirations map[string]IHolder
}

type Result struct {
	Domain      string       `json:"domain"`
	Expirations []Expiration `json:"expirations"`
}

type Expiration struct {
	Name           string     `json:"name"`
	ExpirationDate *time.Time `json:"expiration_date"`
	Duration       *int       `json:"duration"`
	Errors         []error    `json:"errors"`
}

func ExpirationCheckerNew(domain string) *ExpirationChecker {
	return &ExpirationChecker{Domain: domain}
}

func (ex *ExpirationChecker) AddHolder(h IHolder) {
	if ex.Expirations == nil {
		ex.Expirations = make(map[string]IHolder)
	}
	ex.Expirations[h.Name()] = h
}

func (ex *ExpirationChecker) RunAll() *Result {
	var (
		res Result
		err error
	)

	res.Domain = ex.Domain
	for k, v := range ex.Expirations {
		var e Expiration

		e.Name = k
		e.ExpirationDate, err = v.ExpirationDate(res.Domain)
		if err != nil {
			e.Errors = append(e.Errors, err)
		}

		e.Duration, err = CalcDuration(e.ExpirationDate)
		if err != nil {
			e.Errors = append(e.Errors, err)
		}
		res.Expirations = append(res.Expirations, e)
	}
	return &res
}

func CalcDuration(e *time.Time) (*int, error) {
	nowDateStr := timefmt.Format(time.Now().In(time.UTC), "%Y-%m-%d")
	nowDate, err := timefmt.ParseInLocation(nowDateStr, "%Y-%m-%d", time.UTC)
	if err != nil {
		return nil, err
	}

	duration := e.Sub(nowDate)
	durationDay := math.Floor(duration.Hours() / 24.0)
	d := new(int)
	*d = int(durationDay)

	return d, nil
}
