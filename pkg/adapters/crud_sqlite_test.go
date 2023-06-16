// Package adapters are the glue between components and external sources.
// # This manifest was generated by ymir. DO NOT EDIT.
package adapters

import (
	"testing"

	sqlEnt "entgo.io/ent/dialect/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/kubuskotak/king/pkg/infrastructure"
)

func TestWithCrudSQLite(t *testing.T) {
	is := assert.New(t)

	db, mock, err := sqlmock.New()
	if err != nil {
		is.Failf("failed to open stub db", "%v", err)
	}

	is.NotNil(db, "mock db is null")
	is.NotNil(mock, "sqlmock is null")

	CrudSQLiteOpen = func(dialect, source string) (*sqlEnt.Driver, error) {
		return sqlEnt.NewDriver(dialect, sqlEnt.Conn{ExecQuerier: db}), nil
	}

	infrastructure.Configuration(
		infrastructure.WithPath("../.."),
		infrastructure.WithFilename("config.yaml"),
	).Initialize()

	adapter := &Adapter{}
	adapter.Sync(
		WithCrudSQLite(&CrudSQLite{
			File: infrastructure.Envs.CrudSQLite.File,
		}),
	)

	mock.ExpectClose()

	// Asserts
	is.Nil(adapter.UnSync())
	is.Nil(mock.ExpectationsWereMet())
}
