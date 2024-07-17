# CustomFields
(*CustomFields*)

## Overview

View custom contact properties

### Available Operations

* [GetContactsCustomFields](#getcontactscustomfields) - Get a list of custom contact properties

## GetContactsCustomFields

Retrieve a list of your account's custom contact properties.

### Example Usage

```go
package main

import(
	"os"
	loopsgo "github.com/speakeasy-sdks/loops-go"
	"context"
	"log"
)

func main() {
    s := loopsgo.New(
        loopsgo.WithSecurity(os.Getenv("API_KEY")),
    )

    ctx := context.Background()
    res, err := s.CustomFields.GetContactsCustomFields(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomFields != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.GetContactsCustomFieldsResponse](../../models/operations/getcontactscustomfieldsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
