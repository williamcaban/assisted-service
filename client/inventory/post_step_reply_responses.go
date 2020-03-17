// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostStepReplyReader is a Reader for the PostStepReply structure.
type PostStepReplyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStepReplyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostStepReplyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStepReplyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostStepReplyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStepReplyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostStepReplyNoContent creates a PostStepReplyNoContent with default headers values
func NewPostStepReplyNoContent() *PostStepReplyNoContent {
	return &PostStepReplyNoContent{}
}

/*PostStepReplyNoContent handles this case with default header values.

Reply accepted
*/
type PostStepReplyNoContent struct {
}

func (o *PostStepReplyNoContent) Error() string {
	return fmt.Sprintf("[POST /nodes/{node_id}/next-steps/reply][%d] postStepReplyNoContent ", 204)
}

func (o *PostStepReplyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostStepReplyBadRequest creates a PostStepReplyBadRequest with default headers values
func NewPostStepReplyBadRequest() *PostStepReplyBadRequest {
	return &PostStepReplyBadRequest{}
}

/*PostStepReplyBadRequest handles this case with default header values.

Invalid input
*/
type PostStepReplyBadRequest struct {
}

func (o *PostStepReplyBadRequest) Error() string {
	return fmt.Sprintf("[POST /nodes/{node_id}/next-steps/reply][%d] postStepReplyBadRequest ", 400)
}

func (o *PostStepReplyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostStepReplyNotFound creates a PostStepReplyNotFound with default headers values
func NewPostStepReplyNotFound() *PostStepReplyNotFound {
	return &PostStepReplyNotFound{}
}

/*PostStepReplyNotFound handles this case with default header values.

Node not found
*/
type PostStepReplyNotFound struct {
}

func (o *PostStepReplyNotFound) Error() string {
	return fmt.Sprintf("[POST /nodes/{node_id}/next-steps/reply][%d] postStepReplyNotFound ", 404)
}

func (o *PostStepReplyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostStepReplyInternalServerError creates a PostStepReplyInternalServerError with default headers values
func NewPostStepReplyInternalServerError() *PostStepReplyInternalServerError {
	return &PostStepReplyInternalServerError{}
}

/*PostStepReplyInternalServerError handles this case with default header values.

Internal server error
*/
type PostStepReplyInternalServerError struct {
}

func (o *PostStepReplyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /nodes/{node_id}/next-steps/reply][%d] postStepReplyInternalServerError ", 500)
}

func (o *PostStepReplyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
