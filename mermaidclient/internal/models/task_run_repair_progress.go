// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TaskRunRepairProgress task run repair progress
// swagger:model TaskRunRepairProgress
type TaskRunRepairProgress struct {

	// progress
	Progress *RepairProgress `json:"progress,omitempty"`

	// run
	Run *TaskRun `json:"run,omitempty"`
}

// Validate validates this task run repair progress
func (m *TaskRunRepairProgress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProgress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRun(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskRunRepairProgress) validateProgress(formats strfmt.Registry) error {

	if swag.IsZero(m.Progress) { // not required
		return nil
	}

	if m.Progress != nil {
		if err := m.Progress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("progress")
			}
			return err
		}
	}

	return nil
}

func (m *TaskRunRepairProgress) validateRun(formats strfmt.Registry) error {

	if swag.IsZero(m.Run) { // not required
		return nil
	}

	if m.Run != nil {
		if err := m.Run.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("run")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskRunRepairProgress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskRunRepairProgress) UnmarshalBinary(b []byte) error {
	var res TaskRunRepairProgress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
