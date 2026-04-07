package blogrender_test

import (
	"bytes"
	"testing"

	blogrender "github.com/Udit8158/go-learning/20_blog_render"
	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {

	aPost := blogrender.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	t.Run("it converst a signle post into a html", func(t *testing.T) {
		// we will write the html to this buffer
		buf := bytes.Buffer{}
		// var buf *bytes.Buffer
		err := blogrender.Render(&buf, &aPost)

		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())

	})

}
