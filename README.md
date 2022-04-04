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
expiration-check domain example.com  # output => yyyy-mm-dd
expiration-check domain -d example.com  # output => days
```

### Use as package
Import packages

```go
"github.com/hokupod/expiration-check/expchk"
"github.com/hokupod/expiration-check/expchk/domain"
"github.com/hokupod/expiration-check/expchk/ssl"
```

Check the expiration date of supported type.
```go
// If you want to check the domain, use domain.Holder
var sh ssl.Holder

ec := expchk.New("example.com")
ec.AddHolder(sh)
res := ec.Run()
err := res.Expirations[0].Error
if err != nil {
	fmt.Printf("Error: %v: %v\n", res.Expirations[0].Name, err)
}

// Output *time.Time
fmt.Println(res.Expirations[0].ExpirationDate)
// Output the number of days from today to expire in days.
fmt.Println(*res.Expirations[0].Duration)
```

Check the expiration date of all supported types.
```go
var (
	sh ssl.Holder
	dh domain.Holder
)

ec := expchk.New("example.com")
ec.AddHolder(sh)
ec.AddHolder(dh)
res := ec.Run()
for _, ex := range res.Expirations {
	if ex.Error != nil {
		fmt.Printf("Error: %v: %v\n", ex.Name, ex.Error)
	}
}

// Support json output.
jsonStr, err := json.Marshal(res)
if err != nil {
	fmt.Printf("Error: %v", err)
}

var buf bytes.Buffer
err = json.Indent(&buf, []byte(jsonStr), "", "  ")
if err != nil {
	fmt.Printf("Error: %v", err)
}
fmt.Println(buf.String())
```
