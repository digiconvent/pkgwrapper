package migrator

import (
	"io/fs"
	"path"
	"strings"
)

func (d *data) GetPackageMigration(pkg string, version string) (map[string]string, error) {
	files, err := fs.ReadDir(d.embeddedFS, version)
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
		contents, err := fs.ReadFile(d.embeddedFS, path.Join(version, file.Name()))
		if err != nil {
			continue
		}
		migrationScript[file.Name()] = string(contents)
	}
	return migrationScript, nil
}
