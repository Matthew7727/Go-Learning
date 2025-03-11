package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note-app/note"
)

func main() {
	title, content := getNoteData()

	newNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	newNote.DisplayNote()
	err = newNote.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
		return
	} else {
		fmt.Println("Saving the note succeeded")
	}
	

}

func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")

	content := getUserInput("Note Content: ")

	return title, content
}

func getUserInput(prompt string) string {

	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text

}
