package registry_test

import (
	"testing"

	"github.com/digiconvent/pkgwrapper/registry"
	"github.com/digiconvent/pkgwrapper/version"
)

func TestReset_ClosesEveryTargetVersionCachedForName(t *testing.T) {
	dataDir := t.TempDir()
	build := buildEmbeddedFiles(t)

	b1, _ := registry.For[*testServices]("widgets", dataDir, version.ToVersion("0.0.1"), build)
	b1.ServingWith(newTestServices)
	b2, _ := registry.For[*testServices]("widgets", dataDir, version.ToVersion("0.0.2"), build)
	b2.ServingWith(newTestServices)

	registry.Reset("widgets", dataDir)

	b3, err := registry.For[*testServices]("widgets", dataDir, version.ToVersion("0.0.1"), build)
	if err != nil {
		t.Fatalf("unexpected error after reset: %v", err)
	}
	if b3 == b1 {
		t.Error("expected 0.0.1's cached builder to have been reset too, not just 0.0.2's")
	}
}
