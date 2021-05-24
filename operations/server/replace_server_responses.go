// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v2/models"
)

// ReplaceServerOKCode is the HTTP code returned for type ReplaceServerOK
const ReplaceServerOKCode int = 200

/*ReplaceServerOK Server replaced

swagger:response replaceServerOK
*/
type ReplaceServerOK struct {

	/*
	  In: Body
	*/
	Payload *models.Server `json:"body,omitempty"`
}

// NewReplaceServerOK creates ReplaceServerOK with default headers values
func NewReplaceServerOK() *ReplaceServerOK {

	return &ReplaceServerOK{}
}

// WithPayload adds the payload to the replace server o k response
func (o *ReplaceServerOK) WithPayload(payload *models.Server) *ReplaceServerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server o k response
func (o *ReplaceServerOK) SetPayload(payload *models.Server) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceServerAcceptedCode is the HTTP code returned for type ReplaceServerAccepted
const ReplaceServerAcceptedCode int = 202

/*ReplaceServerAccepted Configuration change accepted and reload requested

swagger:response replaceServerAccepted
*/
type ReplaceServerAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Server `json:"body,omitempty"`
}

// NewReplaceServerAccepted creates ReplaceServerAccepted with default headers values
func NewReplaceServerAccepted() *ReplaceServerAccepted {

	return &ReplaceServerAccepted{}
}

// WithReloadID adds the reloadId to the replace server accepted response
func (o *ReplaceServerAccepted) WithReloadID(reloadID string) *ReplaceServerAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace server accepted response
func (o *ReplaceServerAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace server accepted response
func (o *ReplaceServerAccepted) WithPayload(payload *models.Server) *ReplaceServerAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server accepted response
func (o *ReplaceServerAccepted) SetPayload(payload *models.Server) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceServerBadRequestCode is the HTTP code returned for type ReplaceServerBadRequest
const ReplaceServerBadRequestCode int = 400

/*ReplaceServerBadRequest Bad request

swagger:response replaceServerBadRequest
*/
type ReplaceServerBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceServerBadRequest creates ReplaceServerBadRequest with default headers values
func NewReplaceServerBadRequest() *ReplaceServerBadRequest {

	return &ReplaceServerBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace server bad request response
func (o *ReplaceServerBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceServerBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace server bad request response
func (o *ReplaceServerBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace server bad request response
func (o *ReplaceServerBadRequest) WithPayload(payload *models.Error) *ReplaceServerBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server bad request response
func (o *ReplaceServerBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceServerNotFoundCode is the HTTP code returned for type ReplaceServerNotFound
const ReplaceServerNotFoundCode int = 404

/*ReplaceServerNotFound The specified resource was not found

swagger:response replaceServerNotFound
*/
type ReplaceServerNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceServerNotFound creates ReplaceServerNotFound with default headers values
func NewReplaceServerNotFound() *ReplaceServerNotFound {

	return &ReplaceServerNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the replace server not found response
func (o *ReplaceServerNotFound) WithConfigurationVersion(configurationVersion string) *ReplaceServerNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace server not found response
func (o *ReplaceServerNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace server not found response
func (o *ReplaceServerNotFound) WithPayload(payload *models.Error) *ReplaceServerNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server not found response
func (o *ReplaceServerNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ReplaceServerDefault General Error

swagger:response replaceServerDefault
*/
type ReplaceServerDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceServerDefault creates ReplaceServerDefault with default headers values
func NewReplaceServerDefault(code int) *ReplaceServerDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceServerDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace server default response
func (o *ReplaceServerDefault) WithStatusCode(code int) *ReplaceServerDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace server default response
func (o *ReplaceServerDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace server default response
func (o *ReplaceServerDefault) WithConfigurationVersion(configurationVersion string) *ReplaceServerDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace server default response
func (o *ReplaceServerDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace server default response
func (o *ReplaceServerDefault) WithPayload(payload *models.Error) *ReplaceServerDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace server default response
func (o *ReplaceServerDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceServerDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
