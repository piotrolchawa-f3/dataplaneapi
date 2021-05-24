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

package peer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v2/models"
)

// DeletePeerAcceptedCode is the HTTP code returned for type DeletePeerAccepted
const DeletePeerAcceptedCode int = 202

/*DeletePeerAccepted Configuration change accepted and reload requested

swagger:response deletePeerAccepted
*/
type DeletePeerAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeletePeerAccepted creates DeletePeerAccepted with default headers values
func NewDeletePeerAccepted() *DeletePeerAccepted {

	return &DeletePeerAccepted{}
}

// WithReloadID adds the reloadId to the delete peer accepted response
func (o *DeletePeerAccepted) WithReloadID(reloadID string) *DeletePeerAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete peer accepted response
func (o *DeletePeerAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeletePeerAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeletePeerNoContentCode is the HTTP code returned for type DeletePeerNoContent
const DeletePeerNoContentCode int = 204

/*DeletePeerNoContent Peer deleted

swagger:response deletePeerNoContent
*/
type DeletePeerNoContent struct {
}

// NewDeletePeerNoContent creates DeletePeerNoContent with default headers values
func NewDeletePeerNoContent() *DeletePeerNoContent {

	return &DeletePeerNoContent{}
}

// WriteResponse to the client
func (o *DeletePeerNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeletePeerNotFoundCode is the HTTP code returned for type DeletePeerNotFound
const DeletePeerNotFoundCode int = 404

/*DeletePeerNotFound The specified resource was not found

swagger:response deletePeerNotFound
*/
type DeletePeerNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeletePeerNotFound creates DeletePeerNotFound with default headers values
func NewDeletePeerNotFound() *DeletePeerNotFound {

	return &DeletePeerNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete peer not found response
func (o *DeletePeerNotFound) WithConfigurationVersion(configurationVersion string) *DeletePeerNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete peer not found response
func (o *DeletePeerNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete peer not found response
func (o *DeletePeerNotFound) WithPayload(payload *models.Error) *DeletePeerNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete peer not found response
func (o *DeletePeerNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePeerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*DeletePeerDefault General Error

swagger:response deletePeerDefault
*/
type DeletePeerDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeletePeerDefault creates DeletePeerDefault with default headers values
func NewDeletePeerDefault(code int) *DeletePeerDefault {
	if code <= 0 {
		code = 500
	}

	return &DeletePeerDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete peer default response
func (o *DeletePeerDefault) WithStatusCode(code int) *DeletePeerDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete peer default response
func (o *DeletePeerDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete peer default response
func (o *DeletePeerDefault) WithConfigurationVersion(configurationVersion string) *DeletePeerDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete peer default response
func (o *DeletePeerDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete peer default response
func (o *DeletePeerDefault) WithPayload(payload *models.Error) *DeletePeerDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete peer default response
func (o *DeletePeerDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePeerDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
