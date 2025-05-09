package main

import "time"

func CurrentTimeDetail() (int32, int32, int32, int32, int32, int32, int32) {
	curTime := time.Now()
	year := curTime.Year()
	month := curTime.Month()
	day := curTime.Day()
	hour := curTime.Hour()
	min := curTime.Minute()
	sec := curTime.Second()
	milli := curTime.Nanosecond() / 1000000
	return int32(year), int32(month), int32(day), int32(hour), int32(min), int32(sec), int32(milli)
}
