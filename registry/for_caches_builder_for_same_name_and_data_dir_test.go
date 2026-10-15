package registry_test

import (
	"testing"

	"github.com/digiconvent/pkgwrapper/db"
	"github.com/digiconvent/pkgwrapper/registry"
	"github.com/digiconvent/pkgwrapper/version"
)

func TestFor_CachesBuilderForSameNameAndDataDir(t *testing.T) {
	dataDir := t.TempDir()
	target := version.ToVersion("0.0.1")
	build := buildEmbeddedFiles(t)

	b1, err := registry.For[*testServices]("widgets", dataDir, target, build)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	b2, err := registry.For[*testServices]("widgets", dataDir, target, build)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if b1 != b2 {
		t.Error("expected the same builder for the same name+dataDir")
	}

	calls := 0
	factory := func(database db.DatabaseInterface) *testServices { calls++; return newTestServices(database) }
	b1.ServingWith(factory)
	b2.ServingWith(factory)
	if calls != 1 {
		t.Errorf("expected the factory to run once, got %d", calls)
	}
}
