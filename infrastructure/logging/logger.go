package logging

import (
	"go.uber.org/zap"
)

type Logger interface {
	Debug(msg string, fields ...interface{})
	Info(msg string, fields ...interface{})
	Warn(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
	Fatal(msg string, fields ...interface{})
}

type zapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(development bool) (Logger, error) {
	var logger *zap.Logger
	var err error

	if development {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		return nil, err
	}

	return &zapLogger{logger: logger}, nil
}

func (l *zapLogger) Debug(msg string, fields ...interface{}) {
	l.logger.Debug(msg, zap.Any("fields", fields))
}

func (l *zapLogger) Info(msg string, fields ...interface{}) {
	l.logger.Info(msg, zap.Any("fields", fields))
}

func (l *zapLogger) Warn(msg string, fields ...interface{}) {
	l.logger.Warn(msg, zap.Any("fields", fields))
}

func (l *zapLogger) Error(msg string, fields ...interface{}) {
	l.logger.Error(msg, zap.Any("fields", fields))
}

func (l *zapLogger) Fatal(msg string, fields ...interface{}) {
	l.logger.Fatal(msg, zap.Any("fields", fields))
}
