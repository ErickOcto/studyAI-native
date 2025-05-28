package main

func SelectionSort(study *StudyAssistant) {

	for i := 0; i < study.NoteCount-1; i++ {
		minIndex := i
		for j := i + 1; j < study.NoteCount; j++ {
			if study.Notes[j].Difficulty < study.Notes[minIndex].Difficulty {
				minIndex = j
			}
		}
		if minIndex != i {
			temp := study.Notes[i]
			study.Notes[i] = study.Notes[minIndex]
			study.Notes[minIndex] = temp
		}
	}
}

func InsertionSort(study *StudyAssistant) {
	for i := 1; i < study.NoteCount; i++ {
		key := study.Notes[i]
		j := i - 1
		for j >= 0 && CheckDate(study.Notes[j].Date[0], key.Date[0], study.Notes[j].Date[1], key.Date[1], study.Notes[j].Date[2], key.Date[2]) {
			study.Notes[j+1] = study.Notes[j]
			j = j - 1
		}
		study.Notes[j+1] = key
	}
}

func InsertionSortDate(date *StudyAssistant) {
	for i := 1; i < date.ScheduleCount; i++ {
		key := date.Schedules[i]
		j := i - 1
		for j >= 0 && CheckDate(date.Schedules[j].Date[0], key.Date[0], date.Schedules[j].Date[1], key.Date[1], date.Schedules[j].Date[2], key.Date[2]) {
			date.Schedules[j+1] = date.Schedules[j]
			j = j - 1
		}
		date.Schedules[j+1] = key
	}
}

func CheckDate(y1, y2, m1, m2, d1, d2 int) bool {
	if y1 > y2 {
		return true
	} else if y1 < y2 {
		return false
	}
	if m1 > m2 {
		return true
	} else if m1 < m2 {
		return false
	}
	return d1 > d2
}