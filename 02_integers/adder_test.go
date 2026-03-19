package integers

import (
	"fmt"
	"testing"
)

// Test function
func TestAdder(t *testing.T) {
	t.Run("testing adder function", func(t *testing.T) {
		sum := Add(4, 5)
		expected := 9

		if sum != expected {
			t.Errorf("Expected %d but got %d", sum, expected)
		}
	})
}

// Example function
func ExampleAdd() {
	sum := Add(4, 5)
	fmt.Println(sum)
	// Output: 9
}
