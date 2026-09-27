package migrator

import (
	"fmt"
	"os"
	"path"
	"slices"
	"strings"
)

func (d *data) WithPkgDir(dir string) (PackageWrapper, error) {
	if d.verbose {
		fmt.Println("[Migrate Packages] Packages located relative at " + dir)
	}
	downloadFolder := path.Join(os.TempDir(), "migrate_packages")
	d.pkgDir = path.Join(downloadFolder, dir)

	if _, err := os.Stat(d.pkgDir); err != nil {
		fmt.Println("[Migrate Packages] Could not find temporary download folder (" + d.pkgDir + ") for packages")
		return nil, fmt.Errorf("could not find folder %s: expected it to be present in the repository", d.pkgDir)
	}

	segments := slices.DeleteFunc(strings.Split(dir, "/"), func(a string) bool {
		return a == ""
	})
	if d.verbose {
		fmt.Println("[Migrate Packages] removing non-migration files since they are not needed")
	}
	toKeep := ""
	for i := range len(segments) {
		toScan := downloadFolder + toKeep
		toKeep += "/" + segments[i]
		if d.verbose {
			fmt.Println("[Migrate Packages] Scanning " + toScan + " and getting rid of anything that does not start with " + toKeep)
		}
		entries, err := os.ReadDir(toScan)
		if err != nil {
			return nil, err
		}
		for _, entry := range entries {
			uri := path.Join(toScan, entry.Name())
			keep := strings.HasSuffix(uri, toKeep)
			if !keep {
				if d.verbose {
					fmt.Println("[Migrate Packages] Delete " + uri)
				}
				err = os.RemoveAll(uri)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	packages, err := d.GetPackages()
	if err != nil {
		return nil, err
	}

	for _, pkg := range packages {
		pkgDir := path.Join(downloadFolder, "", dir, pkg)
		entries, err := os.ReadDir(pkgDir)
		if err != nil {
			return nil, err
		}
		for _, entry := range entries {
			toRemove := path.Join(pkgDir, entry.Name())
			if entry.Name() != "db" {
				err = os.RemoveAll(toRemove)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return d, nil
}
