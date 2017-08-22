package main

import (
	"os"

	"git.home.foxienet.com/hostnotes/projectservice"
	"git.home.foxienet.com/hostnotes/projectservice/api/handlers"
	apimiddleware "git.home.foxienet.com/hostnotes/projectservice/api/middleware"
	"git.home.foxienet.com/hostnotes/projectservice/gen/restapi"
	"git.home.foxienet.com/hostnotes/projectservice/gen/restapi/operations"

	"github.com/Sirupsen/logrus"
	app "github.com/casualjim/go-app"
	"github.com/casualjim/middlewares"
	"github.com/davecgh/go-spew/spew"
	loads "github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
	"github.com/justinas/alice"
)

func main() {
	app, err := app.New("projectservice")
	if err != nil {
		logrus.Fatalln(err)
	}
	rt, err := projectservice.NewRuntime(app)
	log := app.Logger()
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(rt.Config())
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}
	api := operations.NewProjectAPI(swaggerSpec)
	server := restapi.NewServer(api)
	spew.Dump(swaggerSpec)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Create ans Display available projects"
	parser.LongDescription = "Manage Projects"

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	api.ProjectsGetProjectsHandler = handlers.NewGetProjects(rt)
	api.ProjectsAddOneHandler = handlers.NewAddProject(rt)
	api.ProjectsDestroyOneHandler = handlers.NewDestroyProject(rt)

	handler := alice.New(
		middlewares.GzipMW(middlewares.DefaultCompression),
		middlewares.NewRecoveryMW(app.Info().Name, log),
		middlewares.NewAuditMW(app.Info(), log),
		middlewares.NewProfiler,
		middlewares.NewHealthChecksMW(app.Info().BasePath),
	).Then(api.Serve(nil))

	corsHandle := apimiddleware.AccessControl(handler)
	server.SetHandler(corsHandle)
	server.Port = app.Config().GetInt("ports.http")

	spew.Dump(api)
	// server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
