// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AddSecretsReader is a Reader for the AddSecrets structure.
type AddSecretsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddSecretsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddSecretsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAddSecretsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddSecretsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddSecretsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddSecretsCreated creates a AddSecretsCreated with default headers values
func NewAddSecretsCreated() *AddSecretsCreated {
	return &AddSecretsCreated{}
}

/*
AddSecretsCreated describes a response with status code 201, with default header values.

Secrets successfully added.
*/
type AddSecretsCreated struct {
	Payload *AddSecretsCreatedBody
}

// IsSuccess returns true when this add secrets created response has a 2xx status code
func (o *AddSecretsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add secrets created response has a 3xx status code
func (o *AddSecretsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add secrets created response has a 4xx status code
func (o *AddSecretsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add secrets created response has a 5xx status code
func (o *AddSecretsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add secrets created response a status code equal to that given
func (o *AddSecretsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *AddSecretsCreated) Error() string {
	return fmt.Sprintf("[POST /api/secrets/][%d] addSecretsCreated  %+v", 201, o.Payload)
}

func (o *AddSecretsCreated) String() string {
	return fmt.Sprintf("[POST /api/secrets/][%d] addSecretsCreated  %+v", 201, o.Payload)
}

func (o *AddSecretsCreated) GetPayload() *AddSecretsCreatedBody {
	return o.Payload
}

func (o *AddSecretsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddSecretsCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSecretsForbidden creates a AddSecretsForbidden with default headers values
func NewAddSecretsForbidden() *AddSecretsForbidden {
	return &AddSecretsForbidden{}
}

/*
AddSecretsForbidden describes a response with status code 403, with default header values.

Request failed. Token is not valid.
*/
type AddSecretsForbidden struct {
	Payload *AddSecretsForbiddenBody
}

// IsSuccess returns true when this add secrets forbidden response has a 2xx status code
func (o *AddSecretsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add secrets forbidden response has a 3xx status code
func (o *AddSecretsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add secrets forbidden response has a 4xx status code
func (o *AddSecretsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this add secrets forbidden response has a 5xx status code
func (o *AddSecretsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this add secrets forbidden response a status code equal to that given
func (o *AddSecretsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AddSecretsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/secrets/][%d] addSecretsForbidden  %+v", 403, o.Payload)
}

func (o *AddSecretsForbidden) String() string {
	return fmt.Sprintf("[POST /api/secrets/][%d] addSecretsForbidden  %+v", 403, o.Payload)
}

func (o *AddSecretsForbidden) GetPayload() *AddSecretsForbiddenBody {
	return o.Payload
}

func (o *AddSecretsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddSecretsForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSecretsConflict creates a AddSecretsConflict with default headers values
func NewAddSecretsConflict() *AddSecretsConflict {
	return &AddSecretsConflict{}
}

/*
AddSecretsConflict describes a response with status code 409, with default header values.

Request failed. Secrets could not be added due to a conflict.
*/
type AddSecretsConflict struct {
	Payload *AddSecretsConflictBody
}

// IsSuccess returns true when this add secrets conflict response has a 2xx status code
func (o *AddSecretsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add secrets conflict response has a 3xx status code
func (o *AddSecretsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add secrets conflict response has a 4xx status code
func (o *AddSecretsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this add secrets conflict response has a 5xx status code
func (o *AddSecretsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this add secrets conflict response a status code equal to that given
func (o *AddSecretsConflict) IsCode(code int) bool {
	return code == 409
}

func (o *AddSecretsConflict) Error() string {
	return fmt.Sprintf("[POST /api/secrets/][%d] addSecretsConflict  %+v", 409, o.Payload)
}

func (o *AddSecretsConflict) String() string {
	return fmt.Sprintf("[POST /api/secrets/][%d] addSecretsConflict  %+v", 409, o.Payload)
}

func (o *AddSecretsConflict) GetPayload() *AddSecretsConflictBody {
	return o.Payload
}

func (o *AddSecretsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddSecretsConflictBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSecretsInternalServerError creates a AddSecretsInternalServerError with default headers values
func NewAddSecretsInternalServerError() *AddSecretsInternalServerError {
	return &AddSecretsInternalServerError{}
}

/*
AddSecretsInternalServerError describes a response with status code 500, with default header values.

Request failed. Internal server error.
*/
type AddSecretsInternalServerError struct {
	Payload *AddSecretsInternalServerErrorBody
}

// IsSuccess returns true when this add secrets internal server error response has a 2xx status code
func (o *AddSecretsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add secrets internal server error response has a 3xx status code
func (o *AddSecretsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add secrets internal server error response has a 4xx status code
func (o *AddSecretsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add secrets internal server error response has a 5xx status code
func (o *AddSecretsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add secrets internal server error response a status code equal to that given
func (o *AddSecretsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AddSecretsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/secrets/][%d] addSecretsInternalServerError  %+v", 500, o.Payload)
}

func (o *AddSecretsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/secrets/][%d] addSecretsInternalServerError  %+v", 500, o.Payload)
}

func (o *AddSecretsInternalServerError) GetPayload() *AddSecretsInternalServerErrorBody {
	return o.Payload
}

func (o *AddSecretsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddSecretsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AddSecretsConflictBody add secrets conflict body
swagger:model AddSecretsConflictBody
*/
type AddSecretsConflictBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add secrets conflict body
func (o *AddSecretsConflictBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add secrets conflict body based on context it is used
func (o *AddSecretsConflictBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddSecretsConflictBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddSecretsConflictBody) UnmarshalBinary(b []byte) error {
	var res AddSecretsConflictBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddSecretsCreatedBody add secrets created body
swagger:model AddSecretsCreatedBody
*/
type AddSecretsCreatedBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add secrets created body
func (o *AddSecretsCreatedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add secrets created body based on context it is used
func (o *AddSecretsCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddSecretsCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddSecretsCreatedBody) UnmarshalBinary(b []byte) error {
	var res AddSecretsCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddSecretsForbiddenBody add secrets forbidden body
swagger:model AddSecretsForbiddenBody
*/
type AddSecretsForbiddenBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add secrets forbidden body
func (o *AddSecretsForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add secrets forbidden body based on context it is used
func (o *AddSecretsForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddSecretsForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddSecretsForbiddenBody) UnmarshalBinary(b []byte) error {
	var res AddSecretsForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddSecretsInternalServerErrorBody add secrets internal server error body
swagger:model AddSecretsInternalServerErrorBody
*/
type AddSecretsInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add secrets internal server error body
func (o *AddSecretsInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add secrets internal server error body based on context it is used
func (o *AddSecretsInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddSecretsInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddSecretsInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res AddSecretsInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddSecretsParamsBodyAnon Secret definition.
swagger:model AddSecretsParamsBodyAnon
*/
type AddSecretsParamsBodyAnon struct {

	// Secret name
	Name string `json:"name,omitempty"`

	// How will be the secret assigned to the jobs, either exported as an environment variable or mounted as a file.
	// Enum: [env file]
	Type string `json:"type,omitempty"`

	// Secret value
	Value string `json:"value,omitempty"`
}

// Validate validates this add secrets params body anon
func (o *AddSecretsParamsBodyAnon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addSecretsParamsBodyAnonTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["env","file"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addSecretsParamsBodyAnonTypeTypePropEnum = append(addSecretsParamsBodyAnonTypeTypePropEnum, v)
	}
}

const (

	// AddSecretsParamsBodyAnonTypeEnv captures enum value "env"
	AddSecretsParamsBodyAnonTypeEnv string = "env"

	// AddSecretsParamsBodyAnonTypeFile captures enum value "file"
	AddSecretsParamsBodyAnonTypeFile string = "file"
)

// prop value enum
func (o *AddSecretsParamsBodyAnon) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addSecretsParamsBodyAnonTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddSecretsParamsBodyAnon) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this add secrets params body anon based on context it is used
func (o *AddSecretsParamsBodyAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddSecretsParamsBodyAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddSecretsParamsBodyAnon) UnmarshalBinary(b []byte) error {
	var res AddSecretsParamsBodyAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
