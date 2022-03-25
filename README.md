expiration-check
==

Check expiration dates for delivary server.

Usage
--
### Use as cli command.
Install
```sh
go get github.com/hokupod/expiration-check
```

Check the expiration date of the SSL certificate.
```sh
expiration-check ssl example.com  # output => yyyy-mm-dd
expiration-check ssl -d example.com  # output => days
```

Check the expiration date of the domain.
```sh
expiration-check whois example.com  # output => yyyy-mm-dd
expiration-check whois -d example.com  # output => days
```

### Use as package

```go
package main

import (
	"fmt"

	"github.com/hokupod/expiration-check/lib/ssl"
	"github.com/hokupod/expiration-check/lib/whois"
)

func main() {
	// Check the expiration date of the SSL certificate.
	expirationDate, err := ssl.ExpirationDate("example.com", true)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Println(expirationDate) // output => yyyy-mm-dd

	// Output the number of days from today to expire in days.
	expirationDate, err = ssl.ExpirationDate("example.com", false)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Println(expirationDate) // output => days


	// Check the expiration date of the domain.
	expirationDate, err := whois.ExpirationDate("example.com", true)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Println(expirationDate) // output => yyyy-mm-dd

	// Output the number of days from today to expire in days.
	expirationDate, err = whois.ExpirationDate("example.com", false)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Println(expirationDate) // output => days
}
```
