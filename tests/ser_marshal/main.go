package main 

import (
	"encoding/json"
	"fmt"
)

type Messages struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type ReqBody struct {
    Model      string   `json:"model"`
    Temperature string  `json:"temperature"`
    Messages   Messages `json:"messages"`
}

func main() {
	translateLine()
}

func translateLine() {
    gameName := "Name"
    outputLanguage := "en"
    APIendpoint := "endpoint"
    APIrole := "role"
    LLMmodel := "model"
    temperature := "temperature"
    additional := "additional"
    prompt := "prompt"
    content := fmt.Sprintf(prompt, gameName, outputLanguage, additional, APIendpoint, "text")
    body := &ReqBody{
        Model: LLMmodel,
        Temperature: temperature,
        Messages: Messages{
            Role: APIrole,
            Content: content,
        },
    }

	b, err := json.Marshal(body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
}