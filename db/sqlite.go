package db

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/digiconvent/pkgwrapper/.test/log"
	"github.com/digiconvent/pkgwrapper/.test/utils"
	_ "github.com/mattn/go-sqlite3"
)

func New(uri string) (DatabaseInterface, error) {
	db, err := sql.Open("sqlite3", uri)
	if err != nil {
		return nil, err
	}

	return &SqliteDatabase{
		DB:    db,
		uri:   uri,
		mutex: sync.Mutex{},
	}, nil
}

type SqliteDatabase struct {
	DB    *sql.DB
	uri   string
	mutex sync.Mutex
}

func (s *SqliteDatabase) Exec(query string, args ...any) (sql.Result, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.DB.Exec(query, args...)
}
func (s *SqliteDatabase) ExecDebug(query string, args ...any) (sql.Result, error) {
	log.Info("Exec: " + query)
	log.Info(argsToString(args))
	result, err := s.DB.Exec(query, args...)

	if err != nil {
		log.Error("Exec failed: " + err.Error())
		return nil, err
	}

	if result == nil {
		err := errors.New("result is nil, cannot proceed")
		log.Error(err)
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Error("Could not get rows affected: " + err.Error())
	}

	log.Info("Rows affected: " + strconv.Itoa(int(rowsAffected)))

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		log.Error("Could not get last insert id: " + err.Error())
	}

	log.Info("Last insert id: " + strconv.Itoa(int(lastInsertId)))

	return result, err
}

func (s *SqliteDatabase) Query(query string, args ...any) (*sql.Rows, error) {
	return s.DB.Query(query, args...)
}

func (s *SqliteDatabase) QueryDebug(query string, args ...any) (*sql.Rows, error) {
	log.Info("Query: " + query)
	log.Info(argsToString(args))
	rows, err := s.DB.Query(query, args...)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	columns, err := rows.Columns()

	if err != nil {
		return nil, err
	}

	values := make([]any, len(columns))
	valuePtrs := make([]any, len(columns))

	for i := range values {
		valuePtrs[i] = &values[i]
	}

	var table = utils.NewTable(columns)
	for rows.Next() {
		err = rows.Scan(valuePtrs...)

		if err != nil {
			log.Error("Something went wrong scanning the values from this query, you have different problems")
			return nil, err
		}

		table.AddRow(values...)
	}
	count := len(table.Values)
	fmt.Println("This query has " + strconv.Itoa(count) + " rows")
	fmt.Println(table.Render())

	return rows, nil
}

func (s *SqliteDatabase) QueryRow(query string, args ...any) (*sql.Row, error) {
	row := s.DB.QueryRow(query, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return row, nil
}

func (s *SqliteDatabase) QueryRowDebug(query string, args ...any) (*sql.Row, error) {
	log.Info("Query: " + query)
	log.Info(argsToString(args))
	_, err := s.QueryDebug(query, args...)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	row, _ := s.QueryRow(query, args...)

	return row, nil
}

func (s *SqliteDatabase) Close() {
	s.DB.Close()
}

func argsToString(args []any) string {
	res := []string{}
	for _, x := range args {
		res = append(res, fmt.Sprintf("'%v'", x))
	}
	return strings.Join(res, ",")
}
