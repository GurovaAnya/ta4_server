// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebhookTrigger webhook trigger
//
// swagger:model WebhookTrigger
type WebhookTrigger struct {

	// event id
	// Example: 5e6890cc-c3a1-4c26-b269-b55e715c09b8
	EventID string `json:"event_id,omitempty"`

	// player id
	// Example: 5e6890cc-c3a1-4c26-b269-b55e715c09b8
	PlayerID string `json:"player_id,omitempty"`
}

// Validate validates this webhook trigger
func (m *WebhookTrigger) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this webhook trigger based on context it is used
func (m *WebhookTrigger) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebhookTrigger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebhookTrigger) UnmarshalBinary(b []byte) error {
	var res WebhookTrigger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}