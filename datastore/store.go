package datastore

import (
	"database/sql"
	"errors"

	"github.com/anthonynsimon/parrot/datastore/postgres"
	"github.com/anthonynsimon/parrot/model"
)

var (
	ErrNoDB           = errors.New("couldn't get DB")
	ErrNotImplemented = errors.New("database not implemented")
)

type Datastore struct {
	Store
}

type Store interface {
	model.DocStorer
	Ping() error
	Close() error
}

func NewDatastore(name string, url string) (*Datastore, error) {
	var ds *Datastore

	switch name {
	case "postgres":
		conn, err := sql.Open("postgres", url)
		if err != nil {
			return nil, err
		}
		p := &postgres.PostgresDB{DB: conn}
		ds = &Datastore{p}
	default:
		return nil, ErrNotImplemented
	}

	return ds, nil
}
