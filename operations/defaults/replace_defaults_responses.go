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

package defaults

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v2/models"
)

// ReplaceDefaultsOKCode is the HTTP code returned for type ReplaceDefaultsOK
const ReplaceDefaultsOKCode int = 200

/*ReplaceDefaultsOK Defaults replaced

swagger:response replaceDefaultsOK
*/
type ReplaceDefaultsOK struct {

	/*
	  In: Body
	*/
	Payload *models.Defaults `json:"body,omitempty"`
}

// NewReplaceDefaultsOK creates ReplaceDefaultsOK with default headers values
func NewReplaceDefaultsOK() *ReplaceDefaultsOK {

	return &ReplaceDefaultsOK{}
}

// WithPayload adds the payload to the replace defaults o k response
func (o *ReplaceDefaultsOK) WithPayload(payload *models.Defaults) *ReplaceDefaultsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace defaults o k response
func (o *ReplaceDefaultsOK) SetPayload(payload *models.Defaults) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceDefaultsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceDefaultsAcceptedCode is the HTTP code returned for type ReplaceDefaultsAccepted
const ReplaceDefaultsAcceptedCode int = 202

/*ReplaceDefaultsAccepted Configuration change accepted and reload requested

swagger:response replaceDefaultsAccepted
*/
type ReplaceDefaultsAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Defaults `json:"body,omitempty"`
}

// NewReplaceDefaultsAccepted creates ReplaceDefaultsAccepted with default headers values
func NewReplaceDefaultsAccepted() *ReplaceDefaultsAccepted {

	return &ReplaceDefaultsAccepted{}
}

// WithReloadID adds the reloadId to the replace defaults accepted response
func (o *ReplaceDefaultsAccepted) WithReloadID(reloadID string) *ReplaceDefaultsAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace defaults accepted response
func (o *ReplaceDefaultsAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace defaults accepted response
func (o *ReplaceDefaultsAccepted) WithPayload(payload *models.Defaults) *ReplaceDefaultsAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace defaults accepted response
func (o *ReplaceDefaultsAccepted) SetPayload(payload *models.Defaults) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceDefaultsAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplaceDefaultsBadRequestCode is the HTTP code returned for type ReplaceDefaultsBadRequest
const ReplaceDefaultsBadRequestCode int = 400

/*ReplaceDefaultsBadRequest Bad request

swagger:response replaceDefaultsBadRequest
*/
type ReplaceDefaultsBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceDefaultsBadRequest creates ReplaceDefaultsBadRequest with default headers values
func NewReplaceDefaultsBadRequest() *ReplaceDefaultsBadRequest {

	return &ReplaceDefaultsBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace defaults bad request response
func (o *ReplaceDefaultsBadRequest) WithConfigurationVersion(configurationVersion string) *ReplaceDefaultsBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace defaults bad request response
func (o *ReplaceDefaultsBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace defaults bad request response
func (o *ReplaceDefaultsBadRequest) WithPayload(payload *models.Error) *ReplaceDefaultsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace defaults bad request response
func (o *ReplaceDefaultsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceDefaultsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*ReplaceDefaultsDefault General Error

swagger:response replaceDefaultsDefault
*/
type ReplaceDefaultsDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceDefaultsDefault creates ReplaceDefaultsDefault with default headers values
func NewReplaceDefaultsDefault(code int) *ReplaceDefaultsDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceDefaultsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace defaults default response
func (o *ReplaceDefaultsDefault) WithStatusCode(code int) *ReplaceDefaultsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace defaults default response
func (o *ReplaceDefaultsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace defaults default response
func (o *ReplaceDefaultsDefault) WithConfigurationVersion(configurationVersion string) *ReplaceDefaultsDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace defaults default response
func (o *ReplaceDefaultsDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace defaults default response
func (o *ReplaceDefaultsDefault) WithPayload(payload *models.Error) *ReplaceDefaultsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace defaults default response
func (o *ReplaceDefaultsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceDefaultsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
