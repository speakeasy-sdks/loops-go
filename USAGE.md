<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	loopsgo "github.com/speakeasy-sdks/loops-go"
	"log"
)

func main() {
	s := loopsgo.New(
		loopsgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.APIKey.GetAPIKey(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->