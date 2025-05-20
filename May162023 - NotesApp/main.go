package main

import (
	notes "notesapp/Notes"
	"os"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) < 2 {
		color.Red("Usage: go run main.go [add|list|delete] [arguments...]")
		return
	}

	action := os.Args[1]
	switch action {
	case "add":
		if len(os.Args) < 4 {
			color.Red("Usage: go run main.go add [title] [body]")
			return
		}

		notes.AddNote(notes.Note{
			Title: os.Args[2],
			Desc:  os.Args[3],
		})
	case "list":
		notes.ListNotes()
	case "delete":
		if len(os.Args) < 3 {
			color.Red("Usage: go run main.go delete [title]")
			return
		}

		notes.DeleteNote(os.Args[2])
	default:
		color.Red("Invalid Operation")
		color.Red("Usage: go run main.go [add|list|delete] [arguments...]")
	}
}
