package clockface_test

import (
	"bytes"
	"encoding/xml"
	"math"
	"testing"
	"time"

	clockface "github.com/Udit8158/go-learning/18_using_maths"
)

// SVG structure - for parsing the svg into this struct
// https://xml-to-go.github.io/ used this
type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  struct {
		Text  string `xml:",chardata"`
		Cx    string `xml:"cx,attr"`
		Cy    string `xml:"cy,attr"`
		R     string `xml:"r,attr"`
		Style string `xml:"style,attr"`
	} `xml:"circle"`
	Lines []Line `xml:"line"`
}

type Line struct {
	Text  string  `xml:",chardata"`
	X1    float64 `xml:"x1,attr"`
	Y1    float64 `xml:"y1,attr"`
	X2    float64 `xml:"x2,attr"`
	Y2    float64 `xml:"y2,attr"`
	Style string  `xml:"style,attr"`
}

func TestSvgAtMidnight(t *testing.T) {
	buf := bytes.Buffer{}
	tm := time.Date(1337, time.January, 1, 0, 0, 15, 0, time.UTC)
	clockface.DrawClockFace(tm, &buf)

	var output SVG
	xml.Unmarshal(buf.Bytes(), &output)

	wantedLines := []Line{
		{X1: 150, Y1: 150, X2: 150, Y2: 100}, // hour
		{X1: 150, Y1: 150, X2: 150, Y2: 70},  // min
		{X1: 150, Y1: 150, X2: 240, Y2: 150}, // sec
	}

	// fmt.Println(output.Lines)

	if len(output.Lines) != 3 {
		t.Fatal("All lines of has not been drawn!")
	}

	for i := range output.Lines {
		assertLine(t, output.Lines[i], wantedLines[i])

	}

}

func TestSvgAt114834(t *testing.T) {
	buf := bytes.Buffer{}
	tm := time.Date(1337, time.January, 1, 11, 48, 34, 0, time.UTC)
	clockface.DrawClockFace(tm, &buf)

	var output SVG
	xml.Unmarshal(buf.Bytes(), &output)

	wantedLines := []Line{
		{X1: 150, Y1: 150, X2: 125.000000, Y2: 106.698730}, // hour
		{X1: 150, Y1: 150, X2: 73.915478, Y2: 125.278640},  // minute
		{X1: 150, Y1: 150, X2: 113.393702, Y2: 232.219091}, // second
	}

	if len(output.Lines) != 3 {
		t.Fatal("All lines of has not been drawn!")
	}

	for i := range output.Lines {
		assertLine(t, output.Lines[i], wantedLines[i])
	}
}

func assertLine(t *testing.T, outputLine Line, exptectedLine Line) {
	// if outputLine.X1 != exptectedLine.X1 || outputLine.X2 != exptectedLine.X2 || outputLine.Y1 != exptectedLine.Y1 || outputLine.Y2 != exptectedLine.Y2 {
	// 	t.Helper()
	// 	t.Errorf("not matched expected %v   got %v", exptectedLine, outputLine)
	// }
	tollerance := 1e-4

	if math.Abs(outputLine.X1-exptectedLine.X1) > tollerance || math.Abs(outputLine.X2-exptectedLine.X2) > tollerance || math.Abs(outputLine.Y1-exptectedLine.Y1) > tollerance || math.Abs(outputLine.Y2-exptectedLine.Y2) > tollerance {
		t.Helper()
		t.Errorf("not matched expected %v   got %v", exptectedLine, outputLine)
	}
}
