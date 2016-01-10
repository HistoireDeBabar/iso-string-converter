// Package provides a converter for iso time formatted strings
// into a Time object.
package isoConverter

import (
	"strings"
	"time"
)

const form = "2006-Jan-02"

var (
	months = map[string]string{
		"01": "Jan",
		"02": "Feb",
		"03": "Mar",
		"04": "Apr",
		"05": "May",
		"06": "Jun",
		"07": "Jul",
		"08": "Aug",
		"09": "Sep",
		"10": "Oct",
		"11": "Nov",
		"12": "Dec",
	}
)

// Given a ISO Time formatted string e.g. "2016-Jan-04",
// returns a Time instance if string is valid, else returns a Zero Time.
func IsoStringToDate(isoTimeValue string) (setTime time.Time) {
	if isoTimeValue == "" {
		return time.Time{}
	}

	if !strings.Contains(isoTimeValue, "-") {
		return time.Time{}
	}

	timeSplit := strings.Split(isoTimeValue, "-")

	if len(timeSplit) != 3 {
		return time.Time{}
	}

	year := timeSplit[0]
	month := months[timeSplit[1]]
	day := timeSplit[2]

	shortTimeArray := []string{
		year, month, day,
	}

	shortTime := strings.Join(shortTimeArray, "-")
	setTime, err := time.Parse(form, shortTime)

	if err != nil {
		return time.Time{}
	}
	return setTime
}
