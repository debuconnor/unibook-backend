package unibookBackend

import (
	"log"
	"os"
	"path/filepath"
)

func Log(msg ...interface{}) {
	if !IS_DEBUG {
		return
	}
	msg = append([]interface{}{"UNIBOOK-BACKEND: "}, msg...)
	log.Println(msg...)
}

func SaveLog(filename string, msg ...interface{}) {
	path := "./log/"

	if filename == "" {
		filename = "unibook-backend.log"
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			Log("Failed to create path folder:", err)
			return
		}
	}

	filePath := filepath.Join(path, filename+".log")
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		Log("Failed to open error log file:", err)
		return
	}
	defer file.Close()

	msg = append([]interface{}{"UNIBOOK-BACKEND: "}, msg...)
	log.SetOutput(file)
	log.Println(msg...)
	Log(msg...)
}

func Error(code int) {
	SaveLog("Error: ", code)
	Log("Error: ", code)
}
