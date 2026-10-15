package migrator

func (d *data) GetPackages() ([]string, error) {
	return []string{d.embeddedName}, nil
}
