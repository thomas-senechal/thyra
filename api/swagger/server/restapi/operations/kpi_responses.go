// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// KpiOKCode is the HTTP code returned for type KpiOK
const KpiOKCode int = 200

/*KpiOK kpi message.

swagger:response kpiOK
*/
type KpiOK struct {

	/*
	  In: Body
	*/
	Payload *KpiOKBody `json:"body,omitempty"`
}

// NewKpiOK creates KpiOK with default headers values
func NewKpiOK() *KpiOK {

	return &KpiOK{}
}

// WithPayload adds the payload to the kpi o k response
func (o *KpiOK) WithPayload(payload *KpiOKBody) *KpiOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the kpi o k response
func (o *KpiOK) SetPayload(payload *KpiOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KpiOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}