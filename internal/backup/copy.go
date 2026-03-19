package backup

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func CopyFile(srcPath string, dstPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("problems with opening: %w", err)
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("problems with creating: %w", err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {

		return fmt.Errorf("problems with copying: %w", err)
	}
	return nil
}

func CopyWorld(src string, dst string) error {
	return filepath.WalkDir(src, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err

		}
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(dst, relPath)
		if d.IsDir() {
			return os.MkdirAll(targetPath, 0755)
		} else if d.Name() == "session.lock" {
			return nil
		}
		err = os.MkdirAll(filepath.Dir(targetPath), 0755)
		if err != nil {
			return err
		}
		err = CopyFile(path, targetPath)
		if err != nil {
			if strings.Contains(err.Error(), "problems with opening:") {
				fmt.Println("skipping file:", path, "error:", err)
				return nil
			} else {
				return err
			}
		}
		return nil

	})
}
