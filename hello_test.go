package main

import "testing"

func AssertErrorMessage(got string, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("We got %q but wanted %q", got, want)
	}
}
func TestHello(t *testing.T) {

	// one scenario
	t.Run("saying Hello to people", func(t *testing.T) {
		got := Hello("Udit")
		want := "Hello, Udits"

		AssertErrorMessage(got, want, t)
	})

	// 2nd scenario
	t.Run("saying Hello, world if empty string passed", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		AssertErrorMessage(got, want, t)
	})
}

func TestGreet(t *testing.T) {

	got := Greet("udit")
	want := "Hello, udit"

	AssertErrorMessage(got, want, t)

}