package trace

import (
	"fmt"
	"github.com/soluty/x/logger"
	"time"
)

func noop() {}

func Trace(funcName string, logTime ...time.Duration) func() {
	if !logger.LevelEnable(logger.TraceLevel) {
		return noop
	}
	logger.Trace(fmt.Sprintf("enter %s.", funcName))
	var start time.Time
	if len(logTime) > 0 {
		start = time.Now()
	}
	return func() {
		if len(logTime) > 0 {
			if d := time.Since(start); d > logTime[0] {
				logger.Warn(fmt.Sprintf("exit %s. use large time (%v)", funcName, d))
			} else {
				logger.Trace(fmt.Sprintf("exit %s. use time (%v)", funcName, d))
			}
		} else {
			logger.Trace(fmt.Sprintf("exit %s.", funcName))
		}
	}
}
