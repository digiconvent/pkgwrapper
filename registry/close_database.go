package registry

func (b *PackageBuilder[S]) closeDatabase() {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.database != nil {
		b.database.Close()
	}
	b.database, b.pkg = nil, nil
}
