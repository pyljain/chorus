package agents

import "codeagent/pkg/steps"

type Agent interface {
	ConfigurePipeline(pipeline []steps.Step)
	RunPipeline(prompt string, filePath string) (*PipelineRunResult, error)
}

type PipelineRunResult struct {
	Filename           string
	ConversationRounds int32
	Success            bool
}
