package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

// constants
const (
	SecondHandLength = 90
	MinuteHandLength = 80
	HourHandLength   = 50
	ClockCenterX     = 150.0
	ClockCenterY     = 150.0
)

// for each sec we are rotating by 6 deg
func SecondHand(tm time.Time) Point {
	// extract the second from the time
	second := tm.Second()

	// angles where the end points of the hand will land
	cosOfAngle := math.Cos(SecondsInRadians(second))
	sinOfAngle := math.Sin(SecondsInRadians(second))
	x2 := ClockCenterX + float64(SecondHandLength)*sinOfAngle
	y2 := ClockCenterY - float64(SecondHandLength)*cosOfAngle
	// base point is in 150,150 (x1,x2)

	return Point{x2, y2}
}

func MinuteHand(tm time.Time) Point {
	minute := tm.Minute()

	cosOfAngle := math.Cos(MinutesInRadians(minute))
	sinOfAngle := math.Sin(MinutesInRadians(minute))
	x2 := ClockCenterX + float64(MinuteHandLength)*sinOfAngle
	y2 := ClockCenterY - float64(MinuteHandLength)*cosOfAngle

	return Point{x2, y2}
}

func HourHand(tm time.Time) Point {
	hr := tm.Hour()

	cosOfAngle := math.Cos(HoursInRadians(hr))
	sinOfAngle := math.Sin(HoursInRadians(hr))
	x2 := ClockCenterX + float64(HourHandLength)*sinOfAngle
	y2 := ClockCenterY - float64(HourHandLength)*cosOfAngle

	return Point{x2, y2}
}

// utility functions
func SecondsInRadians(sec int) float64 {
	return (float64(sec) / 60.0) * (2 * math.Pi)
}

func MinutesInRadians(min int) float64 {
	return (float64(min) / 60.0) * (2 * math.Pi)
}

func HoursInRadians(hr int) float64 {
	return (float64(hr) / 12.0) * (2 * math.Pi)
}
