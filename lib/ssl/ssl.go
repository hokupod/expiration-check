package ssl

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"math"
	"strconv"
	"time"

	"github.com/itchyny/timefmt-go"
)

func query(domain string) (*x509.Certificate, error) {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	d := tls.Dialer{
		Config: conf,
	}
	conn, err := d.DialContext(ctx, "tcp", domain+":443")
	cancel()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	tlsConn := conn.(*tls.Conn)
	cert := tlsConn.ConnectionState().PeerCertificates[0]

	return cert, err
}

func ExpirationDate(domain string, isDuration bool) (string, error) {
	cert, err := query(domain)
	if err != nil {
		return "", err
	}

	expDate := cert.NotAfter
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
