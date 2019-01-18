// Code generated by go-swagger; DO NOT EDIT.

package listener

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// DeleteListenerNoContentCode is the HTTP code returned for type DeleteListenerNoContent
const DeleteListenerNoContentCode int = 204

/*DeleteListenerNoContent Listener deleted

swagger:response deleteListenerNoContent
*/
type DeleteListenerNoContent struct {
}

// NewDeleteListenerNoContent creates DeleteListenerNoContent with default headers values
func NewDeleteListenerNoContent() *DeleteListenerNoContent {

	return &DeleteListenerNoContent{}
}

// WriteResponse to the client
func (o *DeleteListenerNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteListenerNotFoundCode is the HTTP code returned for type DeleteListenerNotFound
const DeleteListenerNotFoundCode int = 404

/*DeleteListenerNotFound The specified resource was not found

swagger:response deleteListenerNotFound
*/
type DeleteListenerNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteListenerNotFound creates DeleteListenerNotFound with default headers values
func NewDeleteListenerNotFound() *DeleteListenerNotFound {

	return &DeleteListenerNotFound{}
}

// WithPayload adds the payload to the delete listener not found response
func (o *DeleteListenerNotFound) WithPayload(payload *models.Error) *DeleteListenerNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete listener not found response
func (o *DeleteListenerNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteListenerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteListenerDefault General Error

swagger:response deleteListenerDefault
*/
type DeleteListenerDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteListenerDefault creates DeleteListenerDefault with default headers values
func NewDeleteListenerDefault(code int) *DeleteListenerDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteListenerDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete listener default response
func (o *DeleteListenerDefault) WithStatusCode(code int) *DeleteListenerDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete listener default response
func (o *DeleteListenerDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete listener default response
func (o *DeleteListenerDefault) WithPayload(payload *models.Error) *DeleteListenerDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete listener default response
func (o *DeleteListenerDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteListenerDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}