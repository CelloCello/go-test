package logger

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level = zapcore.Level

type ZapLogger struct {
	logger *zap.SugaredLogger // zap ensure that zap.Logger is safe for concurrent use
	level  zapcore.Level
}

var ZapLevel = map[string]zapcore.Level{
	Debug: zap.DebugLevel,
	Info:  zap.InfoLevel,
	Warn:  zap.WarnLevel,
	Error: zap.ErrorLevel,
	Fatal: zap.FatalLevel,
}

// const (
// 	InfoLevel   Level = zap.InfoLevel   // 0, default level
// 	WarnLevel   Level = zap.WarnLevel   // 1
// 	ErrorLevel  Level = zap.ErrorLevel  // 2
// 	DPanicLevel Level = zap.DPanicLevel // 3, used in development log
// 	// PanicLevel logs a message, then panics
// 	PanicLevel Level = zap.PanicLevel // 4
// 	// FatalLevel logs a message, then calls os.Exit(1).
// 	FatalLevel Level = zap.FatalLevel // 5
// 	DebugLevel Level = zap.DebugLevel // -1
// )

func (z *ZapLogger) Debugf(format string, args ...interface{}) {
	z.logger.Debugf(format, args...)
}

func (z *ZapLogger) Infof(format string, args ...interface{}) {
	z.logger.Infof(format, args...)
}

func (z *ZapLogger) Infow(format string, args ...interface{}) {
	z.logger.Infow(format, args...)
}

func (z *ZapLogger) Warnf(format string, args ...interface{}) {
	z.logger.Warnf(format, args...)
}

func (z *ZapLogger) Errorf(format string, args ...interface{}) {
	z.logger.Errorf(format, args...)
}

func (z *ZapLogger) Fatalf(format string, args ...interface{}) {
	z.logger.Fatalf(format, args...)
}

func (z *ZapLogger) Panicf(format string, args ...interface{}) {
	z.logger.Fatalf(format, args...)
}

// func (z *ZapLogger) WithFields(fields Fields) ZapLogger {
// 	var f = make([]interface{}, 0)
// 	for k, v := range fields {
// 		f = append(f, k)
// 		f = append(f, v)
// 	}
// 	newLogger := z.logger.With(f...)
// 	return &ZapLogger{newLogger}
// }

func (z *ZapLogger) Sync() error {
	return z.logger.Sync()
}

// function variables for all field types
// in github.com/uber-go/zap/field.go

// var (
//     Skip        = zap.Skip
//     Binary      = zap.Binary
//     Bool        = zap.Bool
//     Boolp       = zap.Boolp
//     ByteString  = zap.ByteString
//     ... ...
//     Float64     = zap.Float64
//     Float64p    = zap.Float64p
//     Float32     = zap.Float32
//     Float32p    = zap.Float32p
//     Durationp   = zap.Durationp
//     ... ...
//     Any         = zap.Any

//     Info   = std.Info
//     Warn   = std.Warn
//     Error  = std.Error
//     DPanic = std.DPanic
//     Panic  = std.Panic
//     Fatal  = std.Fatal
//     Debug  = std.Debug
// )

// // not safe for concurrent use
// func ResetDefault(l *ZapLogger) {
//     std = l
//     Info = std.Info
//     Warn = std.Warn
//     Error = std.Error
//     DPanic = std.DPanic
//     Panic = std.Panic
//     Fatal = std.Fatal
//     Debug = std.Debug
// }

// var std = New(os.Stderr, int8(InfoLevel))

// func Default() *ZapLogger {
//     return std
// }

// New create a new logger (not support log rotating).
func newZapLogger(writer io.Writer, level string) *ZapLogger {
	if writer == nil {
		panic("the writer is nil")
	}

	zapLv := ZapLevel[level]

	config := zap.NewDevelopmentConfig()
	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(config.EncoderConfig),
			zapcore.AddSync(os.Stderr),
			zapLv,
		),
		zapcore.NewCore(
			zapcore.NewJSONEncoder(config.EncoderConfig),
			zapcore.AddSync(writer),
			zapLv,
		),
	)
	baseLogger := zap.New(core)
	logger := &ZapLogger{
		logger: baseLogger.Sugar(),
		level:  zap.DebugLevel,
	}
	return logger
}

// func Sync() error {
//     if std != nil {
//         return std.Sync()
//     }
//     return nil
// }
