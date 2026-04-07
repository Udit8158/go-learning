package blogrender

import (
	"embed"
	"html/template"
	"io"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func Render(w io.Writer, p *Post) error {

	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")

	if err != nil {
		return err
	}

	error := templ.ExecuteTemplate(w, "blog.gohtml", p)

	if error != nil {
		return error
	}

	return nil
}
