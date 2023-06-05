// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GitlabConnectReader is a Reader for the GitlabConnect structure.
type GitlabConnectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GitlabConnectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewGitlabConnectFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/gitlab/connect] gitlab_connect", response, response.Code())
	}
}

// NewGitlabConnectFound creates a GitlabConnectFound with default headers values
func NewGitlabConnectFound() *GitlabConnectFound {
	return &GitlabConnectFound{}
}

/*
GitlabConnectFound describes a response with status code 302, with default header values.

Redirection to GitLab site.
*/
type GitlabConnectFound struct {
}

// IsSuccess returns true when this gitlab connect found response has a 2xx status code
func (o *GitlabConnectFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this gitlab connect found response has a 3xx status code
func (o *GitlabConnectFound) IsRedirect() bool {
	return true
}

// IsClientError returns true when this gitlab connect found response has a 4xx status code
func (o *GitlabConnectFound) IsClientError() bool {
	return false
}

// IsServerError returns true when this gitlab connect found response has a 5xx status code
func (o *GitlabConnectFound) IsServerError() bool {
	return false
}

// IsCode returns true when this gitlab connect found response a status code equal to that given
func (o *GitlabConnectFound) IsCode(code int) bool {
	return code == 302
}

// Code gets the status code for the gitlab connect found response
func (o *GitlabConnectFound) Code() int {
	return 302
}

func (o *GitlabConnectFound) Error() string {
	return fmt.Sprintf("[GET /api/gitlab/connect][%d] gitlabConnectFound ", 302)
}

func (o *GitlabConnectFound) String() string {
	return fmt.Sprintf("[GET /api/gitlab/connect][%d] gitlabConnectFound ", 302)
}

func (o *GitlabConnectFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
