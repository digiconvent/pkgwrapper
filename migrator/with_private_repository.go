package migrator

import migrate_packages_internal "github.com/digiconvent/pkgwrapper/internal"

func (d *data) WithPrivateRepository(username string, repository string, token string) (repoPackageWrapper, error) {
	err := migrate_packages_internal.DownloadExtractDeleteZip(username, repository, token, d.verbose)
	if err != nil {
		return nil, err
	}
	return d, nil
}
