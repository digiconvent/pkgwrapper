package registry_test

import (
	"testing"

	"github.com/digiconvent/pkgwrapper/registry"
)

func TestOpenWithConfig_MigratesAndServesUsingLocalFilesInTests(t *testing.T) {
	projectRoot := setupPackageFixture(t)
	dataDir := t.TempDir()

	b, err := registry.OpenWithConfig[*testServices]("widgets", registry.Config{
		ProjectRoot: projectRoot,
		DataRoot:    dataDir,
		PkgDir:      "pkg",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	pkg := b.ServingWith(newTestServices)

	if !pkg.Services.hasColumn("widgets", "name") {
		t.Error("expected widgets table to exist")
	}
}

func TestOpenWithConfig_DefaultsDataRootDuringTests(t *testing.T) {
	projectRoot := setupPackageFixture(t)

	b, err := registry.OpenWithConfig[*testServices]("widgets", registry.Config{
		ProjectRoot: projectRoot,
		PkgDir:      "pkg",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	pkg := b.ServingWith(newTestServices)
	if !pkg.Services.hasColumn("widgets", "name") {
		t.Error("expected widgets table to exist")
	}
}
