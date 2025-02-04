package deepseek

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/EvgeniyGnetnev/Tg-bot_DeepSeek/lib/e"
)

type Client struct {
	apiKey  string
	siteURL string
}

func New(apiKey string, siteUrl string) *Client {
	return &Client{
		apiKey:  apiKey,
		siteURL: siteUrl,
	}
}

func (c *Client) DoRequest(question string) (answer string, err error){
	defer func() { err = e.WrapIfErr("can't do request to DeepSeek", err) }()

	requestBody := RequestBody{
		Model: "deepseek/deepseek-r1-distill-llama-70b:free",
		Messages: []Message{
			{
				Role:    "user",
				Content: question,
			},
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://openrouter.ai/api/v1/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ c.apiKey)
	req.Header.Set("HTTP-Referer", c.siteURL)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()


	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content, nil
	}

	return "Нет ответа от DeepSeek, попробуйте немного позже", nil
}
