package time2

import (
	"time"
)

const (
	TimeFormatDate     = "2006-01-02"
	TimeFormatTime     = "15:04:05"
	TimeFormatDateTime = "2006-01-02 15:04:05"
	TimeFormatDefault  = TimeFormatDateTime
)

// AddDay add day
func AddDay(t time.Time, day int) time.Time {
	return t.AddDate(0, 0, day)
}

// AddHour add hour
func AddHour(t time.Time, hour int) time.Time {
	return t.Add(time.Duration(hour) * time.Hour)
}

// AddMinute add minute
func AddMinute(t time.Time, minutes int) time.Time {
	return t.Add(time.Duration(minutes) * time.Minute)
}

// AddSecond add second
func AddSecond(t time.Time, seconds int) time.Time {
	return t.Add(time.Duration(seconds) * time.Second)
}

// HourStart start of hour
func HourStart(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 0, 0, 0, t.Location())
}

// HourEnd end of hour
func HourEnd(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// DayStart start of day
func DayStart(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// DayEnd end of day
func DayEnd(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// MonthStart start of month
func MonthStart(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

// MonthEnd end of month
func MonthEnd(t time.Time) time.Time {
	return MonthStart(t).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// YearStart start of year
func YearStart(t time.Time) time.Time {
	y, _, _ := t.Date()
	return time.Date(y, time.January, 1, 0, 0, 0, 0, t.Location())
}

// YearEnd end of year
func YearEnd(t time.Time) time.Time {
	return YearStart(t).AddDate(1, 0, 0).Add(-time.Nanosecond)
}
