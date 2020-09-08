package cmd

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"

    "github.com/autobots/touchbase/configs"
    log "github.com/autobots/touchbase/logger"
)

var logger *log.LogInstance

// initLogging initialize logging.
//
// Args:
//   logFormat: the log Format
func initLogging(logFormat string, enableDebugMode bool) {
    if enableDebugMode {
        logger = log.New(logFormat, true, zapcore.DebugLevel.String(), &log.LogInstance{})
    } else {
        logger = log.New(logFormat, true, zapcore.InfoLevel.String(), &log.LogInstance{})
    }

    // Pass the same logger for components
    configs.InitLogger(logger)
}

// getLogger retrieves the logging instance.
//
// Return:
//   the zap logging instance
func getLogger() *zap.Logger {
    return logger.GetLogger()
}
