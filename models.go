package main

type Note struct {
	ID         int
	Title      string
	Content    string
	Date       [3]int
	Difficulty int
}

type Schedule struct {
	Date        [3]int
	Description string
}

type StudyAssistant struct {
	Notes         [100]Note
	NoteCount     int
	Schedules     [20]Schedule
	ScheduleCount int
}