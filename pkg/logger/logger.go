package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)

const (
	logKeysVarCtx = "log_keys"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
}

func SetLogger(newLogger *logrus.Logger) {
	logger = newLogger
}

func getLogKeys(ctx context.Context) []string {
	logKeysFromCtx := ctx.Value(logKeysVarCtx)
	logKeys, ok := logKeysFromCtx.([]string)
	if !ok {
		return nil
	}

	return logKeys
}

func addLogKey(ctx context.Context, key string) context.Context {
	logKeysFromCtx := ctx.Value(logKeysVarCtx)
	logKeys, ok := logKeysFromCtx.([]string)
	if !ok {
		logKeys = make([]string, 0)
	}

	logKeys = append(logKeys, key)

	return context.WithValue(ctx, logKeysVarCtx, logKeys)
}

func With(ctx context.Context, key string, value any) context.Context {
	existValue := ctx.Value(key)
	if existValue == nil {
		ctx = addLogKey(ctx, key)
	}

	return context.WithValue(ctx, key, value)
}

func log(ctx context.Context) *logrus.Entry {
	entry := logger.WithContext(ctx)
	logKeys := getLogKeys(ctx)

	for _, logKey := range logKeys {
		logValue := ctx.Value(logKey)
		entry = entry.WithField(logKey, logValue)
	}

	return entry
}
