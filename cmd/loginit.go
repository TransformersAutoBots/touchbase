package cmd

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"

    log "github.com/autobots/touchbase/logger"
)

var logger *log.LogInstance

// initLogging initialize logging.
//
// Args:
//   logFormat: the log Format
func initLogging(logFormat string) {
    logger = log.New(logFormat, true, zapcore.InfoLevel.String(), &log.LogInstance{})

}

// getLogger retrieves the logging instance.
//
// Return:
//   the zap logging instance
func getLogger() *zap.Logger {
    return logger.GetLogger()
}
