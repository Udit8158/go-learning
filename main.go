package main

import "fmt"

func Change(arr []int) {
	arr[0] = 0
	arr = append(arr, 6)
	fmt.Println("Array changed, ", arr)
}

func main() {

	s := []int{1, 2, 3, 4, 5}
	// Change(s)

	s = append(s, 6)
	fmt.Println(s)
}
