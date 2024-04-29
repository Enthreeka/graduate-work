package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

const serviceName = "[REVERSE-PROXY] "

type Logger struct {
	sugarLogger *zap.SugaredLogger
}

func (l *Logger) Info(format string, v ...any) {
	var builder strings.Builder
	builder.WriteString(serviceName)
	builder.WriteString(format)

	l.sugarLogger.Infof(builder.String(), v...)
}

func (l *Logger) Error(format string, v ...any) {
	var builder strings.Builder
	builder.WriteString(serviceName)
	builder.WriteString(format)

	l.sugarLogger.Errorf(format, v...)
}

func (l *Logger) Fatal(format string, v ...any) {
	var builder strings.Builder
	builder.WriteString(serviceName)
	builder.WriteString(format)

	l.sugarLogger.Fatalf(builder.String(), v...)
}

func New() *Logger {
	config := zap.NewDevelopmentConfig()
	config.DisableStacktrace = true

	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	logger, _ := config.Build(zap.AddCallerSkip(1))
	sugarLogger := logger.Sugar()

	log := &Logger{
		sugarLogger: sugarLogger,
	}

	return log
}

type Log interface {
	Info(format string, v ...any)
	Error(format string, v ...any)
}

var (
	DefaultLogger Log = New()
)

func Info(format string, v ...any) {
	DefaultLogger.Info(format, v...)
}

func Error(format string, v ...any) {
	DefaultLogger.Error(format, v...)
}
