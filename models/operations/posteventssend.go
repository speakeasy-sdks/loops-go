// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/loops-go/models/components"
)

type PostEventsSendResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful send.
	EventSuccessResponse *components.EventSuccessResponse
}

func (o *PostEventsSendResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostEventsSendResponse) GetEventSuccessResponse() *components.EventSuccessResponse {
	if o == nil {
		return nil
	}
	return o.EventSuccessResponse
}
