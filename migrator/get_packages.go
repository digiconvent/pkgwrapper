package migrator

import (
	"fmt"
	"os"
	"path"
)

func (d *data) GetPackages() ([]string, error) {
	if d.verbose {
		fmt.Println("[Migrate Packages] Scanning " + d.pkgDir + " for packages")
	}
	entries, err := os.ReadDir(d.pkgDir)
	if err != nil {
		return nil, err
	}
	var packages = []string{}
	for _, entry := range entries {
		if entry.IsDir() {
			if d.verbose {
				fmt.Println("[Migrate Packages] Found " + entry.Name())
			}
			// only append if the folder is not empty
			dirEntries, err := os.ReadDir(path.Join(d.pkgDir, entry.Name()))
			if err != nil {
				return nil, err
			}
			if len(dirEntries) > 0 {
				if _, err := os.Stat(path.Join(d.pkgDir, entry.Name(), "db")); err != nil {
					fmt.Println("[Migrate Packages] No db folder found for " + entry.Name() + ", skipping...")
					continue
				}
				packages = append(packages, entry.Name())
			} else {
				if d.verbose {
					fmt.Println("[Migrate Packages] No contents found for " + entry.Name() + ", skipping...")
				}
			}
		}
	}
	return packages, nil
}
