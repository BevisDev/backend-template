package utils

import (
	"github.com/BevisDev/backend-template/src/main/logger"
	"time"
)

func ToString(time time.Time, format string) string {
	return time.Format(format)
}

func ToTime(timeStr string, format string) time.Time {
	parsedTime, err := time.Parse(format, timeStr)
	if err != nil {
		logger.Error("Parse time error: " + err.Error())
	}
	return parsedTime
}

func StartOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

func EndOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999999000, date.Location())
}
