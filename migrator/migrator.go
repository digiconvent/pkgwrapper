package migrator

import (
	"io/fs"

	"github.com/digiconvent/pkgwrapper/db"
	"github.com/digiconvent/pkgwrapper/version"
)

type PackageWrapper interface {
	GetPackages() ([]string, error)
	Versions() ([]string, error)
	VersionsToMigrate() ([]string, error)
	GetPackageMigration(pkg, version string) (map[string]string, error)
	// <dir> is the folder where the data of the packages is located.
	// e.g., if package 'iam' has data, the database would be located at <dir>/iam/iam.db
	// so the <dir> usually is a home directory
	MigrateDatabasesIn(dir string) (map[string]db.DatabaseInterface, error)
	GetDatabasesFrom(dir string) (map[string]db.DatabaseInterface, error)
}

type data struct {
	fromVersion *version.Version
	toVersion   *version.Version
	dataDir     string

	embeddedFS   fs.FS
	embeddedName string

	verbose bool
}

type migrateToVersion interface {
	// to migrate to latest, pass on a -1,-1,-1
	ToVersion(ma, mi, pa int) packageManagerChoice
	// to migrate to latest, pass on a nil
	To(version *version.Version) packageManagerChoice
}

type packageManagerChoice interface {
	Verbose() packageManagerChoice
	WithEmbeddedFS(fsys fs.FS, name string) (PackageWrapper, error)
}
