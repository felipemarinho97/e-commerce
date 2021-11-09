package common

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var Logger = NewLog()

type Log struct {
	Logger *log.Logger
}

func NewLog() Log {
	return Log{
		Logger: &log.Logger{
			Out:   os.Stderr,
			Level: log.DebugLevel,
			Formatter: &prefixed.TextFormatter{
				DisableColors:   true,
				TimestampFormat: "2006-01-02 15:04:05",
				FullTimestamp:   true,
				ForceFormatting: true,
			},
		},
	}
}

// formatLogMessage format a message using pipe pattern
func formatLogMessage(msg ...string) string {
	if len(msg) > 0 {
		return "| " + strings.Join(msg, " | ")
	}
	return ""
}

// Log log a message with a given level
func (l Log) Log(logLevel log.Level, msg ...string) string {
	s := formatLogMessage(msg...)
	l.Logger.Logln(logLevel, s)
	return s
}

// LogInfo log a message in INFO level
func (l Log) LogInfo(msg ...string) string {
	return l.Log(log.InfoLevel, msg...)
}

// LogError log a message in ERROR level
func (l Log) LogError(msg ...string) string {
	return l.Log(log.ErrorLevel, msg...)
}

// LogDebug log a message in DEBUG level
func (l Log) LogDebug(msg ...string) string {
	return l.Log(log.DebugLevel, msg...)
}

// LogFatal log a message in FATAL level
func (l Log) LogFatal(msg ...string) string {
	return l.Log(log.FatalLevel, msg...)
}
