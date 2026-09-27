package version

func (a *Version) EarlierThan(b *Version) bool {
	if a.Equals(b) {
		return false
	}

	if a.Major > b.Major {
		return false
	} else if a.Major < b.Major {
		return true
	}

	if a.Minor > b.Minor {
		return false
	} else if a.Minor < b.Minor {
		return true
	}

	if a.Patch > b.Patch {
		return false
	}
	return true
}
