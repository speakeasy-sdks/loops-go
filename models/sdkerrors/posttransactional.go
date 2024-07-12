// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/loops-go/internal/utils"
	"github.com/speakeasy-sdks/loops-go/models/components"
)

type PostTransactionalResponseBodyType string

const (
	PostTransactionalResponseBodyTypeTransactionalFailureResponse  PostTransactionalResponseBodyType = "TransactionalFailureResponse"
	PostTransactionalResponseBodyTypeTransactionalFailure2Response PostTransactionalResponseBodyType = "TransactionalFailure2Response"
	PostTransactionalResponseBodyTypeTransactionalFailure3Response PostTransactionalResponseBodyType = "TransactionalFailure3Response"
)

// PostTransactionalResponseBody - Bad request (e.g. transactional email is not published).
type PostTransactionalResponseBody struct {
	TransactionalFailureResponse  *components.TransactionalFailureResponse
	TransactionalFailure2Response *TransactionalFailure2Response
	TransactionalFailure3Response *components.TransactionalFailure3Response

	Type PostTransactionalResponseBodyType

	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &PostTransactionalResponseBody{}

func CreatePostTransactionalResponseBodyTransactionalFailureResponse(transactionalFailureResponse components.TransactionalFailureResponse) PostTransactionalResponseBody {
	typ := PostTransactionalResponseBodyTypeTransactionalFailureResponse

	return PostTransactionalResponseBody{
		TransactionalFailureResponse: &transactionalFailureResponse,
		Type:                         typ,
	}
}

func CreatePostTransactionalResponseBodyTransactionalFailure2Response(transactionalFailure2Response TransactionalFailure2Response) PostTransactionalResponseBody {
	typ := PostTransactionalResponseBodyTypeTransactionalFailure2Response

	return PostTransactionalResponseBody{
		TransactionalFailure2Response: &transactionalFailure2Response,
		Type:                          typ,
	}
}

func CreatePostTransactionalResponseBodyTransactionalFailure3Response(transactionalFailure3Response components.TransactionalFailure3Response) PostTransactionalResponseBody {
	typ := PostTransactionalResponseBodyTypeTransactionalFailure3Response

	return PostTransactionalResponseBody{
		TransactionalFailure3Response: &transactionalFailure3Response,
		Type:                          typ,
	}
}

func (u *PostTransactionalResponseBody) UnmarshalJSON(data []byte) error {

	var transactionalFailure3Response components.TransactionalFailure3Response = components.TransactionalFailure3Response{}
	if err := utils.UnmarshalJSON(data, &transactionalFailure3Response, "", true, true); err == nil {
		u.TransactionalFailure3Response = &transactionalFailure3Response
		u.Type = PostTransactionalResponseBodyTypeTransactionalFailure3Response
		return nil
	}

	var transactionalFailureResponse components.TransactionalFailureResponse = components.TransactionalFailureResponse{}
	if err := utils.UnmarshalJSON(data, &transactionalFailureResponse, "", true, true); err == nil {
		u.TransactionalFailureResponse = &transactionalFailureResponse
		u.Type = PostTransactionalResponseBodyTypeTransactionalFailureResponse
		return nil
	}

	var transactionalFailure2Response TransactionalFailure2Response = TransactionalFailure2Response{}
	if err := utils.UnmarshalJSON(data, &transactionalFailure2Response, "", true, true); err == nil {
		u.TransactionalFailure2Response = &transactionalFailure2Response
		u.Type = PostTransactionalResponseBodyTypeTransactionalFailure2Response
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for PostTransactionalResponseBody", string(data))
}

func (u PostTransactionalResponseBody) MarshalJSON() ([]byte, error) {
	if u.TransactionalFailureResponse != nil {
		return utils.MarshalJSON(u.TransactionalFailureResponse, "", true)
	}

	if u.TransactionalFailure2Response != nil {
		return utils.MarshalJSON(u.TransactionalFailure2Response, "", true)
	}

	if u.TransactionalFailure3Response != nil {
		return utils.MarshalJSON(u.TransactionalFailure3Response, "", true)
	}

	return nil, errors.New("could not marshal union type PostTransactionalResponseBody: all fields are null")
}

func (u PostTransactionalResponseBody) Error() string {
	switch u.Type {
	case PostTransactionalResponseBodyTypeTransactionalFailureResponse:
		data, _ := json.Marshal(u.TransactionalFailureResponse)
		return string(data)
	case PostTransactionalResponseBodyTypeTransactionalFailure2Response:
		data, _ := json.Marshal(u.TransactionalFailure2Response)
		return string(data)
	case PostTransactionalResponseBodyTypeTransactionalFailure3Response:
		data, _ := json.Marshal(u.TransactionalFailure3Response)
		return string(data)
	default:
		return "unknown error"
	}
}
