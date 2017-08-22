package handlers

import (
	"net/http"

	"git.home.foxienet.com/hostnotes/projectservice"
	"git.home.foxienet.com/hostnotes/projectservice/gen/models"
	"git.home.foxienet.com/hostnotes/projectservice/gen/restapi/operations/projects"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

// NewFindKeys handles a request for finding the known keys
func NewGetProjects(rt *projectservice.Runtime) projects.GetProjectsHandler {
	return &findProjects{rt: rt}
}

func modelsError(err error) *models.Error {
	return &models.Error{
		Message: swag.String(err.Error()),
	}
}

type findProjects struct {
	rt *projectservice.Runtime
}

// Handle the find known keys request
func (d *findProjects) Handle(params projects.GetProjectsParams) middleware.Responder {

	p, err := d.rt.DB().FetchAll()
	if err != nil {
		return projects.NewGetProjectsDefault(http.StatusInternalServerError).WithPayload(modelsError(err))
	}

	return projects.NewGetProjectsOK().WithPayload(p)
}
