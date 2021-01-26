package cmd

import (
    "go.uber.org/zap"

    "github.com/autobots/touchbase/gcpclients"
    log "github.com/autobots/touchbase/logger"
    "github.com/autobots/touchbase/touchbasemanager"
)

var logger *log.LogInstance

// initLogging initialize logging.
//
// Args:
//   logFormat: the log Format
func initLogging(logFormat string, enableDebugMode bool) {
    if enableDebugMode {
        logger = log.New(logFormat, zap.DebugLevel.String(), true)
    } else {
        logger = log.New(logFormat, zap.InfoLevel.String(), true)
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
