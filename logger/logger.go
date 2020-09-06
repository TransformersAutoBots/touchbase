package logger

import (
    "github.com/autobots/touchbase/constants"
    "go.uber.org/zap/zapcore"
    "log"
    "sync"

    "go.uber.org/zap"
)

// LogInstance the log instance for the application.
type LogInstance struct {
    logInitialized sync.Once
    log            *zap.Logger
}

// logConfig config to customize logging.
type logConfig struct {
    level      string
    format     string
    enableTime bool
}

// getLogLevel gets the log level.
//
// Args:
//   level: the logging priority in string format
// Return:
//   the logging priority in zap format
func getLogLevel(level string) zapcore.Level {
    switch level {
    case zap.DebugLevel.String():
        return zap.DebugLevel
    case zap.InfoLevel.String():
        return zap.InfoLevel
    case zap.WarnLevel.String():
        return zap.WarnLevel
    case zap.ErrorLevel.String():
        return zap.ErrorLevel
    case zap.DPanicLevel.String():
        return zap.DPanicLevel
    case zap.PanicLevel.String():
        return zap.PanicLevel
    case zap.FatalLevel.String():
        return zap.FatalLevel
    default:
        return zap.InfoLevel
    }
}

// buildLoggerInstance builds a new zap logger instance.
//
// Args:
//   config: the config to customize logging
// Return:
//   New zap logger instance
func buildLoggerInstance(config logConfig) *zap.Logger {
    if config.format == "" {
        log.Fatalf("Logging format cannot be empty")
    }
    if config.format != constants.ConsoleFormat && config.format != constants.JsonFormat {
        log.Fatalf("Only %v and %v format are supported for logging", constants.ConsoleFormat, constants.JsonFormat)
    }

    zapConfig := zap.Config{
        Level:             zap.NewAtomicLevelAt(getLogLevel(config.level)),
        Development:       false,
        DisableCaller:     false,
        DisableStacktrace: false,
        Encoding:          config.format,
        EncoderConfig: zapcore.EncoderConfig{
            MessageKey:   "logMessage",
            LevelKey:     constants.Level,
            EncodeLevel:  zapcore.CapitalLevelEncoder,
            CallerKey:    constants.Caller,
            EncodeCaller: zapcore.ShortCallerEncoder,
        },
        OutputPaths:      []string{"stderr"},
        ErrorOutputPaths: []string{"stderr"},
        InitialFields:    nil,
    }

    if config.enableTime {
        zapConfig.EncoderConfig.TimeKey = constants.Time
        zapConfig.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
    }

    logger, err := zapConfig.Build()
    if err != nil {
        log.Fatalf("Failed to create new zap logger instance. Reason: %s", err)
    }
    return logger
}

// New creates a new zap logger instance if not initialized.
//
// Args:
//   logFormat: the log format
//   enableTime: enable time in logging
//   logLevel: the log level
//   logInstance: the log instance
// Return:
//   the log instance with new zap logger if not already initialized
func New(logFormat string, enableTime bool, logLevel string, logInstance *LogInstance) *LogInstance {
    logInstance.logInitialized.Do(func() {
        logInstance.log = buildLoggerInstance(logConfig{
            level:      logLevel,
            format:     logFormat,
            enableTime: enableTime,
        })
    })
    return logInstance
}

// GetLogger gets the zap logger instance.
//
// Return:
//   the zap logger instance
func (logInstance *LogInstance) GetLogger() *zap.Logger {
    logInstance.logInitialized.Do(func() {
        log.Fatal("Get logger called before logger initialization!")
    })
    return logInstance.log
}
