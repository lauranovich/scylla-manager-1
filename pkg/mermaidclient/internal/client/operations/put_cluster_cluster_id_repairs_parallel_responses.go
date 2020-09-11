// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/pkg/mermaidclient/internal/models"
)

// PutClusterClusterIDRepairsParallelReader is a Reader for the PutClusterClusterIDRepairsParallel structure.
type PutClusterClusterIDRepairsParallelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutClusterClusterIDRepairsParallelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutClusterClusterIDRepairsParallelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutClusterClusterIDRepairsParallelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutClusterClusterIDRepairsParallelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutClusterClusterIDRepairsParallelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPutClusterClusterIDRepairsParallelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutClusterClusterIDRepairsParallelOK creates a PutClusterClusterIDRepairsParallelOK with default headers values
func NewPutClusterClusterIDRepairsParallelOK() *PutClusterClusterIDRepairsParallelOK {
	return &PutClusterClusterIDRepairsParallelOK{}
}

/*PutClusterClusterIDRepairsParallelOK handles this case with default header values.

OK
*/
type PutClusterClusterIDRepairsParallelOK struct {
}

func (o *PutClusterClusterIDRepairsParallelOK) Error() string {
	return fmt.Sprintf("[PUT /cluster/{cluster_id}/repairs/parallel][%d] putClusterClusterIdRepairsParallelOK ", 200)
}

func (o *PutClusterClusterIDRepairsParallelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutClusterClusterIDRepairsParallelBadRequest creates a PutClusterClusterIDRepairsParallelBadRequest with default headers values
func NewPutClusterClusterIDRepairsParallelBadRequest() *PutClusterClusterIDRepairsParallelBadRequest {
	return &PutClusterClusterIDRepairsParallelBadRequest{}
}

/*PutClusterClusterIDRepairsParallelBadRequest handles this case with default header values.

Bad Request
*/
type PutClusterClusterIDRepairsParallelBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *PutClusterClusterIDRepairsParallelBadRequest) Error() string {
	return fmt.Sprintf("[PUT /cluster/{cluster_id}/repairs/parallel][%d] putClusterClusterIdRepairsParallelBadRequest  %+v", 400, o.Payload)
}

func (o *PutClusterClusterIDRepairsParallelBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PutClusterClusterIDRepairsParallelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutClusterClusterIDRepairsParallelNotFound creates a PutClusterClusterIDRepairsParallelNotFound with default headers values
func NewPutClusterClusterIDRepairsParallelNotFound() *PutClusterClusterIDRepairsParallelNotFound {
	return &PutClusterClusterIDRepairsParallelNotFound{}
}

/*PutClusterClusterIDRepairsParallelNotFound handles this case with default header values.

Not found
*/
type PutClusterClusterIDRepairsParallelNotFound struct {
	Payload *models.ErrorResponse
}

func (o *PutClusterClusterIDRepairsParallelNotFound) Error() string {
	return fmt.Sprintf("[PUT /cluster/{cluster_id}/repairs/parallel][%d] putClusterClusterIdRepairsParallelNotFound  %+v", 404, o.Payload)
}

func (o *PutClusterClusterIDRepairsParallelNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PutClusterClusterIDRepairsParallelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutClusterClusterIDRepairsParallelInternalServerError creates a PutClusterClusterIDRepairsParallelInternalServerError with default headers values
func NewPutClusterClusterIDRepairsParallelInternalServerError() *PutClusterClusterIDRepairsParallelInternalServerError {
	return &PutClusterClusterIDRepairsParallelInternalServerError{}
}

/*PutClusterClusterIDRepairsParallelInternalServerError handles this case with default header values.

Server error
*/
type PutClusterClusterIDRepairsParallelInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *PutClusterClusterIDRepairsParallelInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /cluster/{cluster_id}/repairs/parallel][%d] putClusterClusterIdRepairsParallelInternalServerError  %+v", 500, o.Payload)
}

func (o *PutClusterClusterIDRepairsParallelInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PutClusterClusterIDRepairsParallelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutClusterClusterIDRepairsParallelDefault creates a PutClusterClusterIDRepairsParallelDefault with default headers values
func NewPutClusterClusterIDRepairsParallelDefault(code int) *PutClusterClusterIDRepairsParallelDefault {
	return &PutClusterClusterIDRepairsParallelDefault{
		_statusCode: code,
	}
}

/*PutClusterClusterIDRepairsParallelDefault handles this case with default header values.

Unexpected error
*/
type PutClusterClusterIDRepairsParallelDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the put cluster cluster ID repairs parallel default response
func (o *PutClusterClusterIDRepairsParallelDefault) Code() int {
	return o._statusCode
}

func (o *PutClusterClusterIDRepairsParallelDefault) Error() string {
	return fmt.Sprintf("[PUT /cluster/{cluster_id}/repairs/parallel][%d] PutClusterClusterIDRepairsParallel default  %+v", o._statusCode, o.Payload)
}

func (o *PutClusterClusterIDRepairsParallelDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PutClusterClusterIDRepairsParallelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
