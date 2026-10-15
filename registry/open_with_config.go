package registry

import (
	"fmt"
	"os"

	"github.com/digiconvent/pkgwrapper/version"
)

func OpenWithConfig[S Initialisable](name string, cfg Config) (*PackageBuilder[S], error) {
	dataRoot := cfg.DataRoot
	if dataRoot == "" {
		var err error
		dataRoot, err = resolveDataRoot()
		if err != nil {
			return nil, err
		}
	}
	if err := os.MkdirAll(dataRoot, 0755); err != nil {
		return nil, fmt.Errorf("registry: could not create data root %s: %w", dataRoot, err)
	}

	target := version.ToVersion("-1.-1.-1")
	if cfg.TargetVersion != "" {
		target = version.ToVersion(cfg.TargetVersion)
	}

	return For[S](name, dataRoot, target, migrationFilesFor(cfg, name))
}
