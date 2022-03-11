package ssl

import (
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/itchyny/timefmt-go"
)

func query(domain string) (*http.Response, error) {
	res, err := http.Get("https://" + domain)
	return res, err
}

func ExpirationDate(domain string, isDuration bool) (string, error) {
	res, err := query(domain)
	if err != nil {
		return "", err
	}

	expDate := res.TLS.PeerCertificates[0].NotAfter
	expDateStr := timefmt.Format(expDate, "%Y-%m-%d")

	if isDuration {
		nowDateStr := timefmt.Format(time.Now().In(time.UTC), "%Y-%m-%d")
		nowDate, err := timefmt.ParseInLocation(nowDateStr, "%Y-%m-%d", time.UTC)
		if err != nil {
			return "", err
		}

		duration := expDate.Sub(nowDate)
		durationDay := math.Floor(duration.Hours() / 24.0)
		durationStr := strconv.Itoa(int(durationDay))

		expDateStr = durationStr
	}

	return expDateStr, err
}
