package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	var app StudyAssistant
	app.NoteCount = 0
	app.ExerciseCount = 0
	app.ScheduleCount = 0
	var choice string
	today := time.Now().Format("2006-01-02")

	choice = " "

	for choice != "0" {
		fmt.Println("\n===== AI Learnin-g Assistant Application =====")
		DisplayUpcomingSchedules(app, today)

		fmt.Println("1. Add Note")
		fmt.Println("2. Display All Notes")
		fmt.Println("3. Search Notes")
		fmt.Println("4. Sort Notes by Difficulty")
		fmt.Println("5. Sort Notes by Date")
		fmt.Println("6. Edit Note")
		fmt.Println("7. Delete Note")
		fmt.Println("8. AI-Powered Quiz")
		fmt.Println("9. Save Notes to File")
		fmt.Println("10. Load Notes from File")
		fmt.Println("11. Add Schedule")
		fmt.Println("0. Exit")
		fmt.Print("Choose menu: ")

		fmt.Scan(&choice)

		if choice == "1" {
			app = addNote(app)
		} else if choice == "2" {
			displayNotes(app)
		} else if choice == "3" {
			searchNotes(app)
		} else if choice == "4" {
			SelectionSort(&app.Notes, app.NoteCount)
			fmt.Println("Notes sorted by difficulty")
		} else if choice == "5" {
			InsertionSort(&app.Notes, app.NoteCount)
			fmt.Println("Notes sorted by date")
		} else if choice == "6" {
			app = editNote(app)
		} else if choice == "7" {
			app = deleteNote(app)
		} else if choice == "8" {
			AiQuiz(app)
		} else if choice == "9" {
			fmt.Print("Enter filename: ")
			var filename string
			fmt.Scanln(&filename)
			err := SaveNotes(filename, app.Notes, app.NoteCount)
			if err != nil {
				fmt.Println("error at function save notes")
			}
			fmt.Println("Notes Successfully saved")
		} else if choice == "10" {
			fmt.Print("Enter filename: ")
			var filename string
			fmt.Scanln(&filename)
			notes, count, err := LoadNotes(filename)
			if err != nil {
				fmt.Println("Failed to get notes from your file ><:", err)
			} else {
				app.Notes = notes
				app.NoteCount = count
				fmt.Println("Successfully loaded notes")
			}
		} else if choice == "11" {
			AddSchedule(&app)
		}
	}
}

func addNote(app StudyAssistant) StudyAssistant {
	var newNote Note
	newNote.ID = app.NoteCount + 1

	fmt.Print("Input Title: ")
	newNote.Title = ReadFullLine()

	fmt.Print("Contents: ")
	newNote.Content = ReadFullLine()

	fmt.Print("Date: ")
	fmt.Scanln(&newNote.Date)

	fmt.Print("Difficulty Level (1-5): ")
	fmt.Scan(&newNote.Difficulty)

	app.Notes[app.NoteCount] = newNote
	app.NoteCount++

	fmt.Println("Note added successfully")
	return app
}

func searchNotes(app StudyAssistant) {
	var searchType int
	fmt.Println("\n1. Sequential Search")
	fmt.Println("2. Binary Search in ID)")
	fmt.Print("please choose method: ")
	fmt.Scan(&searchType)
	if searchType == 1 {
		var keyword string
		fmt.Print("Enter keyword: ")
		keyword = ReadFullLine()
		SequentialSearch(app.Notes, app.NoteCount, keyword)

	} else if searchType == 2 {
		var id int
		fmt.Print("Enter ID: ")
		fmt.Scan(&id)
		idx := BinarySearch(app.Notes, app.NoteCount, id)
		if idx != -1 {
			fmt.Printf("Note found:\n%s\n", app.Notes[idx].Title)
		} else {
			fmt.Println("Not found")
		}
	}
}

func editNote(app StudyAssistant) StudyAssistant {
	if app.NoteCount == 0 {
		fmt.Println("Not found")
		return app
	}

	displayNotes(app)

	fmt.Print("Enter ID: ")
	var id int
	fmt.Scan(&id)

	index := -1
	for i := 0; i < app.NoteCount; i++ {
		if app.Notes[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		fmt.Println("Not found")
		return app
	}

	note := app.Notes[index]

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

	fmt.Printf("Current date: %s\n", note.Date)
	fmt.Print("New date (YYYY-MM-DD, press Enter if no change): ")
	var newDate string
	fmt.Scan(&newDate)
	if len(newDate) > 0 {
		note.Date = newDate
	}

	fmt.Printf("Current difficulty: %d\n", note.Difficulty)
	fmt.Print("New difficulty level (1-5, press Enter if no change): ")
	var newDiff int
	fmt.Scan(&newDiff)

	if newDiff != note.Difficulty {
		note.Difficulty = newDiff
	}
	app.Notes[index] = note
	fmt.Println("Note updated")

	return app
}

func displayNotes(app StudyAssistant) {
	if app.NoteCount == 0 {
		fmt.Println("No data")
		return
	}

	fmt.Println("\n===== Notes List =====")
	for i := 0; i < app.NoteCount; i++ {
		fmt.Printf("%d %s (Difficulty: %d, Date: %s)\n", app.Notes[i].ID, app.Notes[i].Title, app.Notes[i].Difficulty, app.Notes[i].Date)
	}
}

func deleteNote(app StudyAssistant) StudyAssistant {
	if app.NoteCount == 0 {
		fmt.Println("No data")
		return app
	}
	displayNotes(app)
	fmt.Print("Enter ID: ")
	var id int
	fmt.Scan(&id)

	index := -1
	for i := 0; i < app.NoteCount; i++ {
		if app.Notes[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		fmt.Println("not found")
		return app
	}

	fmt.Printf("Are you sure? (y/n): ")
	var noteID string
	fmt.Scan(&noteID)
	if noteID != "y" && noteID != "Y" {
		fmt.Println("Delete was cancelled")
		return app
	}

	for i := index; i < app.NoteCount-1; i++ {
		app.Notes[i] = app.Notes[i+1]
	}
	app.NoteCount--

	fmt.Println("Note deleted")
	return app
}

func AddSchedule(app *StudyAssistant) {
	var newSchedule Schedule
	fmt.Print("Enter date (YYYY-MM-DD): ")
	fmt.Scanf("%s", &newSchedule.Date)
	fmt.Print("Enter description: ")
	fmt.Scanln(newSchedule.Description)

	app.Schedules[app.ScheduleCount] = newSchedule
	app.ScheduleCount++

	fmt.Println("Schedule was added")
}

func DisplayUpcomingSchedules(app StudyAssistant, today string) {
	hasSchedule := false
	i := 0

	fmt.Println("\nSchedule:")
	for i < app.ScheduleCount {
		if app.Schedules[i].Date > today {
			fmt.Printf("%s - %s\n", app.Schedules[i].Date, app.Schedules[i].Description)
			hasSchedule = true
		}
		i++
	}

	if !hasSchedule {
		fmt.Println("No schedule, ")
	}
	fmt.Println()
}

func ReadFullLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func StringToInt(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		digit := int(s[i] - '0')
		if digit >= 0 && digit <= 9 {
			result = result*10 + digit
		}
	}
	return result
}
