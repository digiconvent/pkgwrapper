package registry

func ResetWithConfig(name string, cfg Config) error {
	dataRoot := cfg.DataRoot
	if dataRoot == "" {
		var err error
		dataRoot, err = resolveDataRoot()
		if err != nil {
			return err
		}
	}
	Reset(name, dataRoot)
	return nil
}
