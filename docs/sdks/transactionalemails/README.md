# TransactionalEmails
(*TransactionalEmails*)

## Overview

Send transactional emails

### Available Operations

* [PostTransactional](#posttransactional) - Send a transactional email

## PostTransactional

Send a transactional email to a contact.

### Example Usage

```go
package main

import(
	"os"
	loopsgo "github.com/speakeasy-sdks/loops-go"
	"github.com/speakeasy-sdks/loops-go/models/components"
	"context"
	"log"
)

func main() {
    s := loopsgo.New(
        loopsgo.WithSecurity(os.Getenv("API_KEY")),
    )
    request := components.TransactionalRequest{
        Email: "Rollin.Littel98@gmail.com",
        TransactionalID: "<value>",
    }
    ctx := context.Background()
    res, err := s.TransactionalEmails.PostTransactional(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.TransactionalSuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [components.TransactionalRequest](../../models/components/transactionalrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.PostTransactionalResponse](../../models/operations/posttransactionalresponse.md), error**
| Error Object                            | Status Code                             | Content Type                            |
| --------------------------------------- | --------------------------------------- | --------------------------------------- |
| sdkerrors.PostTransactionalResponseBody | 400                                     | application/json                        |
| sdkerrors.TransactionalFailure2Response | 404                                     | application/json                        |
| sdkerrors.SDKError                      | 4xx-5xx                                 | */*                                     |
