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

// CloseInteractiveSessionReader is a Reader for the CloseInteractiveSession structure.
type CloseInteractiveSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloseInteractiveSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloseInteractiveSessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCloseInteractiveSessionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCloseInteractiveSessionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCloseInteractiveSessionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCloseInteractiveSessionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/workflows/{workflow_id_or_name}/close/] close_interactive_session", response, response.Code())
	}
}

// NewCloseInteractiveSessionOK creates a CloseInteractiveSessionOK with default headers values
func NewCloseInteractiveSessionOK() *CloseInteractiveSessionOK {
	return &CloseInteractiveSessionOK{}
}

/*
CloseInteractiveSessionOK describes a response with status code 200, with default header values.

Request succeeded. The interactive session has been closed.
*/
type CloseInteractiveSessionOK struct {
	Payload *CloseInteractiveSessionOKBody
}

// IsSuccess returns true when this close interactive session o k response has a 2xx status code
func (o *CloseInteractiveSessionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this close interactive session o k response has a 3xx status code
func (o *CloseInteractiveSessionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this close interactive session o k response has a 4xx status code
func (o *CloseInteractiveSessionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this close interactive session o k response has a 5xx status code
func (o *CloseInteractiveSessionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this close interactive session o k response a status code equal to that given
func (o *CloseInteractiveSessionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the close interactive session o k response
func (o *CloseInteractiveSessionOK) Code() int {
	return 200
}

func (o *CloseInteractiveSessionOK) Error() string {
	return fmt.Sprintf("[POST /api/workflows/{workflow_id_or_name}/close/][%d] closeInteractiveSessionOK  %+v", 200, o.Payload)
}

func (o *CloseInteractiveSessionOK) String() string {
	return fmt.Sprintf("[POST /api/workflows/{workflow_id_or_name}/close/][%d] closeInteractiveSessionOK  %+v", 200, o.Payload)
}

func (o *CloseInteractiveSessionOK) GetPayload() *CloseInteractiveSessionOKBody {
	return o.Payload
}

func (o *CloseInteractiveSessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CloseInteractiveSessionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloseInteractiveSessionBadRequest creates a CloseInteractiveSessionBadRequest with default headers values
func NewCloseInteractiveSessionBadRequest() *CloseInteractiveSessionBadRequest {
	return &CloseInteractiveSessionBadRequest{}
}

/*
CloseInteractiveSessionBadRequest describes a response with status code 400, with default header values.

Request failed. The incoming payload seems malformed.
*/
type CloseInteractiveSessionBadRequest struct {
	Payload *CloseInteractiveSessionBadRequestBody
}

// IsSuccess returns true when this close interactive session bad request response has a 2xx status code
func (o *CloseInteractiveSessionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this close interactive session bad request response has a 3xx status code
func (o *CloseInteractiveSessionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this close interactive session bad request response has a 4xx status code
func (o *CloseInteractiveSessionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this close interactive session bad request response has a 5xx status code
func (o *CloseInteractiveSessionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this close interactive session bad request response a status code equal to that given
func (o *CloseInteractiveSessionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the close interactive session bad request response
func (o *CloseInteractiveSessionBadRequest) Code() int {
	return 400
}

func (o *CloseInteractiveSessionBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/workflows/{workflow_id_or_name}/close/][%d] closeInteractiveSessionBadRequest  %+v", 400, o.Payload)
}

func (o *CloseInteractiveSessionBadRequest) String() string {
	return fmt.Sprintf("[POST /api/workflows/{workflow_id_or_name}/close/][%d] closeInteractiveSessionBadRequest  %+v", 400, o.Payload)
}

func (o *CloseInteractiveSessionBadRequest) GetPayload() *CloseInteractiveSessionBadRequestBody {
	return o.Payload
}

func (o *CloseInteractiveSessionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CloseInteractiveSessionBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloseInteractiveSessionForbidden creates a CloseInteractiveSessionForbidden with default headers values
func NewCloseInteractiveSessionForbidden() *CloseInteractiveSessionForbidden {
	return &CloseInteractiveSessionForbidden{}
}

/*
CloseInteractiveSessionForbidden describes a response with status code 403, with default header values.

Request failed. User is not allowed to access workflow.
*/
type CloseInteractiveSessionForbidden struct {
	Payload *CloseInteractiveSessionForbiddenBody
}

// IsSuccess returns true when this close interactive session forbidden response has a 2xx status code
func (o *CloseInteractiveSessionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this close interactive session forbidden response has a 3xx status code
func (o *CloseInteractiveSessionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this close interactive session forbidden response has a 4xx status code
func (o *CloseInteractiveSessionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this close interactive session forbidden response has a 5xx status code
func (o *CloseInteractiveSessionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this close interactive session forbidden response a status code equal to that given
func (o *CloseInteractiveSessionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the close interactive session forbidden response
func (o *CloseInteractiveSessionForbidden) Code() int {
	return 403
}

func (o *CloseInteractiveSessionForbidden) Error() string {
	return fmt.Sprintf("[POST /api/workflows/{workflow_id_or_name}/close/][%d] closeInteractiveSessionForbidden  %+v", 403, o.Payload)
}

func (o *CloseInteractiveSessionForbidden) String() string {
	return fmt.Sprintf("[POST /api/workflows/{workflow_id_or_name}/close/][%d] closeInteractiveSessionForbidden  %+v", 403, o.Payload)
}

func (o *CloseInteractiveSessionForbidden) GetPayload() *CloseInteractiveSessionForbiddenBody {
	return o.Payload
}

func (o *CloseInteractiveSessionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CloseInteractiveSessionForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloseInteractiveSessionNotFound creates a CloseInteractiveSessionNotFound with default headers values
func NewCloseInteractiveSessionNotFound() *CloseInteractiveSessionNotFound {
	return &CloseInteractiveSessionNotFound{}
}

/*
CloseInteractiveSessionNotFound describes a response with status code 404, with default header values.

Request failed. Either user or workflow does not exist.
*/
type CloseInteractiveSessionNotFound struct {
	Payload *CloseInteractiveSessionNotFoundBody
}

// IsSuccess returns true when this close interactive session not found response has a 2xx status code
func (o *CloseInteractiveSessionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this close interactive session not found response has a 3xx status code
func (o *CloseInteractiveSessionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this close interactive session not found response has a 4xx status code
func (o *CloseInteractiveSessionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this close interactive session not found response has a 5xx status code
func (o *CloseInteractiveSessionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this close interactive session not found response a status code equal to that given
func (o *CloseInteractiveSessionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the close interactive session not found response
func (o *CloseInteractiveSessionNotFound) Code() int {
	return 404
}

func (o *CloseInteractiveSessionNotFound) Error() string {
	return fmt.Sprintf("[POST /api/workflows/{workflow_id_or_name}/close/][%d] closeInteractiveSessionNotFound  %+v", 404, o.Payload)
}

func (o *CloseInteractiveSessionNotFound) String() string {
	return fmt.Sprintf("[POST /api/workflows/{workflow_id_or_name}/close/][%d] closeInteractiveSessionNotFound  %+v", 404, o.Payload)
}

func (o *CloseInteractiveSessionNotFound) GetPayload() *CloseInteractiveSessionNotFoundBody {
	return o.Payload
}

func (o *CloseInteractiveSessionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CloseInteractiveSessionNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloseInteractiveSessionInternalServerError creates a CloseInteractiveSessionInternalServerError with default headers values
func NewCloseInteractiveSessionInternalServerError() *CloseInteractiveSessionInternalServerError {
	return &CloseInteractiveSessionInternalServerError{}
}

/*
CloseInteractiveSessionInternalServerError describes a response with status code 500, with default header values.

Request failed. Internal controller error.
*/
type CloseInteractiveSessionInternalServerError struct {
	Payload *CloseInteractiveSessionInternalServerErrorBody
}

// IsSuccess returns true when this close interactive session internal server error response has a 2xx status code
func (o *CloseInteractiveSessionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this close interactive session internal server error response has a 3xx status code
func (o *CloseInteractiveSessionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this close interactive session internal server error response has a 4xx status code
func (o *CloseInteractiveSessionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this close interactive session internal server error response has a 5xx status code
func (o *CloseInteractiveSessionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this close interactive session internal server error response a status code equal to that given
func (o *CloseInteractiveSessionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the close interactive session internal server error response
func (o *CloseInteractiveSessionInternalServerError) Code() int {
	return 500
}

func (o *CloseInteractiveSessionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/workflows/{workflow_id_or_name}/close/][%d] closeInteractiveSessionInternalServerError  %+v", 500, o.Payload)
}

func (o *CloseInteractiveSessionInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/workflows/{workflow_id_or_name}/close/][%d] closeInteractiveSessionInternalServerError  %+v", 500, o.Payload)
}

func (o *CloseInteractiveSessionInternalServerError) GetPayload() *CloseInteractiveSessionInternalServerErrorBody {
	return o.Payload
}

func (o *CloseInteractiveSessionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CloseInteractiveSessionInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CloseInteractiveSessionBadRequestBody close interactive session bad request body
swagger:model CloseInteractiveSessionBadRequestBody
*/
type CloseInteractiveSessionBadRequestBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this close interactive session bad request body
func (o *CloseInteractiveSessionBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this close interactive session bad request body based on context it is used
func (o *CloseInteractiveSessionBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CloseInteractiveSessionBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CloseInteractiveSessionBadRequestBody) UnmarshalBinary(b []byte) error {
	var res CloseInteractiveSessionBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CloseInteractiveSessionForbiddenBody close interactive session forbidden body
swagger:model CloseInteractiveSessionForbiddenBody
*/
type CloseInteractiveSessionForbiddenBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this close interactive session forbidden body
func (o *CloseInteractiveSessionForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this close interactive session forbidden body based on context it is used
func (o *CloseInteractiveSessionForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CloseInteractiveSessionForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CloseInteractiveSessionForbiddenBody) UnmarshalBinary(b []byte) error {
	var res CloseInteractiveSessionForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CloseInteractiveSessionInternalServerErrorBody close interactive session internal server error body
swagger:model CloseInteractiveSessionInternalServerErrorBody
*/
type CloseInteractiveSessionInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this close interactive session internal server error body
func (o *CloseInteractiveSessionInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this close interactive session internal server error body based on context it is used
func (o *CloseInteractiveSessionInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CloseInteractiveSessionInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CloseInteractiveSessionInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res CloseInteractiveSessionInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CloseInteractiveSessionNotFoundBody close interactive session not found body
swagger:model CloseInteractiveSessionNotFoundBody
*/
type CloseInteractiveSessionNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this close interactive session not found body
func (o *CloseInteractiveSessionNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this close interactive session not found body based on context it is used
func (o *CloseInteractiveSessionNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CloseInteractiveSessionNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CloseInteractiveSessionNotFoundBody) UnmarshalBinary(b []byte) error {
	var res CloseInteractiveSessionNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CloseInteractiveSessionOKBody close interactive session o k body
swagger:model CloseInteractiveSessionOKBody
*/
type CloseInteractiveSessionOKBody struct {

	// path
	Path string `json:"path,omitempty"`
}

// Validate validates this close interactive session o k body
func (o *CloseInteractiveSessionOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this close interactive session o k body based on context it is used
func (o *CloseInteractiveSessionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CloseInteractiveSessionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CloseInteractiveSessionOKBody) UnmarshalBinary(b []byte) error {
	var res CloseInteractiveSessionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
