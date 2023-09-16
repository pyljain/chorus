package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	chatCompletionsEndpoint = "https://api.openai.com/v1/chat/completions"
)

type OpenAI struct {
	apiKey string
}

func NewOpenAI(apiKey string) *OpenAI {
	return &OpenAI{
		apiKey: apiKey,
	}
}

func (oai *OpenAI) GetChatCompletion(ctx context.Context, req *ChatCompletionRequest) (*ChatCompletionResponse, int, error) {
	openAIRequest := OpenAIChatCompletionRequest{
		Model:    req.Model,
		Messages: req.Messages,
		Stream:   req.Stream,
	}

	requestBytes, err := json.Marshal(&openAIRequest)
	if err != nil {
		// logger.Error("unable to marshal request to OpenAI", zap.Error(err))
		return nil, http.StatusBadRequest, err
	}

	request, err := http.NewRequest(http.MethodPost, chatCompletionsEndpoint, bytes.NewBuffer(requestBytes))
	if err != nil {
		// logger.Error("unable to construct request to OpenAI", zap.Error(err))
		return nil, http.StatusBadRequest, err
	}

	client := http.Client{}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", oai.apiKey))
	resp, err := client.Do(request)
	if err != nil {
		// logger.Error("unable to reach LLM", zap.Error(err))
		return nil, http.StatusInternalServerError, err
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	if resp.StatusCode >= 400 {
		return nil, resp.StatusCode, fmt.Errorf("error received from the OpenAI response: Status Code: %d", resp.StatusCode)
	}

	openAIChatResponse := ChatCompletionResponse{}

	err = json.Unmarshal(respBytes, &openAIChatResponse)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return &openAIChatResponse, http.StatusOK, nil

}

type OpenAIChatCompletionRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}
