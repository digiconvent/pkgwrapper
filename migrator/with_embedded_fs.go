package migrator

import "io/fs"

func (d *data) WithEmbeddedFS(fsys fs.FS, name string) (PackageWrapper, error) {
	d.embeddedFS = fsys
	d.embeddedName = name
	return d, nil
}
