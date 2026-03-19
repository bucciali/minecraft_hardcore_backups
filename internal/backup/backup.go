package backup

import (
	"mcbackup/configs"
	"os"
	"path/filepath"
	"time"
)

func RunSingleBackup(cfg *configs.AppConfig) (string, error) {
	tempPath := filepath.Join("temp", cfg.WorldName)
	err := os.RemoveAll(tempPath)
	if err != nil {
		return "", err
	}
	err = CopyWorld(cfg.WorldPath, tempPath)
	if err != nil {
		return "", err
	}
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	zipFileName := cfg.WorldName + "_" + timestamp + ".zip"
	zipPath := filepath.Join(cfg.WorldBackupDir, zipFileName)
	err = ZipWorld(tempPath, zipPath)
	if err != nil {
		return "", err
	}
	err = os.RemoveAll(tempPath)
	if err != nil {
		return "", err
	}
	err = CleanupOldBackups(cfg.WorldBackupDir, cfg.BackupCount)
	if err != nil {
		return "", err
	}
	return zipPath, nil

}
