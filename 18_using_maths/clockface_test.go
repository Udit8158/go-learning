package clockface_test

// every clock has a centre of (150, 150)
// the hour hand is 50 long
// the minute hand is 80 long
// the second hand is 90 long.

import (
	"math"
	"testing"
	"time"

	clockface "github.com/Udit8158/go-learning/18_using_maths"
)

// not getting total accuracy while dealing with floats
// so accepting this minor fault tollerance
const tollerance = 1e-7

func TestSecondHand(t *testing.T) {
	tests := []struct {
		name string
		time time.Time
		want clockface.Point
	}{
		{
			name: "at midnight points straight up",
			time: time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC),
			want: clockface.Point{X: 150, Y: 60},
		},
		{
			name: "at 30 seconds points straight down",
			time: time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC),
			want: clockface.Point{X: 150, Y: 240},
		},
		{
			name: "at 5 seconds points near one o'clock",
			time: time.Date(1337, time.January, 1, 0, 0, 5, 0, time.UTC),
			want: clockface.Point{X: 195, Y: 72.05771365940052},
		},
		{
			name: "at 15 seconds points right",
			time: time.Date(1337, time.January, 1, 0, 0, 15, 0, time.UTC),
			want: clockface.Point{X: 240, Y: 150},
		},
		{
			name: "at 50 seconds points near ten o'clock",
			time: time.Date(1337, time.January, 1, 0, 0, 50, 0, time.UTC),
			want: clockface.Point{X: 72.0577136594005, Y: 105},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := clockface.SecondHand(test.time)

			assertPoint(t, got, test.want)
		})
	}
}

func TestMinuteHand(t *testing.T) {
	tests := []struct {
		name string
		time time.Time
		want clockface.Point
	}{
		{
			name: "at midnight points straight up",
			time: time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC),
			want: clockface.Point{X: 150, Y: 70},
		},
		{
			name: "at 15 minutes points right",
			time: time.Date(1337, time.January, 1, 0, 15, 0, 0, time.UTC),
			want: clockface.Point{X: 230, Y: 150},
		},
		{
			name: "at 30 minutes points straight down",
			time: time.Date(1337, time.January, 1, 0, 30, 0, 0, time.UTC),
			want: clockface.Point{X: 150, Y: 230},
		},
		{
			name: "at 50 minutes points near ten o'clock",
			time: time.Date(1337, time.January, 1, 0, 50, 0, 0, time.UTC),
			want: clockface.Point{X: 80.7179676972449, Y: 110},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := clockface.MinuteHand(test.time)

			assertPoint(t, got, test.want)
		})
	}
}

func TestHourHand(t *testing.T) {
	tests := []struct {
		name string
		time time.Time
		want clockface.Point
	}{
		{
			name: "at midnight points straight up",
			time: time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC),
			want: clockface.Point{X: 150, Y: 100},
		},
		{
			name: "at three o'clock points right",
			time: time.Date(1337, time.January, 1, 3, 0, 0, 0, time.UTC),
			want: clockface.Point{X: 200, Y: 150},
		},
		{
			name: "at six o'clock points straight down",
			time: time.Date(1337, time.January, 1, 6, 0, 0, 0, time.UTC),
			want: clockface.Point{X: 150, Y: 200},
		},
		{
			name: "at nine o'clock points left",
			time: time.Date(1337, time.January, 1, 9, 0, 0, 0, time.UTC),
			want: clockface.Point{X: 100, Y: 150},
		},
		// not implemented this behaviour for min and hour hand yet
		// {
		// 	name: "at half past three sits halfway between three and four",
		// 	time: time.Date(1337, time.January, 1, 3, 30, 0, 0, time.UTC),
		// 	want: clockface.Point{X: 196.19397662556435, Y: 169.1341716182545},
		// },
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := clockface.HourHand(test.time)

			assertPoint(t, got, test.want)
		})
	}
}

func TestSecondInRadians(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 5, 0, time.UTC)

	want := (5.0 / 60.0) * math.Pi * 2
	got := clockface.SecondsInRadians(tm.Second())

	if math.Abs(got-want) > tollerance {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

func TestMinuteInRadians(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 5, 0, 0, time.UTC)

	want := (5.0 / 60.0) * math.Pi * 2
	got := clockface.MinutesInRadians(tm.Minute())

	if math.Abs(got-want) > tollerance {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

func assertPoint(t testing.TB, got, want clockface.Point) {
	t.Helper()

	if math.Abs(got.X-want.X) > tollerance || math.Abs(got.Y-want.Y) > tollerance {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}
