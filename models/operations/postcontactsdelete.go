// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/loops-go/models/components"
)

type PostContactsDeleteResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful delete.
	ContactDeleteResponse *components.ContactDeleteResponse
}

func (o *PostContactsDeleteResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostContactsDeleteResponse) GetContactDeleteResponse() *components.ContactDeleteResponse {
	if o == nil {
		return nil
	}
	return o.ContactDeleteResponse
}
