package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type ConfigDatabaseStructure struct {
	MainPath  string `json:"main"`
	FilesPath string `json:"files"`
}

type ConfigLogsStructure struct {
	NumberOfDays uint8  `json:"keepNumberOfDays"`
	Path         string `json:"directory"`
}

type ConfigStructure struct {
	Port     uint16                  `json:"port"`
	Logs     ConfigLogsStructure     `json:"logs"`
	Database ConfigDatabaseStructure `json:"database"`
}

var Get ConfigStructure

func Refresh() {
	file, err := os.Open("./data/settings.json")
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	defer file.Close()

	// Read all file content
	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	if err = json.Unmarshal(bytes, &Get); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
}
