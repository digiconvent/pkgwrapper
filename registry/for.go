package registry

import "github.com/digiconvent/pkgwrapper/version"

func For[S Initialisable](name, dataDir string, target *version.Version, buildMgr BuildManager) (*PackageBuilder[S], error) {
	key := registryKey(name, dataDir, target)

	registryMu.Lock()
	defer registryMu.Unlock()

	if existing, ok := registry[key]; ok {
		return existing.(*PackageBuilder[S]), nil
	}

	b := &PackageBuilder[S]{name: name, dataDir: dataDir}
	if err := b.open(target, buildMgr); err != nil {
		return nil, err
	}
	registry[key] = b
	return b, nil
}
