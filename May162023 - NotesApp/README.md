# Notes App (Command-Line Based)

This is a simple command-line app made in **Go (Golang)** to add, list, and delete notes. The notes are saved in a **JSON file** on your local computer.

## 🛠️ What We Used

- **Go Programming Language** — to build the app
- **encoding/json** — to work with JSON data
- **ioutil / os** — to read/write files and handle command-line arguments

## Project Structure
notesapp/
│
├── main.go             → Main file to handle user input (add, list, delete)
├── Notes/
│   └── notes.go        → All logic for adding, listing, deleting notes
└── Data/
    └── notes.json      → Where notes are stored in JSON format

## How to Run

Make sure you are inside the project folder.

### Add a Note:
```bash
go run main.go add "Note Title" "Note Body"
```

### List All Notes:
```bash
go run main.go list
```

### Delete a Note:
```bash
go run main.go delete "Note Title"
```

> Note: Title must match exactly when deleting.

## How It Works

- When you **add a note**, it's stored in `notes.json` file.
- When you **list**, it reads from the same file and shows all notes.
- When you **delete**, it finds the note by title and removes it.

## Example

```bash
go run main.go add "Buy Groceries" "Buy milk, bread, and eggs"
go run main.go list
go run main.go delete "Buy Groceries"
```

---

Built for practice to learn how to use:
- Third-party packages with `go get`
- Reading and writing files
- JSON handling in Go
