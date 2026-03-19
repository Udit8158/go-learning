package api

import (
	"io"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/Udit8158/go-learning/09_demo_db_app/db"
)

func TestAddUserToDB(t *testing.T) {
	details := db.UserDetails{
		Name:  "Udit",
		Email: "udit@email.com",
		Todos: []string{"work", "sleep"},
	}

	t.Run("adds a new user to the db table", func(t *testing.T) {
		clear(db.UserTable)

		AddUserToDB(details)

		got, ok := db.UserTable[details.Email]
		if !ok {
			t.Fatalf("expected user with email %q to be stored", details.Email)
		}

		if !reflect.DeepEqual(got, details) {
			t.Fatalf("stored user mismatch: got %#v want %#v", got, details)
		}
	})

	t.Run("prints duplicate user error details", func(t *testing.T) {
		clear(db.UserTable)

		AddUserToDB(details)
		output := captureStdout(t, func() {
			AddUserToDB(details)
		})

		if !strings.Contains(output, `ERROR occured in "db.go" - "User already exist in the db table"`) {
			t.Fatalf("expected duplicate error details in output, got %q", output)
		}
	})
}

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	oldStdout := os.Stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create stdout pipe: %v", err)
	}

	os.Stdout = writer
	t.Cleanup(func() {
		os.Stdout = oldStdout
	})

	fn()

	if err := writer.Close(); err != nil {
		t.Fatalf("failed to close writer: %v", err)
	}

	out, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("failed to read captured stdout: %v", err)
	}

	return string(out)
}
