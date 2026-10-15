package registry_test

import (
	"testing"

	"github.com/digiconvent/pkgwrapper/registry"
	"github.com/digiconvent/pkgwrapper/version"
)

func TestReset_ClosesAndForgetsSoNextForRebuilds(t *testing.T) {
	dataDir := t.TempDir()
	target := version.ToVersion("0.0.1")
	build := buildEmbeddedFiles(t)

	b1, err := registry.For[*testServices]("widgets", dataDir, target, build)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	b1.ServingWith(newTestServices)

	registry.Reset("widgets", dataDir)

	b2, err := registry.For[*testServices]("widgets", dataDir, target, build)
	if err != nil {
		t.Fatalf("unexpected error after reset: %v", err)
	}
	if b1 == b2 {
		t.Error("expected a new builder after Reset")
	}
	pkg := b2.ServingWith(newTestServices)
	if !pkg.Services.hasColumn("widgets", "name") {
		t.Error("expected the reopened database to still have its schema")
	}
}
