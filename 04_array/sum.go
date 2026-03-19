package array

import "fmt"

// Returning sum of all numbers in a slice
func Sum(numbers []int) int {

	// here numbers is actually a slice not an array
	// array is alawys a fixed sized thing

	var sum int = 0
	// for i := range len(numbers) {
	// 	sum += numbers[i]
	// 	i++
	// }
	// _ -> you can use i too, but if i is not used then it will give error
	// so useing _ for this.
	for _, num := range numbers {
		sum += num
	}

	return sum
}

// taking a slice of numbers -> returning sum of all elements excluding the head (first element)
func SumTails(numbers []int) int {
	sum := 0
	if len(numbers) == 0 {
		return 0
	}
	for i, n := range numbers {
		// ignore the haead (first element)
		if i == 0 {
			continue
		}
		sum += n
	}
	return sum
}

// func SumAll(s1 []int, s2 []int) []int {
// 	output1 := Sum(s1)
// 	output2 := Sum(s2)

// 	return []int{output1, output2}

// }

// Taking multiples slices of numbers and returing a slice of each one's SumAll

// there can be multiple slices in input though
func SumAll(numbersToSum ...[]int) []int {
	var result []int

	for _, n := range numbersToSum {
		result = append(result, Sum(n))
	}

	return result
}

// when you know the slice length you can just make one before and avoid appending
// appending means reallocation but assigning before is better (when you know before hand like this)
func SumAllV2(numbersToSum ...[]int) []int {

	result := make([]int, len(numbersToSum))

	for i, n := range numbersToSum {
		result[i] = Sum(n)
	}

	return result
}

// taking few slices of numbers and returning a slice with each input's tail sum
func SumAllTails(slicesOfNumbersToSum ...[]int) []int {

	result_slice := make([]int, len(slicesOfNumbersToSum))

	for i, n := range slicesOfNumbersToSum {
		result_slice[i] = SumTails(n)
	}

	fmt.Println(result_slice)

	return result_slice
}
