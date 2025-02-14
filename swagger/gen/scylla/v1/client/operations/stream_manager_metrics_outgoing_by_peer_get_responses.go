// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/swagger/gen/scylla/v1/models"
)

// StreamManagerMetricsOutgoingByPeerGetReader is a Reader for the StreamManagerMetricsOutgoingByPeerGet structure.
type StreamManagerMetricsOutgoingByPeerGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StreamManagerMetricsOutgoingByPeerGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStreamManagerMetricsOutgoingByPeerGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStreamManagerMetricsOutgoingByPeerGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStreamManagerMetricsOutgoingByPeerGetOK creates a StreamManagerMetricsOutgoingByPeerGetOK with default headers values
func NewStreamManagerMetricsOutgoingByPeerGetOK() *StreamManagerMetricsOutgoingByPeerGetOK {
	return &StreamManagerMetricsOutgoingByPeerGetOK{}
}

/*StreamManagerMetricsOutgoingByPeerGetOK handles this case with default header values.

StreamManagerMetricsOutgoingByPeerGetOK stream manager metrics outgoing by peer get o k
*/
type StreamManagerMetricsOutgoingByPeerGetOK struct {
	Payload int32
}

func (o *StreamManagerMetricsOutgoingByPeerGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *StreamManagerMetricsOutgoingByPeerGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStreamManagerMetricsOutgoingByPeerGetDefault creates a StreamManagerMetricsOutgoingByPeerGetDefault with default headers values
func NewStreamManagerMetricsOutgoingByPeerGetDefault(code int) *StreamManagerMetricsOutgoingByPeerGetDefault {
	return &StreamManagerMetricsOutgoingByPeerGetDefault{
		_statusCode: code,
	}
}

/*StreamManagerMetricsOutgoingByPeerGetDefault handles this case with default header values.

internal server error
*/
type StreamManagerMetricsOutgoingByPeerGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the stream manager metrics outgoing by peer get default response
func (o *StreamManagerMetricsOutgoingByPeerGetDefault) Code() int {
	return o._statusCode
}

func (o *StreamManagerMetricsOutgoingByPeerGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StreamManagerMetricsOutgoingByPeerGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StreamManagerMetricsOutgoingByPeerGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
