package registry

import (
	"github.com/digiconvent/pkgwrapper/migrator"
	"github.com/digiconvent/pkgwrapper/version"
)

func migrationFilesFor(cfg Config, name string) BuildManager {
	return func(installed, target *version.Version) (migrator.PackageWrapper, error) {
		return migrator.From(installed).To(target).WithEmbeddedFS(cfg.EmbeddedFS, name)
	}
}
