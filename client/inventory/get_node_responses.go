// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/filanov/bm-inventory/models"
)

// GetNodeReader is a Reader for the GetNode structure.
type GetNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetNodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNodeOK creates a GetNodeOK with default headers values
func NewGetNodeOK() *GetNodeOK {
	return &GetNodeOK{}
}

/*GetNodeOK handles this case with default header values.

Node information
*/
type GetNodeOK struct {
	Payload *models.Node
}

func (o *GetNodeOK) Error() string {
	return fmt.Sprintf("[GET /nodes/{node_id}][%d] getNodeOK  %+v", 200, o.Payload)
}

func (o *GetNodeOK) GetPayload() *models.Node {
	return o.Payload
}

func (o *GetNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Node)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeNotFound creates a GetNodeNotFound with default headers values
func NewGetNodeNotFound() *GetNodeNotFound {
	return &GetNodeNotFound{}
}

/*GetNodeNotFound handles this case with default header values.

Node not found
*/
type GetNodeNotFound struct {
}

func (o *GetNodeNotFound) Error() string {
	return fmt.Sprintf("[GET /nodes/{node_id}][%d] getNodeNotFound ", 404)
}

func (o *GetNodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNodeInternalServerError creates a GetNodeInternalServerError with default headers values
func NewGetNodeInternalServerError() *GetNodeInternalServerError {
	return &GetNodeInternalServerError{}
}

/*GetNodeInternalServerError handles this case with default header values.

Internal server error
*/
type GetNodeInternalServerError struct {
}

func (o *GetNodeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /nodes/{node_id}][%d] getNodeInternalServerError ", 500)
}

func (o *GetNodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
