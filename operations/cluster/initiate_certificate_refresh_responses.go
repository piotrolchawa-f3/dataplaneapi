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

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v2/models"
)

// InitiateCertificateRefreshOKCode is the HTTP code returned for type InitiateCertificateRefreshOK
const InitiateCertificateRefreshOKCode int = 200

/*InitiateCertificateRefreshOK refresh activated

swagger:response initiateCertificateRefreshOK
*/
type InitiateCertificateRefreshOK struct {
}

// NewInitiateCertificateRefreshOK creates InitiateCertificateRefreshOK with default headers values
func NewInitiateCertificateRefreshOK() *InitiateCertificateRefreshOK {

	return &InitiateCertificateRefreshOK{}
}

// WriteResponse to the client
func (o *InitiateCertificateRefreshOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// InitiateCertificateRefreshForbiddenCode is the HTTP code returned for type InitiateCertificateRefreshForbidden
const InitiateCertificateRefreshForbiddenCode int = 403

/*InitiateCertificateRefreshForbidden refresh not possible

swagger:response initiateCertificateRefreshForbidden
*/
type InitiateCertificateRefreshForbidden struct {
}

// NewInitiateCertificateRefreshForbidden creates InitiateCertificateRefreshForbidden with default headers values
func NewInitiateCertificateRefreshForbidden() *InitiateCertificateRefreshForbidden {

	return &InitiateCertificateRefreshForbidden{}
}

// WriteResponse to the client
func (o *InitiateCertificateRefreshForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

/*InitiateCertificateRefreshDefault General Error

swagger:response initiateCertificateRefreshDefault
*/
type InitiateCertificateRefreshDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewInitiateCertificateRefreshDefault creates InitiateCertificateRefreshDefault with default headers values
func NewInitiateCertificateRefreshDefault(code int) *InitiateCertificateRefreshDefault {
	if code <= 0 {
		code = 500
	}

	return &InitiateCertificateRefreshDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the initiate certificate refresh default response
func (o *InitiateCertificateRefreshDefault) WithStatusCode(code int) *InitiateCertificateRefreshDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the initiate certificate refresh default response
func (o *InitiateCertificateRefreshDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the initiate certificate refresh default response
func (o *InitiateCertificateRefreshDefault) WithConfigurationVersion(configurationVersion string) *InitiateCertificateRefreshDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the initiate certificate refresh default response
func (o *InitiateCertificateRefreshDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the initiate certificate refresh default response
func (o *InitiateCertificateRefreshDefault) WithPayload(payload *models.Error) *InitiateCertificateRefreshDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the initiate certificate refresh default response
func (o *InitiateCertificateRefreshDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *InitiateCertificateRefreshDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
