package registry

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func resolveDataRoot() (string, error) {
	if testing.Testing() {
		return filepath.Join(os.TempDir(), fmt.Sprintf("migrate_packages_%d", os.Getpid())), nil
	}
	return "", errors.New("registry: DataRoot not set in Config")
}
