package main

import (
	"log"
	"os"
	"time"

	clockface "github.com/Udit8158/go-learning/18_using_maths"
)

func main() {
	tm := time.Now()
	log.Print("Current Time: ", tm)

	clockface.DrawClockFace(tm, os.Stdout)
}
