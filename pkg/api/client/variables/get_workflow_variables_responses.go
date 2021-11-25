// Code generated by go-swagger; DO NOT EDIT.

package variables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetWorkflowVariablesReader is a Reader for the GetWorkflowVariables structure.
type GetWorkflowVariablesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowVariablesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowVariablesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkflowVariablesOK creates a GetWorkflowVariablesOK with default headers values
func NewGetWorkflowVariablesOK() *GetWorkflowVariablesOK {
	return &GetWorkflowVariablesOK{}
}

/* GetWorkflowVariablesOK describes a response with status code 200, with default header values.

successfully got workflow variables
*/
type GetWorkflowVariablesOK struct {
}

func (o *GetWorkflowVariablesOK) Error() string {
	return fmt.Sprintf("[GET /api/namespaces/{namespace}/tree/{workflow}?op=vars][%d] getWorkflowVariablesOK ", 200)
}

func (o *GetWorkflowVariablesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
