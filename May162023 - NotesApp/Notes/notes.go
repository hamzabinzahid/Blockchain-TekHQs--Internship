package notes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/fatih/color"
)

type Note struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

const path string = "Data/notes.json"

// Load Notes from the json file
func LoadNotes() ([]Note, error) {
	var notes []Note
	data, err := ioutil.ReadFile(path)
	if err != nil {
		color.Yellow("Error reading notes file:", err)
		return nil, err
	}

	err = json.Unmarshal(data, &notes) // This function returns an error if any
	if err != nil {
		color.Yellow("Error decoding json:", err)
		return nil, err
	}

	return notes, err
}

// Save the notes to the json file
func SaveNotes(notes []Note) error {
	data, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0644) //This function returns an error
}

func AddNote(note Note) {
	fmt.Println("Title : ", note.Title)
	fmt.Println("Desc  : ", note.Desc)

	// Load Notes
	notes, _ := LoadNotes()
	notes = append(notes, note)

	//Save Notes
	err := SaveNotes(notes)
	if err != nil {
		return
	}
	color.Green("Note added successfully")
}

func ListNotes() {
	notes, _ := LoadNotes()
	for _, note := range notes {
		color.Green("Title : ", note.Title)
		color.Green("Desc  : ", note.Desc)
		fmt.Println()
	}
}

func DeleteNote(title string) {
	// Fetch Existing Notes
	notes, _ := LoadNotes()
	updatedNotes := []Note{}
	var found bool = false

	// Loop Over the Fetched Notes
	for _, note := range notes {
		// If the note is found, remove it from the slice
		if note.Title != title {
			updatedNotes = append(updatedNotes, note)
		} else {
			found = true
		}
	}

	if !found {
		color.Red("Note with title '%s' not found", title)
		return
	}

	// Save the updated notes back to the file
	SaveNotes(updatedNotes)
	color.Green("Note deleted successfully")
}
