package migrate_packages_internal

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/digiconvent/req"
)

// no need to test this it works trust me bro
func DownloadExtractDeleteZip(owner, name, token string, verbose bool) error {
	targetDir := path.Join(os.TempDir(), "migrate_packages")
	targetZip := path.Join(os.TempDir(), "migrate_packages.zip")
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/zipball/main", owner, name)

	r := req.Get(url).Header("Accept", "application/vnd.github.v3+json")
	if token != "" {
		r = r.Header("Authorization", "Bearer "+token)
	}

	body, _, err := r.Do()
	if err != nil {
		return fmt.Errorf("repository %s/%s: %w", owner, name, err)
	}

	if verbose {
		fmt.Println("[Migrate Packages] Downloading " + url)
	}
	outFile, err := os.Create(targetZip)
	if err != nil {
		if verbose {
			fmt.Println("[Migrate Packages] Could not create " + targetZip + ": " + err.Error())
		}
		return err
	}
	defer outFile.Close()
	defer os.Remove(targetZip)

	_, err = outFile.Write(body)
	if err != nil {
		return err
	}

	openCloseReader, _ := zip.OpenReader(outFile.Name())
	defer openCloseReader.Close()

	var prefix string = ""
	for _, file := range openCloseReader.File {
		if prefix == "" {
			prefix = file.Name
			if verbose {
				fmt.Println("[Migrate Packages] Setting prefix to " + prefix + " (we want to have the zipped files, not nested inside another folder)")
			}
		}
		fileName, _ := strings.CutPrefix(file.Name, prefix)
		target := path.Join(targetDir, fileName)
		if file.FileInfo().IsDir() {
			err := os.MkdirAll(target, file.Mode())
			if err != nil {
				return err
			}
		} else {
			srcFile, err := file.Open()
			if err != nil {
				return err
			}
			defer srcFile.Close()

			destFile, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				return err
			}
			defer destFile.Close()

			_, err = io.Copy(destFile, srcFile)
			if err != nil {
				return err
			}
		}
	}

	err = os.Remove(targetZip)

	if err != nil {
		if verbose {
			fmt.Println("[Migrate Packages] Could not clean up " + targetZip)
		}
	}
	return nil
}
