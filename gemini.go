package main

import (
	"context"
	"fmt"
	"google.golang.org/genai"
	"log"
)

func AiQuiz(app StudyAssistant) {
	var topicId int
	var noteContent, finalGenerated string
	var i, index int
	var num1, num2, num3, num4, num5 string

	displayNotes(app)

	fmt.Println("Choose a topic based on ID:")
	fmt.Scan(&topicId)

	index = -1

	for i = 0; i < app.NoteCount; i++ {
		if app.Notes[i].ID == topicId {
			index = i
		}
	}

	if index == -1 {
		NotifyNotFound()
	} else {
		noteContent = app.Notes[index].Content
		fmt.Println("Generating whoosh... 🚀")

		ctx := context.Background()
		client, err := genai.NewClient(ctx, &genai.ClientConfig{
			APIKey:  "AIzaSyDt4n2quqlK8ZEmi10wvRVrufMzmPjpXqw",
			Backend: genai.BackendGeminiAPI,
		})
		if err != nil {
			log.Fatal(err)
		}
		result, err := client.Models.GenerateContent(
			ctx,
			"gemini-2.0-flash",
			genai.Text("Make me 5 question based on this material content, 3 multiple choices and 2 short essays: "+noteContent),
			nil,
		)
		if err != nil {
			log.Fatal(err)
		}

		finalGenerated = result.Text()
		fmt.Println(finalGenerated)
		fmt.Println("Answer for number 1")
		num1 = ReadFullLine()
		fmt.Println("Answer for number 2")
		num2 = ReadFullLine()
		fmt.Println("Answer for number 3")
		num3 = ReadFullLine()
		fmt.Println("Answer for number 4")
		num4 = ReadFullLine()
		fmt.Println("Answer for number 5")
		num5 = ReadFullLine()

		fmt.Printf("\nAnswers submitted!, And your final score is 🥺...\n")

		result, err = client.Models.GenerateContent(
			ctx,
			"gemini-2.0-flash",
			genai.Text("Give the user final score, with questions: "+finalGenerated+"User answers: "+" Num 1="+num1+" Num 2="+num2+" Num 3="+num3+" Num 4="+num4+" Num 5="+num5),
			nil,
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result.Text())
	}
}
