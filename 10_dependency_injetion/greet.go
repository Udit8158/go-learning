package dependencyinjetion

import (
	"fmt"
	"io"
)

func Greet(w io.Writer, name string) {
	output := "Hello, " + name
	fmt.Fprintln(w, output)
}
