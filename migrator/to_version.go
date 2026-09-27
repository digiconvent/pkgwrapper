package migrator

import "github.com/digiconvent/pkgwrapper/version"

func (d *data) ToVersion(ma int, mi int, pa int) packageManagerChoice {
	d.toVersion = &version.Version{Major: ma, Minor: mi, Patch: pa}
	return d
}
