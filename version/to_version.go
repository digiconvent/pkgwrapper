package version

import (
	"strconv"
	"strings"
)

func ToVersion(a string) *Version {
	segments := strings.Split(a, ".")
	if len(segments) != 3 {
		return nil
	}
	c := []int{}
	for _, segment := range segments {
		n, err := strconv.Atoi(segment)
		if err != nil {
			return nil
		}
		c = append(c, n)
	}
	return &Version{
		Major: c[0],
		Minor: c[1],
		Patch: c[2],
	}
}
