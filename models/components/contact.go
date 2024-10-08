// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Contact struct {
	ID         *string `json:"id,omitempty"`
	Email      *string `json:"email,omitempty"`
	FirstName  *string `json:"firstName,omitempty"`
	LastName   *string `json:"lastName,omitempty"`
	Source     *string `json:"source,omitempty"`
	Subscribed *bool   `json:"subscribed,omitempty"`
	UserGroup  *string `json:"userGroup,omitempty"`
	UserID     *string `json:"userId,omitempty"`
}

func (o *Contact) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Contact) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *Contact) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *Contact) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *Contact) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *Contact) GetSubscribed() *bool {
	if o == nil {
		return nil
	}
	return o.Subscribed
}

func (o *Contact) GetUserGroup() *string {
	if o == nil {
		return nil
	}
	return o.UserGroup
}

func (o *Contact) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}
