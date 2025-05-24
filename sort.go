package main

func SelectionSort(notes *[100]Note, count int) {
	for i := 0; i < count-1; i++ {
		minIndex := i
		for j := i + 1; j < count; j++ {
			if notes[j].Difficulty < notes[minIndex].Difficulty {
				minIndex = j
			}
		}
		if minIndex != i {
			temp := notes[i]
			notes[i] = notes[minIndex]
			notes[minIndex] = temp
		}
	}
}

func InsertionSort(notes *[100]Note, count int) {
	for i := 1; i < count; i++ {
		key := notes[i]
		j := i - 1

		for j >= 0 && CheckDate(notes[j].Date, key.Date) {
			notes[j+1] = notes[j]
			j = j - 1
		}

		notes[j+1] = key
	}
}

func CheckDate(date1, date2 string) bool {
	if len(date1) < 10 || len(date2) < 10 {
		return false
	}
	year1 := StringToInt(date1[0:4])
	month1 := StringToInt(date1[5:7])
	day1 := StringToInt(date1[8:10])

	year2 := StringToInt(date2[0:4])
	month2 := StringToInt(date2[5:7])
	day2 := StringToInt(date2[8:10])

	if year1 > year2 {
		return true
	} else if year1 < year2 {
		return false
	}

	if month1 > month2 {
		return true
	} else if month1 < month2 {
		return false
	}

	return day1 > day2
}
