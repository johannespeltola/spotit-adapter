package config

import "gopkg.in/guregu/null.v3"

var (
	basePath   = ""
	fileSuffix = ""
	saveToFile = false
	syslog     = false
	stdOutLog  = true
	verbose    = true
)

type LoggerConfiguration struct {
	LogBasePath    null.String `json:"logBasePath"`
	FileNameSuffix null.String `json:"fileNameSuffix"`
	SaveToFile     null.Bool   `json:"saveToFile"`
	SystemLog      null.Bool   `json:"systemLog"`     // Logging to OS (Windows event log, Linux syslog)
	Verbose        null.Bool   `json:"verbose"`       // Includes info/warning level to log file
	StdOutLogOnly  null.Bool   `json:"stdOutLogOnly"` // Logging to stdout only
}

func GetLogBasePath() string {
	if cfg.LoggerConf.LogBasePath.Valid {
		return cfg.LoggerConf.LogBasePath.String
	}
	return basePath
}

func GetFileNameSuffix() string {
	if cfg.LoggerConf.FileNameSuffix.Valid {
		return cfg.LoggerConf.FileNameSuffix.String
	}
	return fileSuffix
}

func GetSaveToFile() bool {
	if cfg.LoggerConf.SaveToFile.Valid {
		return cfg.LoggerConf.SaveToFile.Bool
	}
	return saveToFile
}

func GetSystemLog() bool {
	if cfg.LoggerConf.SystemLog.Valid {
		return cfg.LoggerConf.SystemLog.Bool
	}
	return syslog
}

func GetStdOutLog() bool {
	if cfg.LoggerConf.StdOutLogOnly.Valid {
		return cfg.LoggerConf.StdOutLogOnly.Bool
	}
	return stdOutLog
}

func GetVerbose() bool {
	if cfg.LoggerConf.Verbose.Valid {
		return cfg.LoggerConf.Verbose.Bool
	}
	return verbose
}
