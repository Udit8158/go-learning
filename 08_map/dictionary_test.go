package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dict := Dictionary{"test": "this is the meaning"}

		got, err := dict.Search("test")
		want := "this is the meaning"

		if err != nil {
			t.Fatal("should find the word, but got error", err)
		}

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dict := Dictionary{"test": "this is the meaning"}

		_, err := dict.Search("btc")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{"test": "this is the meaning"}

		err := dict.Add("mac", "macbook is the best")
		got, searchErr := dict.Search("mac")
		want := "macbook is the best"

		if err != nil {
			t.Fatal("did not expect an error while adding a new word, but got", err)
		}
		if searchErr != nil {
			t.Fatal("should find the added word, but got error", searchErr)
		}

		assertStrings(t, got, want)
	})

	t.Run("existing word", func(t *testing.T) {
		dict := Dictionary{"mac": "macbook is the best"}

		addErr := dict.Add("mac", "hi there")
		got, searchErr := dict.Search("mac")
		want := "macbook is the best"

		if searchErr != nil {
			t.Fatal("there shouldn't be an error while searching, but got", searchErr)
		}
		if addErr == nil {
			t.Fatal("expected an error while adding duplicate key, but didn't get that")
		}

		assertStrings(t, got, want)
		assertError(t, addErr, ErrWordExist)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update an existing word", func(t *testing.T) {
		dict := Dictionary{"test": "this is test"}
		err := dict.Update("test", "this is updated")
		got := dict["test"]
		want := "this is updated"

		if err != nil {
			t.Fatalf("should not get an error while update, but go errr -> %v", err)
		}

		assertStrings(t, got, want)
	})
	t.Run("update an not existing word", func(t *testing.T) {
		dict := Dictionary{"test": "this is test"}
		err := dict.Update("btc", "this is updated")

		got := dict["btc"]
		want := ""
		if err == nil {
			t.Fatalf("should get an error while update, but didn't get any")
		}

		assertError(t, err, ErrNoWordExist)
		assertStrings(t, got, want) // also checking that no change happened while failed update
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing word", func(t *testing.T) {
		dict := Dictionary{"btc": "bitcoin"}
		err := dict.Delete("btc")

		assertError(t, err, nil)

		// search for the word and expecting error to be matched
		_, searchErr := dict.Search("btc")
		assertError(t, searchErr, ErrNotFound)

	})

	t.Run("delete non existing word", func(t *testing.T) {
		dict := Dictionary{"btc": "bitcoin"}
		err := dict.Delete("test") // no existing
		assertError(t, err, ErrNoWordExist)
	})
}
func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %s but got %s", want, got)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q but wanted %q", got, want)
	}
}
