// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteClusterClusterIDRepairUnitUnitIDReader is a Reader for the DeleteClusterClusterIDRepairUnitUnitID structure.
type DeleteClusterClusterIDRepairUnitUnitIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClusterClusterIDRepairUnitUnitIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteClusterClusterIDRepairUnitUnitIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteClusterClusterIDRepairUnitUnitIDOK creates a DeleteClusterClusterIDRepairUnitUnitIDOK with default headers values
func NewDeleteClusterClusterIDRepairUnitUnitIDOK() *DeleteClusterClusterIDRepairUnitUnitIDOK {
	return &DeleteClusterClusterIDRepairUnitUnitIDOK{}
}

/*DeleteClusterClusterIDRepairUnitUnitIDOK handles this case with default header values.

OK
*/
type DeleteClusterClusterIDRepairUnitUnitIDOK struct {
}

func (o *DeleteClusterClusterIDRepairUnitUnitIDOK) Error() string {
	return fmt.Sprintf("[DELETE /cluster/{cluster_id}/repair/unit/{unit_id}][%d] deleteClusterClusterIdRepairUnitUnitIdOK ", 200)
}

func (o *DeleteClusterClusterIDRepairUnitUnitIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
