package migrator

func (d *data) Verbose() packageManagerChoice {
	d.verbose = true
	return d
}
