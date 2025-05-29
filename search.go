package main

import "fmt"

func SequentialSearch(notes [100]Note, count int, keyword string) {
	var i int
	for i = 0; i < count; i++ {
		if keyword == notes[i].Title {
			fmt.Printf("%s - %d-%d-%d\n", notes[i].Title, notes[i].Date[0], notes[i].Date[1], notes[i].Date[2])
		}
	}
}

func BinarySearch(notes [100]Note, count int, id int) int {
	var right, left, mid int
	left = 0
	right = count - 1
	for left <= right {
		mid = (left + right) / 2
		if notes[mid].ID == id {
			return mid
		}
		if notes[mid].ID < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func SelectionSortBinary(study *StudyAssistant) {
	var i, j, minIndex int
	var temp Note
	for i = 0; i < study.NoteCount-1; i++ {
		minIndex = i
		for j = i + 1; j < study.NoteCount; j++ {
			if study.Notes[j].ID < study.Notes[minIndex].ID {
				minIndex = j
			}
		}
		if minIndex != i {
			temp = study.Notes[i]
			study.Notes[i] = study.Notes[minIndex]
			study.Notes[minIndex] = temp
		}
	}
}
