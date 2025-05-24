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

	//   GEMINI DOCSS
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
	fmt.Println(result.Text())
	//END
}
