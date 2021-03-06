// Code generated by go-swagger; DO NOT EDIT.

package metadaservices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/freesky-edward/directory/models"
)

// ApisOKCode is the HTTP code returned for type ApisOK
const ApisOKCode int = 200

/*ApisOK List all APIs the directory has, the result contains a collection of items, each item represents an API metadata service.

swagger:response apisOK
*/
type ApisOK struct {

	/*
	  In: Body
	*/
	Payload *models.Directory `json:"body,omitempty"`
}

// NewApisOK creates ApisOK with default headers values
func NewApisOK() *ApisOK {

	return &ApisOK{}
}

// WithPayload adds the payload to the apis o k response
func (o *ApisOK) WithPayload(payload *models.Directory) *ApisOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the apis o k response
func (o *ApisOK) SetPayload(payload *models.Directory) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ApisOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
