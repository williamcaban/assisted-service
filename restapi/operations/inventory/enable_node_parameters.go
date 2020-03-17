// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewEnableNodeParams creates a new EnableNodeParams object
// no default values defined in spec.
func NewEnableNodeParams() EnableNodeParams {

	return EnableNodeParams{}
}

// EnableNodeParams contains all the bound params for the enable node operation
// typically these are obtained from a http.Request
//
// swagger:parameters EnableNode
type EnableNodeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The ID of the node to enable
	  Required: true
	  In: path
	*/
	NodeID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewEnableNodeParams() beforehand.
func (o *EnableNodeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rNodeID, rhkNodeID, _ := route.Params.GetOK("node_id")
	if err := o.bindNodeID(rNodeID, rhkNodeID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindNodeID binds and validates parameter NodeID from path.
func (o *EnableNodeParams) bindNodeID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("node_id", "path", "strfmt.UUID", raw)
	}
	o.NodeID = *(value.(*strfmt.UUID))

	if err := o.validateNodeID(formats); err != nil {
		return err
	}

	return nil
}

// validateNodeID carries on validations for parameter NodeID
func (o *EnableNodeParams) validateNodeID(formats strfmt.Registry) error {

	if err := validate.FormatOf("node_id", "path", "uuid", o.NodeID.String(), formats); err != nil {
		return err
	}
	return nil
}
