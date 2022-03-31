module github.com/hokupod/expiration-check

go 1.16

require (
	github.com/icamys/whois-parser v0.0.0-20191128111838-6ee28f3c032f
	github.com/itchyny/timefmt-go v0.1.3
	github.com/jpillora/go-tld v1.1.1
	github.com/likexian/whois v1.12.4
	github.com/spf13/cobra v1.3.0
)

replace github.com/icamys/whois-parser => ../whois-parser
