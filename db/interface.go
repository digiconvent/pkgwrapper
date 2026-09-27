package db

import "database/sql"

type DatabaseInterface interface {
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) (*sql.Row, error)

	ExecDebug(query string, args ...any) (sql.Result, error)
	QueryDebug(query string, args ...any) (*sql.Rows, error)
	QueryRowDebug(query string, args ...any) (*sql.Row, error)
	Close()
	Wipe() error
}
