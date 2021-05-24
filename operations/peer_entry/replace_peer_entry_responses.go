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

package peer_entry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v2/models"
)

// ReplacePeerEntryOKCode is the HTTP code returned for type ReplacePeerEntryOK
const ReplacePeerEntryOKCode int = 200

/*ReplacePeerEntryOK PeerEntry replaced

swagger:response replacePeerEntryOK
*/
type ReplacePeerEntryOK struct {

	/*
	  In: Body
	*/
	Payload *models.PeerEntry `json:"body,omitempty"`
}

// NewReplacePeerEntryOK creates ReplacePeerEntryOK with default headers values
func NewReplacePeerEntryOK() *ReplacePeerEntryOK {

	return &ReplacePeerEntryOK{}
}

// WithPayload adds the payload to the replace peer entry o k response
func (o *ReplacePeerEntryOK) WithPayload(payload *models.PeerEntry) *ReplacePeerEntryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace peer entry o k response
func (o *ReplacePeerEntryOK) SetPayload(payload *models.PeerEntry) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplacePeerEntryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplacePeerEntryAcceptedCode is the HTTP code returned for type ReplacePeerEntryAccepted
const ReplacePeerEntryAcceptedCode int = 202

/*ReplacePeerEntryAccepted Configuration change accepted and reload requested

swagger:response replacePeerEntryAccepted
*/
type ReplacePeerEntryAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.PeerEntry `json:"body,omitempty"`
}

// NewReplacePeerEntryAccepted creates ReplacePeerEntryAccepted with default headers values
func NewReplacePeerEntryAccepted() *ReplacePeerEntryAccepted {

	return &ReplacePeerEntryAccepted{}
}

// WithReloadID adds the reloadId to the replace peer entry accepted response
func (o *ReplacePeerEntryAccepted) WithReloadID(reloadID string) *ReplacePeerEntryAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace peer entry accepted response
func (o *ReplacePeerEntryAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace peer entry accepted response
func (o *ReplacePeerEntryAccepted) WithPayload(payload *models.PeerEntry) *ReplacePeerEntryAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace peer entry accepted response
func (o *ReplacePeerEntryAccepted) SetPayload(payload *models.PeerEntry) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplacePeerEntryAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplacePeerEntryBadRequestCode is the HTTP code returned for type ReplacePeerEntryBadRequest
const ReplacePeerEntryBadRequestCode int = 400

/*ReplacePeerEntryBadRequest Bad request

swagger:response replacePeerEntryBadRequest
*/
type ReplacePeerEntryBadRequest struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplacePeerEntryBadRequest creates ReplacePeerEntryBadRequest with default headers values
func NewReplacePeerEntryBadRequest() *ReplacePeerEntryBadRequest {

	return &ReplacePeerEntryBadRequest{}
}

// WithConfigurationVersion adds the configurationVersion to the replace peer entry bad request response
func (o *ReplacePeerEntryBadRequest) WithConfigurationVersion(configurationVersion string) *ReplacePeerEntryBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace peer entry bad request response
func (o *ReplacePeerEntryBadRequest) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace peer entry bad request response
func (o *ReplacePeerEntryBadRequest) WithPayload(payload *models.Error) *ReplacePeerEntryBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace peer entry bad request response
func (o *ReplacePeerEntryBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplacePeerEntryBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ReplacePeerEntryNotFoundCode is the HTTP code returned for type ReplacePeerEntryNotFound
const ReplacePeerEntryNotFoundCode int = 404

/*ReplacePeerEntryNotFound The specified resource was not found

swagger:response replacePeerEntryNotFound
*/
type ReplacePeerEntryNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplacePeerEntryNotFound creates ReplacePeerEntryNotFound with default headers values
func NewReplacePeerEntryNotFound() *ReplacePeerEntryNotFound {

	return &ReplacePeerEntryNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the replace peer entry not found response
func (o *ReplacePeerEntryNotFound) WithConfigurationVersion(configurationVersion string) *ReplacePeerEntryNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace peer entry not found response
func (o *ReplacePeerEntryNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace peer entry not found response
func (o *ReplacePeerEntryNotFound) WithPayload(payload *models.Error) *ReplacePeerEntryNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace peer entry not found response
func (o *ReplacePeerEntryNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplacePeerEntryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*ReplacePeerEntryDefault General Error

swagger:response replacePeerEntryDefault
*/
type ReplacePeerEntryDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplacePeerEntryDefault creates ReplacePeerEntryDefault with default headers values
func NewReplacePeerEntryDefault(code int) *ReplacePeerEntryDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplacePeerEntryDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace peer entry default response
func (o *ReplacePeerEntryDefault) WithStatusCode(code int) *ReplacePeerEntryDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace peer entry default response
func (o *ReplacePeerEntryDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the replace peer entry default response
func (o *ReplacePeerEntryDefault) WithConfigurationVersion(configurationVersion string) *ReplacePeerEntryDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the replace peer entry default response
func (o *ReplacePeerEntryDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the replace peer entry default response
func (o *ReplacePeerEntryDefault) WithPayload(payload *models.Error) *ReplacePeerEntryDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace peer entry default response
func (o *ReplacePeerEntryDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplacePeerEntryDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
