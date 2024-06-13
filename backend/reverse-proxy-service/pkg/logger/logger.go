package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

type Logger struct {
	sugarLogger *zap.SugaredLogger
	serviceName string
}

func (l *Logger) Info(format string, v ...any) {
	//var builder strings.Builder
	//builder.WriteString(l.serviceName)
	//builder.WriteString(format)

	l.sugarLogger.Infof(format, v...)
}

func (l *Logger) Error(format string, v ...any) {
	var builder strings.Builder
	builder.WriteString(l.serviceName)
	builder.WriteString(format)

	l.sugarLogger.Errorf(builder.String(), v...)
}

func (l *Logger) Fatal(format string, v ...any) {
	var builder strings.Builder
	builder.WriteString(l.serviceName)
	builder.WriteString(format)

	l.sugarLogger.Fatalf(builder.String(), v...)
}

func New(serviceName string, isTesting bool) *Logger {
	consoleConfig := zap.NewDevelopmentConfig()
	consoleConfig.DisableStacktrace = true
	consoleConfig.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	consoleConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	fileConfig := zap.NewDevelopmentConfig()
	fileConfig.DisableStacktrace = true
	fileConfig.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	fileConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	fileConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	fileConfig.Encoding = "json"
	fileConfig.OutputPaths = []string{"/Users/smirnovn/go-workspace/university/graduate-work/backend/logfile.log"}

	if isTesting {
		consoleConfig.Level.SetLevel(zapcore.FatalLevel)
		fileConfig.Level.SetLevel(zapcore.FatalLevel)
	}

	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(consoleConfig.EncoderConfig),
		os.Stdout,
		consoleConfig.Level,
	)

	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(fileConfig.EncoderConfig),
		zapcore.AddSync(openFileOrDie(fileConfig.OutputPaths[0])),
		fileConfig.Level,
	)

	core := zapcore.NewTee(consoleCore, fileCore)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	logger = logger.Named(serviceName)
	sugarLogger := logger.Sugar()

	l := &Logger{
		sugarLogger: sugarLogger,
	}

	return l
}

func openFileOrDie(path string) zapcore.WriteSyncer {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return zapcore.AddSync(file)
}

type Log interface {
	Info(format string, v ...any)
	Error(format string, v ...any)
}

var (
	DefaultLogger Log = New("test", false)
)

func Info(format string, v ...any) {
	DefaultLogger.Info(format, v...)
}

func Error(format string, v ...any) {
	DefaultLogger.Error(format, v...)
}
