package countdown

import (
	"fmt"
	"io"
)

type Sleeper interface {
	Sleep()
}

func Countdown(w io.Writer, s Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(w, i)
		// time.Sleep(time.Duration(0) * time.Second)
		s.Sleep()
	}
	fmt.Fprintln(w, "GO")
}
