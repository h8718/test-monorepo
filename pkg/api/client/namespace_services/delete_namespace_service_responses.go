// Code generated by go-swagger; DO NOT EDIT.

package namespace_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteNamespaceServiceReader is a Reader for the DeleteNamespaceService structure.
type DeleteNamespaceServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNamespaceServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNamespaceServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNamespaceServiceOK creates a DeleteNamespaceServiceOK with default headers values
func NewDeleteNamespaceServiceOK() *DeleteNamespaceServiceOK {
	return &DeleteNamespaceServiceOK{}
}

/* DeleteNamespaceServiceOK describes a response with status code 200, with default header values.

successfully deleted service
*/
type DeleteNamespaceServiceOK struct {
}

func (o *DeleteNamespaceServiceOK) Error() string {
	return fmt.Sprintf("[DELETE /api/functions/namespaces/{namespace}/function/{serviceName}][%d] deleteNamespaceServiceOK ", 200)
}

func (o *DeleteNamespaceServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}