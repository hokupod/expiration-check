package ssl

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"time"
)

type IHolder interface {
	Name() string
	ExpirationDate() (*time.Time, error)
}

type Holder struct {
	expirationDate *time.Time
}

func (holder *Holder) query(domain string) (*x509.Certificate, error) {
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

func (holder *Holder) Name() string {
	return "ssl"
}

func (holder *Holder) ExpirationDate(domain string) (*time.Time, error) {
	if holder.expirationDate != nil {
		return holder.expirationDate, nil
	}

	cert, err := holder.query(domain)
	if err != nil {
		return nil, err
	}

	holder.expirationDate = &cert.NotAfter
	return holder.expirationDate, err
}
