# Contacts
(*Contacts*)

## Overview

Manage contacts in your audience

### Available Operations

* [PostContactsCreate](#postcontactscreate) - Create a contact
* [PutContactsUpdate](#putcontactsupdate) - Update a contact
* [GetContactsFind](#getcontactsfind) - Find a contact
* [PostContactsDelete](#postcontactsdelete) - Delete a contact

## PostContactsCreate

Add a contact to your audience.

### Example Usage

```go
package main

import(
	loopsgo "github.com/speakeasy-sdks/loops-go"
	"github.com/speakeasy-sdks/loops-go/models/components"
	"context"
	"log"
)

func main() {
    s := loopsgo.New(
        loopsgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    request := components.ContactRequest{
        Email: "Ashtyn_Beer@gmail.com",
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

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [components.ContactRequest](../../models/components/contactrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |


### Response

**[*operations.PostContactsCreateResponse](../../models/operations/postcontactscreateresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.ContactFailureResponse | 400,405,409                      | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

## PutContactsUpdate

Update a contact by `email` or `userId`.<br>If you want to update a contact’s email address, the contact will first need a `userId` value. You can then make a request containing the userId field along with an updated email address.

### Example Usage

```go
package main

import(
	loopsgo "github.com/speakeasy-sdks/loops-go"
	"github.com/speakeasy-sdks/loops-go/models/components"
	"context"
	"log"
)

func main() {
    s := loopsgo.New(
        loopsgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    request := components.ContactRequest{
        Email: "Giovanna61@yahoo.com",
        MailingLists: &components.MailingLists{},
    }
    ctx := context.Background()
    res, err := s.Contacts.PutContactsUpdate(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.ContactSuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [components.ContactRequest](../../models/components/contactrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |


### Response

**[*operations.PutContactsUpdateResponse](../../models/operations/putcontactsupdateresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.ContactFailureResponse | 400,405                          | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

## GetContactsFind

Search for a contact by `email` or `userId`. Only one parameter is allowed.

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
    res, err := s.Contacts.GetContactsFind(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Contacts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `email`                                                  | **string*                                                | :heavy_minus_sign:                                       | Email address (URI-encoded)                              |
| `userID`                                                 | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.GetContactsFindResponse](../../models/operations/getcontactsfindresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.ContactFailureResponse | 400                              | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

## PostContactsDelete

Delete a contact by `email` or `userId`.

### Example Usage

```go
package main

import(
	loopsgo "github.com/speakeasy-sdks/loops-go"
	"github.com/speakeasy-sdks/loops-go/models/components"
	"context"
	"log"
)

func main() {
    s := loopsgo.New(
        loopsgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    request := components.ContactDeleteRequest{
        Email: "Rebekah34@hotmail.com",
        UserID: "<value>",
    }
    ctx := context.Background()
    res, err := s.Contacts.PostContactsDelete(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.ContactDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [components.ContactDeleteRequest](../../models/components/contactdeleterequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.PostContactsDeleteResponse](../../models/operations/postcontactsdeleteresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.ContactFailureResponse | 400,404                          | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |
