package registry_test

import (
	"embed"
	"io/fs"
	"testing"

	"github.com/digiconvent/pkgwrapper/db"
	"github.com/digiconvent/pkgwrapper/migrator"
	"github.com/digiconvent/pkgwrapper/registry"
	"github.com/digiconvent/pkgwrapper/version"
)

type testServices struct {
	database    db.DatabaseInterface
	initialised bool
}

func (s *testServices) Initialise() { s.initialised = true }

func newTestServices(database db.DatabaseInterface) *testServices {
	return &testServices{database: database}
}

func (s *testServices) hasColumn(table, column string) bool {
	rows, err := s.database.Query("select 1 from pragma_table_info(?) where name = ?", table, column)
	if err != nil {
		return false
	}
	defer rows.Close()
	return rows.Next()
}

//go:embed .test/widgets/db
var embeddedWidgetsDb embed.FS

func widgetsFS(t *testing.T) fs.FS {
	t.Helper()
	sub, err := fs.Sub(embeddedWidgetsDb, ".test/widgets/db")
	if err != nil {
		t.Fatal(err)
	}
	return sub
}

func buildEmbeddedFiles(t *testing.T) registry.BuildManager {
	fsys := widgetsFS(t)
	return func(installed, target *version.Version) (migrator.PackageWrapper, error) {
		return migrator.From(installed).To(target).WithEmbeddedFS(fsys, "widgets")
	}
}
