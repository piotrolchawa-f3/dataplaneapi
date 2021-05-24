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

package backend_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v2/models"
)

// DeleteBackendSwitchingRuleAcceptedCode is the HTTP code returned for type DeleteBackendSwitchingRuleAccepted
const DeleteBackendSwitchingRuleAcceptedCode int = 202

/*DeleteBackendSwitchingRuleAccepted Configuration change accepted and reload requested

swagger:response deleteBackendSwitchingRuleAccepted
*/
type DeleteBackendSwitchingRuleAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteBackendSwitchingRuleAccepted creates DeleteBackendSwitchingRuleAccepted with default headers values
func NewDeleteBackendSwitchingRuleAccepted() *DeleteBackendSwitchingRuleAccepted {

	return &DeleteBackendSwitchingRuleAccepted{}
}

// WithReloadID adds the reloadId to the delete backend switching rule accepted response
func (o *DeleteBackendSwitchingRuleAccepted) WithReloadID(reloadID string) *DeleteBackendSwitchingRuleAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete backend switching rule accepted response
func (o *DeleteBackendSwitchingRuleAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteBackendSwitchingRuleAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteBackendSwitchingRuleNoContentCode is the HTTP code returned for type DeleteBackendSwitchingRuleNoContent
const DeleteBackendSwitchingRuleNoContentCode int = 204

/*DeleteBackendSwitchingRuleNoContent Backend Switching Rule deleted

swagger:response deleteBackendSwitchingRuleNoContent
*/
type DeleteBackendSwitchingRuleNoContent struct {
}

// NewDeleteBackendSwitchingRuleNoContent creates DeleteBackendSwitchingRuleNoContent with default headers values
func NewDeleteBackendSwitchingRuleNoContent() *DeleteBackendSwitchingRuleNoContent {

	return &DeleteBackendSwitchingRuleNoContent{}
}

// WriteResponse to the client
func (o *DeleteBackendSwitchingRuleNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteBackendSwitchingRuleNotFoundCode is the HTTP code returned for type DeleteBackendSwitchingRuleNotFound
const DeleteBackendSwitchingRuleNotFoundCode int = 404

/*DeleteBackendSwitchingRuleNotFound The specified resource was not found

swagger:response deleteBackendSwitchingRuleNotFound
*/
type DeleteBackendSwitchingRuleNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteBackendSwitchingRuleNotFound creates DeleteBackendSwitchingRuleNotFound with default headers values
func NewDeleteBackendSwitchingRuleNotFound() *DeleteBackendSwitchingRuleNotFound {

	return &DeleteBackendSwitchingRuleNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete backend switching rule not found response
func (o *DeleteBackendSwitchingRuleNotFound) WithConfigurationVersion(configurationVersion string) *DeleteBackendSwitchingRuleNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete backend switching rule not found response
func (o *DeleteBackendSwitchingRuleNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete backend switching rule not found response
func (o *DeleteBackendSwitchingRuleNotFound) WithPayload(payload *models.Error) *DeleteBackendSwitchingRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete backend switching rule not found response
func (o *DeleteBackendSwitchingRuleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBackendSwitchingRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*DeleteBackendSwitchingRuleDefault General Error

swagger:response deleteBackendSwitchingRuleDefault
*/
type DeleteBackendSwitchingRuleDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteBackendSwitchingRuleDefault creates DeleteBackendSwitchingRuleDefault with default headers values
func NewDeleteBackendSwitchingRuleDefault(code int) *DeleteBackendSwitchingRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteBackendSwitchingRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete backend switching rule default response
func (o *DeleteBackendSwitchingRuleDefault) WithStatusCode(code int) *DeleteBackendSwitchingRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete backend switching rule default response
func (o *DeleteBackendSwitchingRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete backend switching rule default response
func (o *DeleteBackendSwitchingRuleDefault) WithConfigurationVersion(configurationVersion string) *DeleteBackendSwitchingRuleDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete backend switching rule default response
func (o *DeleteBackendSwitchingRuleDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete backend switching rule default response
func (o *DeleteBackendSwitchingRuleDefault) WithPayload(payload *models.Error) *DeleteBackendSwitchingRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete backend switching rule default response
func (o *DeleteBackendSwitchingRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBackendSwitchingRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
