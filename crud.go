package main

import "fmt"

func AddNote(study *StudyAssistant) {
	var newNote Note
	newNote.ID = study.NoteCount + 1

	fmt.Print("Title: ")
	newNote.Title = ReadFullLine()

	fmt.Print("Content: ")
	newNote.Content = ReadFullLine()

	fmt.Print("Date (YYYY-MM-DD): ")
	fmt.Scan(&newNote.Date[0], &newNote.Date[1], &newNote.Date[2])

	fmt.Print("Difficulty Level (1-5): ")
	fmt.Scan(&newNote.Difficulty)

	study.Notes[study.NoteCount] = newNote
	study.NoteCount++
	NotifyCRUD("Note added")
}

func EditNote(study *StudyAssistant) {
	var i, id, index int
	var note Note

	displayNotes(*study)

	fmt.Print("Enter ID: ")
	fmt.Scan(&id)

	index = -1
	for i = 0; i < study.NoteCount; i++ {
		if study.Notes[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		NotifyNotFound()
	} else {
		note = study.Notes[index]

		fmt.Printf("Current title: %s\n", note.Title)
		fmt.Print("New title (press Enter if no change): ")
		var newTitle string
		newTitle = ReadFullLine()
		if len(newTitle) > 0 {
			note.Title = newTitle
		}

		fmt.Printf("Current content: %s\n", note.Content)
		fmt.Print("New content (press Enter if no change): ")
		newContent := ReadFullLine()
		if len(newContent) > 0 {
			note.Content = newContent
		}

		fmt.Printf("Current date: %d-%d-%d\n", note.Date[0], note.Date[1], note.Date[2])
		fmt.Print("New date (YYYY-MM-DD): ")
		var newDay, newMonth, newYear int
		fmt.Scan(&newYear, &newMonth, &newDay)
		note.Date[0] = newYear
		note.Date[1] = newMonth
		note.Date[2] = newDay

		fmt.Printf("Current difficulty: %d\n", note.Difficulty)
		fmt.Print("New difficulty level (1-5, press Enter if no change): ")
		var newDiff int
		fmt.Scanln(&newDiff)
		if newDiff != note.Difficulty && newDiff != 0 {
			note.Difficulty = newDiff
		}

		study.Notes[index] = note
		NotifyCRUD("Note updated")
	}
}

func DetailNote(study StudyAssistant) {
	var i, id, index int

	displayNotes(study)

	fmt.Print("Enter ID: ")
	fmt.Scan(&id)
	index = -1
	for i = 0; i < study.NoteCount; i++ {
		if study.Notes[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		NotifyNotFound()
	} else {
		note := study.Notes[index]
		fmt.Printf("Topic: %s\n Note: %s\n", note.Title, note.Content)
	}
}

func DeleteNote(study *StudyAssistant) {
	var i, id, index int

	displayNotes(*study)

	fmt.Scanf("Enter ID: %d", &id)

	index = -1
	for i = 0; i < study.NoteCount; i++ {
		if study.Notes[i].ID == id {
			index = i
		}
	}
	if index == -1 {
		NotifyNotFound()
	} else {
		var choice string
		fmt.Scanf("Are you sure? (y/n): %s", &choice)

		if choice != "y" && choice != "Y" {
			fmt.Printf("\n=====================\n[❌] Delete cancelled\n=====================\n\n")
		} else {
			for i = index; i < study.NoteCount-1; i++ {
				study.Notes[i] = study.Notes[i+1]
			}
			study.NoteCount--
			NotifyCRUD("Note deleted")
		}
	}
}
