// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

// DataVariables - An object containing contact data as defined by the data variables added to the transactional email template.
type DataVariables struct {
}

type TransactionalRequest struct {
	Email string `json:"email"`
	// The ID of the transactional email to send.
	TransactionalID string `json:"transactionalId"`
	// If `true`, a contact will be created in your audience using the `email` value (if a matching contact doesn't already exist).
	AddToAudience *bool `json:"addToAudience,omitempty"`
	// An object containing contact data as defined by the data variables added to the transactional email template.
	DataVariables *DataVariables `json:"dataVariables,omitempty"`
}

func (o *TransactionalRequest) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *TransactionalRequest) GetTransactionalID() string {
	if o == nil {
		return ""
	}
	return o.TransactionalID
}

func (o *TransactionalRequest) GetAddToAudience() *bool {
	if o == nil {
		return nil
	}
	return o.AddToAudience
}

func (o *TransactionalRequest) GetDataVariables() *DataVariables {
	if o == nil {
		return nil
	}
	return o.DataVariables
}
