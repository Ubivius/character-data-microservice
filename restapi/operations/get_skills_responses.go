// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"character-data-microservice/models"
)

// GetSkillsOKCode is the HTTP code returned for type GetSkillsOK
const GetSkillsOKCode int = 200

/*GetSkillsOK Returns list of skills

swagger:response getSkillsOK
*/
type GetSkillsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Skill `json:"body,omitempty"`
}

// NewGetSkillsOK creates GetSkillsOK with default headers values
func NewGetSkillsOK() *GetSkillsOK {

	return &GetSkillsOK{}
}

// WithPayload adds the payload to the get skills o k response
func (o *GetSkillsOK) WithPayload(payload []*models.Skill) *GetSkillsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get skills o k response
func (o *GetSkillsOK) SetPayload(payload []*models.Skill) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSkillsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Skill, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetSkillsDefault Generic error response

swagger:response getSkillsDefault
*/
type GetSkillsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSkillsDefault creates GetSkillsDefault with default headers values
func NewGetSkillsDefault(code int) *GetSkillsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSkillsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get skills default response
func (o *GetSkillsDefault) WithStatusCode(code int) *GetSkillsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get skills default response
func (o *GetSkillsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get skills default response
func (o *GetSkillsDefault) WithPayload(payload *models.Error) *GetSkillsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get skills default response
func (o *GetSkillsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSkillsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}