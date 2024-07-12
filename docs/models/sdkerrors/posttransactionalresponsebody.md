# PostTransactionalResponseBody

Bad request (e.g. transactional email is not published).


## Supported Types

### TransactionalFailureResponse

```go
postTransactionalResponseBody := sdkerrors.CreatePostTransactionalResponseBodyTransactionalFailureResponse(components.TransactionalFailureResponse{/* values here */})
```

### TransactionalFailure2Response

```go
postTransactionalResponseBody := sdkerrors.CreatePostTransactionalResponseBodyTransactionalFailure2Response(sdkerrors.TransactionalFailure2Response{/* values here */})
```

### TransactionalFailure3Response

```go
postTransactionalResponseBody := sdkerrors.CreatePostTransactionalResponseBodyTransactionalFailure3Response(components.TransactionalFailure3Response{/* values here */})
```

