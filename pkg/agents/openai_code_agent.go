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
	maxTurns = 5
)

type openAICodeAgent struct {
	api             llm.LLM
	pipeline        []steps.Step
	outputDirectory string
}

func NewOpenAICodeAgent(apiKey, outputDirectory string) Agent {
	return &openAICodeAgent{
		api:             llm.NewOpenAI(apiKey),
		outputDirectory: outputDirectory,
	}
}

func (ca *openAICodeAgent) ConfigurePipeline(steps []steps.Step) {
	ca.pipeline = steps
}

func (ca *openAICodeAgent) RunPipeline(startingPrompt string, filenameWithoutExtension string) error {
	ctx := context.Background()

	memory := []llm.Message{
		{
			Role:    "user",
			Content: startingPrompt,
		},
	}

	for i := 0; i < maxTurns; i++ {
		log.Write(memory)
		resp, _, err := ca.api.GetChatCompletion(ctx, &llm.ChatCompletionRequest{
			LLM:      "",
			Model:    "gpt-4",
			Messages: memory,
			Stream:   false,
		})
		if err != nil {
			return err
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
			log.Write(memory)
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
				log.Write(memory)
				break
			}
		}

		if !errorInStep {
			break
		}
	}

	return nil
}
