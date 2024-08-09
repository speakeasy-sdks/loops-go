<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	loopsgo "github.com/speakeasy-sdks/loops-go"
	"github.com/speakeasy-sdks/loops-go/models/components"
	"log"
)

func main() {
	s := loopsgo.New(
		loopsgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)
	request := components.ContactRequest{
		Email:        "Ashtyn_Beer@gmail.com",
		MailingLists: &components.MailingLists{},
	}
	ctx := context.Background()
	res, err := s.Contacts.PostContactsCreate(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.ContactSuccessResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->