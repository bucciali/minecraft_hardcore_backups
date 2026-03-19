package configs

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type AppConfig struct {
	WorldName      string
	SavesPath      string
	WorldPath      string
	BackupRoot     string
	WorldBackupDir string
	Frequency      int
	BackupCount    int
}

const backupRoot = "backups"

func GetMinecraftSavesPath() (string, error) {
	appdata := os.Getenv("APPDATA")
	if appdata == "" {
		return "", errors.New("APPDATA environment variable not found")
	}
	fullPath := filepath.Join(appdata, ".minecraft", "saves")
	_, err := os.Stat(fullPath)
	if err != nil {
		return "", err
	}
	return fullPath, nil

}

func AskInfoAboutWorld() (string, int, int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Write name of the world you would like to save :)")
	name, err := reader.ReadString('\n')
	if err != nil {
		return "", 0, 0, err
	}
	name = strings.TrimSpace(name)
	if name == "" {
		return "", 0, 0, errors.New("world name cannot be empty")
	}
	fmt.Println("Write the backup frequency in minutes:)")
	flagConv := false
	var frequency int
	for flagConv != true {
		frequencystr, err := reader.ReadString('\n')
		if err != nil {
			return "", 0, 0, err
		}
		frequencystr = strings.TrimSpace(frequencystr)
		frequency, err = strconv.Atoi(frequencystr)
		if err == nil && frequency > 0 {
			flagConv = true
		} else {
			fmt.Println("enter the correct integer number of minutes")
		}

	}

	fmt.Println("write how many recent backups to keep \n (at least 2 is recommended because the program can make a backup at the time of death):)")

	flagConv = false
	var backupCount int
	for flagConv != true {
		backupCountstr, err := reader.ReadString('\n')
		if err != nil {
			return "", 0, 0, err
		}
		backupCountstr = strings.TrimSpace(backupCountstr)
		backupCount, err = strconv.Atoi(backupCountstr)
		if err == nil && backupCount > 0 {
			flagConv = true
		} else {
			fmt.Println("enter the correct integer number of backups")
		}

	}

	return name, frequency, backupCount, nil
}

func BuildConfig() (*AppConfig, error) {
	path, err := GetMinecraftSavesPath()
	if err != nil {
		return nil, err
	}
	worldName, frequency, backupCount, err := AskInfoAboutWorld()
	if err != nil {
		return nil, err
	}
	worldPath := filepath.Join(path, worldName)
	_, err = os.Stat(worldPath)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("world folder %q not found", worldName)
	}
	if err != nil {
		return nil, err
	}
	worldBackupDir := filepath.Join(backupRoot, worldName)
	err = os.MkdirAll(worldBackupDir, 0755)
	if err != nil {
		return nil, err
	}
	cfg := AppConfig{
		WorldName:      worldName,
		SavesPath:      path,
		WorldPath:      worldPath,
		BackupRoot:     backupRoot,
		WorldBackupDir: worldBackupDir,
		Frequency:      frequency,
		BackupCount:    backupCount,
	}
	return &cfg, nil
}
