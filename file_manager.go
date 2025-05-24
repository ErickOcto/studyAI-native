package main

import (
	"bufio"
	"fmt"
	"os"
)

func LoadNotes(filename string) ([100]Note, int, error) {
	var notes [100]Note
	var count int
	var continueReading bool
	count = 0
	file, err := os.Open(filename)
	if err != nil {
		return notes, count, err
	}
	scanner := bufio.NewScanner(file)
	continueReading = true

	for scanner.Scan() && count < 100 && continueReading {
		var idStr string
		idStr = scanner.Text()

		continueReading = scanner.Scan()
		if !continueReading {
			return notes, count, nil
		}
		title := scanner.Text()

		continueReading = scanner.Scan()
		if !continueReading {
			return notes, count, nil
		}
		content := scanner.Text()

		continueReading = scanner.Scan()
		if !continueReading {
			return notes, count, nil
		}
		date := scanner.Text()

		continueReading = scanner.Scan()
		if !continueReading {
			return notes, count, nil
		}
		difficultyStr := scanner.Text()
		continueReading = scanner.Scan()
		if !continueReading || scanner.Text() != "---" {
			return notes, count, nil
		}

		notes[count].ID = StringToInt(idStr)
		notes[count].Title = title
		notes[count].Content = content
		notes[count].Date = date
		notes[count].Difficulty = StringToInt(difficultyStr)

		count++
	}

	if err := scanner.Err(); err != nil {
		return notes, count, err
	}

	defer file.Close()
	return notes, count, nil
}

func SaveNotes(filename string, notes [100]Note, count int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < count; i++ {
		writer.WriteString(fmt.Sprint(notes[i].ID) + "\n")
		writer.WriteString(notes[i].Title + "\n")
		writer.WriteString(notes[i].Content + "\n")
		writer.WriteString(notes[i].Date + "\n")
		writer.WriteString(fmt.Sprint(notes[i].Difficulty) + "\n")
		writer.WriteString("---\n")
	}
	return writer.Flush()
}
