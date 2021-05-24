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

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v2/models"
)

// GetStorageEndpointsOKCode is the HTTP code returned for type GetStorageEndpointsOK
const GetStorageEndpointsOKCode int = 200

/*GetStorageEndpointsOK Success

swagger:response getStorageEndpointsOK
*/
type GetStorageEndpointsOK struct {

	/*
	  In: Body
	*/
	Payload models.Endpoints `json:"body,omitempty"`
}

// NewGetStorageEndpointsOK creates GetStorageEndpointsOK with default headers values
func NewGetStorageEndpointsOK() *GetStorageEndpointsOK {

	return &GetStorageEndpointsOK{}
}

// WithPayload adds the payload to the get storage endpoints o k response
func (o *GetStorageEndpointsOK) WithPayload(payload models.Endpoints) *GetStorageEndpointsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get storage endpoints o k response
func (o *GetStorageEndpointsOK) SetPayload(payload models.Endpoints) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStorageEndpointsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Endpoints{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetStorageEndpointsDefault General Error

swagger:response getStorageEndpointsDefault
*/
type GetStorageEndpointsDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetStorageEndpointsDefault creates GetStorageEndpointsDefault with default headers values
func NewGetStorageEndpointsDefault(code int) *GetStorageEndpointsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetStorageEndpointsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get storage endpoints default response
func (o *GetStorageEndpointsDefault) WithStatusCode(code int) *GetStorageEndpointsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get storage endpoints default response
func (o *GetStorageEndpointsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get storage endpoints default response
func (o *GetStorageEndpointsDefault) WithConfigurationVersion(configurationVersion string) *GetStorageEndpointsDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get storage endpoints default response
func (o *GetStorageEndpointsDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get storage endpoints default response
func (o *GetStorageEndpointsDefault) WithPayload(payload *models.Error) *GetStorageEndpointsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get storage endpoints default response
func (o *GetStorageEndpointsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStorageEndpointsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
