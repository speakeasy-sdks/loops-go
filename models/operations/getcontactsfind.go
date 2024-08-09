// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/loops-go/models/components"
)

type GetContactsFindRequest struct {
	// Email address (URI-encoded)
	Email  *string `queryParam:"style=form,explode=true,name=email"`
	UserID *string `queryParam:"style=form,explode=true,name=userId"`
}

func (o *GetContactsFindRequest) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *GetContactsFindRequest) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

type GetContactsFindResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List of contacts (or an empty array if no contact was found). Contact objects will include any custom properties.
	Contacts []components.Contact
}

func (o *GetContactsFindResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetContactsFindResponse) GetContacts() []components.Contact {
	if o == nil {
		return nil
	}
	return o.Contacts
}
