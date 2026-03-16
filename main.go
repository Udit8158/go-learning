package main

import (
	"fmt"

	custom_append "github.com/Udit8158/go-learning/custom_append"
)

func Change(arr []int) {
	arr[0] = 0
	arr = append(arr, 6)
	fmt.Println("Array changed, ", arr)
}

func main() {

	s := []int{1, 2, 3, 4, 5} // this a way to define a slice
	s2 := make([]int, 5, 6)   // len -> 5, cap -> 8 slice (but this is better ig, more control)

	// interesting
	// s2 = s2[:6]
	// s2[5] = 12

	fmt.Println(custom_append.CustomAppend(s2, 12))
	fmt.Printf("After custom append %p\n", &s2[0])

	fmt.Println("Current cap", cap(s2))
	s2 = append(s2, 100)
	s2 = append(s2, 100)
	s2 = append(s2, 100)
	s2 = append(s2, 100)
	s2 = append(s2, 100)
	fmt.Println(s2, len(s2), cap(s2))
	// Change(s)

	fmt.Println(cap(s)) // capacity -> 5
	s = append(s, 6)
	fmt.Println(cap(s)) // capacity -> 10
	s = append(s, 7)
	fmt.Println(cap(s)) // capacity -> 10
	fmt.Println(s)
}
