package db

import (
	"errors"
	"reflect"
	"testing"
)

func TestAddUser(t *testing.T) {
	details := UserDetails{
		Name:  "Udit",
		Email: "udit@email.com",
		Todos: []string{"work", "sleep"},
	}

	t.Run("stores a new user", func(t *testing.T) {
		table := make(users)

		err := table.AddUser(details)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		got, ok := table[details.Email]
		if !ok {
			t.Fatalf("expected user with email %q to be stored", details.Email)
		}

		if !reflect.DeepEqual(got, details) {
			t.Fatalf("stored user mismatch: got %#v want %#v", got, details)
		}
	})

	t.Run("returns duplicate user error", func(t *testing.T) {
		table := make(users)

		if err := table.AddUser(details); err != nil {
			t.Fatalf("setup failed: expected first insert to succeed, got %v", err)
		}

		err := table.AddUser(details)
		if err == nil {
			t.Fatal("expected duplicate insert to return an error")
		}

		var userExistsErr *ErrUserExist
		if !errors.As(err, &userExistsErr) {
			t.Fatalf("expected ErrUserExist, got %T", err)
		}

		if userExistsErr.File != "db.go" {
			t.Fatalf("unexpected error file: got %q want %q", userExistsErr.File, "db.go")
		}

		if userExistsErr.Msg != "User already exist in the db table" {
			t.Fatalf("unexpected error message: got %q", userExistsErr.Msg)
		}
	})
}
