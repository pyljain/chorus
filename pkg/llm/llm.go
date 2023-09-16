package llm

import "context"

type LLM interface {
	GetChatCompletion(context.Context, *ChatCompletionRequest) (*ChatCompletionResponse, int, error)
}

type ChatCompletionRequest struct {
	LLM      string    `json:"llm"`
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionResponse struct {
	Choices []ChatCompletionResponseChoice `json:"choices"`
	Id      string                         `json:"id"`
}

type ChatCompletionResponseChoice struct {
	Message Message `json:"message"`
}
