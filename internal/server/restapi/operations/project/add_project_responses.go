// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"ta4/mod/internal/server/models"
)

// AddProjectCreatedCode is the HTTP code returned for type AddProjectCreated
const AddProjectCreatedCode int = 201

/*
AddProjectCreated Project created

swagger:response addProjectCreated
*/
type AddProjectCreated struct {

	/*
	  In: Body
	*/
	Payload *models.ProjectResponse `json:"body,omitempty"`
}

// NewAddProjectCreated creates AddProjectCreated with default headers values
func NewAddProjectCreated() *AddProjectCreated {

	return &AddProjectCreated{}
}

// WithPayload adds the payload to the add project created response
func (o *AddProjectCreated) WithPayload(payload *models.ProjectResponse) *AddProjectCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add project created response
func (o *AddProjectCreated) SetPayload(payload *models.ProjectResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddProjectCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddProjectMethodNotAllowedCode is the HTTP code returned for type AddProjectMethodNotAllowed
const AddProjectMethodNotAllowedCode int = 405

/*
AddProjectMethodNotAllowed Invalid input

swagger:response addProjectMethodNotAllowed
*/
type AddProjectMethodNotAllowed struct {
}

// NewAddProjectMethodNotAllowed creates AddProjectMethodNotAllowed with default headers values
func NewAddProjectMethodNotAllowed() *AddProjectMethodNotAllowed {

	return &AddProjectMethodNotAllowed{}
}

// WriteResponse to the client
func (o *AddProjectMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}
