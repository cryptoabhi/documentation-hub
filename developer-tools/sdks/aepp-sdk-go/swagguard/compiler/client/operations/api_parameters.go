// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAPIParams creates a new APIParams object
// with the default values initialized.
func NewAPIParams() *APIParams {

	return &APIParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIParamsWithTimeout creates a new APIParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIParamsWithTimeout(timeout time.Duration) *APIParams {

	return &APIParams{

		timeout: timeout,
	}
}

// NewAPIParamsWithContext creates a new APIParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIParamsWithContext(ctx context.Context) *APIParams {

	return &APIParams{

		Context: ctx,
	}
}

// NewAPIParamsWithHTTPClient creates a new APIParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIParamsWithHTTPClient(client *http.Client) *APIParams {

	return &APIParams{
		HTTPClient: client,
	}
}

/*APIParams contains all the parameters to send to the API endpoint
for the Api operation typically these are written to a http.Request
*/
type APIParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Api params
func (o *APIParams) WithTimeout(timeout time.Duration) *APIParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Api params
func (o *APIParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Api params
func (o *APIParams) WithContext(ctx context.Context) *APIParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Api params
func (o *APIParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Api params
func (o *APIParams) WithHTTPClient(client *http.Client) *APIParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Api params
func (o *APIParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *APIParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
