package backup

import (
	"os"
	"path/filepath"
	"sort"
	"time"
)

type backupFile struct {
	path    string
	modTime time.Time
}

func CleanupOldBackups(worldBackupDir string, keepLast int) error {
	entries, err := os.ReadDir(worldBackupDir)
	if err != nil {
		return err
	}
	slice := make([]backupFile, 0, 0)
	for _, values := range entries {
		if values.IsDir() || filepath.Ext(values.Name()) != ".zip" {
			continue
		}
		filePath := filepath.Join(worldBackupDir, values.Name())
		info, err := values.Info()
		if err != nil {
			return err
		}
		modTime := info.ModTime()
		backup := backupFile{
			path:    filePath,
			modTime: modTime,
		}
		slice = append(slice, backup)

	}
	sort.Slice(slice, func(a, b int) bool {
		return slice[a].modTime.After(slice[b].modTime)
	})
	if len(slice) <= keepLast {
		return nil
	}
	for _, backup := range slice[keepLast:] {
		err := os.Remove(backup.path)
		if err != nil {
			return err
		}
	}
	return nil
}
