package handlers

import (
	"net/http"

	"git.home.foxienet.com/hostnotes/projectservice"
	"git.home.foxienet.com/hostnotes/projectservice/gen/restapi/operations/projects"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-openapi/runtime/middleware"
)

// NewFindKeys handles a request for finding the known keys
func NewDestroyProject(rt *projectservice.Runtime) projects.DestroyOneHandler {
	return &destroyProject{rt: rt}
}

type destroyProject struct {
	rt *projectservice.Runtime
}

// Handle the find known keys request
func (d *destroyProject) Handle(params projects.DestroyOneParams) middleware.Responder {
	spew.Dump(params)
	err := d.rt.DB().Delete(params.UUID)
	if err != nil {
		projects.NewDestroyOneDefault(http.StatusInternalServerError).WithPayload(modelsError(err))
	}

	return projects.NewDestroyOneNoContent()
	// return projects.NewDestroyOneDefault(http.StatusInternalServerError).WithPayload(modelsError(fmt.Errorf("blah")))
	// p, err := d.rt.DB().FetchAll()
	// if err != nil {
	// 	return projects.NewGetProjectsDefault(http.StatusInternalServerError).WithPayload(modelsError(err))
	// }

	// return projects.NewGetProjectsOK().WithPayload(p)
}
