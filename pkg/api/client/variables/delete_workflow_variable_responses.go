// Code generated by go-swagger; DO NOT EDIT.

package variables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteWorkflowVariableReader is a Reader for the DeleteWorkflowVariable structure.
type DeleteWorkflowVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWorkflowVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteWorkflowVariableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteWorkflowVariableOK creates a DeleteWorkflowVariableOK with default headers values
func NewDeleteWorkflowVariableOK() *DeleteWorkflowVariableOK {
	return &DeleteWorkflowVariableOK{}
}

/* DeleteWorkflowVariableOK describes a response with status code 200, with default header values.

successfully deleted workflow variable
*/
type DeleteWorkflowVariableOK struct {
}

func (o *DeleteWorkflowVariableOK) Error() string {
	return fmt.Sprintf("[DELETE /api/namespaces/{namespace}/tree/{workflow}?op=delete-var][%d] deleteWorkflowVariableOK ", 200)
}

func (o *DeleteWorkflowVariableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
