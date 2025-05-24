package main

import "fmt"

func SequentialSearch(notes [100]Note, count int, keyword string) {
	counter := 0
	for i := 0; i < count; i++ {
		if keyword == notes[i].Title {
			fmt.Printf("%s %s\n", notes[i].Title, notes[i].Date)
			counter++
		}
	}
	if counter == 0 {
		fmt.Print("Note not found")
	}
}
func BinarySearch(notes [100]Note, count int, id int) int {
	left := 0
	right := count - 1
	for left <= right {
		mid := (left + right) / 2
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
