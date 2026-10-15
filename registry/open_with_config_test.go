package registry_test

import (
	"testing"

	"github.com/digiconvent/pkgwrapper/registry"
)

func TestOpenWithConfig_MigratesAndServesUsingAnEmbeddedFS(t *testing.T) {
	b, err := registry.OpenWithConfig[*testServices]("widgets", registry.Config{
		DataRoot:   t.TempDir(),
		EmbeddedFS: widgetsFS(t),
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
	b, err := registry.OpenWithConfig[*testServices]("widgets", registry.Config{
		EmbeddedFS: widgetsFS(t),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	pkg := b.ServingWith(newTestServices)
	if !pkg.Services.hasColumn("widgets", "name") {
		t.Error("expected widgets table to exist")
	}
}

func TestOpenWithConfig_WorksOutsideAnyProjectRoot(t *testing.T) {
	fsys := widgetsFS(t)
	t.Chdir(t.TempDir())

	b, err := registry.OpenWithConfig[*testServices]("widgets", registry.Config{
		DataRoot:   t.TempDir(),
		EmbeddedFS: fsys,
	})
	if err != nil {
		t.Fatalf("expected no error outside any project root, got: %v", err)
	}
	pkg := b.ServingWith(newTestServices)
	if !pkg.Services.hasColumn("widgets", "name") {
		t.Error("expected widgets table to exist")
	}
}
