package logs

import (
	"log"
	"os"
)

func SetupLogger(logFilePath string) *log.Logger {
	file, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	return log.New(file, "ERROR: ", log.LstdFlags|log.Lshortfile)
}
