package migrator

import "github.com/digiconvent/pkgwrapper/version"

func From(v *version.Version) migrateToVersion {
	return &data{
		fromVersion: v,
	}
}
