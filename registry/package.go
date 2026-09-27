package registry

import (
	"sync"

	"github.com/digiconvent/pkgwrapper/db"
	"github.com/digiconvent/pkgwrapper/migrator"
	"github.com/digiconvent/pkgwrapper/version"
)

type Initialisable interface {
	Initialise()
}

type Package[S Initialisable] struct {
	Name     string
	Services S
}

var (
	registry   = map[string]any{}
	registryMu sync.Mutex
)

type BuildManager func(installed, target *version.Version) (migrator.PackageWrapper, error)

type PackageBuilder[S Initialisable] struct {
	name    string
	dataDir string

	mu       sync.Mutex
	database db.DatabaseInterface
	pkg      *Package[S]
}

type closer interface{ closeDatabase() }
