package dependencyinjetion

import (
	"bytes"
	"testing"
)

type Buf []byte

func (b *Buf) Write(p []byte) (n int, err error) {
	*b = append(*b, p...)
	return len(*b), nil
}

func TestGreet(t *testing.T) {
	// var buffer Buf
	var buffer bytes.Buffer
	Greet(&buffer, "Udit")
	// got := string(buffer)
	got := buffer.String()
	want := "Hello, Udit\n"

	if got != want {
		t.Errorf("Expected %q but got %q", want, got)
	}
}
