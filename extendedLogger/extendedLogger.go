package extendedlogger

import (
	"log"
	"os"
)

type LogLevel int

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
)

type LogExtended struct {
	*log.Logger
	LogLevel
}

func (l *LogExtended) SetLogLevel(logLevel LogLevel) {
	l.LogLevel = logLevel
}

func NewLogExtended() *LogExtended {
	return &LogExtended{
		LogLevel: LogLevelError,
		Logger:   log.New(os.Stderr, "", log.LstdFlags),
	}
}

func (l *LogExtended) Infoln(message string) {
	l.println(LogLevelInfo, "INFO", message)

}
func (l *LogExtended) Warnln(message string) {
	l.println(LogLevelWarning, "WARN", message)

}

func (l *LogExtended) Errorln(message string) {
	l.println(LogLevelError, "ERROR", message)

}
func (l *LogExtended) println(srcLogLvl LogLevel, prefix, msg string) {
	if l.LogLevel < srcLogLvl {
		return
	}

	l.Logger.Println(prefix, msg)
}
