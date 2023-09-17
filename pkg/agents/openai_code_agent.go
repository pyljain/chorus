package agents

import (
	"codeagent/pkg/llm"
	"codeagent/pkg/log"
	"codeagent/pkg/output"
	"codeagent/pkg/steps"
	"context"
	"fmt"
	"path"
)

const (
	maxTurns = 10
)

type openAICodeAgent struct {
	api             llm.LLM
	pipeline        []steps.Step
	outputDirectory string
	model           string
}

func NewOpenAICodeAgent(apiKey, outputDirectory string, model string) Agent {
	return &openAICodeAgent{
		api:             llm.NewOpenAI(apiKey),
		outputDirectory: outputDirectory,
		model:           model,
	}
}

func (ca *openAICodeAgent) ConfigurePipeline(steps []steps.Step) {
	ca.pipeline = steps
}

func (ca *openAICodeAgent) RunPipeline(startingPrompt string, filenameWithoutExtension string) (*PipelineRunResult, error) {
	ctx := context.Background()

	memory := []llm.Message{
		{
			Role:    "user",
			Content: startingPrompt,
		},
	}

	var rounds int32
	var isSuccessful = false

	for rounds = 1; rounds <= maxTurns; rounds++ {
		log.Write(memory)
		resp, _, err := ca.api.GetChatCompletion(ctx, &llm.ChatCompletionRequest{
			LLM:      "",
			Model:    ca.model,
			Messages: memory,
			Stream:   false,
		})
		if err != nil {
			return nil, err
		}

		memory = append(memory, llm.Message{
			Role:    "assistant",
			Content: resp.Choices[0].Message.Content,
		})

		log.Write(memory)

		err = output.Generate(ca.outputDirectory, filenameWithoutExtension, resp.Choices[0].Message.Content)
		if err != nil {
			memory = append(memory, llm.Message{
				Role:    "user",
				Content: fmt.Sprintf("Could not parse the response. Error is %s. Could you please correct and follow the same response format (JSON) as before.", err),
			})
			continue
		}

		fullFilePath := path.Join(ca.outputDirectory, filenameWithoutExtension)

		errorInStep := false
		for _, step := range ca.pipeline {
			err := step.Execute(fullFilePath)
			if err != nil {
				errorInStep = true
				memory = append(memory, llm.Message{
					Role:    "user",
					Content: fmt.Sprintf("Could not run the code. Error is %s. Could you please correct and follow the same response format (JSON) as before.", err),
				})
				// log.Write(memory)
				break
			}
		}

		if !errorInStep {
			isSuccessful = true
			break
		}
	}

	pr := PipelineRunResult{
		Filename:           filenameWithoutExtension,
		ConversationRounds: rounds,
		Success:            isSuccessful,
	}
	return &pr, nil
}
