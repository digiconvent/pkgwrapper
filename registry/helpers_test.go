package registry_test

import (
	"os"
	"path/filepath"
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

func setupPackageFixture(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()

	write := func(path, sql string) {
		t.Helper()
		full := filepath.Join(dir, path)
		os.MkdirAll(filepath.Dir(full), 0755)
		if err := os.WriteFile(full, []byte(sql), 0644); err != nil {
			t.Fatalf("failed to write %s: %v", path, err)
		}
	}

	write("pkg/widgets/db/0.0.1/001_create_widgets.sql", `create table widgets (id integer primary key, name text not null);`)
	write("pkg/widgets/db/0.0.2/001_add_color.sql", `alter table widgets add column color text;`)

	return dir
}

func buildLocalFiles(projectRoot string) registry.BuildManager {
	return func(installed, target *version.Version) (migrator.PackageWrapper, error) {
		return migrator.From(installed).To(target).WithLocalFilesAt(projectRoot, "pkg")
	}
}
