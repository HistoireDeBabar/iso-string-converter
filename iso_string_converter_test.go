package isoConverter

import (
	"testing"
	"time"
)

func TestISOStringConverterEmptyStringShouldEqualZeroTime(t *testing.T) {
	timeString := ""
	result := IsoStringToDate(timeString)
	if !(result.IsZero()) {
		t.Error("Result should be default time")
	}
}

func TestISOStringConverterInvalidStringShouldEqualZeroTime(t *testing.T) {
	timeString := "fhsoudfhsf"
	result := IsoStringToDate(timeString)
	if !(result.IsZero()) {
		t.Error("Result should be default time")
	}
}

func TestISOStringConverter(t *testing.T) {
	timeString := "2016-02-21"
	result := IsoStringToDate(timeString)
	if result.IsZero() {
		t.Error("Result should be not be Zero time")
	}
	if result.Month() != time.February {
		t.Errorf("result Month should be Febuary not %v", result.Month())
	}

	if result.Day() != 21 {
		t.Errorf("result Day should be 21 not %v", result.Day())
	}

	if result.Year() != 2016 {
		t.Errorf("result Day should be 2016 not %v", result.Year())
	}
}
