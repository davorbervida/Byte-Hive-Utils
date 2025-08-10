package catlog

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/davorbervida/Byte-Hive-Utils/v2/config"
)

func Info(infoMessage string) {

	saveLog(infoMessage, "INFO")
	deleteOldLogs()
}

func Warning(warningMessage string) {

	saveLog(warningMessage, "WARNING")
	deleteOldLogs()
}

func Debug(debugMessage string) {

	saveLog(debugMessage, "DEBUG")
	deleteOldLogs()
}

func Error(errorMessage interface{}) {

	switch errorMessage.(type) {
	case error:
		saveLog(fmt.Sprintf("%s", errorMessage), "ERROR")

	case string:
		saveLog(errorMessage.(string), "ERROR")

	default:
		return
	}

	deleteOldLogs()
}

func deleteOldLogs() {

	now := time.Now()
	keeptFilenames := make(map[string]bool, config.Get.Logs.NumberOfDays)
	keeptFilenames[now.Format("2006-01-02.log")] = true

	var uint8i uint8 = 1
	for i := uint8i; i < config.Get.Logs.NumberOfDays; i++ {

		now = now.AddDate(0, 0, -1)
		keeptFilenames[now.Format("2006-01-02.log")] = true
	}

	entries, err := os.ReadDir(config.Get.Logs.Path)
	if err != nil {

		saveLog(err.Error(), "ERROR")
		return
	}

	for _, entry := range entries {

		if _, ok := keeptFilenames[entry.Name()]; ok {

			continue
		}

		filePath := fmt.Sprintf("%s/%s", config.Get.Logs.Path, entry.Name())
		if err := os.Remove(filePath); err != nil {

			saveLog(err.Error(), "ERROR")
			return
		}
	}
}

func saveLog(message string, logType string) {

	now := time.Now()
	dateNow := now.Format("2006-01-02")
	timeNow := now.Format("15:04:05")

	_, fileName, lineNumber, ok := runtime.Caller(2)
	if !ok {

		return
	}

	fileName = filepath.Base(fileName)
	logMessage := fmt.Sprintf("%s %s %s:%d\n%s\n", logType, timeNow, fileName, lineNumber, message)
	logFilePath := fmt.Sprintf("%s/%s.log", config.Get.Logs.Path, dateNow)

	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {

		fmt.Println("ERROR WHILE WRITING LOG MESSAGE:", logMessage)
		fmt.Println(err)
		return
	}

	defer f.Close()
	f.WriteString(logMessage)
}
