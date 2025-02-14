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

	"github.com/scylladb/scylla-manager/swagger/gen/scylla-manager/models"
)

// NewPostClustersParams creates a new PostClustersParams object
// with the default values initialized.
func NewPostClustersParams() *PostClustersParams {
	var ()
	return &PostClustersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostClustersParamsWithTimeout creates a new PostClustersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostClustersParamsWithTimeout(timeout time.Duration) *PostClustersParams {
	var ()
	return &PostClustersParams{

		timeout: timeout,
	}
}

// NewPostClustersParamsWithContext creates a new PostClustersParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostClustersParamsWithContext(ctx context.Context) *PostClustersParams {
	var ()
	return &PostClustersParams{

		Context: ctx,
	}
}

// NewPostClustersParamsWithHTTPClient creates a new PostClustersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostClustersParamsWithHTTPClient(client *http.Client) *PostClustersParams {
	var ()
	return &PostClustersParams{
		HTTPClient: client,
	}
}

/*PostClustersParams contains all the parameters to send to the API endpoint
for the post clusters operation typically these are written to a http.Request
*/
type PostClustersParams struct {

	/*Cluster*/
	Cluster *models.Cluster

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post clusters params
func (o *PostClustersParams) WithTimeout(timeout time.Duration) *PostClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post clusters params
func (o *PostClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post clusters params
func (o *PostClustersParams) WithContext(ctx context.Context) *PostClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post clusters params
func (o *PostClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post clusters params
func (o *PostClustersParams) WithHTTPClient(client *http.Client) *PostClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post clusters params
func (o *PostClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCluster adds the cluster to the post clusters params
func (o *PostClustersParams) WithCluster(cluster *models.Cluster) *PostClustersParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the post clusters params
func (o *PostClustersParams) SetCluster(cluster *models.Cluster) {
	o.Cluster = cluster
}

// WriteToRequest writes these params to a swagger request
func (o *PostClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cluster != nil {
		if err := r.SetBodyParam(o.Cluster); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
