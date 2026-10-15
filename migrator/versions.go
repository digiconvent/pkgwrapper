package migrator

import (
	"io/fs"
	"slices"

	"github.com/digiconvent/pkgwrapper/version"
)

func (d *data) Versions() ([]string, error) {
	var versions = []string{}

	entries, err := fs.ReadDir(d.embeddedFS, ".")
	if err != nil {
		return versions, nil
	}
	for _, entry := range entries {
		if version.ToVersion(entry.Name()) == nil {
			continue
		}

		if !slices.Contains(versions, entry.Name()) {
			versions = append(versions, entry.Name())
		}
	}
	return versions, nil
}
