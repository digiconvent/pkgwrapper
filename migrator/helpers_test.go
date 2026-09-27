package migrator_test

import (
	"fmt"
	"os"
	"path"
	"slices"
	"testing"

	"github.com/digiconvent/pkgwrapper/migrator"
	"github.com/digiconvent/pkgwrapper/version"
)

var currentVersion = version.ToVersion("-1.0.0")
var targetVersion = version.ToVersion("1.0.1")

// [ 0.1.0 / 1.0.0 / 1.0.1 ]
var pkgFolder = "migrator/.test/pkg" // fixture's path within the downloaded repo's root
var username = "digiconvent"
var reponame = "pkgwrapper"

func cleanup() {
	err := os.RemoveAll(path.Join(os.TempDir(), "migrate_packages"))
	if err != nil {
		fmt.Println(err)
	}
}

func testPackageWrapper(manager migrator.PackageWrapper, t *testing.T) {
	isPackages, err := manager.GetPackages()
	shouldPackages := []string{"iam", "post", "sys"}
	if err != nil {
		t.Fatal("did not expect err but got " + err.Error())
	}
	slices.Sort(isPackages)
	if !slices.Equal(isPackages, shouldPackages) {
		t.Fatal("expected", shouldPackages, "instead got", isPackages)
	}

	isVersions, err := manager.Versions()
	shouldVersions := []string{"0.0.0", "0.1.0", "1.0.0", "1.0.1", "12.0.4"}
	if err != nil {
		t.Fatal("expected err to be nil, instead got ", err.Error())
	}

	if !slices.Equal(isVersions, shouldVersions) {
		t.Fatal("expected", shouldVersions, "instead got", isVersions)
	}
	isVersions, err = manager.VersionsToMigrate()
	if err != nil {
		t.Fatal("did not expect err")
	}
	shouldVersions = []string{"0.1.0", "1.0.0", "1.0.1"}
	if !slices.Equal(isVersions, shouldVersions) {
		t.Fatal("expected", shouldVersions, "instead got", isVersions)
	}

	// todo: implement check migration to make sure the correct scripts are returned for the test cases
	cleanup()
}
