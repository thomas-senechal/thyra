// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/massalabs/thyra/api/swagger/server/models"
)

// BrowseOKCode is the HTTP code returned for type BrowseOK
const BrowseOKCode int = 200

/*BrowseOK Resource retrieved.

swagger:response browseOK
*/
type BrowseOK struct {
}

// NewBrowseOK creates BrowseOK with default headers values
func NewBrowseOK() *BrowseOK {

	return &BrowseOK{}
}

// WriteResponse to the client
func (o *BrowseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// BrowseBadRequestCode is the HTTP code returned for type BrowseBadRequest
const BrowseBadRequestCode int = 400

/*BrowseBadRequest Bad request.

swagger:response browseBadRequest
*/
type BrowseBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewBrowseBadRequest creates BrowseBadRequest with default headers values
func NewBrowseBadRequest() *BrowseBadRequest {

	return &BrowseBadRequest{}
}

// WithPayload adds the payload to the browse bad request response
func (o *BrowseBadRequest) WithPayload(payload *models.Error) *BrowseBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the browse bad request response
func (o *BrowseBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BrowseBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BrowseNotFoundCode is the HTTP code returned for type BrowseNotFound
const BrowseNotFoundCode int = 404

/*BrowseNotFound Resource not found.

swagger:response browseNotFound
*/
type BrowseNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewBrowseNotFound creates BrowseNotFound with default headers values
func NewBrowseNotFound() *BrowseNotFound {

	return &BrowseNotFound{}
}

// WithPayload adds the payload to the browse not found response
func (o *BrowseNotFound) WithPayload(payload *models.Error) *BrowseNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the browse not found response
func (o *BrowseNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BrowseNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BrowseInternalServerErrorCode is the HTTP code returned for type BrowseInternalServerError
const BrowseInternalServerErrorCode int = 500

/*BrowseInternalServerError Internal Server Error - The server has encountered a situation it does not know how to handle.

swagger:response browseInternalServerError
*/
type BrowseInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewBrowseInternalServerError creates BrowseInternalServerError with default headers values
func NewBrowseInternalServerError() *BrowseInternalServerError {

	return &BrowseInternalServerError{}
}

// WithPayload adds the payload to the browse internal server error response
func (o *BrowseInternalServerError) WithPayload(payload *models.Error) *BrowseInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the browse internal server error response
func (o *BrowseInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BrowseInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}