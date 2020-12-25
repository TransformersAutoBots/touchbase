package logger

import (
    "log"
    "sync"

    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"

    "github.com/autobots/touchbase/constants"
)

// LogInstance the log instance for the application.
type LogInstance struct {
    logInitialized sync.Once
    log            *zap.Logger

    configUpdateDetails sync.Once
    // senderDetailsInitialized sync.Once
}

// logConfig holds information necessary for customizing the logger.
type logConfig struct {
    level      string
    format     string
    enableTime bool
}

// AddConfigUpdateDetails adds the updated config details to root logging.
//
// Args:
//   v: the sender details
//   logInstance: the log instance for the current event
// Return:
//   the log instance with the updated config details
func AddConfigUpdateDetails(v interface{}, logInstance *LogInstance) *LogInstance {
    logInstance.configUpdateDetails.Do(func() {
        logInstance.log = logInstance.log.With(
            convertToField("configUpdateDetails", v),
        )
    })
    return logInstance
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
        log.Fatal("Logging format cannot be empty")
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
            LineEnding:   zapcore.DefaultLineEnding,
            MessageKey:   constants.MessageKey,
            LevelKey:     constants.LevelKey,
            EncodeLevel:  zapcore.CapitalLevelEncoder,
            CallerKey:    constants.CallerKey,
            EncodeCaller: zapcore.ShortCallerEncoder,
        },
        OutputPaths:      []string{"stderr"},
        ErrorOutputPaths: []string{"stderr"},
        InitialFields:    nil,
    }

    if config.enableTime {
        zapConfig.EncoderConfig.TimeKey = constants.TimeKey
        zapConfig.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
    }

    logger, err := zapConfig.Build()
    if err != nil {
        log.Fatalf("Failed to create new zap logger instance. Reason: %s", err)
    }
    return logger
}

// newLogger creates a new configured Zap logger.
//
// Args:
//   logFormat: the log format
//   logLevel: the log level
//   enableTime: enable time in logging
// Return:
//   New Zap logger instance
func newLogger(logFormat, logLevel string, enableTime bool) *zap.Logger {
    return buildLoggerInstance(
        logConfig{
            level:      logLevel,
            format:     logFormat,
            enableTime: enableTime,
        },
    )
}

// New initiates a new Zap logger instance if not initialized.
//
// Args:
//   logFormat: the log format
//   logLevel: the log level
//   enableTime: flag to enable time in logging
// Return:
//   Log Instance with new zap logger instance if not already initialized
func New(logFormat, logLevel string, enableTime bool) *LogInstance {
    logInstance := &LogInstance{}
    logInstance.logInitialized.Do(func() {
        logInstance.log = newLogger(logFormat, logLevel, enableTime)
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
