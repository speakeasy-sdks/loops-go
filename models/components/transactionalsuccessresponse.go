// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type TransactionalSuccessResponse struct {
	Success *bool `json:"success,omitempty"`
}

func (o *TransactionalSuccessResponse) GetSuccess() *bool {
	if o == nil {
		return nil
	}
	return o.Success
}
