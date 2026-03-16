package custom_append

import (
	"slices"
	"testing"
)

func TestCustomAppend(t *testing.T) {

	checkResult := func(t *testing.T, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	test_cases := []struct {
		name           string
		inputSlice     []int
		numberToAppend int
		want           []int
	}{
		{
			name:           "Noramal scenario",
			inputSlice:     []int{34, 32},
			numberToAppend: 12,
			want:           []int{34, 32, 12},
		},
		{
			name:           "empty slice (len=0, cap=0)",
			inputSlice:     []int{},
			numberToAppend: 45,
			want:           []int{45},
		},
		{
			name:           "nil slice",
			inputSlice:     nil,
			numberToAppend: 10,
			want:           []int{10},
		},
		{
			name:           "slice is full (len == cap), triggers reallocation",
			inputSlice:     []int{1, 2, 3},
			numberToAppend: 4,
			want:           []int{1, 2, 3, 4},
		},
		{
			name:           "slice has spare capacity (len < cap), no reallocation",
			inputSlice:     append(make([]int, 0, 5), 1, 2, 3),
			numberToAppend: 4,
			want:           []int{1, 2, 3, 4},
		},
		{
			name:           "single element slice, full",
			inputSlice:     []int{99},
			numberToAppend: 7,
			want:           []int{99, 7},
		},
		{
			name:           "appending zero value",
			inputSlice:     []int{1, 2},
			numberToAppend: 0,
			want:           []int{1, 2, 0},
		},
		{
			name:           "appending negative number",
			inputSlice:     []int{1, 2},
			numberToAppend: -5,
			want:           []int{1, 2, -5},
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.name, func(t *testing.T) {
			checkResult(t, CustomAppend(test_case.inputSlice, test_case.numberToAppend), test_case.want)
		})
	}
}
