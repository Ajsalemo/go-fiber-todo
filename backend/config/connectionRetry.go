package config

import (
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
)

func ConnectionRetry(err error) {
	retryCount := 0
	var backoffSchedule = []time.Duration{
		1 * time.Second,
		3 * time.Second,
		10 * time.Second,
	}

	for _, backoff := range backoffSchedule {
		if err == nil {
			// Reset the retry count
			retryCount = 0
			zap.L().Info("Successfully connected to PostgreSQL..")
			break
			// If the error cants a connection failure, retry the connection
		} else if err.Error() != "" && strings.Contains(err.Error(), "failed to connect") {
			// Increment the retry count
			retryCount++
			zap.L().Error(err.Error())
			zap.L().Warn("Retry count " + strconv.Itoa(retryCount))
			zap.L().Warn("Retrying in: " + backoff.String())
			time.Sleep(backoff)
			// If retry count is greater than 2 and err is not nil, return the error - Fatal() implicitly calls os.Exit(1)
			if retryCount > 2 && err.Error() != "" {
				zap.L().Fatal(err.Error())
			}
		}
	}
}
