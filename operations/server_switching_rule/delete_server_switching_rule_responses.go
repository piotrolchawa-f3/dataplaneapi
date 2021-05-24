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

package server_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v2/models"
)

// DeleteServerSwitchingRuleAcceptedCode is the HTTP code returned for type DeleteServerSwitchingRuleAccepted
const DeleteServerSwitchingRuleAcceptedCode int = 202

/*DeleteServerSwitchingRuleAccepted Configuration change accepted and reload requested

swagger:response deleteServerSwitchingRuleAccepted
*/
type DeleteServerSwitchingRuleAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteServerSwitchingRuleAccepted creates DeleteServerSwitchingRuleAccepted with default headers values
func NewDeleteServerSwitchingRuleAccepted() *DeleteServerSwitchingRuleAccepted {

	return &DeleteServerSwitchingRuleAccepted{}
}

// WithReloadID adds the reloadId to the delete server switching rule accepted response
func (o *DeleteServerSwitchingRuleAccepted) WithReloadID(reloadID string) *DeleteServerSwitchingRuleAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete server switching rule accepted response
func (o *DeleteServerSwitchingRuleAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteServerSwitchingRuleAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteServerSwitchingRuleNoContentCode is the HTTP code returned for type DeleteServerSwitchingRuleNoContent
const DeleteServerSwitchingRuleNoContentCode int = 204

/*DeleteServerSwitchingRuleNoContent Server Switching Rule deleted

swagger:response deleteServerSwitchingRuleNoContent
*/
type DeleteServerSwitchingRuleNoContent struct {
}

// NewDeleteServerSwitchingRuleNoContent creates DeleteServerSwitchingRuleNoContent with default headers values
func NewDeleteServerSwitchingRuleNoContent() *DeleteServerSwitchingRuleNoContent {

	return &DeleteServerSwitchingRuleNoContent{}
}

// WriteResponse to the client
func (o *DeleteServerSwitchingRuleNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteServerSwitchingRuleNotFoundCode is the HTTP code returned for type DeleteServerSwitchingRuleNotFound
const DeleteServerSwitchingRuleNotFoundCode int = 404

/*DeleteServerSwitchingRuleNotFound The specified resource was not found

swagger:response deleteServerSwitchingRuleNotFound
*/
type DeleteServerSwitchingRuleNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteServerSwitchingRuleNotFound creates DeleteServerSwitchingRuleNotFound with default headers values
func NewDeleteServerSwitchingRuleNotFound() *DeleteServerSwitchingRuleNotFound {

	return &DeleteServerSwitchingRuleNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete server switching rule not found response
func (o *DeleteServerSwitchingRuleNotFound) WithConfigurationVersion(configurationVersion string) *DeleteServerSwitchingRuleNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete server switching rule not found response
func (o *DeleteServerSwitchingRuleNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete server switching rule not found response
func (o *DeleteServerSwitchingRuleNotFound) WithPayload(payload *models.Error) *DeleteServerSwitchingRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete server switching rule not found response
func (o *DeleteServerSwitchingRuleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServerSwitchingRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*DeleteServerSwitchingRuleDefault General Error

swagger:response deleteServerSwitchingRuleDefault
*/
type DeleteServerSwitchingRuleDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteServerSwitchingRuleDefault creates DeleteServerSwitchingRuleDefault with default headers values
func NewDeleteServerSwitchingRuleDefault(code int) *DeleteServerSwitchingRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteServerSwitchingRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete server switching rule default response
func (o *DeleteServerSwitchingRuleDefault) WithStatusCode(code int) *DeleteServerSwitchingRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete server switching rule default response
func (o *DeleteServerSwitchingRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete server switching rule default response
func (o *DeleteServerSwitchingRuleDefault) WithConfigurationVersion(configurationVersion string) *DeleteServerSwitchingRuleDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete server switching rule default response
func (o *DeleteServerSwitchingRuleDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete server switching rule default response
func (o *DeleteServerSwitchingRuleDefault) WithPayload(payload *models.Error) *DeleteServerSwitchingRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete server switching rule default response
func (o *DeleteServerSwitchingRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServerSwitchingRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
