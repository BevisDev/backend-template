package datetime

import (
	"github.com/BevisDev/backend-template/src/main/consts"
	"time"
)

func ToString(time time.Time, format string) string {
	return time.Format(format)
}

func ToTime(timeStr string, format string) (time.Time, error) {
	parsedTime, err := time.Parse(format, timeStr)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func StartOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

func EndOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999999000, date.Location())
}

func AddTime(date time.Time, v int, kind string) time.Time {
	switch kind {
	case consts.Second:
		return date.Add(time.Duration(v) * time.Second)
	case consts.Minute:
		return date.Add(time.Duration(v) * time.Minute)
	case consts.Hour:
		return date.Add(time.Duration(v) * time.Hour)
	case consts.Day:
		return date.AddDate(0, 0, v)
	case consts.Month:
		return date.AddDate(0, v, 0)
	case consts.Year:
		return date.AddDate(v, 0, 0)
	default:
		return date
	}
}
