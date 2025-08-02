package loggers

import (
	"auth_service/pkg/settings"
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
	env string
}

func (l *LoggerZap) Error(
	errorMessage string,
	requestId string,
	fields ...zap.Field,
) {
	if l.env == "dev" {
		println("Error:", errorMessage)
		println("Request ID:", requestId)
		for _, field := range fields {
			println(fmt.Sprintf("%v", field))
		}
		println("Stack Trace:", string(debug.Stack()))
		return
	}

	l.Logger.Error(
		errorMessage,
		append([]zap.Field{
			zap.String("request_id", requestId),
			zap.Stack("stack_trace"),
		}, fields...)...,
	)
}

// Overwrite the Error method to accept a LogString
func (l *LoggerZap) Info(
	message string,
	requestId string,
	fields ...zap.Field,
) {
	l.Logger.Info(
		message,
		append([]zap.Field{
			zap.String("request_id", requestId),
		}, fields...)...,
	)
}

func (l *LoggerZap) Warn(
	message string,
	requestId string,
	fields ...zap.Field,
) {
	l.Logger.Warn(
		message,
		append([]zap.Field{
			zap.String("request_id", requestId),
		}, fields...)...,
	)
}

// Create a new LoggerZap instance with the provided configuration
func NewLogger(cfg settings.Log) *LoggerZap {
	logLevel := cfg.Level
	var level zapcore.Level = getLogLevelFromConfig(logLevel)

	encoder := getEncoder()
	hook := lumberjack.Logger{
		Filename:   cfg.FileLogPath + time.Now().Format("2006010215") + "_auth.log",
		MaxSize:    cfg.MaxSize, // megabytes
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,   //days
		Compress:   cfg.Compress, // disabled by default
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level,
	)

	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "dev"
	}

	return &LoggerZap{
		Logger: zap.New(core, zap.AddCaller()),
		env:    env,
	}
}

// create a new zap encoder with custom configuration
func getEncoder() zapcore.Encoder {
	// Set the encoder configuration
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getLogLevelFromConfig returns the zapcore.Level based on the log level string from the config
func getLogLevelFromConfig(logLevel string) zapcore.Level {
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}
	return level
}
