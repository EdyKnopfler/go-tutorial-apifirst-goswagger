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

// Measurement measurement
//
// swagger:model measurement
type Measurement struct {

	// sensor
	// Required: true
	// Min Length: 8
	Sensor *string `json:"sensor"`

	// time taken
	// Required: true
	// Format: date-time
	TimeTaken *strfmt.DateTime `json:"timeTaken"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this measurement
func (m *Measurement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSensor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeTaken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Measurement) validateSensor(formats strfmt.Registry) error {

	if err := validate.Required("sensor", "body", m.Sensor); err != nil {
		return err
	}

	if err := validate.MinLength("sensor", "body", *m.Sensor, 8); err != nil {
		return err
	}

	return nil
}

func (m *Measurement) validateTimeTaken(formats strfmt.Registry) error {

	if err := validate.Required("timeTaken", "body", m.TimeTaken); err != nil {
		return err
	}

	if err := validate.FormatOf("timeTaken", "body", "date-time", m.TimeTaken.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Measurement) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this measurement based on context it is used
func (m *Measurement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Measurement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Measurement) UnmarshalBinary(b []byte) error {
	var res Measurement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
