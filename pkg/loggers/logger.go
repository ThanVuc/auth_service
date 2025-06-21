package loggers

import (
	"auth_service/internal/models"
	"auth_service/pkg/response"
	"auth_service/pkg/settings"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func (l *LoggerZap) Error(
	errRes response.ErrorResponse,
	requestId string,
	stack []byte,
) {
	l.Logger.Error(errRes.Message,
		zap.Int("status_code", errRes.StatusCode),
		zap.String("error_reason", errRes.CodeReason),
		zap.String("message", errRes.Message),
		zap.String("request_id", requestId),
		zap.String("stack", string(stack)),
	)
}

func (l *LoggerZap) ErrorString(message string, fields ...zap.Field) {
	l.Logger.Error(message, fields...)
}

// Overwrite the Error method to accept a LogString
func (l *LoggerZap) Info(message models.LogString, fields ...zap.Field) {
	l.Logger.Info(message.String(), fields...)
}

// Overwrite the Info method to accept a string message
func (l *LoggerZap) InfoString(message string, fields ...zap.Field) {
	l.Logger.Info(message, fields...)
}

// Create a new LoggerZap instance with the provided configuration
func NewLogger(cfg settings.Log) *LoggerZap {
	logLevel := cfg.Level
	var level zapcore.Level = getLogLevelFromConfig(logLevel)

	encoder := getEncoder()
	hook := lumberjack.Logger{
		Filename:   cfg.FileLogPath + time.Now().Format("2006010215") + "_author.log",
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
	return &LoggerZap{
		Logger: zap.New(core, zap.AddCaller()),
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
	return zapcore.NewJSONEncoder(encoderConfig)
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
