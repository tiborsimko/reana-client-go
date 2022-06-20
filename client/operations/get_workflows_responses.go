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

// GetWorkflowsReader is a Reader for the GetWorkflows structure.
type GetWorkflowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkflowsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkflowsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkflowsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkflowsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkflowsOK creates a GetWorkflowsOK with default headers values
func NewGetWorkflowsOK() *GetWorkflowsOK {
	return &GetWorkflowsOK{}
}

/* GetWorkflowsOK describes a response with status code 200, with default header values.

Request succeeded. The response contains the list of all workflows.
*/
type GetWorkflowsOK struct {
	Payload *GetWorkflowsOKBody
}

func (o *GetWorkflowsOK) Error() string {
	return fmt.Sprintf("[GET /api/workflows][%d] getWorkflowsOK  %+v", 200, o.Payload)
}
func (o *GetWorkflowsOK) GetPayload() *GetWorkflowsOKBody {
	return o.Payload
}

func (o *GetWorkflowsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowsBadRequest creates a GetWorkflowsBadRequest with default headers values
func NewGetWorkflowsBadRequest() *GetWorkflowsBadRequest {
	return &GetWorkflowsBadRequest{}
}

/* GetWorkflowsBadRequest describes a response with status code 400, with default header values.

Request failed. The incoming payload seems malformed.
*/
type GetWorkflowsBadRequest struct {
}

func (o *GetWorkflowsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/workflows][%d] getWorkflowsBadRequest ", 400)
}

func (o *GetWorkflowsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWorkflowsForbidden creates a GetWorkflowsForbidden with default headers values
func NewGetWorkflowsForbidden() *GetWorkflowsForbidden {
	return &GetWorkflowsForbidden{}
}

/* GetWorkflowsForbidden describes a response with status code 403, with default header values.

Request failed. User is not allowed to access workflow.
*/
type GetWorkflowsForbidden struct {
}

func (o *GetWorkflowsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/workflows][%d] getWorkflowsForbidden ", 403)
}

func (o *GetWorkflowsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWorkflowsNotFound creates a GetWorkflowsNotFound with default headers values
func NewGetWorkflowsNotFound() *GetWorkflowsNotFound {
	return &GetWorkflowsNotFound{}
}

/* GetWorkflowsNotFound describes a response with status code 404, with default header values.

Request failed. User does not exist.
*/
type GetWorkflowsNotFound struct {
}

func (o *GetWorkflowsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/workflows][%d] getWorkflowsNotFound ", 404)
}

func (o *GetWorkflowsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWorkflowsInternalServerError creates a GetWorkflowsInternalServerError with default headers values
func NewGetWorkflowsInternalServerError() *GetWorkflowsInternalServerError {
	return &GetWorkflowsInternalServerError{}
}

/* GetWorkflowsInternalServerError describes a response with status code 500, with default header values.

Request failed. Internal controller error.
*/
type GetWorkflowsInternalServerError struct {
}

func (o *GetWorkflowsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/workflows][%d] getWorkflowsInternalServerError ", 500)
}

func (o *GetWorkflowsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetWorkflowsOKBody get workflows o k body
swagger:model GetWorkflowsOKBody
*/
type GetWorkflowsOKBody struct {

	// items
	Items []*GetWorkflowsOKBodyItemsItems0 `json:"items"`

	// total
	Total int64 `json:"total,omitempty"`
}

// Validate validates this get workflows o k body
func (o *GetWorkflowsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowsOKBody) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(o.Items) { // not required
		return nil
	}

	for i := 0; i < len(o.Items); i++ {
		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {
			if err := o.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getWorkflowsOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getWorkflowsOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get workflows o k body based on the context it is used
func (o *GetWorkflowsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowsOKBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {
			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getWorkflowsOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getWorkflowsOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowsOKBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWorkflowsOKBodyItemsItems0 get workflows o k body items items0
swagger:model GetWorkflowsOKBodyItemsItems0
*/
type GetWorkflowsOKBodyItemsItems0 struct {

	// created
	Created string `json:"created,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// launcher url
	LauncherURL *string `json:"launcher_url,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// progress
	Progress *GetWorkflowsOKBodyItemsItems0Progress `json:"progress,omitempty"`

	// size
	Size *GetWorkflowsOKBodyItemsItems0Size `json:"size,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// user
	User string `json:"user,omitempty"`
}

// Validate validates this get workflows o k body items items0
func (o *GetWorkflowsOKBodyItemsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProgress(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowsOKBodyItemsItems0) validateProgress(formats strfmt.Registry) error {
	if swag.IsZero(o.Progress) { // not required
		return nil
	}

	if o.Progress != nil {
		if err := o.Progress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("progress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("progress")
			}
			return err
		}
	}

	return nil
}

func (o *GetWorkflowsOKBodyItemsItems0) validateSize(formats strfmt.Registry) error {
	if swag.IsZero(o.Size) { // not required
		return nil
	}

	if o.Size != nil {
		if err := o.Size.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get workflows o k body items items0 based on the context it is used
func (o *GetWorkflowsOKBodyItemsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateProgress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowsOKBodyItemsItems0) contextValidateProgress(ctx context.Context, formats strfmt.Registry) error {

	if o.Progress != nil {
		if err := o.Progress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("progress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("progress")
			}
			return err
		}
	}

	return nil
}

func (o *GetWorkflowsOKBodyItemsItems0) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

	if o.Size != nil {
		if err := o.Size.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowsOKBodyItemsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowsOKBodyItemsItems0) UnmarshalBinary(b []byte) error {
	var res GetWorkflowsOKBodyItemsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWorkflowsOKBodyItemsItems0Progress get workflows o k body items items0 progress
swagger:model GetWorkflowsOKBodyItemsItems0Progress
*/
type GetWorkflowsOKBodyItemsItems0Progress struct {

	// current command
	CurrentCommand *string `json:"current_command,omitempty"`

	// current step name
	CurrentStepName *string `json:"current_step_name,omitempty"`

	// failed
	Failed *GetWorkflowsOKBodyItemsItems0ProgressFailed `json:"failed,omitempty"`

	// finished
	Finished *GetWorkflowsOKBodyItemsItems0ProgressFinished `json:"finished,omitempty"`

	// run finished at
	RunFinishedAt *string `json:"run_finished_at,omitempty"`

	// run started at
	RunStartedAt *string `json:"run_started_at,omitempty"`

	// running
	Running *GetWorkflowsOKBodyItemsItems0ProgressRunning `json:"running,omitempty"`

	// total
	Total *GetWorkflowsOKBodyItemsItems0ProgressTotal `json:"total,omitempty"`
}

// Validate validates this get workflows o k body items items0 progress
func (o *GetWorkflowsOKBodyItemsItems0Progress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFailed(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFinished(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRunning(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowsOKBodyItemsItems0Progress) validateFailed(formats strfmt.Registry) error {
	if swag.IsZero(o.Failed) { // not required
		return nil
	}

	if o.Failed != nil {
		if err := o.Failed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("progress" + "." + "failed")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("progress" + "." + "failed")
			}
			return err
		}
	}

	return nil
}

func (o *GetWorkflowsOKBodyItemsItems0Progress) validateFinished(formats strfmt.Registry) error {
	if swag.IsZero(o.Finished) { // not required
		return nil
	}

	if o.Finished != nil {
		if err := o.Finished.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("progress" + "." + "finished")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("progress" + "." + "finished")
			}
			return err
		}
	}

	return nil
}

func (o *GetWorkflowsOKBodyItemsItems0Progress) validateRunning(formats strfmt.Registry) error {
	if swag.IsZero(o.Running) { // not required
		return nil
	}

	if o.Running != nil {
		if err := o.Running.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("progress" + "." + "running")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("progress" + "." + "running")
			}
			return err
		}
	}

	return nil
}

func (o *GetWorkflowsOKBodyItemsItems0Progress) validateTotal(formats strfmt.Registry) error {
	if swag.IsZero(o.Total) { // not required
		return nil
	}

	if o.Total != nil {
		if err := o.Total.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("progress" + "." + "total")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("progress" + "." + "total")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get workflows o k body items items0 progress based on the context it is used
func (o *GetWorkflowsOKBodyItemsItems0Progress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFailed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateFinished(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRunning(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateTotal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowsOKBodyItemsItems0Progress) contextValidateFailed(ctx context.Context, formats strfmt.Registry) error {

	if o.Failed != nil {
		if err := o.Failed.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("progress" + "." + "failed")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("progress" + "." + "failed")
			}
			return err
		}
	}

	return nil
}

func (o *GetWorkflowsOKBodyItemsItems0Progress) contextValidateFinished(ctx context.Context, formats strfmt.Registry) error {

	if o.Finished != nil {
		if err := o.Finished.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("progress" + "." + "finished")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("progress" + "." + "finished")
			}
			return err
		}
	}

	return nil
}

func (o *GetWorkflowsOKBodyItemsItems0Progress) contextValidateRunning(ctx context.Context, formats strfmt.Registry) error {

	if o.Running != nil {
		if err := o.Running.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("progress" + "." + "running")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("progress" + "." + "running")
			}
			return err
		}
	}

	return nil
}

func (o *GetWorkflowsOKBodyItemsItems0Progress) contextValidateTotal(ctx context.Context, formats strfmt.Registry) error {

	if o.Total != nil {
		if err := o.Total.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("progress" + "." + "total")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("progress" + "." + "total")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowsOKBodyItemsItems0Progress) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowsOKBodyItemsItems0Progress) UnmarshalBinary(b []byte) error {
	var res GetWorkflowsOKBodyItemsItems0Progress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWorkflowsOKBodyItemsItems0ProgressFailed get workflows o k body items items0 progress failed
swagger:model GetWorkflowsOKBodyItemsItems0ProgressFailed
*/
type GetWorkflowsOKBodyItemsItems0ProgressFailed struct {

	// job ids
	JobIds []string `json:"job_ids"`

	// total
	Total float64 `json:"total,omitempty"`
}

// Validate validates this get workflows o k body items items0 progress failed
func (o *GetWorkflowsOKBodyItemsItems0ProgressFailed) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflows o k body items items0 progress failed based on context it is used
func (o *GetWorkflowsOKBodyItemsItems0ProgressFailed) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowsOKBodyItemsItems0ProgressFailed) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowsOKBodyItemsItems0ProgressFailed) UnmarshalBinary(b []byte) error {
	var res GetWorkflowsOKBodyItemsItems0ProgressFailed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWorkflowsOKBodyItemsItems0ProgressFinished get workflows o k body items items0 progress finished
swagger:model GetWorkflowsOKBodyItemsItems0ProgressFinished
*/
type GetWorkflowsOKBodyItemsItems0ProgressFinished struct {

	// job ids
	JobIds []string `json:"job_ids"`

	// total
	Total float64 `json:"total,omitempty"`
}

// Validate validates this get workflows o k body items items0 progress finished
func (o *GetWorkflowsOKBodyItemsItems0ProgressFinished) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflows o k body items items0 progress finished based on context it is used
func (o *GetWorkflowsOKBodyItemsItems0ProgressFinished) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowsOKBodyItemsItems0ProgressFinished) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowsOKBodyItemsItems0ProgressFinished) UnmarshalBinary(b []byte) error {
	var res GetWorkflowsOKBodyItemsItems0ProgressFinished
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWorkflowsOKBodyItemsItems0ProgressRunning get workflows o k body items items0 progress running
swagger:model GetWorkflowsOKBodyItemsItems0ProgressRunning
*/
type GetWorkflowsOKBodyItemsItems0ProgressRunning struct {

	// job ids
	JobIds []string `json:"job_ids"`

	// total
	Total float64 `json:"total,omitempty"`
}

// Validate validates this get workflows o k body items items0 progress running
func (o *GetWorkflowsOKBodyItemsItems0ProgressRunning) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflows o k body items items0 progress running based on context it is used
func (o *GetWorkflowsOKBodyItemsItems0ProgressRunning) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowsOKBodyItemsItems0ProgressRunning) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowsOKBodyItemsItems0ProgressRunning) UnmarshalBinary(b []byte) error {
	var res GetWorkflowsOKBodyItemsItems0ProgressRunning
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWorkflowsOKBodyItemsItems0ProgressTotal get workflows o k body items items0 progress total
swagger:model GetWorkflowsOKBodyItemsItems0ProgressTotal
*/
type GetWorkflowsOKBodyItemsItems0ProgressTotal struct {

	// job ids
	JobIds []string `json:"job_ids"`

	// total
	Total float64 `json:"total,omitempty"`
}

// Validate validates this get workflows o k body items items0 progress total
func (o *GetWorkflowsOKBodyItemsItems0ProgressTotal) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflows o k body items items0 progress total based on context it is used
func (o *GetWorkflowsOKBodyItemsItems0ProgressTotal) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowsOKBodyItemsItems0ProgressTotal) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowsOKBodyItemsItems0ProgressTotal) UnmarshalBinary(b []byte) error {
	var res GetWorkflowsOKBodyItemsItems0ProgressTotal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWorkflowsOKBodyItemsItems0Size get workflows o k body items items0 size
swagger:model GetWorkflowsOKBodyItemsItems0Size
*/
type GetWorkflowsOKBodyItemsItems0Size struct {

	// human readable
	HumanReadable string `json:"human_readable,omitempty"`

	// raw
	Raw float64 `json:"raw,omitempty"`
}

// Validate validates this get workflows o k body items items0 size
func (o *GetWorkflowsOKBodyItemsItems0Size) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflows o k body items items0 size based on context it is used
func (o *GetWorkflowsOKBodyItemsItems0Size) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowsOKBodyItemsItems0Size) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowsOKBodyItemsItems0Size) UnmarshalBinary(b []byte) error {
	var res GetWorkflowsOKBodyItemsItems0Size
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
