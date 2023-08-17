package utils

import (
	"log"
	"time"
)

func TimeParse(ti string) time.Time {
	format := "2006-01-02"
	t, _ := time.Parse(format, ti)
	return t
}

func TimeParseBool(ti string) bool {
	format := "2006-01-02"
	if _, err := time.Parse(format, ti); err != nil {
		log.Println(ti, err)
		return false
	}
	return true
}
func ComPareTimeHour(t1 time.Time, t2 time.Time) bool {
	truncatedTime1 := t1.Truncate(time.Hour)
	truncatedTime2 := t2.Truncate(time.Hour)
	return truncatedTime1.Equal(truncatedTime2)
}
