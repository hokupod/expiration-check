package expchk

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
	domain  string
	holders map[string]IHolder
}

type Result struct {
	Domain      string       `json:"domain"`
	Expirations []Expiration `json:"expirations"`
}

type Expiration struct {
	Name           string     `json:"name"`
	ExpirationDate *time.Time `json:"expiration_date"`
	Duration       *int       `json:"duration"`
	Error          error      `json:"error,omitempty"`
}

func New(domain string) *ExpirationChecker {
	return &ExpirationChecker{domain: domain}
}

func (ex *ExpirationChecker) AddHolder(h IHolder) {
	if ex.holders == nil {
		ex.holders = make(map[string]IHolder)
	}
	ex.holders[h.Name()] = h
}

func (ex *ExpirationChecker) Run() *Result {
	var (
		res Result
		err error
	)

	res.Domain = ex.domain
	for n, h := range ex.holders {
		var e Expiration

		e.Name = n
		e.ExpirationDate, err = h.ExpirationDate(res.Domain)
		if err != nil {
			e.Error = err
			continue
		}

		e.Duration, err = calcDuration(e.ExpirationDate)
		if err != nil {
			e.Error = err
		}
		res.Expirations = append(res.Expirations, e)
	}
	return &res
}

func calcDuration(e *time.Time) (*int, error) {
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
