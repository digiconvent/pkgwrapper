package version

func (v *Version) IsUninitialised() bool {
	return v.Major == -1 && v.Minor == -1 && v.Patch == -1
}
