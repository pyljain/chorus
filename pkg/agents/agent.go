package agents

import "codeagent/pkg/steps"

type Agent interface {
	ConfigurePipeline(pipeline []steps.Step)
	RunPipeline(prompt string, filePath string) error
}
