// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StatusReader is a Reader for the Status structure.
type StatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStatusOK creates a StatusOK with default headers values
func NewStatusOK() *StatusOK {
	return &StatusOK{}
}

/* StatusOK describes a response with status code 200, with default header values.

Cluster health status information.
*/
type StatusOK struct {
	Payload *StatusOKBody
}

func (o *StatusOK) Error() string {
	return fmt.Sprintf("[GET /api/status][%d] statusOK  %+v", 200, o.Payload)
}
func (o *StatusOK) GetPayload() *StatusOKBody {
	return o.Payload
}

func (o *StatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StatusOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StatusOKBody status o k body
swagger:model StatusOKBody
*/
type StatusOKBody struct {

	// job
	Job *StatusOKBodyJob `json:"job,omitempty"`

	// node
	Node *StatusOKBodyNode `json:"node,omitempty"`

	// session
	Session *StatusOKBodySession `json:"session,omitempty"`

	// workflow
	Workflow *StatusOKBodyWorkflow `json:"workflow,omitempty"`
}

// Validate validates this status o k body
func (o *StatusOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateJob(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSession(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWorkflow(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StatusOKBody) validateJob(formats strfmt.Registry) error {
	if swag.IsZero(o.Job) { // not required
		return nil
	}

	if o.Job != nil {
		if err := o.Job.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statusOK" + "." + "job")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statusOK" + "." + "job")
			}
			return err
		}
	}

	return nil
}

func (o *StatusOKBody) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(o.Node) { // not required
		return nil
	}

	if o.Node != nil {
		if err := o.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statusOK" + "." + "node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statusOK" + "." + "node")
			}
			return err
		}
	}

	return nil
}

func (o *StatusOKBody) validateSession(formats strfmt.Registry) error {
	if swag.IsZero(o.Session) { // not required
		return nil
	}

	if o.Session != nil {
		if err := o.Session.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statusOK" + "." + "session")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statusOK" + "." + "session")
			}
			return err
		}
	}

	return nil
}

func (o *StatusOKBody) validateWorkflow(formats strfmt.Registry) error {
	if swag.IsZero(o.Workflow) { // not required
		return nil
	}

	if o.Workflow != nil {
		if err := o.Workflow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statusOK" + "." + "workflow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statusOK" + "." + "workflow")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this status o k body based on the context it is used
func (o *StatusOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateJob(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSession(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateWorkflow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StatusOKBody) contextValidateJob(ctx context.Context, formats strfmt.Registry) error {

	if o.Job != nil {
		if err := o.Job.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statusOK" + "." + "job")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statusOK" + "." + "job")
			}
			return err
		}
	}

	return nil
}

func (o *StatusOKBody) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if o.Node != nil {
		if err := o.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statusOK" + "." + "node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statusOK" + "." + "node")
			}
			return err
		}
	}

	return nil
}

func (o *StatusOKBody) contextValidateSession(ctx context.Context, formats strfmt.Registry) error {

	if o.Session != nil {
		if err := o.Session.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statusOK" + "." + "session")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statusOK" + "." + "session")
			}
			return err
		}
	}

	return nil
}

func (o *StatusOKBody) contextValidateWorkflow(ctx context.Context, formats strfmt.Registry) error {

	if o.Workflow != nil {
		if err := o.Workflow.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statusOK" + "." + "workflow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statusOK" + "." + "workflow")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *StatusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StatusOKBody) UnmarshalBinary(b []byte) error {
	var res StatusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StatusOKBodyJob status o k body job
swagger:model StatusOKBodyJob
*/
type StatusOKBodyJob struct {

	// available
	Available float64 `json:"available,omitempty"`

	// health
	Health string `json:"health,omitempty"`

	// pending
	Pending float64 `json:"pending,omitempty"`

	// percentage
	Percentage float64 `json:"percentage,omitempty"`

	// running
	Running float64 `json:"running,omitempty"`

	// sort
	Sort float64 `json:"sort,omitempty"`

	// total
	Total float64 `json:"total,omitempty"`
}

// Validate validates this status o k body job
func (o *StatusOKBodyJob) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this status o k body job based on context it is used
func (o *StatusOKBodyJob) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StatusOKBodyJob) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StatusOKBodyJob) UnmarshalBinary(b []byte) error {
	var res StatusOKBodyJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StatusOKBodyNode status o k body node
swagger:model StatusOKBodyNode
*/
type StatusOKBodyNode struct {

	// available
	Available float64 `json:"available,omitempty"`

	// health
	Health string `json:"health,omitempty"`

	// percentage
	Percentage float64 `json:"percentage,omitempty"`

	// sort
	Sort float64 `json:"sort,omitempty"`

	// total
	Total float64 `json:"total,omitempty"`

	// unschedulable
	Unschedulable float64 `json:"unschedulable,omitempty"`
}

// Validate validates this status o k body node
func (o *StatusOKBodyNode) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this status o k body node based on context it is used
func (o *StatusOKBodyNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StatusOKBodyNode) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StatusOKBodyNode) UnmarshalBinary(b []byte) error {
	var res StatusOKBodyNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StatusOKBodySession status o k body session
swagger:model StatusOKBodySession
*/
type StatusOKBodySession struct {

	// active
	Active float64 `json:"active,omitempty"`

	// sort
	Sort float64 `json:"sort,omitempty"`
}

// Validate validates this status o k body session
func (o *StatusOKBodySession) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this status o k body session based on context it is used
func (o *StatusOKBodySession) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StatusOKBodySession) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StatusOKBodySession) UnmarshalBinary(b []byte) error {
	var res StatusOKBodySession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StatusOKBodyWorkflow status o k body workflow
swagger:model StatusOKBodyWorkflow
*/
type StatusOKBodyWorkflow struct {

	// available
	Available float64 `json:"available,omitempty"`

	// health
	Health string `json:"health,omitempty"`

	// pending
	Pending float64 `json:"pending,omitempty"`

	// percentage
	Percentage float64 `json:"percentage,omitempty"`

	// queued
	Queued float64 `json:"queued,omitempty"`

	// running
	Running float64 `json:"running,omitempty"`

	// sort
	Sort float64 `json:"sort,omitempty"`

	// total
	Total float64 `json:"total,omitempty"`
}

// Validate validates this status o k body workflow
func (o *StatusOKBodyWorkflow) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this status o k body workflow based on context it is used
func (o *StatusOKBodyWorkflow) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StatusOKBodyWorkflow) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StatusOKBodyWorkflow) UnmarshalBinary(b []byte) error {
	var res StatusOKBodyWorkflow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
