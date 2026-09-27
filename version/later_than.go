package version

func (a *Version) LaterThan(b *Version) bool {
	return b.EarlierThan(a)
}
