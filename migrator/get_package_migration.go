package migrator

import (
	"os"
	"path"
	"strings"
)

func (d *data) GetPackageMigration(pkg string, version string) (map[string]string, error) {
	versionInPkgDir := path.Join(d.pkgDir, pkg, "db", version)
	files, err := os.ReadDir(versionInPkgDir)
	if err != nil {
		return nil, err
	}

	var migrationScript map[string]string = map[string]string{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}
		// these files cannot be too big to handle so reading the entire contents of an sql file should not be a problem
		contents, err := os.ReadFile(path.Join(versionInPkgDir, file.Name()))
		if err != nil {
			continue
		}
		migrationScript[file.Name()] = string(contents)
	}
	return migrationScript, nil
}
