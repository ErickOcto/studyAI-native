package main

import (
	"fmt"
	"time"
)

func AddSchedule(study *StudyAssistant) {
	var newSchedule Schedule
	var y, m, d int

	fmt.Println("Enter date (YYYY-MM-DD): ")
	fmt.Scan(&y, &m, &d)
	newSchedule.Date[0] = y
	newSchedule.Date[1] = m
	newSchedule.Date[2] = d

	fmt.Print("Enter description: ")
	newSchedule.Description = ReadFullLine()
	study.Schedules[study.ScheduleCount] = newSchedule
	study.ScheduleCount++
	NotifyCRUD("Schedule added")
}

func DisplayCurrentSchedules(study StudyAssistant) {
	var i, currentYear, currentMonth, currentDay int
	if study.ScheduleCount == 0 {
		fmt.Printf("\n===============================================================\n[📅🤷🏻] No schedule available, please make at least one schedule\n===============================================================\n\n")
	} else {
		currentYear = time.Now().Year()
		currentMonth = int(time.Now().Month())
		currentDay = time.Now().Day()
		fmt.Println("Your schedule from now on:")
		InsertionSortDate(&study)
		for i = 0; i < study.ScheduleCount; i++ {
			if study.Schedules[i].Date[0] == currentYear && study.Schedules[i].Date[1] == currentMonth && study.Schedules[i].Date[2] == currentDay {
				fmt.Printf("Today, %d-%d-%d - %s\n", study.Schedules[i].Date[0], study.Schedules[i].Date[1], study.Schedules[i].Date[2], study.Schedules[i].Description)
			} else if study.Schedules[i].Date[0] > currentYear || (study.Schedules[i].Date[0] == currentYear && study.Schedules[i].Date[1] > currentMonth) || (study.Schedules[i].Date[0] == currentYear && study.Schedules[i].Date[1] == currentMonth && study.Schedules[i].Date[2] > currentDay) {
				fmt.Printf("%d-%d-%d - %s\n", study.Schedules[i].Date[0], study.Schedules[i].Date[1], study.Schedules[i].Date[2], study.Schedules[i].Description)
			}
		}
	}
}
