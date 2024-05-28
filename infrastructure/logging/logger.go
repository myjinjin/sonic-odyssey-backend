package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger Logger

type Logger interface {
	Debug(msg string, fields ...zapcore.Field)
	Info(msg string, fields ...zapcore.Field)
	Warn(msg string, fields ...zapcore.Field)
	Error(msg string, fields ...zapcore.Field)
	Fatal(msg string, fields ...zapcore.Field)
	Sync()
}

type zapLogger struct {
	logger *zap.Logger
}

func Log() Logger {
	if logger == nil {
		logger, _ = NewZapLogger(true)
	}
	return logger
}

func NewZapLogger(development bool) (Logger, error) {
	var logger *zap.Logger
	var err error

	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	if development {
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	logger, err = config.Build()
	if err != nil {
		return nil, err
	}

	return &zapLogger{logger: logger}, nil
}

func (l *zapLogger) Debug(msg string, fields ...zapcore.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *zapLogger) Info(msg string, fields ...zapcore.Field) {
	l.logger.Info(msg, fields...)
}

func (l *zapLogger) Warn(msg string, fields ...zapcore.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *zapLogger) Error(msg string, fields ...zapcore.Field) {
	l.logger.Error(msg, fields...)
}

func (l *zapLogger) Fatal(msg string, fields ...zapcore.Field) {
	l.logger.Fatal(msg, fields...)
}

func (l *zapLogger) Sync() {
	l.logger.Sync()
}
