package log

import (
	"os"
	"strconv"
)

type logModeType string

var (
	developmentMode logModeType = "DevelopmentMode"
	productionMode  logModeType = "ProductionMode"
	testMode        logModeType = "testMode"
)

var logMode = productionMode

func initConfigMode() {
	devStr := os.Getenv("LOG_USE_DEVELOPMENT_CONFIG")
	if isDev, err := strconv.ParseBool(devStr); err == nil {
		if isDev {
			logMode = developmentMode
		}
	}

	testStr := os.Getenv("TEST")
	if isTest, err := strconv.ParseBool(testStr); err == nil {
		if isTest {
			logMode = testMode
		}
	}
}
