// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewLogoutParams creates a new LogoutParams object
// with the default values initialized.
func NewLogoutParams() LogoutParams {
	var ()
	return LogoutParams{}
}

// LogoutParams contains all the bound params for the logout operation
// typically these are obtained from a http.Request
//
// swagger:parameters Logout
type LogoutParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*
	  Required: true
	  In: query
	*/
	Jwt string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *LogoutParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qJwt, qhkJwt, _ := qs.GetOK("jwt")
	if err := o.bindJwt(qJwt, qhkJwt, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LogoutParams) bindJwt(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("jwt", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("jwt", "query", raw); err != nil {
		return err
	}

	o.Jwt = raw

	return nil
}