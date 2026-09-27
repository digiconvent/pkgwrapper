package migrator

import "github.com/digiconvent/pkgwrapper/version"

func (d *data) VersionsToMigrate() ([]string, error) {
	versions, err := d.Versions()
	if err != nil {
		return nil, err
	}
	if len(versions) == 0 {
		return []string{}, nil
	}

	// we don't need to do this with minVersion since
	var maxVersion *version.Version = d.toVersion
	if maxVersion.IsUninitialised() {
		// there is no "max version", it is -1.-1.-1 so we take the latest version as max
		maxVersion = version.ToVersion(versions[len(versions)-1])
	}
	var versionsToMigrate = []string{}
	for _, v := range versions {
		ver := version.ToVersion(v)
		if ver.EarlierThan(d.fromVersion) {
			continue
		}
		if d.fromVersion.Equals(ver) {
			continue
		}
		if ver.LaterThan(maxVersion) {
			continue
		}
		versionsToMigrate = append(versionsToMigrate, ver.String())
	}
	return versionsToMigrate, nil
}
