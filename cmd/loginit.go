package cmd

import (
    "github.com/autobots/touchbase/gcpclients"
    "github.com/autobots/touchbase/touchbasemanager"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"

    log "github.com/autobots/touchbase/logger"
)

var logger *log.LogInstance

// initLogging initialize logging.
//
// Args:
//   logFormat: the log Format
func initLogging(logFormat string, enableDebugMode bool) {
    if enableDebugMode {
        logger = log.New(logFormat, zapcore.DebugLevel.String(), true)
    } else {
        logger = log.New(logFormat, zapcore.InfoLevel.String(), true)
    }

    // Pass the same logger for components
    touchbasemanager.InitLogger(logger)
    gcpclients.InitLogger(logger)
    // configs.InitLogger(logger)
}

// getLogger retrieves the logging instance.
//
// Return:
//   the zap logging instance
func getLogger() *zap.Logger {
    return logger.GetLogger()
}
