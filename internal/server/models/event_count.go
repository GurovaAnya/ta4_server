// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EventCount event count
//
// swagger:model EventCount
type EventCount struct {

	// count
	Count int64 `json:"count,omitempty"`

	// event id
	// Example: f74687fa-2df7-450e-8c31-993695dcebf7
	// Format: uuid
	EventID strfmt.UUID `json:"event_id,omitempty"`
}

// Validate validates this event count
func (m *EventCount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventCount) validateEventID(formats strfmt.Registry) error {
	if swag.IsZero(m.EventID) { // not required
		return nil
	}

	if err := validate.FormatOf("event_id", "body", "uuid", m.EventID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this event count based on context it is used
func (m *EventCount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventCount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventCount) UnmarshalBinary(b []byte) error {
	var res EventCount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
