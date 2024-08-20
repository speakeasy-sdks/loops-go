# APIKey
(*APIKey*)

### Available Operations

* [GetAPIKey](#getapikey) - Test your API key

## GetAPIKey

Test your API key

### Example Usage

```go
package main

import(
	loopsgo "github.com/speakeasy-sdks/loops-go"
	"context"
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



### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.GetAPIKeyResponse](../../models/operations/getapikeyresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.GetAPIKeyResponseBody | 401                             | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |
