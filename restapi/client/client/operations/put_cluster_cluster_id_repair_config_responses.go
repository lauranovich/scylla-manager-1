// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutClusterClusterIDRepairConfigReader is a Reader for the PutClusterClusterIDRepairConfig structure.
type PutClusterClusterIDRepairConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutClusterClusterIDRepairConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutClusterClusterIDRepairConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPutClusterClusterIDRepairConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutClusterClusterIDRepairConfigOK creates a PutClusterClusterIDRepairConfigOK with default headers values
func NewPutClusterClusterIDRepairConfigOK() *PutClusterClusterIDRepairConfigOK {
	return &PutClusterClusterIDRepairConfigOK{}
}

/*PutClusterClusterIDRepairConfigOK handles this case with default header values.

OK
*/
type PutClusterClusterIDRepairConfigOK struct {
}

func (o *PutClusterClusterIDRepairConfigOK) Error() string {
	return fmt.Sprintf("[PUT /cluster/{cluster_id}/repair/config][%d] putClusterClusterIdRepairConfigOK ", 200)
}

func (o *PutClusterClusterIDRepairConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutClusterClusterIDRepairConfigNotFound creates a PutClusterClusterIDRepairConfigNotFound with default headers values
func NewPutClusterClusterIDRepairConfigNotFound() *PutClusterClusterIDRepairConfigNotFound {
	return &PutClusterClusterIDRepairConfigNotFound{}
}

/*PutClusterClusterIDRepairConfigNotFound handles this case with default header values.

not found
*/
type PutClusterClusterIDRepairConfigNotFound struct {
}

func (o *PutClusterClusterIDRepairConfigNotFound) Error() string {
	return fmt.Sprintf("[PUT /cluster/{cluster_id}/repair/config][%d] putClusterClusterIdRepairConfigNotFound ", 404)
}

func (o *PutClusterClusterIDRepairConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
