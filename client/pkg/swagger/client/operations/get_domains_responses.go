// Copyright: This file is part of korrel8r, released under https://github.com/korrel8r/korrel8r/blob/main/LICENSE

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/korrel8r/korrel8r/client/pkg/swagger/models"
)

// GetDomainsReader is a Reader for the GetDomains structure.
type GetDomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDomainsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDomainsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDomainsOK creates a GetDomainsOK with default headers values
func NewGetDomainsOK() *GetDomainsOK {
	return &GetDomainsOK{}
}

/*
GetDomainsOK describes a response with status code 200, with default header values.

OK
*/
type GetDomainsOK struct {
	Payload []*models.Domain
}

// IsSuccess returns true when this get domains o k response has a 2xx status code
func (o *GetDomainsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get domains o k response has a 3xx status code
func (o *GetDomainsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domains o k response has a 4xx status code
func (o *GetDomainsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get domains o k response has a 5xx status code
func (o *GetDomainsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get domains o k response a status code equal to that given
func (o *GetDomainsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get domains o k response
func (o *GetDomainsOK) Code() int {
	return 200
}

func (o *GetDomainsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /domains][%d] getDomainsOK %s", 200, payload)
}

func (o *GetDomainsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /domains][%d] getDomainsOK %s", 200, payload)
}

func (o *GetDomainsOK) GetPayload() []*models.Domain {
	return o.Payload
}

func (o *GetDomainsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainsDefault creates a GetDomainsDefault with default headers values
func NewGetDomainsDefault(code int) *GetDomainsDefault {
	return &GetDomainsDefault{
		_statusCode: code,
	}
}

/*
GetDomainsDefault describes a response with status code -1, with default header values.

GetDomainsDefault get domains default
*/
type GetDomainsDefault struct {
	_statusCode int

	Payload string
}

// IsSuccess returns true when this get domains default response has a 2xx status code
func (o *GetDomainsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get domains default response has a 3xx status code
func (o *GetDomainsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get domains default response has a 4xx status code
func (o *GetDomainsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get domains default response has a 5xx status code
func (o *GetDomainsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get domains default response a status code equal to that given
func (o *GetDomainsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get domains default response
func (o *GetDomainsDefault) Code() int {
	return o._statusCode
}

func (o *GetDomainsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /domains][%d] GetDomains default %s", o._statusCode, payload)
}

func (o *GetDomainsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /domains][%d] GetDomains default %s", o._statusCode, payload)
}

func (o *GetDomainsDefault) GetPayload() string {
	return o.Payload
}

func (o *GetDomainsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
