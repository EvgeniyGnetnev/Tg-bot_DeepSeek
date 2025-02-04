package deepseek


type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type RequestBody struct {
    Model    string    `json:"model,omitempty"` 
    Messages []Message `json:"messages"`
}

type Response struct {
    ID      string `json:"id"`
    Model   string `json:"model"`
    Choices []struct {
        Message Message `json:"message"`
    } `json:"choices"`
}