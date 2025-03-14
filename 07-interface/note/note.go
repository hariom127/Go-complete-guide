package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

// It is defined with a receiver: (note Note). This means Display() can be called on instances of Note.
func (note Note) Display() {
	fmt.Printf("This is your note title: %v and this is content: \n\n %v \n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")

	//Used for assigning a new value to an existing variable.
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)

	fmt.Printf("Json==== %v", json)

	if err != nil {
		fmt.Println(err)
	}
	return os.WriteFile(fileName, json, 0644)
}
