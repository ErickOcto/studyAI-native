package main

import (
	"context"
	"fmt"
	"google.golang.org/genai"
	"log"
)

func AiQuiz(app StudyAssistant) {
	if app.NoteCount == 0 {
		fmt.Println("No notes, make it first")
		return
	}
	displayNotes(app)
	fmt.Println("Choose a topic based on ID:")
	var topicId int
	fmt.Scan(&topicId)

	index := -1
	for i := 0; i < app.NoteCount; i++ {
		if app.Notes[i].ID == topicId {
			index = i
		}
	}
	if index == -1 {
		fmt.Println("Not found")
		return
	}

	noteContent := app.Notes[index].Content

	fmt.Println("Generating whushhhhh 🚀")

	//   GEMINI DOCS, Kak tolong APIKey nya jangan di apa apain :)
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
	finalGenerated := result.Text()
	fmt.Println(finalGenerated)
	var num1, num2, num3, num4, num5 string
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
	fmt.Println()
	fmt.Println("Answers submitted!, And your final score is🥺...")

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
	//END
}