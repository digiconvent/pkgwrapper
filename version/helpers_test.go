package version_test

import "github.com/digiconvent/pkgwrapper/version"

func v(a string) *version.Version {
	return version.ToVersion(a)
}
