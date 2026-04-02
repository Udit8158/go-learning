package main

import (
	"fmt"
	"time"
)

type NumChan struct {
	index int
	value int
}

// non concurrent
func multipliedBy1000(s []int) {
	for i := range s {
		// main work
		time.Sleep(10 * time.Millisecond)
		s[i] = s[i] * 1000
	}
}

// concurrent
func multipliedBy1000Eff(s []int, c chan NumChan) {
	for i := range s {

		go func(i int) {
			// main work
			time.Sleep(10 * time.Millisecond)
			c <- NumChan{index: i, value: s[i] * 1000}
		}(i)
	}
}

func main() {

	startTime := time.Now()
	// create the big array
	numbers := make([]int, 100)
	numberch := make(chan NumChan)
	for i := range numbers {
		numbers[i] = i
	}

	multipliedBy1000Eff(numbers, numberch)
	// multipliedBy1000(numbers)

	for range numbers {
		r := <-numberch
		numbers[r.index] = r.value
	}

	fmt.Println(numbers[0:10])
	fmt.Println(time.Since(startTime))

}
