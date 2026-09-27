package registry

import "github.com/digiconvent/pkgwrapper/db"

func recordVersion(database db.DatabaseInterface, name, version string) error {
	if _, err := database.Exec(`create table if not exists _meta (
		key   text primary key,
		value text not null
	)`); err != nil {
		return err
	}
	if _, err := database.Exec(`create table if not exists _migrations (
		version    text     not null,
		applied_at datetime not null default (datetime('now')),
		primary key (version)
	)`); err != nil {
		return err
	}
	if _, err := database.Exec(`insert or replace into _meta (key, value) values ('package', ?)`, name); err != nil {
		return err
	}
	if _, err := database.Exec(`insert or replace into _meta (key, value) values ('version', ?)`, version); err != nil {
		return err
	}
	_, err := database.Exec(`insert or ignore into _migrations (version) values (?)`, version)
	return err
}
