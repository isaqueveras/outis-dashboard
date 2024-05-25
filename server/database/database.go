package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/isaqueveras/outis-dashboard/server/config"
)

type database interface {
	transaction(ctx context.Context) (*sql.Tx, error)
	open(cfg *config.Database) error
	close()
}

var connection database

type Tx struct {
	postgres *sql.Tx
}

func OpenConnections(cfg *config.Database) {
	if connection != nil {
		return
	}

	connection = &postgres{}
	if err := connection.open(cfg); err != nil {
		log.Fatal("Unable to open connections to database: ", err)
	}
}

func CloseConnections() {
	connection.close()
}

func NewTx(ctx context.Context) (tx *Tx, err error) {
	tx = new(Tx)
	if tx.postgres, err = connection.transaction(ctx); err != nil {
		return nil, err
	}
	return tx, nil
}

func (t *Tx) Commit() (erro error) {
	return t.postgres.Commit()
}

func (t *Tx) Rollback() {
	_ = t.postgres.Rollback()
}

func (t *Tx) Query(query string, args ...any) (*sql.Rows, error) {
	return t.postgres.Query(query, args...)
}

func (t *Tx) QueryRow(query string, args ...any) *sql.Row {
	return t.postgres.QueryRow(query, args...)
}

func (t *Tx) Execute(query string, args ...any) (sql.Result, error) {
	return t.postgres.Exec(query, args...)
}
