package registry

import "os"

func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return findProjectRootFrom(dir)
}
