// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeregisterNodeNoContentCode is the HTTP code returned for type DeregisterNodeNoContent
const DeregisterNodeNoContentCode int = 204

/*DeregisterNodeNoContent Node deregistered

swagger:response deregisterNodeNoContent
*/
type DeregisterNodeNoContent struct {
}

// NewDeregisterNodeNoContent creates DeregisterNodeNoContent with default headers values
func NewDeregisterNodeNoContent() *DeregisterNodeNoContent {

	return &DeregisterNodeNoContent{}
}

// WriteResponse to the client
func (o *DeregisterNodeNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeregisterNodeBadRequestCode is the HTTP code returned for type DeregisterNodeBadRequest
const DeregisterNodeBadRequestCode int = 400

/*DeregisterNodeBadRequest Node in use

swagger:response deregisterNodeBadRequest
*/
type DeregisterNodeBadRequest struct {
}

// NewDeregisterNodeBadRequest creates DeregisterNodeBadRequest with default headers values
func NewDeregisterNodeBadRequest() *DeregisterNodeBadRequest {

	return &DeregisterNodeBadRequest{}
}

// WriteResponse to the client
func (o *DeregisterNodeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// DeregisterNodeNotFoundCode is the HTTP code returned for type DeregisterNodeNotFound
const DeregisterNodeNotFoundCode int = 404

/*DeregisterNodeNotFound Node not found

swagger:response deregisterNodeNotFound
*/
type DeregisterNodeNotFound struct {
}

// NewDeregisterNodeNotFound creates DeregisterNodeNotFound with default headers values
func NewDeregisterNodeNotFound() *DeregisterNodeNotFound {

	return &DeregisterNodeNotFound{}
}

// WriteResponse to the client
func (o *DeregisterNodeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DeregisterNodeInternalServerErrorCode is the HTTP code returned for type DeregisterNodeInternalServerError
const DeregisterNodeInternalServerErrorCode int = 500

/*DeregisterNodeInternalServerError Internal server error

swagger:response deregisterNodeInternalServerError
*/
type DeregisterNodeInternalServerError struct {
}

// NewDeregisterNodeInternalServerError creates DeregisterNodeInternalServerError with default headers values
func NewDeregisterNodeInternalServerError() *DeregisterNodeInternalServerError {

	return &DeregisterNodeInternalServerError{}
}

// WriteResponse to the client
func (o *DeregisterNodeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
