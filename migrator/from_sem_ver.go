package migrator

import "github.com/digiconvent/pkgwrapper/version"

// if fresh, pass on a -1.-1.-1
func FromSemVer(ma, mi, pa int) migrateToVersion {
	return From(&version.Version{Major: ma, Minor: mi, Patch: pa})
}
