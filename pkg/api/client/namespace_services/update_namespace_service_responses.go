// Code generated by go-swagger; DO NOT EDIT.

package namespace_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateNamespaceServiceReader is a Reader for the UpdateNamespaceService structure.
type UpdateNamespaceServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNamespaceServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNamespaceServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNamespaceServiceOK creates a UpdateNamespaceServiceOK with default headers values
func NewUpdateNamespaceServiceOK() *UpdateNamespaceServiceOK {
	return &UpdateNamespaceServiceOK{}
}

/* UpdateNamespaceServiceOK describes a response with status code 200, with default header values.

successfully created service revision
*/
type UpdateNamespaceServiceOK struct {
}

func (o *UpdateNamespaceServiceOK) Error() string {
	return fmt.Sprintf("[POST /api/functions/namespaces/{namespace}/function/{serviceName}][%d] updateNamespaceServiceOK ", 200)
}

func (o *UpdateNamespaceServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*UpdateNamespaceServiceBody update namespace service body
// Example: {"cmd":"","image":"direktiv/request:v10","minScale":"1","size":"small","trafficPercent":50}
swagger:model UpdateNamespaceServiceBody
*/
type UpdateNamespaceServiceBody struct {

	// cmd
	// Required: true
	Cmd *string `json:"cmd"`

	// Target image a service will use
	// Required: true
	Image *string `json:"image"`

	// Minimum amount of service pods to be live
	// Required: true
	MinScale *int64 `json:"minScale"`

	// Size of created service pods
	// Required: true
	// Enum: [small medium large]
	Size *string `json:"size"`

	// Traffic percentage new revision will use
	// Required: true
	TrafficPercent *int64 `json:"trafficPercent"`
}

// Validate validates this update namespace service body
func (o *UpdateNamespaceServiceBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCmd(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMinScale(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTrafficPercent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNamespaceServiceBody) validateCmd(formats strfmt.Registry) error {

	if err := validate.Required("Service"+"."+"cmd", "body", o.Cmd); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNamespaceServiceBody) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("Service"+"."+"image", "body", o.Image); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNamespaceServiceBody) validateMinScale(formats strfmt.Registry) error {

	if err := validate.Required("Service"+"."+"minScale", "body", o.MinScale); err != nil {
		return err
	}

	return nil
}

var updateNamespaceServiceBodyTypeSizePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["small","medium","large"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNamespaceServiceBodyTypeSizePropEnum = append(updateNamespaceServiceBodyTypeSizePropEnum, v)
	}
}

const (

	// UpdateNamespaceServiceBodySizeSmall captures enum value "small"
	UpdateNamespaceServiceBodySizeSmall string = "small"

	// UpdateNamespaceServiceBodySizeMedium captures enum value "medium"
	UpdateNamespaceServiceBodySizeMedium string = "medium"

	// UpdateNamespaceServiceBodySizeLarge captures enum value "large"
	UpdateNamespaceServiceBodySizeLarge string = "large"
)

// prop value enum
func (o *UpdateNamespaceServiceBody) validateSizeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNamespaceServiceBodyTypeSizePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNamespaceServiceBody) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("Service"+"."+"size", "body", o.Size); err != nil {
		return err
	}

	// value enum
	if err := o.validateSizeEnum("Service"+"."+"size", "body", *o.Size); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNamespaceServiceBody) validateTrafficPercent(formats strfmt.Registry) error {

	if err := validate.Required("Service"+"."+"trafficPercent", "body", o.TrafficPercent); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update namespace service body based on context it is used
func (o *UpdateNamespaceServiceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNamespaceServiceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNamespaceServiceBody) UnmarshalBinary(b []byte) error {
	var res UpdateNamespaceServiceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}