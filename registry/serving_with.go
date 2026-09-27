package registry

import "github.com/digiconvent/pkgwrapper/db"

func (b *PackageBuilder[S]) ServingWith(factory func(db.DatabaseInterface) S) *Package[S] {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.pkg != nil {
		return b.pkg
	}
	services := factory(b.database)
	services.Initialise()
	b.pkg = &Package[S]{Name: b.name, Services: services}
	return b.pkg
}
