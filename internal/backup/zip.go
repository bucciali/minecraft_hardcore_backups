package backup

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func ZipWorld(sourceDir string, zipPath string) error {
	zipFile, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	return filepath.WalkDir(sourceDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}

		zipName := filepath.ToSlash(relPath)

		writer, err := zipWriter.Create(zipName)
		if err != nil {
			return err
		}

		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		_, err = io.Copy(writer, srcFile)
		if err != nil {
			return err
		}

		return nil
	})
}
