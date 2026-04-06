package blogpost

import (
	"bufio"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPostFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")

	// handeling error
	if err != nil {
		return nil, err
	}

	posts := make([]Post, 0, len(dir)) // len 0 but cap - what we needed

	for _, file := range dir {
		post, err := getPost(fileSystem, file.Name())

		if err != nil {
			return nil, err // one file open fails the entire thing now
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// utility functions

// reading a file and creating a post from that data
func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	return newPost(postFile)
}

const (
	titleSeperator       = "Title: "
	descriptionSeperator = "Description: "
	tageSeperator        = "Tags: "
)

// creating a post from the file data
func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLine := func(seperator string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), seperator)
	}

	// titleLine - starts after Title: -> then
	// same for descriptionLine
	titleLine := readLine(titleSeperator)
	descriptionLine := readLine(descriptionSeperator)
	tagsLine := readLine(tageSeperator)
	tags := strings.Split(tagsLine, ", ")

	return Post{Title: titleLine, Description: descriptionLine, Tags: tags, Body: readBody(readLine)}, nil
}

func readBody(readLine func(s string) string) string {
	// for body writting the lines (after --- format) in an arr
	// then joining them with new line as seperator in the post
	var bodyLines []string

	for {
		output := readLine("") // no seperator - so whole line output
		if output == "---" {
			continue
		}
		if output == "" {
			break
		}

		bodyLines = append(bodyLines, output)
	}

	postBody := strings.Join(bodyLines, "\n")

	return postBody
}
