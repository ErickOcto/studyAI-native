package main

type Note struct {
	ID         int
	Title      string
	Content    string
	Date       string
	Difficulty int
}

type Schedule struct {
	Date        string
	Description string
}

type StudyAssistant struct {
	Notes         [100]Note
	NoteCount     int
	ExerciseCount int
	Schedules     [20]Schedule
	ScheduleCount int
}
