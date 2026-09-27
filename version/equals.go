package version

func (a *Version) Equals(b *Version) bool {
	return a.Major == b.Major && a.Minor == b.Minor && a.Patch == b.Patch
}
