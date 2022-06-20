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

// GetSecretsReader is a Reader for the GetSecrets structure.
type GetSecretsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecretsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSecretsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetSecretsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSecretsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSecretsOK creates a GetSecretsOK with default headers values
func NewGetSecretsOK() *GetSecretsOK {
	return &GetSecretsOK{}
}

/* GetSecretsOK describes a response with status code 200, with default header values.

List of user secrets.
*/
type GetSecretsOK struct {
	Payload []*GetSecretsOKBodyItems0
}

func (o *GetSecretsOK) Error() string {
	return fmt.Sprintf("[GET /api/secrets][%d] getSecretsOK  %+v", 200, o.Payload)
}
func (o *GetSecretsOK) GetPayload() []*GetSecretsOKBodyItems0 {
	return o.Payload
}

func (o *GetSecretsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecretsForbidden creates a GetSecretsForbidden with default headers values
func NewGetSecretsForbidden() *GetSecretsForbidden {
	return &GetSecretsForbidden{}
}

/* GetSecretsForbidden describes a response with status code 403, with default header values.

Request failed. Token is not valid.
*/
type GetSecretsForbidden struct {
}

func (o *GetSecretsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/secrets][%d] getSecretsForbidden ", 403)
}

func (o *GetSecretsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSecretsInternalServerError creates a GetSecretsInternalServerError with default headers values
func NewGetSecretsInternalServerError() *GetSecretsInternalServerError {
	return &GetSecretsInternalServerError{}
}

/* GetSecretsInternalServerError describes a response with status code 500, with default header values.

Request failed. Internal server error.
*/
type GetSecretsInternalServerError struct {
}

func (o *GetSecretsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/secrets][%d] getSecretsInternalServerError ", 500)
}

func (o *GetSecretsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetSecretsOKBodyItems0 get secrets o k body items0
swagger:model GetSecretsOKBodyItems0
*/
type GetSecretsOKBodyItems0 struct {

	// Secret name
	Name string `json:"name,omitempty"`

	// How will be the secret assigned to the jobs, either exported as an environment variable or mounted as a file.
	// Enum: [env file]
	Type string `json:"type,omitempty"`
}

// Validate validates this get secrets o k body items0
func (o *GetSecretsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getSecretsOKBodyItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["env","file"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getSecretsOKBodyItems0TypeTypePropEnum = append(getSecretsOKBodyItems0TypeTypePropEnum, v)
	}
}

const (

	// GetSecretsOKBodyItems0TypeEnv captures enum value "env"
	GetSecretsOKBodyItems0TypeEnv string = "env"

	// GetSecretsOKBodyItems0TypeFile captures enum value "file"
	GetSecretsOKBodyItems0TypeFile string = "file"
)

// prop value enum
func (o *GetSecretsOKBodyItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getSecretsOKBodyItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetSecretsOKBodyItems0) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get secrets o k body items0 based on context it is used
func (o *GetSecretsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSecretsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSecretsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetSecretsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
