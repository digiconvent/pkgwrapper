package registry

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/digiconvent/pkgwrapper/db"
	"github.com/digiconvent/pkgwrapper/version"
)

func (b *PackageBuilder[S]) open(target *version.Version, buildMgr BuildManager) error {
	dbPath := filepath.Join(b.dataDir, b.name, b.name+".db")
	_, dbExists := os.Stat(dbPath)

	installed := readInstalledVersion(dbPath)

	mgr, err := buildMgr(installed, target)
	if err != nil {
		return fmt.Errorf("migrate_packages: migration setup failed for %q: %w", b.name, err)
	}

	var databases map[string]db.DatabaseInterface
	if installed.Equals(target) && dbExists == nil {
		databases, err = mgr.GetDatabasesFrom(b.dataDir)
	} else {
		databases, err = mgr.MigrateDatabasesIn(b.dataDir)
	}
	if err != nil {
		return fmt.Errorf("migrate_packages: failed to open database for %q: %w", b.name, err)
	}

	database, ok := databases[b.name]
	if !ok {
		return fmt.Errorf("migrate_packages: no database found for package %q", b.name)
	}

	metaConn, err := db.New(dbPath)
	if err != nil {
		return fmt.Errorf("migrate_packages: failed to open meta connection for %q: %w", b.name, err)
	}
	defer metaConn.Close()
	if err := recordVersion(metaConn, b.name, target.String()); err != nil {
		return fmt.Errorf("migrate_packages: failed to record version for %q: %w", b.name, err)
	}

	b.database = database
	return nil
}
