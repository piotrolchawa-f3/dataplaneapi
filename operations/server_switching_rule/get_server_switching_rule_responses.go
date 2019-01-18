// Code generated by go-swagger; DO NOT EDIT.

package server_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// GetServerSwitchingRuleOKCode is the HTTP code returned for type GetServerSwitchingRuleOK
const GetServerSwitchingRuleOKCode int = 200

/*GetServerSwitchingRuleOK Successful operation

swagger:response getServerSwitchingRuleOK
*/
type GetServerSwitchingRuleOK struct {

	/*
	  In: Body
	*/
	Payload *models.GetServerSwitchingRuleOKBody `json:"body,omitempty"`
}

// NewGetServerSwitchingRuleOK creates GetServerSwitchingRuleOK with default headers values
func NewGetServerSwitchingRuleOK() *GetServerSwitchingRuleOK {

	return &GetServerSwitchingRuleOK{}
}

// WithPayload adds the payload to the get server switching rule o k response
func (o *GetServerSwitchingRuleOK) WithPayload(payload *models.GetServerSwitchingRuleOKBody) *GetServerSwitchingRuleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get server switching rule o k response
func (o *GetServerSwitchingRuleOK) SetPayload(payload *models.GetServerSwitchingRuleOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServerSwitchingRuleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetServerSwitchingRuleNotFoundCode is the HTTP code returned for type GetServerSwitchingRuleNotFound
const GetServerSwitchingRuleNotFoundCode int = 404

/*GetServerSwitchingRuleNotFound The specified resource was not found

swagger:response getServerSwitchingRuleNotFound
*/
type GetServerSwitchingRuleNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetServerSwitchingRuleNotFound creates GetServerSwitchingRuleNotFound with default headers values
func NewGetServerSwitchingRuleNotFound() *GetServerSwitchingRuleNotFound {

	return &GetServerSwitchingRuleNotFound{}
}

// WithPayload adds the payload to the get server switching rule not found response
func (o *GetServerSwitchingRuleNotFound) WithPayload(payload *models.Error) *GetServerSwitchingRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get server switching rule not found response
func (o *GetServerSwitchingRuleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServerSwitchingRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetServerSwitchingRuleDefault General Error

swagger:response getServerSwitchingRuleDefault
*/
type GetServerSwitchingRuleDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetServerSwitchingRuleDefault creates GetServerSwitchingRuleDefault with default headers values
func NewGetServerSwitchingRuleDefault(code int) *GetServerSwitchingRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &GetServerSwitchingRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get server switching rule default response
func (o *GetServerSwitchingRuleDefault) WithStatusCode(code int) *GetServerSwitchingRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get server switching rule default response
func (o *GetServerSwitchingRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get server switching rule default response
func (o *GetServerSwitchingRuleDefault) WithPayload(payload *models.Error) *GetServerSwitchingRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get server switching rule default response
func (o *GetServerSwitchingRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServerSwitchingRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}