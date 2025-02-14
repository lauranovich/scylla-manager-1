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

// StorageServiceNodesMovingGetReader is a Reader for the StorageServiceNodesMovingGet structure.
type StorageServiceNodesMovingGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceNodesMovingGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceNodesMovingGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceNodesMovingGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceNodesMovingGetOK creates a StorageServiceNodesMovingGetOK with default headers values
func NewStorageServiceNodesMovingGetOK() *StorageServiceNodesMovingGetOK {
	return &StorageServiceNodesMovingGetOK{}
}

/*StorageServiceNodesMovingGetOK handles this case with default header values.

StorageServiceNodesMovingGetOK storage service nodes moving get o k
*/
type StorageServiceNodesMovingGetOK struct {
	Payload []string
}

func (o *StorageServiceNodesMovingGetOK) GetPayload() []string {
	return o.Payload
}

func (o *StorageServiceNodesMovingGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceNodesMovingGetDefault creates a StorageServiceNodesMovingGetDefault with default headers values
func NewStorageServiceNodesMovingGetDefault(code int) *StorageServiceNodesMovingGetDefault {
	return &StorageServiceNodesMovingGetDefault{
		_statusCode: code,
	}
}

/*StorageServiceNodesMovingGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceNodesMovingGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service nodes moving get default response
func (o *StorageServiceNodesMovingGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceNodesMovingGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceNodesMovingGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceNodesMovingGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
