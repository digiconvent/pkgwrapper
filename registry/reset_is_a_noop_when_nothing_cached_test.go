package registry_test

import (
	"testing"

	"github.com/digiconvent/pkgwrapper/registry"
)

func TestReset_IsANoopWhenNothingCached(t *testing.T) {
	registry.Reset("never-built", t.TempDir())
}
