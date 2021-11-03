// Code generated by go-swagger; DO NOT EDIT.

package variables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SetWorkflowVariableReader is a Reader for the SetWorkflowVariable structure.
type SetWorkflowVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetWorkflowVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetWorkflowVariableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetWorkflowVariableOK creates a SetWorkflowVariableOK with default headers values
func NewSetWorkflowVariableOK() *SetWorkflowVariableOK {
	return &SetWorkflowVariableOK{}
}

/* SetWorkflowVariableOK describes a response with status code 200, with default header values.

successfully set workflow variable
*/
type SetWorkflowVariableOK struct {
}

func (o *SetWorkflowVariableOK) Error() string {
	return fmt.Sprintf("[PUT /api/namespaces/{namespace}/tree/{workflow}?op=set-var][%d] setWorkflowVariableOK ", 200)
}

func (o *SetWorkflowVariableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}