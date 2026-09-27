package registry

import "github.com/digiconvent/pkgwrapper/version"

func registryKey(name, dataDir string, target *version.Version) string {
	return registryKeyPrefix(name, dataDir) + target.String()
}
