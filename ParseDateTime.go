package main

import (
	"fmt"
	"time"
)

const (
	layout = "2006-01-02 15:04:05"
)

//ref: https://developpaper.com/the-use-of-goang-time-time-zone-and-format/
func main() {
	str := "2021-10-16 20:45:55"

	timestamp, err := StringToTimestamp(str, layout, "Local")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(timestamp)

	utcTimestamp, err := ConvertToUTC(timestamp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(utcTimestamp)
}

func StringToTimestamp(dateTimeStr, layout, locationStr string) (time.Time, error) {
	locale, err := time.LoadLocation(locationStr) // Time zone set by the server
	if err != nil {
		return time.Time{}, err
	}
	timestamp, err := time.ParseInLocation(layout, dateTimeStr, locale)
	if err != nil {
		return timestamp, err
	}

	return timestamp, nil
}

func ConvertToUTC(timestamp time.Time) (time.Time, error) {
	utc, err := time.LoadLocation("UTC")
	if err != nil {
		return timestamp, err
	}
	utcTimestamp := timestamp.In(utc)
	return utcTimestamp, err
}

