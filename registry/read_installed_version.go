package registry

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/digiconvent/pkgwrapper/version"
)

func readInstalledVersion(dbPath string) *version.Version {
	uninitialised := version.ToVersion("-1.-1.-1")
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		return uninitialised
	}

	sqlDB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return uninitialised
	}
	defer sqlDB.Close()

	var versionStr string
	if err := sqlDB.QueryRow(`select value from _meta where key = 'version'`).Scan(&versionStr); err != nil {
		return uninitialised
	}
	return version.ToVersion(versionStr)
}
