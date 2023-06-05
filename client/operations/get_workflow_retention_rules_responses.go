// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetWorkflowRetentionRulesReader is a Reader for the GetWorkflowRetentionRules structure.
type GetWorkflowRetentionRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowRetentionRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowRetentionRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetWorkflowRetentionRulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkflowRetentionRulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkflowRetentionRulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkflowRetentionRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/workflows/{workflow_id_or_name}/retention_rules] get_workflow_retention_rules", response, response.Code())
	}
}

// NewGetWorkflowRetentionRulesOK creates a GetWorkflowRetentionRulesOK with default headers values
func NewGetWorkflowRetentionRulesOK() *GetWorkflowRetentionRulesOK {
	return &GetWorkflowRetentionRulesOK{}
}

/*
GetWorkflowRetentionRulesOK describes a response with status code 200, with default header values.

Request succeeded. The response contains the list of all the retention rules.
*/
type GetWorkflowRetentionRulesOK struct {
	Payload *GetWorkflowRetentionRulesOKBody
}

// IsSuccess returns true when this get workflow retention rules o k response has a 2xx status code
func (o *GetWorkflowRetentionRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workflow retention rules o k response has a 3xx status code
func (o *GetWorkflowRetentionRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow retention rules o k response has a 4xx status code
func (o *GetWorkflowRetentionRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workflow retention rules o k response has a 5xx status code
func (o *GetWorkflowRetentionRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow retention rules o k response a status code equal to that given
func (o *GetWorkflowRetentionRulesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get workflow retention rules o k response
func (o *GetWorkflowRetentionRulesOK) Code() int {
	return 200
}

func (o *GetWorkflowRetentionRulesOK) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/retention_rules][%d] getWorkflowRetentionRulesOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowRetentionRulesOK) String() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/retention_rules][%d] getWorkflowRetentionRulesOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowRetentionRulesOK) GetPayload() *GetWorkflowRetentionRulesOKBody {
	return o.Payload
}

func (o *GetWorkflowRetentionRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowRetentionRulesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowRetentionRulesUnauthorized creates a GetWorkflowRetentionRulesUnauthorized with default headers values
func NewGetWorkflowRetentionRulesUnauthorized() *GetWorkflowRetentionRulesUnauthorized {
	return &GetWorkflowRetentionRulesUnauthorized{}
}

/*
GetWorkflowRetentionRulesUnauthorized describes a response with status code 401, with default header values.

Request failed. User not signed in.
*/
type GetWorkflowRetentionRulesUnauthorized struct {
	Payload *GetWorkflowRetentionRulesUnauthorizedBody
}

// IsSuccess returns true when this get workflow retention rules unauthorized response has a 2xx status code
func (o *GetWorkflowRetentionRulesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow retention rules unauthorized response has a 3xx status code
func (o *GetWorkflowRetentionRulesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow retention rules unauthorized response has a 4xx status code
func (o *GetWorkflowRetentionRulesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workflow retention rules unauthorized response has a 5xx status code
func (o *GetWorkflowRetentionRulesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow retention rules unauthorized response a status code equal to that given
func (o *GetWorkflowRetentionRulesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get workflow retention rules unauthorized response
func (o *GetWorkflowRetentionRulesUnauthorized) Code() int {
	return 401
}

func (o *GetWorkflowRetentionRulesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/retention_rules][%d] getWorkflowRetentionRulesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkflowRetentionRulesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/retention_rules][%d] getWorkflowRetentionRulesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkflowRetentionRulesUnauthorized) GetPayload() *GetWorkflowRetentionRulesUnauthorizedBody {
	return o.Payload
}

func (o *GetWorkflowRetentionRulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowRetentionRulesUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowRetentionRulesForbidden creates a GetWorkflowRetentionRulesForbidden with default headers values
func NewGetWorkflowRetentionRulesForbidden() *GetWorkflowRetentionRulesForbidden {
	return &GetWorkflowRetentionRulesForbidden{}
}

/*
GetWorkflowRetentionRulesForbidden describes a response with status code 403, with default header values.

Request failed. Credentials are invalid or revoked.
*/
type GetWorkflowRetentionRulesForbidden struct {
	Payload *GetWorkflowRetentionRulesForbiddenBody
}

// IsSuccess returns true when this get workflow retention rules forbidden response has a 2xx status code
func (o *GetWorkflowRetentionRulesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow retention rules forbidden response has a 3xx status code
func (o *GetWorkflowRetentionRulesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow retention rules forbidden response has a 4xx status code
func (o *GetWorkflowRetentionRulesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workflow retention rules forbidden response has a 5xx status code
func (o *GetWorkflowRetentionRulesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow retention rules forbidden response a status code equal to that given
func (o *GetWorkflowRetentionRulesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get workflow retention rules forbidden response
func (o *GetWorkflowRetentionRulesForbidden) Code() int {
	return 403
}

func (o *GetWorkflowRetentionRulesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/retention_rules][%d] getWorkflowRetentionRulesForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkflowRetentionRulesForbidden) String() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/retention_rules][%d] getWorkflowRetentionRulesForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkflowRetentionRulesForbidden) GetPayload() *GetWorkflowRetentionRulesForbiddenBody {
	return o.Payload
}

func (o *GetWorkflowRetentionRulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowRetentionRulesForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowRetentionRulesNotFound creates a GetWorkflowRetentionRulesNotFound with default headers values
func NewGetWorkflowRetentionRulesNotFound() *GetWorkflowRetentionRulesNotFound {
	return &GetWorkflowRetentionRulesNotFound{}
}

/*
GetWorkflowRetentionRulesNotFound describes a response with status code 404, with default header values.

Request failed. Workflow does not exist.
*/
type GetWorkflowRetentionRulesNotFound struct {
	Payload *GetWorkflowRetentionRulesNotFoundBody
}

// IsSuccess returns true when this get workflow retention rules not found response has a 2xx status code
func (o *GetWorkflowRetentionRulesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow retention rules not found response has a 3xx status code
func (o *GetWorkflowRetentionRulesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow retention rules not found response has a 4xx status code
func (o *GetWorkflowRetentionRulesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workflow retention rules not found response has a 5xx status code
func (o *GetWorkflowRetentionRulesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow retention rules not found response a status code equal to that given
func (o *GetWorkflowRetentionRulesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get workflow retention rules not found response
func (o *GetWorkflowRetentionRulesNotFound) Code() int {
	return 404
}

func (o *GetWorkflowRetentionRulesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/retention_rules][%d] getWorkflowRetentionRulesNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkflowRetentionRulesNotFound) String() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/retention_rules][%d] getWorkflowRetentionRulesNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkflowRetentionRulesNotFound) GetPayload() *GetWorkflowRetentionRulesNotFoundBody {
	return o.Payload
}

func (o *GetWorkflowRetentionRulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowRetentionRulesNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowRetentionRulesInternalServerError creates a GetWorkflowRetentionRulesInternalServerError with default headers values
func NewGetWorkflowRetentionRulesInternalServerError() *GetWorkflowRetentionRulesInternalServerError {
	return &GetWorkflowRetentionRulesInternalServerError{}
}

/*
GetWorkflowRetentionRulesInternalServerError describes a response with status code 500, with default header values.

Request failed. Internal server error.
*/
type GetWorkflowRetentionRulesInternalServerError struct {
	Payload *GetWorkflowRetentionRulesInternalServerErrorBody
}

// IsSuccess returns true when this get workflow retention rules internal server error response has a 2xx status code
func (o *GetWorkflowRetentionRulesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow retention rules internal server error response has a 3xx status code
func (o *GetWorkflowRetentionRulesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow retention rules internal server error response has a 4xx status code
func (o *GetWorkflowRetentionRulesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workflow retention rules internal server error response has a 5xx status code
func (o *GetWorkflowRetentionRulesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get workflow retention rules internal server error response a status code equal to that given
func (o *GetWorkflowRetentionRulesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get workflow retention rules internal server error response
func (o *GetWorkflowRetentionRulesInternalServerError) Code() int {
	return 500
}

func (o *GetWorkflowRetentionRulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/retention_rules][%d] getWorkflowRetentionRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkflowRetentionRulesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/retention_rules][%d] getWorkflowRetentionRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkflowRetentionRulesInternalServerError) GetPayload() *GetWorkflowRetentionRulesInternalServerErrorBody {
	return o.Payload
}

func (o *GetWorkflowRetentionRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowRetentionRulesInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetWorkflowRetentionRulesForbiddenBody get workflow retention rules forbidden body
swagger:model GetWorkflowRetentionRulesForbiddenBody
*/
type GetWorkflowRetentionRulesForbiddenBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get workflow retention rules forbidden body
func (o *GetWorkflowRetentionRulesForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow retention rules forbidden body based on context it is used
func (o *GetWorkflowRetentionRulesForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowRetentionRulesForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowRetentionRulesForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowRetentionRulesForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowRetentionRulesInternalServerErrorBody get workflow retention rules internal server error body
swagger:model GetWorkflowRetentionRulesInternalServerErrorBody
*/
type GetWorkflowRetentionRulesInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get workflow retention rules internal server error body
func (o *GetWorkflowRetentionRulesInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow retention rules internal server error body based on context it is used
func (o *GetWorkflowRetentionRulesInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowRetentionRulesInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowRetentionRulesInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowRetentionRulesInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowRetentionRulesNotFoundBody get workflow retention rules not found body
swagger:model GetWorkflowRetentionRulesNotFoundBody
*/
type GetWorkflowRetentionRulesNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get workflow retention rules not found body
func (o *GetWorkflowRetentionRulesNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow retention rules not found body based on context it is used
func (o *GetWorkflowRetentionRulesNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowRetentionRulesNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowRetentionRulesNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowRetentionRulesNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowRetentionRulesOKBody get workflow retention rules o k body
swagger:model GetWorkflowRetentionRulesOKBody
*/
type GetWorkflowRetentionRulesOKBody struct {

	// retention rules
	RetentionRules []*GetWorkflowRetentionRulesOKBodyRetentionRulesItems0 `json:"retention_rules"`

	// workflow id
	WorkflowID string `json:"workflow_id,omitempty"`

	// workflow name
	WorkflowName string `json:"workflow_name,omitempty"`
}

// Validate validates this get workflow retention rules o k body
func (o *GetWorkflowRetentionRulesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRetentionRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowRetentionRulesOKBody) validateRetentionRules(formats strfmt.Registry) error {
	if swag.IsZero(o.RetentionRules) { // not required
		return nil
	}

	for i := 0; i < len(o.RetentionRules); i++ {
		if swag.IsZero(o.RetentionRules[i]) { // not required
			continue
		}

		if o.RetentionRules[i] != nil {
			if err := o.RetentionRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getWorkflowRetentionRulesOK" + "." + "retention_rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getWorkflowRetentionRulesOK" + "." + "retention_rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get workflow retention rules o k body based on the context it is used
func (o *GetWorkflowRetentionRulesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRetentionRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowRetentionRulesOKBody) contextValidateRetentionRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.RetentionRules); i++ {

		if o.RetentionRules[i] != nil {

			if swag.IsZero(o.RetentionRules[i]) { // not required
				return nil
			}

			if err := o.RetentionRules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getWorkflowRetentionRulesOK" + "." + "retention_rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getWorkflowRetentionRulesOK" + "." + "retention_rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowRetentionRulesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowRetentionRulesOKBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowRetentionRulesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowRetentionRulesOKBodyRetentionRulesItems0 get workflow retention rules o k body retention rules items0
swagger:model GetWorkflowRetentionRulesOKBodyRetentionRulesItems0
*/
type GetWorkflowRetentionRulesOKBodyRetentionRulesItems0 struct {

	// apply on
	ApplyOn *string `json:"apply_on,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// retention days
	RetentionDays int64 `json:"retention_days,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// workspace files
	WorkspaceFiles string `json:"workspace_files,omitempty"`
}

// Validate validates this get workflow retention rules o k body retention rules items0
func (o *GetWorkflowRetentionRulesOKBodyRetentionRulesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow retention rules o k body retention rules items0 based on context it is used
func (o *GetWorkflowRetentionRulesOKBodyRetentionRulesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowRetentionRulesOKBodyRetentionRulesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowRetentionRulesOKBodyRetentionRulesItems0) UnmarshalBinary(b []byte) error {
	var res GetWorkflowRetentionRulesOKBodyRetentionRulesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowRetentionRulesUnauthorizedBody get workflow retention rules unauthorized body
swagger:model GetWorkflowRetentionRulesUnauthorizedBody
*/
type GetWorkflowRetentionRulesUnauthorizedBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get workflow retention rules unauthorized body
func (o *GetWorkflowRetentionRulesUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow retention rules unauthorized body based on context it is used
func (o *GetWorkflowRetentionRulesUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowRetentionRulesUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowRetentionRulesUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowRetentionRulesUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
