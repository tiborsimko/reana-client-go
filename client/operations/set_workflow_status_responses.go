// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SetWorkflowStatusReader is a Reader for the SetWorkflowStatus structure.
type SetWorkflowStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetWorkflowStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetWorkflowStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetWorkflowStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetWorkflowStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetWorkflowStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewSetWorkflowStatusConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetWorkflowStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewSetWorkflowStatusNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /api/workflows/{workflow_id_or_name}/status] set_workflow_status", response, response.Code())
	}
}

// NewSetWorkflowStatusOK creates a SetWorkflowStatusOK with default headers values
func NewSetWorkflowStatusOK() *SetWorkflowStatusOK {
	return &SetWorkflowStatusOK{}
}

/*
SetWorkflowStatusOK describes a response with status code 200, with default header values.

Request succeeded. Info about a workflow, including the status is returned.
*/
type SetWorkflowStatusOK struct {
	Payload *SetWorkflowStatusOKBody
}

// IsSuccess returns true when this set workflow status o k response has a 2xx status code
func (o *SetWorkflowStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set workflow status o k response has a 3xx status code
func (o *SetWorkflowStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set workflow status o k response has a 4xx status code
func (o *SetWorkflowStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set workflow status o k response has a 5xx status code
func (o *SetWorkflowStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set workflow status o k response a status code equal to that given
func (o *SetWorkflowStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the set workflow status o k response
func (o *SetWorkflowStatusOK) Code() int {
	return 200
}

func (o *SetWorkflowStatusOK) Error() string {
	return fmt.Sprintf("[PUT /api/workflows/{workflow_id_or_name}/status][%d] setWorkflowStatusOK  %+v", 200, o.Payload)
}

func (o *SetWorkflowStatusOK) String() string {
	return fmt.Sprintf("[PUT /api/workflows/{workflow_id_or_name}/status][%d] setWorkflowStatusOK  %+v", 200, o.Payload)
}

func (o *SetWorkflowStatusOK) GetPayload() *SetWorkflowStatusOKBody {
	return o.Payload
}

func (o *SetWorkflowStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SetWorkflowStatusOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetWorkflowStatusBadRequest creates a SetWorkflowStatusBadRequest with default headers values
func NewSetWorkflowStatusBadRequest() *SetWorkflowStatusBadRequest {
	return &SetWorkflowStatusBadRequest{}
}

/*
SetWorkflowStatusBadRequest describes a response with status code 400, with default header values.

Request failed. The incoming payload seems malformed.
*/
type SetWorkflowStatusBadRequest struct {
	Payload *SetWorkflowStatusBadRequestBody
}

// IsSuccess returns true when this set workflow status bad request response has a 2xx status code
func (o *SetWorkflowStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set workflow status bad request response has a 3xx status code
func (o *SetWorkflowStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set workflow status bad request response has a 4xx status code
func (o *SetWorkflowStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this set workflow status bad request response has a 5xx status code
func (o *SetWorkflowStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this set workflow status bad request response a status code equal to that given
func (o *SetWorkflowStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the set workflow status bad request response
func (o *SetWorkflowStatusBadRequest) Code() int {
	return 400
}

func (o *SetWorkflowStatusBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/workflows/{workflow_id_or_name}/status][%d] setWorkflowStatusBadRequest  %+v", 400, o.Payload)
}

func (o *SetWorkflowStatusBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/workflows/{workflow_id_or_name}/status][%d] setWorkflowStatusBadRequest  %+v", 400, o.Payload)
}

func (o *SetWorkflowStatusBadRequest) GetPayload() *SetWorkflowStatusBadRequestBody {
	return o.Payload
}

func (o *SetWorkflowStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SetWorkflowStatusBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetWorkflowStatusForbidden creates a SetWorkflowStatusForbidden with default headers values
func NewSetWorkflowStatusForbidden() *SetWorkflowStatusForbidden {
	return &SetWorkflowStatusForbidden{}
}

/*
SetWorkflowStatusForbidden describes a response with status code 403, with default header values.

Request failed. User is not allowed to access workflow.
*/
type SetWorkflowStatusForbidden struct {
	Payload *SetWorkflowStatusForbiddenBody
}

// IsSuccess returns true when this set workflow status forbidden response has a 2xx status code
func (o *SetWorkflowStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set workflow status forbidden response has a 3xx status code
func (o *SetWorkflowStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set workflow status forbidden response has a 4xx status code
func (o *SetWorkflowStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this set workflow status forbidden response has a 5xx status code
func (o *SetWorkflowStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this set workflow status forbidden response a status code equal to that given
func (o *SetWorkflowStatusForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the set workflow status forbidden response
func (o *SetWorkflowStatusForbidden) Code() int {
	return 403
}

func (o *SetWorkflowStatusForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/workflows/{workflow_id_or_name}/status][%d] setWorkflowStatusForbidden  %+v", 403, o.Payload)
}

func (o *SetWorkflowStatusForbidden) String() string {
	return fmt.Sprintf("[PUT /api/workflows/{workflow_id_or_name}/status][%d] setWorkflowStatusForbidden  %+v", 403, o.Payload)
}

func (o *SetWorkflowStatusForbidden) GetPayload() *SetWorkflowStatusForbiddenBody {
	return o.Payload
}

func (o *SetWorkflowStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SetWorkflowStatusForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetWorkflowStatusNotFound creates a SetWorkflowStatusNotFound with default headers values
func NewSetWorkflowStatusNotFound() *SetWorkflowStatusNotFound {
	return &SetWorkflowStatusNotFound{}
}

/*
SetWorkflowStatusNotFound describes a response with status code 404, with default header values.

Request failed. Either User or Workflow does not exist.
*/
type SetWorkflowStatusNotFound struct {
	Payload *SetWorkflowStatusNotFoundBody
}

// IsSuccess returns true when this set workflow status not found response has a 2xx status code
func (o *SetWorkflowStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set workflow status not found response has a 3xx status code
func (o *SetWorkflowStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set workflow status not found response has a 4xx status code
func (o *SetWorkflowStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this set workflow status not found response has a 5xx status code
func (o *SetWorkflowStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this set workflow status not found response a status code equal to that given
func (o *SetWorkflowStatusNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the set workflow status not found response
func (o *SetWorkflowStatusNotFound) Code() int {
	return 404
}

func (o *SetWorkflowStatusNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/workflows/{workflow_id_or_name}/status][%d] setWorkflowStatusNotFound  %+v", 404, o.Payload)
}

func (o *SetWorkflowStatusNotFound) String() string {
	return fmt.Sprintf("[PUT /api/workflows/{workflow_id_or_name}/status][%d] setWorkflowStatusNotFound  %+v", 404, o.Payload)
}

func (o *SetWorkflowStatusNotFound) GetPayload() *SetWorkflowStatusNotFoundBody {
	return o.Payload
}

func (o *SetWorkflowStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SetWorkflowStatusNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetWorkflowStatusConflict creates a SetWorkflowStatusConflict with default headers values
func NewSetWorkflowStatusConflict() *SetWorkflowStatusConflict {
	return &SetWorkflowStatusConflict{}
}

/*
SetWorkflowStatusConflict describes a response with status code 409, with default header values.

Request failed. The workflow could not be started due to a conflict.
*/
type SetWorkflowStatusConflict struct {
	Payload *SetWorkflowStatusConflictBody
}

// IsSuccess returns true when this set workflow status conflict response has a 2xx status code
func (o *SetWorkflowStatusConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set workflow status conflict response has a 3xx status code
func (o *SetWorkflowStatusConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set workflow status conflict response has a 4xx status code
func (o *SetWorkflowStatusConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this set workflow status conflict response has a 5xx status code
func (o *SetWorkflowStatusConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this set workflow status conflict response a status code equal to that given
func (o *SetWorkflowStatusConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the set workflow status conflict response
func (o *SetWorkflowStatusConflict) Code() int {
	return 409
}

func (o *SetWorkflowStatusConflict) Error() string {
	return fmt.Sprintf("[PUT /api/workflows/{workflow_id_or_name}/status][%d] setWorkflowStatusConflict  %+v", 409, o.Payload)
}

func (o *SetWorkflowStatusConflict) String() string {
	return fmt.Sprintf("[PUT /api/workflows/{workflow_id_or_name}/status][%d] setWorkflowStatusConflict  %+v", 409, o.Payload)
}

func (o *SetWorkflowStatusConflict) GetPayload() *SetWorkflowStatusConflictBody {
	return o.Payload
}

func (o *SetWorkflowStatusConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SetWorkflowStatusConflictBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetWorkflowStatusInternalServerError creates a SetWorkflowStatusInternalServerError with default headers values
func NewSetWorkflowStatusInternalServerError() *SetWorkflowStatusInternalServerError {
	return &SetWorkflowStatusInternalServerError{}
}

/*
SetWorkflowStatusInternalServerError describes a response with status code 500, with default header values.

Request failed. Internal controller error.
*/
type SetWorkflowStatusInternalServerError struct {
	Payload *SetWorkflowStatusInternalServerErrorBody
}

// IsSuccess returns true when this set workflow status internal server error response has a 2xx status code
func (o *SetWorkflowStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set workflow status internal server error response has a 3xx status code
func (o *SetWorkflowStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set workflow status internal server error response has a 4xx status code
func (o *SetWorkflowStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this set workflow status internal server error response has a 5xx status code
func (o *SetWorkflowStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this set workflow status internal server error response a status code equal to that given
func (o *SetWorkflowStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the set workflow status internal server error response
func (o *SetWorkflowStatusInternalServerError) Code() int {
	return 500
}

func (o *SetWorkflowStatusInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/workflows/{workflow_id_or_name}/status][%d] setWorkflowStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *SetWorkflowStatusInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/workflows/{workflow_id_or_name}/status][%d] setWorkflowStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *SetWorkflowStatusInternalServerError) GetPayload() *SetWorkflowStatusInternalServerErrorBody {
	return o.Payload
}

func (o *SetWorkflowStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SetWorkflowStatusInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetWorkflowStatusNotImplemented creates a SetWorkflowStatusNotImplemented with default headers values
func NewSetWorkflowStatusNotImplemented() *SetWorkflowStatusNotImplemented {
	return &SetWorkflowStatusNotImplemented{}
}

/*
SetWorkflowStatusNotImplemented describes a response with status code 501, with default header values.

Request failed. The specified status change is not implemented.
*/
type SetWorkflowStatusNotImplemented struct {
	Payload *SetWorkflowStatusNotImplementedBody
}

// IsSuccess returns true when this set workflow status not implemented response has a 2xx status code
func (o *SetWorkflowStatusNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set workflow status not implemented response has a 3xx status code
func (o *SetWorkflowStatusNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set workflow status not implemented response has a 4xx status code
func (o *SetWorkflowStatusNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this set workflow status not implemented response has a 5xx status code
func (o *SetWorkflowStatusNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this set workflow status not implemented response a status code equal to that given
func (o *SetWorkflowStatusNotImplemented) IsCode(code int) bool {
	return code == 501
}

// Code gets the status code for the set workflow status not implemented response
func (o *SetWorkflowStatusNotImplemented) Code() int {
	return 501
}

func (o *SetWorkflowStatusNotImplemented) Error() string {
	return fmt.Sprintf("[PUT /api/workflows/{workflow_id_or_name}/status][%d] setWorkflowStatusNotImplemented  %+v", 501, o.Payload)
}

func (o *SetWorkflowStatusNotImplemented) String() string {
	return fmt.Sprintf("[PUT /api/workflows/{workflow_id_or_name}/status][%d] setWorkflowStatusNotImplemented  %+v", 501, o.Payload)
}

func (o *SetWorkflowStatusNotImplemented) GetPayload() *SetWorkflowStatusNotImplementedBody {
	return o.Payload
}

func (o *SetWorkflowStatusNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SetWorkflowStatusNotImplementedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SetWorkflowStatusBadRequestBody set workflow status bad request body
swagger:model SetWorkflowStatusBadRequestBody
*/
type SetWorkflowStatusBadRequestBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this set workflow status bad request body
func (o *SetWorkflowStatusBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this set workflow status bad request body based on context it is used
func (o *SetWorkflowStatusBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SetWorkflowStatusBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetWorkflowStatusBadRequestBody) UnmarshalBinary(b []byte) error {
	var res SetWorkflowStatusBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SetWorkflowStatusBody set workflow status body
swagger:model SetWorkflowStatusBody
*/
type SetWorkflowStatusBody struct {

	// c a c h e
	CACHE string `json:"CACHE,omitempty"`

	// all runs
	AllRuns bool `json:"all_runs,omitempty"`

	// workspace
	Workspace bool `json:"workspace,omitempty"`
}

// Validate validates this set workflow status body
func (o *SetWorkflowStatusBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this set workflow status body based on context it is used
func (o *SetWorkflowStatusBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SetWorkflowStatusBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetWorkflowStatusBody) UnmarshalBinary(b []byte) error {
	var res SetWorkflowStatusBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SetWorkflowStatusConflictBody set workflow status conflict body
swagger:model SetWorkflowStatusConflictBody
*/
type SetWorkflowStatusConflictBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this set workflow status conflict body
func (o *SetWorkflowStatusConflictBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this set workflow status conflict body based on context it is used
func (o *SetWorkflowStatusConflictBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SetWorkflowStatusConflictBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetWorkflowStatusConflictBody) UnmarshalBinary(b []byte) error {
	var res SetWorkflowStatusConflictBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SetWorkflowStatusForbiddenBody set workflow status forbidden body
swagger:model SetWorkflowStatusForbiddenBody
*/
type SetWorkflowStatusForbiddenBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this set workflow status forbidden body
func (o *SetWorkflowStatusForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this set workflow status forbidden body based on context it is used
func (o *SetWorkflowStatusForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SetWorkflowStatusForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetWorkflowStatusForbiddenBody) UnmarshalBinary(b []byte) error {
	var res SetWorkflowStatusForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SetWorkflowStatusInternalServerErrorBody set workflow status internal server error body
swagger:model SetWorkflowStatusInternalServerErrorBody
*/
type SetWorkflowStatusInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this set workflow status internal server error body
func (o *SetWorkflowStatusInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this set workflow status internal server error body based on context it is used
func (o *SetWorkflowStatusInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SetWorkflowStatusInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetWorkflowStatusInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res SetWorkflowStatusInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SetWorkflowStatusNotFoundBody set workflow status not found body
swagger:model SetWorkflowStatusNotFoundBody
*/
type SetWorkflowStatusNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this set workflow status not found body
func (o *SetWorkflowStatusNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this set workflow status not found body based on context it is used
func (o *SetWorkflowStatusNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SetWorkflowStatusNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetWorkflowStatusNotFoundBody) UnmarshalBinary(b []byte) error {
	var res SetWorkflowStatusNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SetWorkflowStatusNotImplementedBody set workflow status not implemented body
swagger:model SetWorkflowStatusNotImplementedBody
*/
type SetWorkflowStatusNotImplementedBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this set workflow status not implemented body
func (o *SetWorkflowStatusNotImplementedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this set workflow status not implemented body based on context it is used
func (o *SetWorkflowStatusNotImplementedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SetWorkflowStatusNotImplementedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetWorkflowStatusNotImplementedBody) UnmarshalBinary(b []byte) error {
	var res SetWorkflowStatusNotImplementedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SetWorkflowStatusOKBody set workflow status o k body
swagger:model SetWorkflowStatusOKBody
*/
type SetWorkflowStatusOKBody struct {

	// message
	Message string `json:"message,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// user
	User string `json:"user,omitempty"`

	// workflow id
	WorkflowID string `json:"workflow_id,omitempty"`

	// workflow name
	WorkflowName string `json:"workflow_name,omitempty"`
}

// Validate validates this set workflow status o k body
func (o *SetWorkflowStatusOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this set workflow status o k body based on context it is used
func (o *SetWorkflowStatusOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SetWorkflowStatusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetWorkflowStatusOKBody) UnmarshalBinary(b []byte) error {
	var res SetWorkflowStatusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
