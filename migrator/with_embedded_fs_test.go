package migrator_test

import (
	"embed"
	"io/fs"
	"slices"
	"testing"

	"github.com/digiconvent/pkgwrapper/migrator"
	"github.com/digiconvent/pkgwrapper/version"
)

//go:embed .test/pkg/iam/db
var embeddedIamDb embed.FS

func rootedAtDb(t *testing.T) migrator.PackageWrapper {
	t.Helper()
	currentVersion := &version.Version{Major: 0, Minor: 0, Patch: 0}
	targetVersion := &version.Version{Major: 1, Minor: 0, Patch: 1}

	sub, err := fs.Sub(embeddedIamDb, ".test/pkg/iam/db")
	if err != nil {
		t.Fatal(err)
	}

	manager, err := migrator.From(currentVersion).To(targetVersion).WithEmbeddedFS(sub, "iam")
	if err != nil {
		t.Fatal(err)
	}
	return manager
}

func TestWithEmbeddedFS_GetPackagesReturnsTheGivenName(t *testing.T) {
	manager := rootedAtDb(t)
	packages, err := manager.GetPackages()
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(packages, []string{"iam"}) {
		t.Fatalf("expected [iam], got %v", packages)
	}
}

func TestWithEmbeddedFS_VersionsListsEmbeddedVersionFolders(t *testing.T) {
	manager := rootedAtDb(t)
	versions, err := manager.Versions()
	if err != nil {
		t.Fatal(err)
	}
	slices.Sort(versions)
	if !slices.Equal(versions, []string{"0.0.0", "0.1.0", "1.0.0"}) {
		t.Fatalf("expected [0.0.0 0.1.0 1.0.0], got %v", versions)
	}
}

func TestWithEmbeddedFS_GetPackageMigrationReadsEmbeddedSQL(t *testing.T) {
	manager := rootedAtDb(t)
	scripts, err := manager.GetPackageMigration("iam", "0.1.0")
	if err != nil {
		t.Fatal(err)
	}
	if len(scripts) != 2 {
		t.Fatalf("expected 2 scripts for 0.1.0, got %d: %v", len(scripts), scripts)
	}
	if _, ok := scripts["add_first_name.sql"]; !ok {
		t.Errorf("expected add_first_name.sql in %v", scripts)
	}
}

func TestWithEmbeddedFS_VersionsToMigrateRespectsFromAndTo(t *testing.T) {
	manager := rootedAtDb(t)
	versions, err := manager.VersionsToMigrate()
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(versions, []string{"0.1.0", "1.0.0"}) {
		t.Fatalf("expected [0.1.0 1.0.0], got %v", versions)
	}
}
