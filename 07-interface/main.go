package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note-app/note"
)

type sever interface {
	Save() error
}

func main() {
	title, content := getNoteData()

	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	note.Display()

	err = note.Save()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Note has been saved successfully !")
}

func getNoteData() (string, string) {
	title := getUserInput("Enter note title: ")
	content := getUserInput("Enter note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var text string
	// fmt.Scanln(&note)
	/**
	* NewReader is a constructor function
	* os.Stdin  is the source from which we get input its meaning is standard input which mean input from terminal
	*
	 */

	reader := bufio.NewReader(os.Stdin)

	//use single code here because its special value type in GO called Rune
	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
