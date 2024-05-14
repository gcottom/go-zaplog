package zaplog

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func CreateAndInject(ctx context.Context) context.Context {
	logger := zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(zapcore.EncoderConfig{TimeKey: "ts", LevelKey: "level", NameKey: "logger", CallerKey: "caller", MessageKey: "message", StacktraceKey: "stacktrace",
		LineEnding: zapcore.DefaultLineEnding, EncodeLevel: zapcore.CapitalLevelEncoder, EncodeDuration: zapcore.StringDurationEncoder, EncodeTime: zapcore.EpochMillisTimeEncoder, EncodeCaller: zapcore.ShortCallerEncoder}),
		os.Stdout, zap.NewAtomicLevelAt(zap.InfoLevel)), zap.WithCaller(true), zap.AddCallerSkip(1))
	zap.ReplaceGlobals(logger)
	return InjectIntoContext(ctx, logger)
}
func InjectIntoContext(ctx context.Context, logger *zap.Logger) context.Context {
	//lint:ignore SA1029 ignore
	ctx = context.WithValue(ctx, "logger", logger)
	return ctx
}
func GetLoggerFromContext(ctx context.Context) *zap.Logger {
	logger := ctx.Value("logger")
	return logger.(*zap.Logger)
}
func InfoC(ctx context.Context, msg string, fields ...zapcore.Field) {
	logger := GetLoggerFromContext(ctx)
	logger.Info(msg, fields...)
}
func DebugC(ctx context.Context, msg string, fields ...zapcore.Field) {
	logger := GetLoggerFromContext(ctx)
	logger.Debug(msg, fields...)
}
func WarnC(ctx context.Context, msg string, fields ...zapcore.Field) {
	logger := GetLoggerFromContext(ctx)
	logger.Warn(msg, fields...)
}
func ErrorC(ctx context.Context, msg string, fields ...zapcore.Field) {
	logger := GetLoggerFromContext(ctx)
	logger.Error(msg, fields...)
}
func FatalC(ctx context.Context, msg string, fields ...zapcore.Field) {
	logger := GetLoggerFromContext(ctx)
	logger.Fatal(msg, fields...)
}
func PanicC(ctx context.Context, msg string, fields ...zapcore.Field) {
	logger := GetLoggerFromContext(ctx)
	logger.Panic(msg, fields...)
}
func LogC(ctx context.Context, lvl zapcore.Level, msg string, fields ...zapcore.Field) {
	logger := GetLoggerFromContext(ctx)
	logger.Log(lvl, msg, fields...)
}
