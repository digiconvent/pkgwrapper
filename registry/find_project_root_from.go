package registry

import (
	"errors"
	"os"
	"path/filepath"
)

func findProjectRootFrom(dir string) (string, error) {
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("registry: could not find project root: no go.mod in any parent directory")
		}
		dir = parent
	}
}
