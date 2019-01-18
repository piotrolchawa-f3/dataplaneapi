// Code generated by go-swagger; DO NOT EDIT.

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// GetConfigurationEndpointsOKCode is the HTTP code returned for type GetConfigurationEndpointsOK
const GetConfigurationEndpointsOKCode int = 200

/*GetConfigurationEndpointsOK Success

swagger:response getConfigurationEndpointsOK
*/
type GetConfigurationEndpointsOK struct {

	/*
	  In: Body
	*/
	Payload models.Endpoints `json:"body,omitempty"`
}

// NewGetConfigurationEndpointsOK creates GetConfigurationEndpointsOK with default headers values
func NewGetConfigurationEndpointsOK() *GetConfigurationEndpointsOK {

	return &GetConfigurationEndpointsOK{}
}

// WithPayload adds the payload to the get configuration endpoints o k response
func (o *GetConfigurationEndpointsOK) WithPayload(payload models.Endpoints) *GetConfigurationEndpointsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get configuration endpoints o k response
func (o *GetConfigurationEndpointsOK) SetPayload(payload models.Endpoints) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigurationEndpointsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.Endpoints, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetConfigurationEndpointsDefault General Error

swagger:response getConfigurationEndpointsDefault
*/
type GetConfigurationEndpointsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigurationEndpointsDefault creates GetConfigurationEndpointsDefault with default headers values
func NewGetConfigurationEndpointsDefault(code int) *GetConfigurationEndpointsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetConfigurationEndpointsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get configuration endpoints default response
func (o *GetConfigurationEndpointsDefault) WithStatusCode(code int) *GetConfigurationEndpointsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get configuration endpoints default response
func (o *GetConfigurationEndpointsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get configuration endpoints default response
func (o *GetConfigurationEndpointsDefault) WithPayload(payload *models.Error) *GetConfigurationEndpointsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get configuration endpoints default response
func (o *GetConfigurationEndpointsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigurationEndpointsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}