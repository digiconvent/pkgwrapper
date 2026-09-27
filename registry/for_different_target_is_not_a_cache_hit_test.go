package registry_test

import (
	"testing"

	"github.com/digiconvent/pkgwrapper/registry"
	"github.com/digiconvent/pkgwrapper/version"
)

func TestFor_DifferentTargetIsNotACacheHit(t *testing.T) {
	projectRoot := setupPackageFixture(t)
	dataDir := t.TempDir()

	b1, err := registry.For[*testServices]("widgets", dataDir, version.ToVersion("0.0.1"), buildLocalFiles(projectRoot))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	pkg1 := b1.ServingWith(newTestServices)
	if pkg1.Services.hasColumn("widgets", "color") {
		t.Fatal("color column should not exist yet at 0.0.1")
	}

	b2, err := registry.For[*testServices]("widgets", dataDir, version.ToVersion("0.0.2"), buildLocalFiles(projectRoot))
	if err != nil {
		t.Fatalf("unexpected error migrating to 0.0.2: %v", err)
	}
	if b1 == b2 {
		t.Fatal("expected a different builder for a different target version")
	}
	pkg2 := b2.ServingWith(newTestServices)
	if !pkg2.Services.hasColumn("widgets", "color") {
		t.Error("expected the second call to actually migrate forward to 0.0.2")
	}
}
