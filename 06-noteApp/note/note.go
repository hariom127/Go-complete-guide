package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}
	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}

// It is defined with a receiver: (note Note). This means Display() can be called on instances of Note.
func (note Note) Display() {
	fmt.Printf("This is your note title: %v and this is content: \n\n %v \n", note.title, note.content)
}
