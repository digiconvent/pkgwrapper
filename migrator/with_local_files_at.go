package migrator

import "path"

func (d *data) WithLocalFilesAt(projectRoot, pkgDir string) (PackageWrapper, error) {
	d.pkgDir = path.Join(projectRoot, pkgDir)
	return d, nil
}
