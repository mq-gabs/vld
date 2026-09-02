package utils

import "time"

func IsMidnight(t time.Time) bool {
	return t.Hour() == 0 &&
		t.Minute() == 0 &&
		t.Second() == 0 &&
		t.Nanosecond() == 0
}

func IsSameDate(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day()
}

func IsWeekday(t time.Time) bool {
	return t.Weekday() >= time.Monday && t.Weekday() <= time.Friday
}

func IsWeekend(t time.Time) bool {
	return !IsWeekday(t)
}
