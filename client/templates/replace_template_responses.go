package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/rackn/rocket-skates/models"
)

// ReplaceTemplateReader is a Reader for the ReplaceTemplate structure.
type ReplaceTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewReplaceTemplateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewReplaceTemplateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewReplaceTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewReplaceTemplateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 417:
		result := NewReplaceTemplateExpectationFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewReplaceTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceTemplateCreated creates a ReplaceTemplateCreated with default headers values
func NewReplaceTemplateCreated() *ReplaceTemplateCreated {
	return &ReplaceTemplateCreated{}
}

/*ReplaceTemplateCreated handles this case with default header values.

ReplaceTemplateCreated replace template created
*/
type ReplaceTemplateCreated struct {
	Payload *models.TemplateOutput
}

func (o *ReplaceTemplateCreated) Error() string {
	return fmt.Sprintf("[POST /templates/{uuid}][%d] replaceTemplateCreated  %+v", 201, o.Payload)
}

func (o *ReplaceTemplateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TemplateOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceTemplateAccepted creates a ReplaceTemplateAccepted with default headers values
func NewReplaceTemplateAccepted() *ReplaceTemplateAccepted {
	return &ReplaceTemplateAccepted{}
}

/*ReplaceTemplateAccepted handles this case with default header values.

ReplaceTemplateAccepted replace template accepted
*/
type ReplaceTemplateAccepted struct {
	Payload *models.TemplateOutput
}

func (o *ReplaceTemplateAccepted) Error() string {
	return fmt.Sprintf("[POST /templates/{uuid}][%d] replaceTemplateAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceTemplateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TemplateOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceTemplateUnauthorized creates a ReplaceTemplateUnauthorized with default headers values
func NewReplaceTemplateUnauthorized() *ReplaceTemplateUnauthorized {
	return &ReplaceTemplateUnauthorized{}
}

/*ReplaceTemplateUnauthorized handles this case with default header values.

ReplaceTemplateUnauthorized replace template unauthorized
*/
type ReplaceTemplateUnauthorized struct {
	Payload *models.Result
}

func (o *ReplaceTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /templates/{uuid}][%d] replaceTemplateUnauthorized  %+v", 401, o.Payload)
}

func (o *ReplaceTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceTemplateConflict creates a ReplaceTemplateConflict with default headers values
func NewReplaceTemplateConflict() *ReplaceTemplateConflict {
	return &ReplaceTemplateConflict{}
}

/*ReplaceTemplateConflict handles this case with default header values.

ReplaceTemplateConflict replace template conflict
*/
type ReplaceTemplateConflict struct {
	Payload *models.Result
}

func (o *ReplaceTemplateConflict) Error() string {
	return fmt.Sprintf("[POST /templates/{uuid}][%d] replaceTemplateConflict  %+v", 409, o.Payload)
}

func (o *ReplaceTemplateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceTemplateExpectationFailed creates a ReplaceTemplateExpectationFailed with default headers values
func NewReplaceTemplateExpectationFailed() *ReplaceTemplateExpectationFailed {
	return &ReplaceTemplateExpectationFailed{}
}

/*ReplaceTemplateExpectationFailed handles this case with default header values.

ReplaceTemplateExpectationFailed replace template expectation failed
*/
type ReplaceTemplateExpectationFailed struct {
	Payload *models.Result
}

func (o *ReplaceTemplateExpectationFailed) Error() string {
	return fmt.Sprintf("[POST /templates/{uuid}][%d] replaceTemplateExpectationFailed  %+v", 417, o.Payload)
}

func (o *ReplaceTemplateExpectationFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceTemplateInternalServerError creates a ReplaceTemplateInternalServerError with default headers values
func NewReplaceTemplateInternalServerError() *ReplaceTemplateInternalServerError {
	return &ReplaceTemplateInternalServerError{}
}

/*ReplaceTemplateInternalServerError handles this case with default header values.

ReplaceTemplateInternalServerError replace template internal server error
*/
type ReplaceTemplateInternalServerError struct {
	Payload *models.Result
}

func (o *ReplaceTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /templates/{uuid}][%d] replaceTemplateInternalServerError  %+v", 500, o.Payload)
}

func (o *ReplaceTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}