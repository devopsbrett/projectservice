package projectservice

import (
	"git.home.foxienet.com/hostnotes/projectservice/store"
	app "github.com/casualjim/go-app"
	"github.com/spf13/viper"
)

func NewRuntime(app app.Application) (*Runtime, error) {
	db, err := store.NewRethinkDBStore(app.Config().Get("db"))
	if err != nil {
		return nil, err
	}
	return &Runtime{
		db:  db,
		app: app,
	}, nil
}

// Runtime encapsulates the shared services for this application
type Runtime struct {
	db  store.Store
	app app.Application
}

func (r *Runtime) Config() *viper.Viper {
	return r.app.Config()
}

func (r *Runtime) DB() store.Store {
	return r.db
}
