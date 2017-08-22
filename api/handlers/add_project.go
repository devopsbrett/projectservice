package handlers

import (
	"net/http"

	"git.home.foxienet.com/hostnotes/projectservice"
	"git.home.foxienet.com/hostnotes/projectservice/gen/restapi/operations/projects"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-openapi/runtime/middleware"
)

// NewFindKeys handles a request for finding the known keys
func NewAddProject(rt *projectservice.Runtime) projects.AddOneHandler {
	return &newProject{rt: rt}
}

type newProject struct {
	rt *projectservice.Runtime
}

// Handle the find known keys request
func (d *newProject) Handle(params projects.AddOneParams) middleware.Responder {
	proj := params.Body

	p, err := d.rt.DB().Add(proj)
	if err != nil {
		return projects.NewAddOneDefault(http.StatusInternalServerError).WithPayload(modelsError(err))
	}

	// name := params.
	// p, err := d.rt.DB().FetchAll()
	// if err != nil {
	// 	return projects.NewGetProjectsDefault(http.StatusInternalServerError).WithPayload(modelsError(err))
	// }

	spew.Dump(p)
	return projects.NewAddOneCreated().WithPayload(p)

	// return middleware.NotImplemented("operation projects.AddOne has not yet been implemented")
}
