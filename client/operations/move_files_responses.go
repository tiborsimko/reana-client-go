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

// MoveFilesReader is a Reader for the MoveFiles structure.
type MoveFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MoveFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMoveFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMoveFilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMoveFilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMoveFilesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewMoveFilesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewMoveFilesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /api/workflows/move_files/{workflow_id_or_name}] move_files", response, response.Code())
	}
}

// NewMoveFilesOK creates a MoveFilesOK with default headers values
func NewMoveFilesOK() *MoveFilesOK {
	return &MoveFilesOK{}
}

/*
MoveFilesOK describes a response with status code 200, with default header values.

Request succeeded. Message about successfully moved files is returned.
*/
type MoveFilesOK struct {
	Payload *MoveFilesOKBody
}

// IsSuccess returns true when this move files o k response has a 2xx status code
func (o *MoveFilesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this move files o k response has a 3xx status code
func (o *MoveFilesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this move files o k response has a 4xx status code
func (o *MoveFilesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this move files o k response has a 5xx status code
func (o *MoveFilesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this move files o k response a status code equal to that given
func (o *MoveFilesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the move files o k response
func (o *MoveFilesOK) Code() int {
	return 200
}

func (o *MoveFilesOK) Error() string {
	return fmt.Sprintf("[PUT /api/workflows/move_files/{workflow_id_or_name}][%d] moveFilesOK  %+v", 200, o.Payload)
}

func (o *MoveFilesOK) String() string {
	return fmt.Sprintf("[PUT /api/workflows/move_files/{workflow_id_or_name}][%d] moveFilesOK  %+v", 200, o.Payload)
}

func (o *MoveFilesOK) GetPayload() *MoveFilesOKBody {
	return o.Payload
}

func (o *MoveFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(MoveFilesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMoveFilesBadRequest creates a MoveFilesBadRequest with default headers values
func NewMoveFilesBadRequest() *MoveFilesBadRequest {
	return &MoveFilesBadRequest{}
}

/*
MoveFilesBadRequest describes a response with status code 400, with default header values.

Request failed. The incoming payload seems malformed.
*/
type MoveFilesBadRequest struct {
	Payload *MoveFilesBadRequestBody
}

// IsSuccess returns true when this move files bad request response has a 2xx status code
func (o *MoveFilesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this move files bad request response has a 3xx status code
func (o *MoveFilesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this move files bad request response has a 4xx status code
func (o *MoveFilesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this move files bad request response has a 5xx status code
func (o *MoveFilesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this move files bad request response a status code equal to that given
func (o *MoveFilesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the move files bad request response
func (o *MoveFilesBadRequest) Code() int {
	return 400
}

func (o *MoveFilesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/workflows/move_files/{workflow_id_or_name}][%d] moveFilesBadRequest  %+v", 400, o.Payload)
}

func (o *MoveFilesBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/workflows/move_files/{workflow_id_or_name}][%d] moveFilesBadRequest  %+v", 400, o.Payload)
}

func (o *MoveFilesBadRequest) GetPayload() *MoveFilesBadRequestBody {
	return o.Payload
}

func (o *MoveFilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(MoveFilesBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMoveFilesForbidden creates a MoveFilesForbidden with default headers values
func NewMoveFilesForbidden() *MoveFilesForbidden {
	return &MoveFilesForbidden{}
}

/*
MoveFilesForbidden describes a response with status code 403, with default header values.

Request failed. User is not allowed to access workflow.
*/
type MoveFilesForbidden struct {
	Payload *MoveFilesForbiddenBody
}

// IsSuccess returns true when this move files forbidden response has a 2xx status code
func (o *MoveFilesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this move files forbidden response has a 3xx status code
func (o *MoveFilesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this move files forbidden response has a 4xx status code
func (o *MoveFilesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this move files forbidden response has a 5xx status code
func (o *MoveFilesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this move files forbidden response a status code equal to that given
func (o *MoveFilesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the move files forbidden response
func (o *MoveFilesForbidden) Code() int {
	return 403
}

func (o *MoveFilesForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/workflows/move_files/{workflow_id_or_name}][%d] moveFilesForbidden  %+v", 403, o.Payload)
}

func (o *MoveFilesForbidden) String() string {
	return fmt.Sprintf("[PUT /api/workflows/move_files/{workflow_id_or_name}][%d] moveFilesForbidden  %+v", 403, o.Payload)
}

func (o *MoveFilesForbidden) GetPayload() *MoveFilesForbiddenBody {
	return o.Payload
}

func (o *MoveFilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(MoveFilesForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMoveFilesNotFound creates a MoveFilesNotFound with default headers values
func NewMoveFilesNotFound() *MoveFilesNotFound {
	return &MoveFilesNotFound{}
}

/*
MoveFilesNotFound describes a response with status code 404, with default header values.

Request failed. Either User or Workflow does not exist.
*/
type MoveFilesNotFound struct {
	Payload *MoveFilesNotFoundBody
}

// IsSuccess returns true when this move files not found response has a 2xx status code
func (o *MoveFilesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this move files not found response has a 3xx status code
func (o *MoveFilesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this move files not found response has a 4xx status code
func (o *MoveFilesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this move files not found response has a 5xx status code
func (o *MoveFilesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this move files not found response a status code equal to that given
func (o *MoveFilesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the move files not found response
func (o *MoveFilesNotFound) Code() int {
	return 404
}

func (o *MoveFilesNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/workflows/move_files/{workflow_id_or_name}][%d] moveFilesNotFound  %+v", 404, o.Payload)
}

func (o *MoveFilesNotFound) String() string {
	return fmt.Sprintf("[PUT /api/workflows/move_files/{workflow_id_or_name}][%d] moveFilesNotFound  %+v", 404, o.Payload)
}

func (o *MoveFilesNotFound) GetPayload() *MoveFilesNotFoundBody {
	return o.Payload
}

func (o *MoveFilesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(MoveFilesNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMoveFilesConflict creates a MoveFilesConflict with default headers values
func NewMoveFilesConflict() *MoveFilesConflict {
	return &MoveFilesConflict{}
}

/*
MoveFilesConflict describes a response with status code 409, with default header values.

Request failed. The files could not be moved due to a conflict.
*/
type MoveFilesConflict struct {
	Payload *MoveFilesConflictBody
}

// IsSuccess returns true when this move files conflict response has a 2xx status code
func (o *MoveFilesConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this move files conflict response has a 3xx status code
func (o *MoveFilesConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this move files conflict response has a 4xx status code
func (o *MoveFilesConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this move files conflict response has a 5xx status code
func (o *MoveFilesConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this move files conflict response a status code equal to that given
func (o *MoveFilesConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the move files conflict response
func (o *MoveFilesConflict) Code() int {
	return 409
}

func (o *MoveFilesConflict) Error() string {
	return fmt.Sprintf("[PUT /api/workflows/move_files/{workflow_id_or_name}][%d] moveFilesConflict  %+v", 409, o.Payload)
}

func (o *MoveFilesConflict) String() string {
	return fmt.Sprintf("[PUT /api/workflows/move_files/{workflow_id_or_name}][%d] moveFilesConflict  %+v", 409, o.Payload)
}

func (o *MoveFilesConflict) GetPayload() *MoveFilesConflictBody {
	return o.Payload
}

func (o *MoveFilesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(MoveFilesConflictBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMoveFilesInternalServerError creates a MoveFilesInternalServerError with default headers values
func NewMoveFilesInternalServerError() *MoveFilesInternalServerError {
	return &MoveFilesInternalServerError{}
}

/*
MoveFilesInternalServerError describes a response with status code 500, with default header values.

Request failed. Internal controller error.
*/
type MoveFilesInternalServerError struct {
	Payload *MoveFilesInternalServerErrorBody
}

// IsSuccess returns true when this move files internal server error response has a 2xx status code
func (o *MoveFilesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this move files internal server error response has a 3xx status code
func (o *MoveFilesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this move files internal server error response has a 4xx status code
func (o *MoveFilesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this move files internal server error response has a 5xx status code
func (o *MoveFilesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this move files internal server error response a status code equal to that given
func (o *MoveFilesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the move files internal server error response
func (o *MoveFilesInternalServerError) Code() int {
	return 500
}

func (o *MoveFilesInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/workflows/move_files/{workflow_id_or_name}][%d] moveFilesInternalServerError  %+v", 500, o.Payload)
}

func (o *MoveFilesInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/workflows/move_files/{workflow_id_or_name}][%d] moveFilesInternalServerError  %+v", 500, o.Payload)
}

func (o *MoveFilesInternalServerError) GetPayload() *MoveFilesInternalServerErrorBody {
	return o.Payload
}

func (o *MoveFilesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(MoveFilesInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
MoveFilesBadRequestBody move files bad request body
swagger:model MoveFilesBadRequestBody
*/
type MoveFilesBadRequestBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this move files bad request body
func (o *MoveFilesBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this move files bad request body based on context it is used
func (o *MoveFilesBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *MoveFilesBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MoveFilesBadRequestBody) UnmarshalBinary(b []byte) error {
	var res MoveFilesBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
MoveFilesConflictBody move files conflict body
swagger:model MoveFilesConflictBody
*/
type MoveFilesConflictBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this move files conflict body
func (o *MoveFilesConflictBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this move files conflict body based on context it is used
func (o *MoveFilesConflictBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *MoveFilesConflictBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MoveFilesConflictBody) UnmarshalBinary(b []byte) error {
	var res MoveFilesConflictBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
MoveFilesForbiddenBody move files forbidden body
swagger:model MoveFilesForbiddenBody
*/
type MoveFilesForbiddenBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this move files forbidden body
func (o *MoveFilesForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this move files forbidden body based on context it is used
func (o *MoveFilesForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *MoveFilesForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MoveFilesForbiddenBody) UnmarshalBinary(b []byte) error {
	var res MoveFilesForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
MoveFilesInternalServerErrorBody move files internal server error body
swagger:model MoveFilesInternalServerErrorBody
*/
type MoveFilesInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this move files internal server error body
func (o *MoveFilesInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this move files internal server error body based on context it is used
func (o *MoveFilesInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *MoveFilesInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MoveFilesInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res MoveFilesInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
MoveFilesNotFoundBody move files not found body
swagger:model MoveFilesNotFoundBody
*/
type MoveFilesNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this move files not found body
func (o *MoveFilesNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this move files not found body based on context it is used
func (o *MoveFilesNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *MoveFilesNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MoveFilesNotFoundBody) UnmarshalBinary(b []byte) error {
	var res MoveFilesNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
MoveFilesOKBody move files o k body
swagger:model MoveFilesOKBody
*/
type MoveFilesOKBody struct {

	// message
	Message string `json:"message,omitempty"`

	// workflow id
	WorkflowID string `json:"workflow_id,omitempty"`

	// workflow name
	WorkflowName string `json:"workflow_name,omitempty"`
}

// Validate validates this move files o k body
func (o *MoveFilesOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this move files o k body based on context it is used
func (o *MoveFilesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *MoveFilesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MoveFilesOKBody) UnmarshalBinary(b []byte) error {
	var res MoveFilesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
