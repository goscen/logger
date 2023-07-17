package main

import extendedlogger "github.com/goscen/logger/extendedLogger"

func main() {
	logger := extendedlogger.NewLogExtended()
	logger.SetLogLevel(extendedlogger.LogLevelWarning)
	logger.Infoln("Hi")
	logger.Warnln("Hello")
	logger.Errorln("World")
}
