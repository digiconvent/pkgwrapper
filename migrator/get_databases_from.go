package migrator

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/digiconvent/pkgwrapper/db"
)

func (d *data) GetDatabasesFrom(dir string) (map[string]db.DatabaseInterface, error) {
	if _, err := os.Stat(dir); err != nil {
		return nil, errors.New("data folder does not exist")
	}

	d.dataDir = dir
	packages, err := d.GetPackages()
	if err != nil {
		return nil, err
	}

	var databases map[string]db.DatabaseInterface = make(map[string]db.DatabaseInterface)
	for _, pkg := range packages {
		pkgDir := path.Join(dir, pkg)

		pragmas := []string{
			// persistent pragmas
			"journal_mode = WAL",
			"foreign_keys = ON",
			// non-persistent pragmas
			"synchronous = NORMAL",
			"busy_timeout = 1000",
			"cache_size = -500",
			"mmap_size = 16777216",
		}
		dbUri := path.Join(pkgDir, pkg)

		dbConn, err := db.New(dbUri)
		if err != nil {
			return nil, err
		}
		databases[pkg] = dbConn

		for _, pragma := range pragmas {
			_, err := dbConn.Exec("PRAGMA " + pragma)
			if err != nil {
				return nil, fmt.Errorf("%s: could not execute pragma %q: %w", dbUri, pragma, err)
			}
		}
	}
	return databases, nil
}
