package registry_test

import (
	"testing"

	"github.com/digiconvent/pkgwrapper/registry"
	"github.com/digiconvent/pkgwrapper/version"
)

func TestFor_ReopensAtLaterVersionWithoutRerunningEarlierMigrations(t *testing.T) {
	projectRoot := setupPackageFixture(t)
	dataDir := t.TempDir()

	registry.Reset("widgets", dataDir)
	b, err := registry.For[*testServices]("widgets", dataDir, version.ToVersion("0.0.1"), buildLocalFiles(projectRoot))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	b.ServingWith(newTestServices)

	// a later process, same data dir, asking for a later target version -
	// this is a fresh For call (name+dataDir not in this run's cache), so it
	// must read the *already installed* version back from disk rather than
	// re-running 0.0.1's migration against a table it already created.
	registry.Reset("widgets", dataDir)
	b2, err := registry.For[*testServices]("widgets", dataDir, version.ToVersion("0.0.2"), buildLocalFiles(projectRoot))
	if err != nil {
		t.Fatalf("unexpected error migrating forward: %v", err)
	}
	pkg := b2.ServingWith(newTestServices)
	if !pkg.Services.hasColumn("widgets", "color") {
		t.Error("expected widgets to have the color column after migrating to 0.0.2")
	}
}
