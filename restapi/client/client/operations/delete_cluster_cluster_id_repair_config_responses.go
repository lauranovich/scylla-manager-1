// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteClusterClusterIDRepairConfigReader is a Reader for the DeleteClusterClusterIDRepairConfig structure.
type DeleteClusterClusterIDRepairConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClusterClusterIDRepairConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteClusterClusterIDRepairConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteClusterClusterIDRepairConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteClusterClusterIDRepairConfigOK creates a DeleteClusterClusterIDRepairConfigOK with default headers values
func NewDeleteClusterClusterIDRepairConfigOK() *DeleteClusterClusterIDRepairConfigOK {
	return &DeleteClusterClusterIDRepairConfigOK{}
}

/*DeleteClusterClusterIDRepairConfigOK handles this case with default header values.

OK
*/
type DeleteClusterClusterIDRepairConfigOK struct {
}

func (o *DeleteClusterClusterIDRepairConfigOK) Error() string {
	return fmt.Sprintf("[DELETE /cluster/{cluster_id}/repair/config][%d] deleteClusterClusterIdRepairConfigOK ", 200)
}

func (o *DeleteClusterClusterIDRepairConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClusterClusterIDRepairConfigNotFound creates a DeleteClusterClusterIDRepairConfigNotFound with default headers values
func NewDeleteClusterClusterIDRepairConfigNotFound() *DeleteClusterClusterIDRepairConfigNotFound {
	return &DeleteClusterClusterIDRepairConfigNotFound{}
}

/*DeleteClusterClusterIDRepairConfigNotFound handles this case with default header values.

not found
*/
type DeleteClusterClusterIDRepairConfigNotFound struct {
}

func (o *DeleteClusterClusterIDRepairConfigNotFound) Error() string {
	return fmt.Sprintf("[DELETE /cluster/{cluster_id}/repair/config][%d] deleteClusterClusterIdRepairConfigNotFound ", 404)
}

func (o *DeleteClusterClusterIDRepairConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
