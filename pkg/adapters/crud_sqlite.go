// Package adapters are the glue between components and external sources.
// # This manifest was generated by ymir. DO NOT EDIT.
package adapters

import (
	"fmt"

	"entgo.io/ent/dialect"
	sqlEnt "entgo.io/ent/dialect/sql"
	"github.com/rs/zerolog/log"
	_ "modernc.org/sqlite"
)

var CrudSQLiteOpen = sqlEnt.Open // CrudSQLiteOpen will invoke to test case.

// CrudSQLite is data of instances.
type CrudSQLite struct {
	File   string `json:"file"`
	driver *sqlEnt.Driver
}

// Open is open the connection of sqlite.
func (c *CrudSQLite) Open() (*sqlEnt.Driver, error) {
	if c.driver == nil {
		return nil, fmt.Errorf("driver was failed to connected")
	}
	return c.driver, nil
}

// Connect is connected the connection of sqlite.
func (c *CrudSQLite) Connect() (err error) {
	c.driver, err = CrudSQLiteOpen(dialect.SQLite,
		c.File)
	if err != nil {
		log.Error().Err(err).Msg("CrudSQLiteOpen is failed to open")
		return err
	}
	pool := c.driver.DB()
	pool.SetMaxOpenConns(1)

	return nil
}

// Disconnect is disconnect the connection of sqlite.
func (c *CrudSQLite) Disconnect() error {
	return c.driver.Close()
}

// WithCrudSQLite option function to assign on adapters.
func WithCrudSQLite(driver Driver[*sqlEnt.Driver]) Option {
	return func(a *Adapter) {
		if err := driver.Connect(); err != nil {
			panic(err)
		}
		open, err := driver.Open()
		if err != nil {
			panic(err)
		}
		a.CrudSQLite = open
	}
}
