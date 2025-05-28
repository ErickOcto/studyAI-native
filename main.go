package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	var study StudyAssistant
	study.NoteCount = 0
	study.ScheduleCount = 0
	var choice string
	choice = " "
	for choice != "0" {
		fmt.Println("\n===== AI Learning Assistant Application =====")
		displayCurrentSchedules(study)
		fmt.Println("1.  [📝] Add Note")
		fmt.Println("2.  [📝] Edit Note")
		fmt.Println("3.  [🚮] Delete Note")
		fmt.Println("4.  [🗒️] View Note")
		fmt.Println("5.  [📶] Sort Notes by Difficulty")
		fmt.Println("6.  [📶] Sort Notes by Date")
		fmt.Println("7.  [📖] Display All Notes")
		fmt.Println("8.  [🔍] Search Notes")
		fmt.Println("9.  [🧠] AI-Powered Quiz")
		fmt.Println("10. [📅] Add Schedule")
		fmt.Println("0.  [⛔️] Exit")
		fmt.Print("Choose menu: ")

		fmt.Scan(&choice)

		switch choice {
		case "1":
			addNote(&study)
		case "2":
			editNote(&study)
		case "3":
			deleteNote(&study)
		case "4":
			detailNote(study)
		case "5":
			SelectionSort(&study)
			fmt.Println()
			fmt.Println("Notes sorted by difficulty")
			fmt.Println()
			displayNotes(study)
		case "6":
			InsertionSort(&study)
			fmt.Println()
			fmt.Println("Notes sorted by date")
			displayNotes(study)
		case "7":
			displayNotes(study)
		case "8":
			searchNotes(study)
		case "9":
			AiQuiz(study)
		case "10":
			addSchedule(&study)
		}
	}
}

func displayNotes(study StudyAssistant) {
	if study.NoteCount == 0 {
		fmt.Println("======================================================")
		fmt.Println("[🤷🏻] No notes available, please make at least one note")
		fmt.Println("======================================================")
	}

	fmt.Println()
	fmt.Println("====== Notes List ======")
	for i := 0; i < study.NoteCount; i++ {
		fmt.Printf("%d %s (Difficulty: %d, Date: %d-%d-%d)\n", study.Notes[i].ID, study.Notes[i].Title, study.Notes[i].Difficulty, study.Notes[i].Date[0], study.Notes[i].Date[1], study.Notes[i].Date[2])
	}
	fmt.Println("========================")
}

func addNote(study *StudyAssistant) {
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

	fmt.Println()
	fmt.Println("============================")
	fmt.Println("[✅] Note added successfully")
	fmt.Println("============================")
}

func searchNotes(study StudyAssistant) {
	if study.NoteCount == 0 {
		fmt.Println("======================================================")
		fmt.Println("[🤷🏻] No notes available, please make at least one note")
		fmt.Println("======================================================")
	}

	var searchType int
	fmt.Println()
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search in ID)")
	fmt.Print("Please choose method: ")
	fmt.Scan(&searchType)
	if searchType == 1 {
		var keyword string
		fmt.Print("Enter keyword: ")
		keyword = ReadFullLine()
		SequentialSearch(study.Notes, study.NoteCount, keyword)

	} else if searchType == 2 {
		var id int
		fmt.Print("Enter ID: ")
		fmt.Scan(&id)
		idx := BinarySearch(study.Notes, study.NoteCount, id)
		if idx != -1 {
			fmt.Printf("Note found:\n%s\n", study.Notes[idx].Title)
		} else {
			fmt.Println("===================")
			fmt.Println("[🥀] Ups, Not Found")
			fmt.Println("===================")
		}
	} else {
		fmt.Println("==============")
		fmt.Println("[🤷🏻] No method")
		fmt.Println("==============")
	}
}

func editNote(study *StudyAssistant) {
	if study.NoteCount == 0 {
		fmt.Println("======================================================")
		fmt.Println("[🤷🏻] No notes available, please make at least one note")
		fmt.Println("======================================================")
	}

	displayNotes(*study)

	fmt.Print("Enter ID: ")

	var id int
	fmt.Scan(&id)

	index := -1
	for i := 0; i < study.NoteCount; i++ {
		if study.Notes[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		fmt.Println("===================")
		fmt.Println("[🥀] Ups, Not Found")
		fmt.Println("===================")
	}

	note := study.Notes[index]

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
	fmt.Println("==============================")
	fmt.Println("[✅] Note updated successfully")
	fmt.Println("==============================")
}

func detailNote(study StudyAssistant) {
	if study.NoteCount == 0 {
		fmt.Println("========================================================")
		fmt.Println("[📝🤷🏻] No notes available, please make at least one note")
		fmt.Println("========================================================")
	}
	displayNotes(study)

	fmt.Print("Enter ID: ")
	var id int
	fmt.Scan(&id)

	index := -1
	for i := 0; i < study.NoteCount; i++ {
		if study.Notes[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		fmt.Println("===================")
		fmt.Println("[🥀] Ups, Not Found")
		fmt.Println("===================")
	}

	note := study.Notes[index]
	fmt.Printf("Topic: %s\n Note: %s\n", note.Title, note.Content)
}

func deleteNote(study *StudyAssistant) {
	if study.NoteCount == 0 {
		fmt.Println("========================================================")
		fmt.Println("[📝🤷🏻] No notes available, please make at least one note")
		fmt.Println("========================================================")
	}

	displayNotes(*study)

	fmt.Print("Enter ID: ")
	var id int
	fmt.Scan(&id)

	index := -1
	for i := 0; i < study.NoteCount; i++ {
		if study.Notes[i].ID == id {
			index = i
		}
	}
	if index == -1 {
		fmt.Println("=====================")
		fmt.Println("[🥀] Ups, Not Found")
		fmt.Println("=====================")
	}

	fmt.Printf("Are you sure? (y/n): ")
	var choice string
	fmt.Scan(&choice)
	if choice != "y" && choice != "Y" {
		fmt.Println("=====================")
		fmt.Println("[❌] Delete cancelled")
		fmt.Println("=====================")
	} else {
		for i := index; i < study.NoteCount-1; i++ {
			study.Notes[i] = study.Notes[i+1]
		}
		study.NoteCount--
		fmt.Println("==============================")
		fmt.Println("[✅] Note deleted successfully")
		fmt.Println("==============================")
	}

}

func addSchedule(study *StudyAssistant) {
	var newSchedule Schedule
	fmt.Print("Enter date (YYYY-MM-DD): ")
	var y, m, d int
	fmt.Scan(&y, &m, &d)
	newSchedule.Date[0] = y
	newSchedule.Date[1] = m
	newSchedule.Date[2] = d
	fmt.Print("Enter description: ")
	newSchedule.Description = ReadFullLine()
	study.Schedules[study.ScheduleCount] = newSchedule
	study.ScheduleCount++

	fmt.Println("================================")
	fmt.Println("[✅] Schedule added successfully")
	fmt.Println("================================")
}

func displayCurrentSchedules(study StudyAssistant) {
	fmt.Println()
	currentYear := time.Now().Year()
	currentMonth := int(time.Now().Month())
	currentDay := time.Now().Day()
	fmt.Println("Schedule in this month:")
	InsertionSortDate(&study)
	for i := 0; i < study.ScheduleCount; i++ {
		if study.Schedules[i].Date[0] == currentYear && study.Schedules[i].Date[1] == currentMonth && study.Schedules[i].Date[2] == currentDay {
			fmt.Printf("Today - %s\n", study.Schedules[i].Description)
		} else if study.Schedules[i].Date[0] >= currentYear && study.Schedules[i].Date[1] >= currentMonth && study.Schedules[i].Date[2] >= currentDay {
			fmt.Printf("%d-%d-%d - %s\n", study.Schedules[i].Date[0], study.Schedules[i].Date[1], study.Schedules[i].Date[2], study.Schedules[i].Description)
		}
		i++
	}

	if study.ScheduleCount == 0 {
		fmt.Println("===============================================================")
		fmt.Println("[📅🤷🏻] No schedule available, please make at least one schedule")
		fmt.Println("===============================================================")
	}
	fmt.Println()
}

func ReadFullLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}