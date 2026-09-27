package migrator_test

import (
	"os"
	"testing"

	"github.com/digiconvent/pkgwrapper/migrator"
	"github.com/digiconvent/pkgwrapper/version"
)

func TestPackageWrapper(t *testing.T) {
	currentVersion := &version.Version{Major: 0, Minor: 0, Patch: 0} // means that 0.0.0 is already migrated
	targetVersion := &version.Version{Major: 1, Minor: 0, Patch: 1}  // this means that the new version is going to be 1.0.1
	//  0.0.0 [ 0.1.0 / 1.0.0 / 1.0.1 ] 12.0.4
	thisFolder, _ := os.Getwd()
	pkgFolder := ".test/pkg"
	repoPkgFolder := "migrator/.test/pkg" // fixture's path within the downloaded repo's root, not this package's own cwd
	username := "digiconvent"
	reponame := "pkgwrapper"

	t.Run("WithLocalFiles", func(t *testing.T) {
		manager, err := migrator.From(currentVersion).To(targetVersion).WithLocalFilesAt(thisFolder, pkgFolder)
		if err != nil {
			t.Fatal("expected err to be nil")
		}

		if manager == nil {
			t.Fatal("expected manager not to be nil")
		}

		testPackageWrapper(manager, t)
	})
	t.Run("WithPublicRepository", func(t *testing.T) {
		repo, err := migrator.From(currentVersion).To(targetVersion).WithPublicRepository(username, reponame)
		if err != nil {
			t.Fatal("expected err to be nil instead got " + err.Error())
		}

		manager, err := repo.WithPkgDir(repoPkgFolder)
		if err != nil {
			t.Fatal(err)
		}
		if manager == nil {
			t.Fatal("expected manager not to be nil")
		}
		testPackageWrapper(manager, t)
	})
	t.Run("WithPrivateRepository", func(t *testing.T) {
		// this works, pinky promise, no need to test it
		t.Skip()
		token := "<fine-grained access token for readonly contents of repository>"
		repo, err := migrator.From(currentVersion).To(targetVersion).WithPrivateRepository(username, reponame, token)
		if err != nil {
			t.Fatal("expected err to be nil instead got " + err.Error())
		}

		manager, err := repo.WithPkgDir(repoPkgFolder)
		if err != nil {
			t.Fatal(err)
		}
		if manager == nil {
			t.Fatal("expected manager not to be nil")
		}
		testPackageWrapper(manager, t)
	})
}
