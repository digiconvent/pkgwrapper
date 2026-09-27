package registry

import (
	"testing"

	"github.com/digiconvent/pkgwrapper/migrator"
	"github.com/digiconvent/pkgwrapper/version"
)

func migrationFilesFor(cfg Config, projectRoot string) BuildManager {
	return func(installed, target *version.Version) (migrator.PackageWrapper, error) {
		if testing.Testing() || target.Major == -1 {
			return migrator.From(installed).To(target).WithLocalFilesAt(projectRoot, cfg.PkgDir)
		}
		repoMgr, err := migrator.From(installed).To(target).WithPublicRepository(cfg.RepoOwner, cfg.RepoName)
		if err != nil {
			return nil, err
		}
		return repoMgr.WithPkgDir(cfg.PkgDir)
	}
}
