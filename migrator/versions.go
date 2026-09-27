package migrator

import (
	"os"
	"path"
	"slices"

	"github.com/digiconvent/pkgwrapper/version"
)

func (d *data) Versions() ([]string, error) {
	var versions = []string{}
	packages, err := d.GetPackages()
	if err != nil {
		return nil, err
	}

	for _, pkg := range packages {
		entries, err := os.ReadDir(path.Join(d.pkgDir, pkg, "db"))
		if err != nil {
			continue // if db doesn't exist, don't do anything
		}
		for _, entry := range entries {
			if version.ToVersion(entry.Name()) == nil {
				continue
			}

			if !slices.Contains(versions, entry.Name()) {
				versions = append(versions, entry.Name())
			}
		}
	}
	return versions, nil
}
