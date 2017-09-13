package main

import (
	"fmt"
	"time"
)

// Simple Date format
type Date struct {
	Year  int
	Month int
	Day   int
}

// DateFormat Layout
const (
	F1 = "2017-09-21"
	F2 = "21-Sep-17"
	F3 = "2017년 09월 21일"
)

// Today
// return Today's date
func Toady() Date {

	year, month, day := time.Now().Date()
	return Date{Year: year, Month: int(month), Day: day}
}

// Tomorrow
// return Tomorrow's date
func Tomorrow() Date {
	year, month, day := time.Now().Date()
	day = day + 1
	return Date{Year: year, Month: int(month), Day: day}
}

// Yesterday
// return Yesterday's date
func Yesterday() Date {
	year, month, day := time.Now().Date()
	day = day - 1
	return Date{Year: year, Month: int(month), Day: day}
}

// IsAWeek
// return it's weekend or not
func IsAWeek(t time.Time) bool {
	day := t.Weekday().String()

	switch day {

	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		return false

	case "Saturday", "Sunday":
		return true

	default:
		return false

	}

}

// CalcDate
// add or minus date.
func (d Date) CalcDate(years, months, days int) Date {

	return Date{Year: d.Year + years, Month: d.Month + months, Day: d.Day + days}

}

func main() {
	a := Toady()
	fmt.Println(a)
}
