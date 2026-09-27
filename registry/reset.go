package registry

import "strings"

func Reset(name, dataDir string) {
	prefix := registryKeyPrefix(name, dataDir)
	registryMu.Lock()
	defer registryMu.Unlock()
	for key, entry := range registry {
		if !strings.HasPrefix(key, prefix) {
			continue
		}
		delete(registry, key)
		entry.(closer).closeDatabase()
	}
}
