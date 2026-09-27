package db_test

import (
	"path/filepath"
	"testing"

	"github.com/digiconvent/pkgwrapper/db"
)

func TestWipe_ClearsDataWithoutChangingTheDatabaseInterfaceIdentity(t *testing.T) {
	path := filepath.Join(t.TempDir(), "widgets")
	conn, err := db.New(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, err := conn.Exec("create table widgets (id integer primary key, name text not null)"); err != nil {
		t.Fatalf("failed to create table: %v", err)
	}
	if _, err := conn.Exec("insert into widgets (name) values ('frodo')"); err != nil {
		t.Fatalf("failed to insert a row: %v", err)
	}

	if err := conn.Wipe(); err != nil {
		t.Fatalf("unexpected error from Wipe: %v", err)
	}

	// the same DatabaseInterface value keeps working - no caller needs to
	// know a new connection now backs it.
	if _, err := conn.Exec("create table widgets (id integer primary key, name text not null)"); err != nil {
		t.Fatalf("expected the schema to be re-creatable after wipe: %v", err)
	}
	rows, err := conn.Query("select 1 from widgets")
	if err != nil {
		t.Fatalf("unexpected error querying widgets: %v", err)
	}
	defer rows.Close()
	if rows.Next() {
		t.Error("expected no rows left after wipe")
	}
}
