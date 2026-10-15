package registry_test

import (
	"testing"

	"github.com/digiconvent/pkgwrapper/registry"
	"github.com/digiconvent/pkgwrapper/version"
)

func TestFor_DifferentDataDirProducesDifferentBuilder(t *testing.T) {
	target := version.ToVersion("0.0.1")
	build := buildEmbeddedFiles(t)

	b1, err := registry.For[*testServices]("widgets", t.TempDir(), target, build)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	b2, err := registry.For[*testServices]("widgets", t.TempDir(), target, build)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if b1 == b2 {
		t.Error("expected different builders for different data directories")
	}
}
