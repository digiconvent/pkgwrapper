package registry_test

import (
	"errors"
	"testing"

	"github.com/digiconvent/pkgwrapper/migrator"
	"github.com/digiconvent/pkgwrapper/registry"
	"github.com/digiconvent/pkgwrapper/version"
)

func TestFor_PropagatesBuildManagerError(t *testing.T) {
	dataDir := t.TempDir()
	wantErr := "forced build manager error"
	buildMgr := func(_, _ *version.Version) (migrator.PackageWrapper, error) {
		return nil, errors.New(wantErr)
	}
	_, err := registry.For[*testServices]("widgets", dataDir, version.ToVersion("0.0.1"), buildMgr)
	if err == nil {
		t.Fatal("expected an error when buildMgr fails")
	}
}
