package config

import (
	"codeagent/pkg/steps"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Prompt        string `yaml:"prompt"`
	PipelineSteps []Step `yaml:"pipeline_steps"`
	Model         string `yaml:"model"`
}

type Step struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

func New(path string) (*Config, error) {
	configBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := Config{}
	err = yaml.Unmarshal(configBytes, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil

}

func (c *Config) GetPipelineSteps() ([]steps.Step, error) {
	result := []steps.Step{}

	for _, s := range c.PipelineSteps {
		var step steps.Step

		switch s.Type {
		case "py_unittests":
			step = steps.NewPyUnitTest()
		case "pylint":
			step = steps.NewPylint()
		default:
			return nil, fmt.Errorf("invalid step type %s", s.Type)
		}

		result = append(result, step)
	}

	return result, nil
}
