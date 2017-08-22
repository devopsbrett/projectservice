package store

import (
	"git.home.foxienet.com/hostnotes/projectservice/gen/models"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-openapi/strfmt"
	r "gopkg.in/gorethink/gorethink.v3"
)

type Store interface {
	FetchAll() ([]*models.Project, error)
	Add(*models.Project) (*models.Project, error)
	Delete(string) error
}

type projectsStore struct {
	session *r.Session
	table   r.Term
}

func NewRethinkDBStore(config interface{}) (Store, error) {
	r.SetTags("gorethink", "json")
	c := config.(map[string]interface{})
	session, err := r.Connect(r.ConnectOpts{
		Address:  c["host"].(string),
		Database: "hostnotes",
	})
	if err != nil {
		return nil, err
	}
	s := &projectsStore{
		session: session,
		table:   r.Table("projects"),
	}

	spew.Dump(s)
	return s, nil
}

func (p *projectsStore) FetchAll() ([]*models.Project, error) {
	var projects []*models.Project
	cur, err := p.table.Run(p.session)
	if err != nil {
		return projects, err
	}
	cur.All(&projects)
	return projects, nil
}

func (p *projectsStore) Add(proj *models.Project) (*models.Project, error) {
	resp, err := p.table.Insert(proj).RunWrite(p.session)
	if err != nil {
		return nil, err
	}
	var id strfmt.UUID
	id.Scan(resp.GeneratedKeys[0])
	proj.UUID = id
	spew.Dump(proj)
	return proj, err
}

func (p *projectsStore) Delete(id string) error {
	err := p.table.Get(id).Delete().Exec(p.session)
	return err
}
