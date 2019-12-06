// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/pkg/scyllaclient/internal/agent/models"
)

// CoreBwlimitReader is a Reader for the CoreBwlimit structure.
type CoreBwlimitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CoreBwlimitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCoreBwlimitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCoreBwlimitDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCoreBwlimitOK creates a CoreBwlimitOK with default headers values
func NewCoreBwlimitOK() *CoreBwlimitOK {
	return &CoreBwlimitOK{}
}

/*CoreBwlimitOK handles this case with default header values.

bandwidth rate
*/
type CoreBwlimitOK struct {
	Payload *models.Bandwidth
	JobID   int64
}

func (o *CoreBwlimitOK) GetPayload() *models.Bandwidth {
	return o.Payload
}

func (o *CoreBwlimitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Bandwidth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	if jobIDHeader := response.GetHeader("x-rclone-jobid"); jobIDHeader != "" {
		jobID, err := strconv.ParseInt(jobIDHeader, 10, 64)
		if err != nil {
			return err
		}

		o.JobID = jobID
	}
	return nil
}

// NewCoreBwlimitDefault creates a CoreBwlimitDefault with default headers values
func NewCoreBwlimitDefault(code int) *CoreBwlimitDefault {
	return &CoreBwlimitDefault{
		_statusCode: code,
	}
}

/*CoreBwlimitDefault handles this case with default header values.

Server error
*/
type CoreBwlimitDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
	JobID   int64
}

// Code gets the status code for the core bwlimit default response
func (o *CoreBwlimitDefault) Code() int {
	return o._statusCode
}

func (o *CoreBwlimitDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CoreBwlimitDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	if jobIDHeader := response.GetHeader("x-rclone-jobid"); jobIDHeader != "" {
		jobID, err := strconv.ParseInt(jobIDHeader, 10, 64)
		if err != nil {
			return err
		}

		o.JobID = jobID
	}
	return nil
}

func (o *CoreBwlimitDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
