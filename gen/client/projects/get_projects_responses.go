// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"git.home.foxienet.com/hostnotes/projectservice/gen/models"
)

// GetProjectsReader is a Reader for the GetProjects structure.
type GetProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetProjectsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProjectsOK creates a GetProjectsOK with default headers values
func NewGetProjectsOK() *GetProjectsOK {
	return &GetProjectsOK{}
}

/*GetProjectsOK handles this case with default header values.

list all the current projects
*/
type GetProjectsOK struct {
	Payload []*models.Project
}

func (o *GetProjectsOK) Error() string {
	return fmt.Sprintf("[GET /projects][%d] getProjectsOK  %+v", 200, o.Payload)
}

func (o *GetProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsDefault creates a GetProjectsDefault with default headers values
func NewGetProjectsDefault(code int) *GetProjectsDefault {
	return &GetProjectsDefault{
		_statusCode: code,
	}
}

/*GetProjectsDefault handles this case with default header values.

generic error response
*/
type GetProjectsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get projects default response
func (o *GetProjectsDefault) Code() int {
	return o._statusCode
}

func (o *GetProjectsDefault) Error() string {
	return fmt.Sprintf("[GET /projects][%d] GetProjects default  %+v", o._statusCode, o.Payload)
}

func (o *GetProjectsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}