package main

type Note struct {
	ID         int
	Title      string
	Content    string
	Date       string
	Difficulty int
}

type Exercise struct {
	ID       int
	Question string
	Options  [4]string
	Answer   int
	TopicID  int
}

type Schedule struct {
	Date        string
	Description string
}

type StudyAssistant struct {
	Notes         [100]Note
	NoteCount     int
	Exercises     [200]Exercise
	ExerciseCount int
	Schedules     [20]Schedule
	ScheduleCount int
}
