package array

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("test sum function - which will take array of numbers and return the sum of them", func(t *testing.T) {
		inputs := []int{1, 2, 3, 4, 5}
		expected_output := 15
		output := Sum(inputs)

		if expected_output != output {
			t.Errorf("Expected %d but got %d while given %v", expected_output, output, inputs)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("test sum all - which take 2 slices and return a slice with sum", func(t *testing.T) {
		output := SumAll([]int{1, 2}, []int{3, 4, 5})
		expected_output := []int{3, 12}

		if !slices.Equal(output, expected_output) {
			t.Errorf("Expected %v but got %v where inputs were %v %v", expected_output, output, []int{1, 2}, []int{3, 4, 5})
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, output, expected_output []int) {
		t.Helper()
		if !reflect.DeepEqual(output, expected_output) {
			t.Errorf("Expected %v but got %v where inputs were %v %v", expected_output, output, []int{1, 2}, []int{3, 4, 5})
		}

	}

	t.Run("test sum all (tails) - which take 2 slices and return a slice with sum of their tails (tails means elements except the head or first element)",
		func(t *testing.T) {
			output := SumAllTails([]int{1, 2}, []int{3, 4, 5})
			expected_output := []int{2, 9}
			checkSums(t, output, expected_output)
		})

	t.Run("if one of the input slices length is 0", func(t *testing.T) {
		output := SumAllTails([]int{}, []int{1, 2, 3})
		expected_output := []int{0, 5}
		checkSums(t, output, expected_output)
	})
}
