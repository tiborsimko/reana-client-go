// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteGitlabWebhookReader is a Reader for the DeleteGitlabWebhook structure.
type DeleteGitlabWebhookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGitlabWebhookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteGitlabWebhookNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteGitlabWebhookForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteGitlabWebhookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteGitlabWebhookInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteGitlabWebhookNoContent creates a DeleteGitlabWebhookNoContent with default headers values
func NewDeleteGitlabWebhookNoContent() *DeleteGitlabWebhookNoContent {
	return &DeleteGitlabWebhookNoContent{}
}

/* DeleteGitlabWebhookNoContent describes a response with status code 204, with default header values.

The webhook was properly deleted.
*/
type DeleteGitlabWebhookNoContent struct {
}

func (o *DeleteGitlabWebhookNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/gitlab/webhook][%d] deleteGitlabWebhookNoContent ", 204)
}

func (o *DeleteGitlabWebhookNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteGitlabWebhookForbidden creates a DeleteGitlabWebhookForbidden with default headers values
func NewDeleteGitlabWebhookForbidden() *DeleteGitlabWebhookForbidden {
	return &DeleteGitlabWebhookForbidden{}
}

/* DeleteGitlabWebhookForbidden describes a response with status code 403, with default header values.

Request failed. User token not valid.
*/
type DeleteGitlabWebhookForbidden struct {
}

func (o *DeleteGitlabWebhookForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/gitlab/webhook][%d] deleteGitlabWebhookForbidden ", 403)
}

func (o *DeleteGitlabWebhookForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteGitlabWebhookNotFound creates a DeleteGitlabWebhookNotFound with default headers values
func NewDeleteGitlabWebhookNotFound() *DeleteGitlabWebhookNotFound {
	return &DeleteGitlabWebhookNotFound{}
}

/* DeleteGitlabWebhookNotFound describes a response with status code 404, with default header values.

No webhook found with provided id.
*/
type DeleteGitlabWebhookNotFound struct {
}

func (o *DeleteGitlabWebhookNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/gitlab/webhook][%d] deleteGitlabWebhookNotFound ", 404)
}

func (o *DeleteGitlabWebhookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteGitlabWebhookInternalServerError creates a DeleteGitlabWebhookInternalServerError with default headers values
func NewDeleteGitlabWebhookInternalServerError() *DeleteGitlabWebhookInternalServerError {
	return &DeleteGitlabWebhookInternalServerError{}
}

/* DeleteGitlabWebhookInternalServerError describes a response with status code 500, with default header values.

Request failed. Internal controller error.
*/
type DeleteGitlabWebhookInternalServerError struct {
}

func (o *DeleteGitlabWebhookInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/gitlab/webhook][%d] deleteGitlabWebhookInternalServerError ", 500)
}

func (o *DeleteGitlabWebhookInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
