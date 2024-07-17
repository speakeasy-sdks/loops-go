# Events
(*Events*)

## Overview

Trigger email sending with events

### Available Operations

* [PostEventsSend](#posteventssend) - Send an event

## PostEventsSend

Send events to trigger emails in Loops.

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
    request := components.EventRequest{
        EventName: "<value>",
        MailingLists: &components.EventRequestMailingLists{},
    }
    ctx := context.Background()
    res, err := s.Events.PostEventsSend(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.EventSuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [components.EventRequest](../../models/components/eventrequest.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |


### Response

**[*operations.PostEventsSendResponse](../../models/operations/posteventssendresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.EventFailureResponse | 400                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |
