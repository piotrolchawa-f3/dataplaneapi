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

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v2/models"
)

// DeleteStorageMapNoContentCode is the HTTP code returned for type DeleteStorageMapNoContent
const DeleteStorageMapNoContentCode int = 204

/*DeleteStorageMapNoContent Map file deleted

swagger:response deleteStorageMapNoContent
*/
type DeleteStorageMapNoContent struct {
}

// NewDeleteStorageMapNoContent creates DeleteStorageMapNoContent with default headers values
func NewDeleteStorageMapNoContent() *DeleteStorageMapNoContent {

	return &DeleteStorageMapNoContent{}
}

// WriteResponse to the client
func (o *DeleteStorageMapNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteStorageMapNotFoundCode is the HTTP code returned for type DeleteStorageMapNotFound
const DeleteStorageMapNotFoundCode int = 404

/*DeleteStorageMapNotFound The specified resource was not found

swagger:response deleteStorageMapNotFound
*/
type DeleteStorageMapNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteStorageMapNotFound creates DeleteStorageMapNotFound with default headers values
func NewDeleteStorageMapNotFound() *DeleteStorageMapNotFound {

	return &DeleteStorageMapNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete storage map not found response
func (o *DeleteStorageMapNotFound) WithConfigurationVersion(configurationVersion string) *DeleteStorageMapNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete storage map not found response
func (o *DeleteStorageMapNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete storage map not found response
func (o *DeleteStorageMapNotFound) WithPayload(payload *models.Error) *DeleteStorageMapNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete storage map not found response
func (o *DeleteStorageMapNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteStorageMapNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*DeleteStorageMapDefault General Error

swagger:response deleteStorageMapDefault
*/
type DeleteStorageMapDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteStorageMapDefault creates DeleteStorageMapDefault with default headers values
func NewDeleteStorageMapDefault(code int) *DeleteStorageMapDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteStorageMapDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete storage map default response
func (o *DeleteStorageMapDefault) WithStatusCode(code int) *DeleteStorageMapDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete storage map default response
func (o *DeleteStorageMapDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete storage map default response
func (o *DeleteStorageMapDefault) WithConfigurationVersion(configurationVersion string) *DeleteStorageMapDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete storage map default response
func (o *DeleteStorageMapDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete storage map default response
func (o *DeleteStorageMapDefault) WithPayload(payload *models.Error) *DeleteStorageMapDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete storage map default response
func (o *DeleteStorageMapDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteStorageMapDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
