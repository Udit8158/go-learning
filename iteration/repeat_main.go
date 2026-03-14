package iteration

import (
	"strings"
)

func RepeatOld(char string, repeat int) string {
	var output string

	for i := range repeat {
		output += char // warning - not efficient
		i++
	}

	return output
}

// Efficient way using string builder
func Repeat(char string, repeat int) string {
	var output strings.Builder

	for i := range repeat {
		output.WriteString(char)
		i++
	}

	// fmt.Println(&output)

	str := output.String()
	return str
}
