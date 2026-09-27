package migrator

import migrate_packages_internal "github.com/digiconvent/pkgwrapper/internal"

func (d *data) WithPublicRepository(username string, repository string) (repoPackageWrapper, error) {
	err := migrate_packages_internal.DownloadExtractDeleteZip(username, repository, "", d.verbose)
	if err != nil {
		return nil, err
	}
	return d, nil
}
