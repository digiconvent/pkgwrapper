package registry

import "io/fs"

type Config struct {
	DataRoot      string
	TargetVersion string
	EmbeddedFS    fs.FS
}
