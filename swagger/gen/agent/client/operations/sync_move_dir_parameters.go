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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/scylladb/scylla-manager/swagger/gen/agent/models"
)

// NewSyncMoveDirParams creates a new SyncMoveDirParams object
// with the default values initialized.
func NewSyncMoveDirParams() *SyncMoveDirParams {
	var (
		asyncDefault = bool(true)
	)
	return &SyncMoveDirParams{
		Async: asyncDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSyncMoveDirParamsWithTimeout creates a new SyncMoveDirParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSyncMoveDirParamsWithTimeout(timeout time.Duration) *SyncMoveDirParams {
	var (
		asyncDefault = bool(true)
	)
	return &SyncMoveDirParams{
		Async: asyncDefault,

		timeout: timeout,
	}
}

// NewSyncMoveDirParamsWithContext creates a new SyncMoveDirParams object
// with the default values initialized, and the ability to set a context for a request
func NewSyncMoveDirParamsWithContext(ctx context.Context) *SyncMoveDirParams {
	var (
		asyncDefault = bool(true)
	)
	return &SyncMoveDirParams{
		Async: asyncDefault,

		Context: ctx,
	}
}

// NewSyncMoveDirParamsWithHTTPClient creates a new SyncMoveDirParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSyncMoveDirParamsWithHTTPClient(client *http.Client) *SyncMoveDirParams {
	var (
		asyncDefault = bool(true)
	)
	return &SyncMoveDirParams{
		Async:      asyncDefault,
		HTTPClient: client,
	}
}

/*SyncMoveDirParams contains all the parameters to send to the API endpoint
for the sync move dir operation typically these are written to a http.Request
*/
type SyncMoveDirParams struct {

	/*Options
	  Options

	*/
	Options *models.MoveOrCopyFileOptions
	/*Async
	  Async request

	*/
	Async bool
	/*Group
	  Place this operation under this stat group

	*/
	Group string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sync move dir params
func (o *SyncMoveDirParams) WithTimeout(timeout time.Duration) *SyncMoveDirParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sync move dir params
func (o *SyncMoveDirParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sync move dir params
func (o *SyncMoveDirParams) WithContext(ctx context.Context) *SyncMoveDirParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sync move dir params
func (o *SyncMoveDirParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sync move dir params
func (o *SyncMoveDirParams) WithHTTPClient(client *http.Client) *SyncMoveDirParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sync move dir params
func (o *SyncMoveDirParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOptions adds the options to the sync move dir params
func (o *SyncMoveDirParams) WithOptions(options *models.MoveOrCopyFileOptions) *SyncMoveDirParams {
	o.SetOptions(options)
	return o
}

// SetOptions adds the options to the sync move dir params
func (o *SyncMoveDirParams) SetOptions(options *models.MoveOrCopyFileOptions) {
	o.Options = options
}

// WithAsync adds the async to the sync move dir params
func (o *SyncMoveDirParams) WithAsync(async bool) *SyncMoveDirParams {
	o.SetAsync(async)
	return o
}

// SetAsync adds the async to the sync move dir params
func (o *SyncMoveDirParams) SetAsync(async bool) {
	o.Async = async
}

// WithGroup adds the group to the sync move dir params
func (o *SyncMoveDirParams) WithGroup(group string) *SyncMoveDirParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the sync move dir params
func (o *SyncMoveDirParams) SetGroup(group string) {
	o.Group = group
}

// WriteToRequest writes these params to a swagger request
func (o *SyncMoveDirParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Options != nil {
		if err := r.SetBodyParam(o.Options); err != nil {
			return err
		}
	}

	// query param _async
	qrAsync := o.Async
	qAsync := swag.FormatBool(qrAsync)
	if qAsync != "" {
		if err := r.SetQueryParam("_async", qAsync); err != nil {
			return err
		}
	}

	// query param _group
	qrGroup := o.Group
	qGroup := qrGroup
	if qGroup != "" {
		if err := r.SetQueryParam("_group", qGroup); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
