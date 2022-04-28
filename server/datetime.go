package server

import (
	"fmt"
	"strconv"
	"time"
)

func NowLoc(loc *time.Location) time.Time {
	return time.Now().In(loc)
}

func Date() string {
	year, month, day := time.Now().Date()
	return fmt.Sprintf("%02d.%02d.%d", day, month, year)
}

func DateLoc(loc *time.Location) string {
	year, month, day := NowLoc(loc).Date()
	return fmt.Sprintf("%02d.%02d.%d", day, month, year)
}

func Time() string {
	hour, min, sec := time.Now().Clock()
	return fmt.Sprintf("%02d:%02d:%02d", hour, min, sec)
}

func TimeLoc(loc *time.Location) string {
	hour, min, sec := NowLoc(loc).Clock()
	return fmt.Sprintf("%02d:%02d:%02d", hour, min, sec)
}

func Weekday() string {
	return time.Now().Weekday().String()
}

func WeekdayLoc(loc *time.Location) string {
	return NowLoc(loc).Weekday().String()
}

func Day() string {
	return strconv.Itoa(time.Now().Day())
}

func DayLoc(loc *time.Location) string {
	return strconv.Itoa(NowLoc(loc).Day())
}

func Month() string {
	return strconv.Itoa(int(time.Now().Month()))
}

func MonthLoc(loc *time.Location) string {
	return strconv.Itoa(int(NowLoc(loc).Month()))
}

func MonthText() string {
	return time.Now().Month().String()
}

func Year() string {
	return strconv.Itoa(time.Now().Year())
}

func YearLoc(loc *time.Location) string {
	return strconv.Itoa(NowLoc(loc).Year())
}
