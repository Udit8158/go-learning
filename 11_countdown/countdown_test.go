package countdown

import (
	"bytes"
	"testing"
)

type MockSleeper struct {
	RunCount int32
}

func (ms *MockSleeper) Sleep() {
	ms.RunCount++
}

func TestCountdown(t *testing.T) {
	var buffer bytes.Buffer
	var testMockSleeper MockSleeper

	Countdown(&buffer, &testMockSleeper)

	got := buffer.String()
	want := "3\n2\n1\nGO\n"

	expectedCountDownSleepertToRun := 3

	if got != want {
		t.Errorf("Expected %q but got %q\n", want, got)
	}

	if testMockSleeper.RunCount != int32(expectedCountDownSleepertToRun) {
		t.Errorf("Expected sleeper to run %d times but ran for %d times", expectedCountDownSleepertToRun, testMockSleeper.RunCount)
	}
}
