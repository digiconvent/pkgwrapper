package registry_test

import (
	"testing"

	"github.com/digiconvent/pkgwrapper/registry"
	"github.com/digiconvent/pkgwrapper/version"
)

func TestFor_MigratesAndServesFreshDatabase(t *testing.T) {
	dataDir := t.TempDir()
	target := version.ToVersion("0.0.1")

	b, err := registry.For[*testServices]("widgets", dataDir, target, buildEmbeddedFiles(t))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	pkg := b.ServingWith(newTestServices)

	if !pkg.Services.initialised {
		t.Error("expected services to be initialised")
	}
	if !pkg.Services.hasColumn("widgets", "name") {
		t.Error("expected widgets table to exist at 0.0.1")
	}
	if pkg.Services.hasColumn("widgets", "color") {
		t.Error("color column should not exist yet at 0.0.1")
	}
}
