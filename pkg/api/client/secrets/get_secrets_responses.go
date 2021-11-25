// Code generated by go-swagger; DO NOT EDIT.

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/direktiv/direktiv/pkg/api/models"
)

// GetSecretsReader is a Reader for the GetSecrets structure.
type GetSecretsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecretsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSecretsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSecretsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSecretsOK creates a GetSecretsOK with default headers values
func NewGetSecretsOK() *GetSecretsOK {
	return &GetSecretsOK{}
}

/* GetSecretsOK describes a response with status code 200, with default header values.

successfully got namespace nodes
*/
type GetSecretsOK struct {
	Payload models.OkBody
}

func (o *GetSecretsOK) Error() string {
	return fmt.Sprintf("[GET /api/namespaces/{namespace}/secrets][%d] getSecretsOK  %+v", 200, o.Payload)
}
func (o *GetSecretsOK) GetPayload() models.OkBody {
	return o.Payload
}

func (o *GetSecretsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecretsDefault creates a GetSecretsDefault with default headers values
func NewGetSecretsDefault(code int) *GetSecretsDefault {
	return &GetSecretsDefault{
		_statusCode: code,
	}
}

/* GetSecretsDefault describes a response with status code -1, with default header values.

an error has occurred
*/
type GetSecretsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get secrets default response
func (o *GetSecretsDefault) Code() int {
	return o._statusCode
}

func (o *GetSecretsDefault) Error() string {
	return fmt.Sprintf("[GET /api/namespaces/{namespace}/secrets][%d] getSecrets default  %+v", o._statusCode, o.Payload)
}
func (o *GetSecretsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetSecretsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
