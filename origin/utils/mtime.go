package utils

import "time"

func Format(time time.Time) string  {
	return time.Format("2006-01-02 15:04:05")
}

func FormatMonth(time time.Time) string  {
	return time.Format("2006-01")
}