// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/massalabs/thyra/api/swagger/server/models"
)

// PluginManagerExecuteCommandNoContentCode is the HTTP code returned for type PluginManagerExecuteCommandNoContent
const PluginManagerExecuteCommandNoContentCode int = 204

/*
PluginManagerExecuteCommandNoContent Command successfuly executed.

swagger:response pluginManagerExecuteCommandNoContent
*/
type PluginManagerExecuteCommandNoContent struct {
}

// NewPluginManagerExecuteCommandNoContent creates PluginManagerExecuteCommandNoContent with default headers values
func NewPluginManagerExecuteCommandNoContent() *PluginManagerExecuteCommandNoContent {

	return &PluginManagerExecuteCommandNoContent{}
}

// WriteResponse to the client
func (o *PluginManagerExecuteCommandNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PluginManagerExecuteCommandBadRequestCode is the HTTP code returned for type PluginManagerExecuteCommandBadRequest
const PluginManagerExecuteCommandBadRequestCode int = 400

/*
PluginManagerExecuteCommandBadRequest Bad request.

swagger:response pluginManagerExecuteCommandBadRequest
*/
type PluginManagerExecuteCommandBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPluginManagerExecuteCommandBadRequest creates PluginManagerExecuteCommandBadRequest with default headers values
func NewPluginManagerExecuteCommandBadRequest() *PluginManagerExecuteCommandBadRequest {

	return &PluginManagerExecuteCommandBadRequest{}
}

// WithPayload adds the payload to the plugin manager execute command bad request response
func (o *PluginManagerExecuteCommandBadRequest) WithPayload(payload *models.Error) *PluginManagerExecuteCommandBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the plugin manager execute command bad request response
func (o *PluginManagerExecuteCommandBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PluginManagerExecuteCommandBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PluginManagerExecuteCommandNotFoundCode is the HTTP code returned for type PluginManagerExecuteCommandNotFound
const PluginManagerExecuteCommandNotFoundCode int = 404

/*
PluginManagerExecuteCommandNotFound Not found.

swagger:response pluginManagerExecuteCommandNotFound
*/
type PluginManagerExecuteCommandNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPluginManagerExecuteCommandNotFound creates PluginManagerExecuteCommandNotFound with default headers values
func NewPluginManagerExecuteCommandNotFound() *PluginManagerExecuteCommandNotFound {

	return &PluginManagerExecuteCommandNotFound{}
}

// WithPayload adds the payload to the plugin manager execute command not found response
func (o *PluginManagerExecuteCommandNotFound) WithPayload(payload *models.Error) *PluginManagerExecuteCommandNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the plugin manager execute command not found response
func (o *PluginManagerExecuteCommandNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PluginManagerExecuteCommandNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PluginManagerExecuteCommandUnprocessableEntityCode is the HTTP code returned for type PluginManagerExecuteCommandUnprocessableEntity
const PluginManagerExecuteCommandUnprocessableEntityCode int = 422

/*
PluginManagerExecuteCommandUnprocessableEntity Unprocessable Entity - The syntax is correct, but the server was unable to process the contained instructions.

swagger:response pluginManagerExecuteCommandUnprocessableEntity
*/
type PluginManagerExecuteCommandUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPluginManagerExecuteCommandUnprocessableEntity creates PluginManagerExecuteCommandUnprocessableEntity with default headers values
func NewPluginManagerExecuteCommandUnprocessableEntity() *PluginManagerExecuteCommandUnprocessableEntity {

	return &PluginManagerExecuteCommandUnprocessableEntity{}
}

// WithPayload adds the payload to the plugin manager execute command unprocessable entity response
func (o *PluginManagerExecuteCommandUnprocessableEntity) WithPayload(payload *models.Error) *PluginManagerExecuteCommandUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the plugin manager execute command unprocessable entity response
func (o *PluginManagerExecuteCommandUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PluginManagerExecuteCommandUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PluginManagerExecuteCommandInternalServerErrorCode is the HTTP code returned for type PluginManagerExecuteCommandInternalServerError
const PluginManagerExecuteCommandInternalServerErrorCode int = 500

/*
PluginManagerExecuteCommandInternalServerError Internal Server Error - The server has encountered a situation it does not know how to handle.

swagger:response pluginManagerExecuteCommandInternalServerError
*/
type PluginManagerExecuteCommandInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPluginManagerExecuteCommandInternalServerError creates PluginManagerExecuteCommandInternalServerError with default headers values
func NewPluginManagerExecuteCommandInternalServerError() *PluginManagerExecuteCommandInternalServerError {

	return &PluginManagerExecuteCommandInternalServerError{}
}

// WithPayload adds the payload to the plugin manager execute command internal server error response
func (o *PluginManagerExecuteCommandInternalServerError) WithPayload(payload *models.Error) *PluginManagerExecuteCommandInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the plugin manager execute command internal server error response
func (o *PluginManagerExecuteCommandInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PluginManagerExecuteCommandInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PluginManagerExecuteCommandNotImplementedCode is the HTTP code returned for type PluginManagerExecuteCommandNotImplemented
const PluginManagerExecuteCommandNotImplementedCode int = 501

/*
PluginManagerExecuteCommandNotImplemented Not Implemented - the server does not support the functionality required to fulfill the request.

swagger:response pluginManagerExecuteCommandNotImplemented
*/
type PluginManagerExecuteCommandNotImplemented struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPluginManagerExecuteCommandNotImplemented creates PluginManagerExecuteCommandNotImplemented with default headers values
func NewPluginManagerExecuteCommandNotImplemented() *PluginManagerExecuteCommandNotImplemented {

	return &PluginManagerExecuteCommandNotImplemented{}
}

// WithPayload adds the payload to the plugin manager execute command not implemented response
func (o *PluginManagerExecuteCommandNotImplemented) WithPayload(payload *models.Error) *PluginManagerExecuteCommandNotImplemented {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the plugin manager execute command not implemented response
func (o *PluginManagerExecuteCommandNotImplemented) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PluginManagerExecuteCommandNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
