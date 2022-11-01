package syslogger

import (
	"entsoe/src/config"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/logger"
)

// NewLogger Create a new logger from
// ALWAYS REMEMBER TO CLOSE FILE POINTER MANUALLY
func NewLogger(path string, verbose bool, sysLog bool, writeTofiles bool) (*logger.Logger, *os.File) {

	lf, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}

	newLogger := logger.Init("LoggerExample", verbose, sysLog, lf)
	if !writeTofiles {
		newLogger = logger.Init("LoggerExample", verbose, sysLog, ioutil.Discard)
	}

	return newLogger, lf
}

// NewDayLogger:
func NewDayLogger(lg *logger.Logger, filePointer *os.File) (*logger.Logger, *os.File) {

	// Close current logger
	if filePointer != nil {
		filePointer.Close()
	}
	if lg != nil {
		lg.Close()
	}

	loggingPath := fmt.Sprintf("%v/%v", config.GetLogBasePath(), config.GetFileNameSuffix())
	newLogger, lf := NewLogger(loggingPath, config.GetVerbose(), config.GetSystemLog(), config.GetSaveToFile())
	return newLogger, lf
}

func Close(lg *logger.Logger) {
	lg.Close()
}

func CreatePath() string {
	return ""
}
