package migrator_test

import (
	"slices"
	"testing"

	"github.com/digiconvent/pkgwrapper/migrator"
)

func TestDatabaseHandling(t *testing.T) {
	dataFolder := t.TempDir()

	repo, err := migrator.From(currentVersion).To(targetVersion).WithPublicRepository(username, reponame)
	if err != nil {
		t.Fatal(err)
	}
	manager, err := repo.WithPkgDir(pkgFolder)

	if err != nil {
		t.Fatal(err)
	}
	_, err = manager.GetDatabasesFrom(dataFolder)
	if err == nil {
		t.Fatal("expected an err since databases are not initialised")
	}

	databases, err := manager.MigrateDatabasesIn(dataFolder)
	if err != nil {
		t.Fatal(err)
	}

	isKeys := make([]string, 0, len(databases))
	for k := range databases {
		isKeys = append(isKeys, k)
	}

	shouldKeys := []string{"iam", "post", "sys"}
	if !slices.Equal(shouldKeys, isKeys) {
		t.Fatal("expected", shouldKeys, "instead got", isKeys)
	}
	// this package is for separate sqlite databases for every package under /pkg
	// so every pkg needs to handle their own database
	// create live/test
	// - this needs a folder where the database would be (/home/user/data/<pkgName>/database)
	// migrate
	// - database exists, migration script should work on it
	//	delete
	// - test databases should be deleted after every run so there will be a fresh start for the next time
	// - production databases should NOT be deleted once they are created

	databases2, err := manager.GetDatabasesFrom(dataFolder)
	if err != nil {
		t.Fatal("did not expect err instead got", err.Error())
	}
	keys := make([]string, 0, len(databases))
	for k := range databases {
		keys = append(keys, k)
	}
	keys2 := make([]string, 0, len(databases2))
	for k := range databases2 {
		keys2 = append(keys2, k)
	}
	slices.Sort(keys)
	slices.Sort(keys2)
	if !slices.Equal(keys, keys2) {
		t.Fatalf("should have the same databases")
	}
}
