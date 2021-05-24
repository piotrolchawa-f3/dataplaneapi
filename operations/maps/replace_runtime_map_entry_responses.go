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

package maps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v2/models"
)

// ReplaceRuntimeMapEntryOKCode is the HTTP code returned for type ReplaceRuntimeMapEntryOK
const ReplaceRuntimeMapEntryOKCode int = 200

/*ReplaceRuntimeMapEntryOK Map value replaced

swagger:response replaceRuntimeMapEntryOK
*/
type ReplaceRuntimeMapEntryOK struct {

	/*
	  In: Body
	*/
	Payload *models.MapEntry `json:"body,omitempty"`
}

// NewReplaceRuntimeMapEntryOK creates ReplaceRuntimeMapEntryOK with default headers values
func NewReplaceRuntimeMapEntryOK() *ReplaceRuntimeMapEntryOK {

	return &ReplaceRuntimeMapEntryOK{}
}

// WithPayload adds the payload to the replace runtime map entry o k response
func (o *ReplaceRuntimeMapEntryOK) WithPayload(payload *models.MapEntry) *ReplaceRuntimeMapEntryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace runtime map entry o k response
func (o *ReplaceRuntimeMapEntryOK) SetPayload(payload *models.MapEntry) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceRuntimeMapEntryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceRuntimeMapEntryBadRequestCode is the HTTP code returned for type ReplaceRuntimeMapEntryBadRequest
const ReplaceRuntimeMapEntryBadRequestCode int = 400

/*ReplaceRuntimeMapEntryBadRequest Bad request

swagger:response replaceRuntimeMapEntryBadRequest
*/
type ReplaceRuntimeMapEntryBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceRuntimeMapEntryBadRequest creates ReplaceRuntimeMapEntryBadRequest with default headers values
func NewReplaceRuntimeMapEntryBadRequest() *ReplaceRuntimeMapEntryBadRequest {

	return &ReplaceRuntimeMapEntryBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace runtime map entry bad request response
func (o *ReplaceRuntimeMapEntryBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceRuntimeMapEntryBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace runtime map entry bad request response
func (o *ReplaceRuntimeMapEntryBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace runtime map entry bad request response
func (o *ReplaceRuntimeMapEntryBadRequest) WithPayload(payload *models.Error) *ReplaceRuntimeMapEntryBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace runtime map entry bad request response
func (o *ReplaceRuntimeMapEntryBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceRuntimeMapEntryBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceRuntimeMapEntryNotFoundCode is the HTTP code returned for type ReplaceRuntimeMapEntryNotFound
const ReplaceRuntimeMapEntryNotFoundCode int = 404

/*ReplaceRuntimeMapEntryNotFound The specified resource was not found

swagger:response replaceRuntimeMapEntryNotFound
*/
type ReplaceRuntimeMapEntryNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceRuntimeMapEntryNotFound creates ReplaceRuntimeMapEntryNotFound with default headers values
func NewReplaceRuntimeMapEntryNotFound() *ReplaceRuntimeMapEntryNotFound {

	return &ReplaceRuntimeMapEntryNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the replace runtime map entry not found response
func (o *ReplaceRuntimeMapEntryNotFound) WithConfigurationVersion(configurationVersion string) *ReplaceRuntimeMapEntryNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace runtime map entry not found response
func (o *ReplaceRuntimeMapEntryNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace runtime map entry not found response
func (o *ReplaceRuntimeMapEntryNotFound) WithPayload(payload *models.Error) *ReplaceRuntimeMapEntryNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace runtime map entry not found response
func (o *ReplaceRuntimeMapEntryNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceRuntimeMapEntryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*ReplaceRuntimeMapEntryDefault General Error

swagger:response replaceRuntimeMapEntryDefault
*/
type ReplaceRuntimeMapEntryDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceRuntimeMapEntryDefault creates ReplaceRuntimeMapEntryDefault with default headers values
func NewReplaceRuntimeMapEntryDefault(code int) *ReplaceRuntimeMapEntryDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceRuntimeMapEntryDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace runtime map entry default response
func (o *ReplaceRuntimeMapEntryDefault) WithStatusCode(code int) *ReplaceRuntimeMapEntryDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace runtime map entry default response
func (o *ReplaceRuntimeMapEntryDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace runtime map entry default response
func (o *ReplaceRuntimeMapEntryDefault) WithConfigurationVersion(configurationVersion string) *ReplaceRuntimeMapEntryDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace runtime map entry default response
func (o *ReplaceRuntimeMapEntryDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace runtime map entry default response
func (o *ReplaceRuntimeMapEntryDefault) WithPayload(payload *models.Error) *ReplaceRuntimeMapEntryDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace runtime map entry default response
func (o *ReplaceRuntimeMapEntryDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceRuntimeMapEntryDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
