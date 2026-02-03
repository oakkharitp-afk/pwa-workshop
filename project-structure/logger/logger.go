package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger      *zap.SugaredLogger
	once        sync.Once
	atomicLevel = zap.NewAtomicLevelAt(zap.DebugLevel) // default is Debug
)

// Init initializes the logger depending on environment: local, dev, prod
func Init(env string) {
	once.Do(func() {
		var config zap.Config
		switch env {
		case "local":
			config = zap.NewDevelopmentConfig()
			config.Level = atomicLevel
			config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
			config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
			config.EncoderConfig.FunctionKey = zapcore.OmitKey
			config.EncoderConfig.TimeKey = "time"
			config.EncoderConfig.CallerKey = "caller"
			config.EncoderConfig.MessageKey = "msg"
			config.EncoderConfig.LevelKey = "level"

			baseLogger, _ := config.Build(
				zap.AddCaller(),
				zap.AddCallerSkip(1),
				zap.AddStacktrace(zap.PanicLevel), // stacktrace only for PanicLevel+
			)
			logger = baseLogger.Sugar()

		case "development":
			config = zap.NewDevelopmentConfig()
			config.Level = zap.NewAtomicLevelAt(zap.InfoLevel) // only Info+
			config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
			config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
			config.EncoderConfig.FunctionKey = zapcore.OmitKey
			config.EncoderConfig.TimeKey = "time"
			config.EncoderConfig.CallerKey = "caller"
			config.EncoderConfig.MessageKey = "msg"
			config.EncoderConfig.LevelKey = "level"

			baseLogger, _ := config.Build(
				zap.AddCaller(),
				zap.AddCallerSkip(1),
				zap.AddStacktrace(zap.DPanicLevel), // stacktrace from DPanicLevel+
			)
			logger = baseLogger.Sugar()

		case "production":

			config = zap.NewProductionConfig()
			config.Level = zap.NewAtomicLevelAt(zap.InfoLevel) // only Info+
			config.EncoderConfig.TimeKey = "ts"
			config.EncoderConfig.LevelKey = "level"
			config.EncoderConfig.CallerKey = "caller"
			config.EncoderConfig.MessageKey = "msg"

			baseLogger, _ := config.Build(
				zap.AddCaller(),
				zap.AddCallerSkip(1),
				zap.AddStacktrace(zap.ErrorLevel), // stacktrace from ErrorLevel+
			)
			logger = baseLogger.Sugar()

		default:
			// fallback to dev if not set
			Init("development")
			return
		}
	})
}

func getLogger() *zap.SugaredLogger {
	if logger == nil {
		// default: read from ENV if Init not explicitly called
		env := os.Getenv("APP_ENV")
		if env == "" {
			env = "development"
		}
		Init(env)
	}
	return logger
}

// SetLevel allows dynamic log level change
func SetLevel(level string) {
	var lvl zapcore.Level
	if err := lvl.UnmarshalText([]byte(level)); err != nil {
		panic("invalid log level: " + level)
	}
	atomicLevel.SetLevel(lvl)
}

// --- Level wrappers ---
func Debug(args ...any)              { getLogger().Debug(args...) }
func Debugf(msg string, args ...any) { getLogger().Debugf(msg, args...) }
func Info(args ...any)               { getLogger().Info(args...) }
func Infof(msg string, args ...any)  { getLogger().Infof(msg, args...) }
func Warn(args ...any)               { getLogger().Warn(args...) }
func Warnf(msg string, args ...any)  { getLogger().Warnf(msg, args...) }
func Error(args ...any)              { getLogger().Error(args...) }
func Errorf(msg string, args ...any) { getLogger().Errorf(msg, args...) }
func Fatal(args ...any)              { getLogger().Fatal(args...) }
func Fatalf(msg string, args ...any) { getLogger().Fatalf(msg, args...) }
