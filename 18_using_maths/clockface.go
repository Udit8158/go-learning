package clockface

import (
	"fmt"
	"io"
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

// function to draw the clockface in a buffer/output
func DrawClockFace(tm time.Time, w io.Writer) {
	shPoint := SecondHand(tm)
	mhPoint := MinuteHand(tm)
	hhPoint := HourHand(tm)

	// writting the svg in the io writer
	// can be used by buffer (in test) and stdout (in main)
	fmt.Fprint(w, svgStart)
	fmt.Fprint(w, bezel)

	fmt.Fprint(w, hourHandTag(hhPoint))
	fmt.Fprint(w, minuteHandTag(mhPoint))
	fmt.Fprint(w, secondHandTag(shPoint))

	fmt.Fprint(w, svgEnd)
}

// end point for second hand (x2,y2)
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

// end point for min hand (x2,y2)
func MinuteHand(tm time.Time) Point {
	minute := tm.Minute()

	cosOfAngle := math.Cos(MinutesInRadians(minute))
	sinOfAngle := math.Sin(MinutesInRadians(minute))
	x2 := ClockCenterX + float64(MinuteHandLength)*sinOfAngle
	y2 := ClockCenterY - float64(MinuteHandLength)*cosOfAngle

	return Point{x2, y2}
}

// end point for hour hand (x2,y2)
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

// utility functions for drawing hands into the svg
func secondHandTag(p Point) string {
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func minuteHandTag(p Point) string {
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#3443eb;stroke-width:3px;"/>`, p.X, p.Y)
}

func hourHandTag(p Point) string {
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#141414;stroke-width:3px;"/>`, p.X, p.Y)
}

// svg specific constants
const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
