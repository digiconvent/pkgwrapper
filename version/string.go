package version

import "fmt"

func (a *Version) String() string {
	return fmt.Sprintf("%v.%v.%v", a.Major, a.Minor, a.Patch)
}
