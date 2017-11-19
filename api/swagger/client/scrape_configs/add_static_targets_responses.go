// Code generated by go-swagger; DO NOT EDIT.

package scrape_configs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/percona/pmm-managed/api/swagger/models"
)

// AddStaticTargetsReader is a Reader for the AddStaticTargets structure.
type AddStaticTargetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddStaticTargetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddStaticTargetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddStaticTargetsOK creates a AddStaticTargetsOK with default headers values
func NewAddStaticTargetsOK() *AddStaticTargetsOK {
	return &AddStaticTargetsOK{}
}

/*AddStaticTargetsOK handles this case with default header values.

(empty)
*/
type AddStaticTargetsOK struct {
	Payload models.APIScrapeConfigsAddStaticTargetsResponse
}

func (o *AddStaticTargetsOK) Error() string {
	return fmt.Sprintf("[POST /v0/scrape-configs/{job_name}/static-targets][%d] addStaticTargetsOK  %+v", 200, o.Payload)
}

func (o *AddStaticTargetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}