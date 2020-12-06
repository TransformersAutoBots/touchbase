package touchbasemanager

import (
    "go.uber.org/zap"

    log "github.com/autobots/touchbase/logger"
)

var logger *log.LogInstance

// InitLogger initialize logging for component.
//
// Args:
//   logInstance: the log instance
func InitLogger(logInstance *log.LogInstance) {
    logger = logInstance
}

// getLogger retrieves the logging instance.
//
// Return:
//   the zap logging instance
func getLogger() *zap.Logger {
    return logger.GetLogger()
}
