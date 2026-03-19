package main

import (
	"fmt"
	"mcbackup/configs"
	"mcbackup/internal/backup"
	"time"
)

func main() {
	fmt.Println("=== Minecraft Backup Tool ===")
	cfg, err := configs.BuildConfig()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	zipPath, err := backup.RunSingleBackup(cfg)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("backup created successfully")
		fmt.Println("archive:", zipPath)

	}
	fmt.Println("Starting backups...")
	ticker := time.NewTicker(time.Duration(cfg.Frequency) * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		zipPath, err := backup.RunSingleBackup(cfg)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}

		fmt.Println("["+time.Now().Format("15:04:05")+"] backup created:", zipPath)
		fmt.Println("archive:", zipPath)
	}

}
