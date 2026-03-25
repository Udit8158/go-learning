package walk

import (
	"reflect"
	"slices"
	"testing"
)

type Details struct {
	Address Address
	Name    string
}
type Address struct {
	Home string
	City string
}

type Person struct {
	Name    string
	Profile struct {
		Age  int
		City string
	}
}

type ProfileType struct {
	Age  int
	City string
}

type RandomType struct {
	Amount int
	Age    int
}

func TestWalk(t *testing.T) {

	cases := []struct {
		name          string
		expectedCalls []string
		input         any
	}{
		{
			name:          "struct with one string filed",
			input:         struct{ Name string }{Name: "Chris"},
			expectedCalls: []string{"Chris"},
		},
		{
			name:          "struct with two string filed",
			input:         Address{"New Home", "London"},
			expectedCalls: []string{"New Home", "London"},
		},
		{
			name:          "struct with two number filed",
			input:         RandomType{23, 33},
			expectedCalls: []string{},
		},
		{
			name:          "stuct with nested fields",
			input:         Details{Address: Address{Home: "New Home", City: "London"}, Name: "Chris"},
			expectedCalls: []string{"New Home", "London", "Chris"},
		},
		{
			name:          "pointers to things",
			input:         &Person{"Udit", ProfileType{22, "London"}},
			expectedCalls: []string{"Udit", "London"},
		},
		{
			name:          "slices",
			input:         []struct{ Name string }{{"Udit"}, {"Aman"}},
			expectedCalls: []string{"Udit", "Aman"},
		},
		{
			name:          "arrays",
			input:         [2][2]string{{"London", "Delhi"}, {"India", "US"}},
			expectedCalls: []string{"London", "Delhi", "India", "US"},
		},
		// {
		// 	name:          "maps",
		// 	input:         map[string]string{"State": "DK", "Country": "DKsj"},
		// 	expectedCalls: []string{"DK", "DKsj"},
		// },
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := []string{}

			// calling the function
			walk(c.input, func(s string) {
				got = append(got, s)
			})

			// checking
			if !reflect.DeepEqual(got, c.expectedCalls) {
				t.Errorf("Wanted %v, but got %v\n", c.expectedCalls, got)
			}
		})
	}

	t.Run("maps", func(t *testing.T) {
		input := map[string]string{
			"State":   "DK",
			"Country": "DKsj",
		}
		var got []string
		want := []string{"DK", "DKsj"}
		walk(input, func(s string) {
			got = append(got, s)
		})

		// now got should contains DK and DKsj
		if !(slices.Contains(got, "DK") && slices.Contains(got, "DKsj") && len(got) == len(want)) {
			t.Errorf("Got %v wanted %v", got, want)
		}
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan ProfileType)

		go func() {
			aChannel <- ProfileType{33, "Berlin"}
			aChannel <- ProfileType{34, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (ProfileType, ProfileType) {
			return ProfileType{33, "Berlin"}, ProfileType{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}
