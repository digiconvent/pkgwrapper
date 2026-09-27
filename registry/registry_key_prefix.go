package registry

func registryKeyPrefix(name, dataDir string) string {
	return dataDir + "|" + name + "|"
}
