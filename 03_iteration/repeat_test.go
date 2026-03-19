package iteration

import (
	"fmt"
	"testing"
)

func TestRepeater(t *testing.T) {
	t.Run("repeat a character 5 times", func(t *testing.T) {
		ouput := Repeat("u", 5)
		expected_output := "uuuuu"

		if ouput != expected_output {
			t.Errorf("Expected %q but got %q", expected_output, ouput)
		}
	})
}

func BenchmarkRepeater(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5000)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("a", 5))
	// Output: aaaaa
}
