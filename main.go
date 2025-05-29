package main

import (
	"fmt"
)

func main() {
	var study StudyAssistant
	study.NoteCount = 0
	study.ScheduleCount = 0
	var choice string
	choice = " "
	for choice != "0" {
		fmt.Println("\n===== AI Learning Assistant Application =====")
		DisplayCurrentSchedules(study)
		displayNotes(study)
		fmt.Println("1.  [📝] Add Note")
		fmt.Println("2.  [📝] Edit Note")
		fmt.Println("3.  [🚮] Delete Note")
		fmt.Println("4.  [🗒️] View Note")
		fmt.Println("5.  [📶] Sort Notes by Difficulty")
		fmt.Println("6.  [📶] Sort Notes by Date")
		fmt.Println("7.  [🔍] Search Notes")
		fmt.Println("8.  [🧠] AI-Powered Quiz")
		fmt.Println("9. [📅] Add Schedule")
		fmt.Println("0.  [⛔️] Exit")
		fmt.Print("Choose menu: ")
		fmt.Scan(&choice)
		switch choice {
		case "1":
			AddNote(&study)
		case "2":
			if study.NoteCount == 0 {
				Notify()
			} else {
				EditNote(&study)
			}
		case "3":
			if study.NoteCount == 0 {
				Notify()
			} else {
				DeleteNote(&study)
			}
		case "4":
			if study.NoteCount == 0 {
				Notify()
			} else {
				DetailNote(study)
			}
		case "5":
			SelectionSort(&study)
			fmt.Printf("\nNotes sorted by difficulty\n\n")
		case "6":
			InsertionSort(&study)
			fmt.Printf("\nNotes sorted by date\n\n")
		case "7":
			searchNotes(study)
		case "8":
			AiQuiz(study)
		case "9":
			AddSchedule(&study)
		}
	}
}

func displayNotes(study StudyAssistant) {
	if study.NoteCount == 0 {
		Notify()
	} else {
		fmt.Println()
		fmt.Printf("\n====== Notes List ======\n\n")
		for i := 0; i < study.NoteCount; i++ {
			fmt.Printf("%d %s (Difficulty: %d, Date: %d-%d-%d)\n", study.Notes[i].ID, study.Notes[i].Title, study.Notes[i].Difficulty, study.Notes[i].Date[0], study.Notes[i].Date[1], study.Notes[i].Date[2])
		}
		fmt.Printf("\n========================\n\n")
	}
}

func searchNotes(study StudyAssistant) {
	if study.NoteCount == 0 {
		Notify()
	} else {
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
				NotifyNotFound()
			}
		} else {
			fmt.Println("==============\n[🤷🏻] Wrong method\n==============")
		}
	}
}
