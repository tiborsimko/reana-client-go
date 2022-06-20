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

// OpenInteractiveSessionReader is a Reader for the OpenInteractiveSession structure.
type OpenInteractiveSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpenInteractiveSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpenInteractiveSessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOpenInteractiveSessionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOpenInteractiveSessionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpenInteractiveSessionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOpenInteractiveSessionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpenInteractiveSessionOK creates a OpenInteractiveSessionOK with default headers values
func NewOpenInteractiveSessionOK() *OpenInteractiveSessionOK {
	return &OpenInteractiveSessionOK{}
}

/* OpenInteractiveSessionOK describes a response with status code 200, with default header values.

Request succeeded. The interactive session has been opened.
*/
type OpenInteractiveSessionOK struct {
	Payload *OpenInteractiveSessionOKBody
}

func (o *OpenInteractiveSessionOK) Error() string {
	return fmt.Sprintf("[POST /api/workflows/{workflow_id_or_name}/open/{interactive_session_type}][%d] openInteractiveSessionOK  %+v", 200, o.Payload)
}
func (o *OpenInteractiveSessionOK) GetPayload() *OpenInteractiveSessionOKBody {
	return o.Payload
}

func (o *OpenInteractiveSessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(OpenInteractiveSessionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenInteractiveSessionBadRequest creates a OpenInteractiveSessionBadRequest with default headers values
func NewOpenInteractiveSessionBadRequest() *OpenInteractiveSessionBadRequest {
	return &OpenInteractiveSessionBadRequest{}
}

/* OpenInteractiveSessionBadRequest describes a response with status code 400, with default header values.

Request failed. The incoming payload seems malformed.
*/
type OpenInteractiveSessionBadRequest struct {
}

func (o *OpenInteractiveSessionBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/workflows/{workflow_id_or_name}/open/{interactive_session_type}][%d] openInteractiveSessionBadRequest ", 400)
}

func (o *OpenInteractiveSessionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOpenInteractiveSessionForbidden creates a OpenInteractiveSessionForbidden with default headers values
func NewOpenInteractiveSessionForbidden() *OpenInteractiveSessionForbidden {
	return &OpenInteractiveSessionForbidden{}
}

/* OpenInteractiveSessionForbidden describes a response with status code 403, with default header values.

Request failed. User is not allowed to access workflow.
*/
type OpenInteractiveSessionForbidden struct {
}

func (o *OpenInteractiveSessionForbidden) Error() string {
	return fmt.Sprintf("[POST /api/workflows/{workflow_id_or_name}/open/{interactive_session_type}][%d] openInteractiveSessionForbidden ", 403)
}

func (o *OpenInteractiveSessionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOpenInteractiveSessionNotFound creates a OpenInteractiveSessionNotFound with default headers values
func NewOpenInteractiveSessionNotFound() *OpenInteractiveSessionNotFound {
	return &OpenInteractiveSessionNotFound{}
}

/* OpenInteractiveSessionNotFound describes a response with status code 404, with default header values.

Request failed. Either user or workflow does not exist.
*/
type OpenInteractiveSessionNotFound struct {
}

func (o *OpenInteractiveSessionNotFound) Error() string {
	return fmt.Sprintf("[POST /api/workflows/{workflow_id_or_name}/open/{interactive_session_type}][%d] openInteractiveSessionNotFound ", 404)
}

func (o *OpenInteractiveSessionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOpenInteractiveSessionInternalServerError creates a OpenInteractiveSessionInternalServerError with default headers values
func NewOpenInteractiveSessionInternalServerError() *OpenInteractiveSessionInternalServerError {
	return &OpenInteractiveSessionInternalServerError{}
}

/* OpenInteractiveSessionInternalServerError describes a response with status code 500, with default header values.

Request failed. Internal controller error.
*/
type OpenInteractiveSessionInternalServerError struct {
}

func (o *OpenInteractiveSessionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/workflows/{workflow_id_or_name}/open/{interactive_session_type}][%d] openInteractiveSessionInternalServerError ", 500)
}

func (o *OpenInteractiveSessionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*OpenInteractiveSessionBody open interactive session body
swagger:model OpenInteractiveSessionBody
*/
type OpenInteractiveSessionBody struct {

	// Replaces the default Docker image of an interactive session.
	Image string `json:"image,omitempty"`
}

// Validate validates this open interactive session body
func (o *OpenInteractiveSessionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this open interactive session body based on context it is used
func (o *OpenInteractiveSessionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *OpenInteractiveSessionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *OpenInteractiveSessionBody) UnmarshalBinary(b []byte) error {
	var res OpenInteractiveSessionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*OpenInteractiveSessionOKBody open interactive session o k body
swagger:model OpenInteractiveSessionOKBody
*/
type OpenInteractiveSessionOKBody struct {

	// path
	Path string `json:"path,omitempty"`
}

// Validate validates this open interactive session o k body
func (o *OpenInteractiveSessionOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this open interactive session o k body based on context it is used
func (o *OpenInteractiveSessionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *OpenInteractiveSessionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *OpenInteractiveSessionOKBody) UnmarshalBinary(b []byte) error {
	var res OpenInteractiveSessionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
