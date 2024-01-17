package logger

import "context"

func Info(ctx context.Context, args ...interface{}) {
	log(ctx).Info(args...)
}

func Debug(ctx context.Context, args ...interface{}) {
	log(ctx).Debug(args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	log(ctx).Warn(args...)
}

func Error(ctx context.Context, args ...interface{}) {
	log(ctx).Error(args...)
}

func Fatal(ctx context.Context, args ...interface{}) {
	logger.Fatal(args...)
}
