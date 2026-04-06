package blogpost_test

import (
	"errors"
	"fmt"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogpost "github.com/Udit8158/go-learning/19_reading_files"
)

// our fake fs type only to create err
// this will satisfy like fs or mapfs
type StubFailingFs struct{}

func (fs StubFailingFs) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, I always fail")
}

func TestNewPost(t *testing.T) {
	t.Run("should create posts from a dir correctly", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1
Tags: rust, memory
---
Rust is rusty!
But that is Rust.
`
			secondBody = `Title: Post 2
Description: Description 2 hi there
Tags: go, concurrency
---
Go is gowine
`
		)

		// it's just a mock representation of a file system
		// which currently has 2 blogs (md files)
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		// this will create a new post from the file and return the posts array
		// if getting err, return that
		posts, err := blogpost.NewPostFromFS(fs)

		if err != nil {
			t.Fatal("ERROR Occured In TEST: ", err)
		}

		got := posts[0]
		want := blogpost.Post{
			Title: "Post 1", Description: "Description 1", Tags: []string{"rust", "memory"},
			Body: "Rust is rusty!\nBut that is Rust.",
		}

		fmt.Printf("%d", len(got.Body))
		fmt.Printf("%d", len(want.Body))

		assertPost(t, got, want)

		if len(posts) != len(fs) {
			t.Errorf("length of directory is %d but go %d", len(fs), len(posts))
		}
	})

	t.Run("should give error", func(t *testing.T) {
		fs := StubFailingFs{} // it will always give error
		_, err := blogpost.NewPostFromFS(fs)

		if err == nil {
			t.Errorf("Expected an error")
		}
	})
}

func assertPost(t *testing.T, got blogpost.Post, want blogpost.Post) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted %+v but got %+v", want, got)
	}
}
