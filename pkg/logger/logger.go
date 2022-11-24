package logger

import "os"

// type Fields map[string]interface{}

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	// Infow(format string, args ...interface{})
	// Warnf(format string, args ...interface{})
	// Errorf(format string, args ...interface{})
	// Fatalf(format string, args ...interface{})
	// Panicf(format string, args ...interface{})
	// // WithFields(keyValues Fields) Logger
	// Sync() error
}

// Log Level
const (
	// Debug has verbose message
	Debug = "debug"
	// Info is default log level
	Info = "info"
	// Warn is for logging messages about possible issues
	Warn = "warn"
	// Error is for logging errors
	Error = "error"
	// Fatal is for logging fatal messages. The system shutsdown after logging the message.
	Fatal = "fatal"
)

// global logger
var gLogger Logger

// func init() {

// }

func New() Logger {
	f, err := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	logger := newZapLogger(f, Debug)
	gLogger = logger
	return gLogger
}

func GetLogger() Logger {
	return gLogger
}
