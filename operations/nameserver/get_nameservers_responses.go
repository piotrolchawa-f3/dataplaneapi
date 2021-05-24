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

package nameserver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v2/models"
)

// GetNameserversOKCode is the HTTP code returned for type GetNameserversOK
const GetNameserversOKCode int = 200

/*GetNameserversOK Successful operation

swagger:response getNameserversOK
*/
type GetNameserversOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetNameserversOKBody `json:"body,omitempty"`
}

// NewGetNameserversOK creates GetNameserversOK with default headers values
func NewGetNameserversOK() *GetNameserversOK {

	return &GetNameserversOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get nameservers o k response
func (o *GetNameserversOK) WithConfigurationVersion(configurationVersion string) *GetNameserversOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get nameservers o k response
func (o *GetNameserversOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get nameservers o k response
func (o *GetNameserversOK) WithPayload(payload *GetNameserversOKBody) *GetNameserversOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get nameservers o k response
func (o *GetNameserversOK) SetPayload(payload *GetNameserversOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNameserversOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetNameserversDefault General Error

swagger:response getNameserversDefault
*/
type GetNameserversDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNameserversDefault creates GetNameserversDefault with default headers values
func NewGetNameserversDefault(code int) *GetNameserversDefault {
	if code <= 0 {
		code = 500
	}

	return &GetNameserversDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get nameservers default response
func (o *GetNameserversDefault) WithStatusCode(code int) *GetNameserversDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get nameservers default response
func (o *GetNameserversDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get nameservers default response
func (o *GetNameserversDefault) WithConfigurationVersion(configurationVersion string) *GetNameserversDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get nameservers default response
func (o *GetNameserversDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get nameservers default response
func (o *GetNameserversDefault) WithPayload(payload *models.Error) *GetNameserversDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get nameservers default response
func (o *GetNameserversDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNameserversDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
