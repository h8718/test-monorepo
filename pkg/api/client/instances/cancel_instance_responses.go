// Code generated by go-swagger; DO NOT EDIT.

package instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CancelInstanceReader is a Reader for the CancelInstance structure.
type CancelInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCancelInstanceOK creates a CancelInstanceOK with default headers values
func NewCancelInstanceOK() *CancelInstanceOK {
	return &CancelInstanceOK{}
}

/* CancelInstanceOK describes a response with status code 200, with default header values.

successfully cancelled instance
*/
type CancelInstanceOK struct {
}

func (o *CancelInstanceOK) Error() string {
	return fmt.Sprintf("[POST /api/namespaces/{namespace}/instances/{instance}/cancel][%d] cancelInstanceOK ", 200)
}

func (o *CancelInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}