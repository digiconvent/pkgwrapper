package migrator

import "github.com/digiconvent/pkgwrapper/version"

func (d *data) To(v *version.Version) packageManagerChoice {
	d.toVersion = v
	return d
}
