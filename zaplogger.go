// Package zap_open5gs_format defines custom zap configuration for effective logging with Open5GS style
package zap_open5gs_format

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// customTimeEncoder formats the time as "MM/DD HH:mm:ss"
func customTimeEncoder(timeFormat string) zapcore.TimeEncoder {
	return func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(timeFormat))
	}
}

// customCallerEncoder formats the caller as "(file:line)"
func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("(" + caller.TrimmedPath() + ")")
}

// NewLogger creates a logger with the custom format
func NewLogger() *zap.Logger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		EncodeTime:     customTimeEncoder("01/02 15:04:05"),
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeCaller:   customCallerEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		zap.DebugLevel,
	)

	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}
