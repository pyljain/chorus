package main

import (
	"codeagent/pkg/agents"
	"codeagent/pkg/config"
	"codeagent/pkg/formatters"
	"codeagent/pkg/prompts"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {

	configLoc := flag.String("config", "./chorus.yaml", "Location of the config file\n")
	flag.Parse()

	// Get directory with incoming files to convert
	dir := flag.Arg(0)
	if dir == "" {
		fmt.Printf("Make sure to provide the root directory for incoming code files to process\n")
		os.Exit(-1)
	}

	if *configLoc == "" {
		fmt.Printf("Make sure to provide the config file location\n")
		os.Exit(-1)
	}

	cfg, err := config.New(*configLoc)
	if err != nil {
		fmt.Printf("Could not read the provided config file %s\n", err)
		os.Exit(-1)
	}

	templateManager, err := prompts.NewPromptTemplate(cfg.Prompt)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	outputDirectory := path.Join(dir, "output")

	openAIAPIKey := os.Getenv("OPENAI_API_KEY")
	agent := agents.NewOpenAICodeAgent(openAIAPIKey, outputDirectory, cfg.Model)
	steps, err := cfg.GetPipelineSteps()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	agent.ConfigurePipeline(steps)

	// Read files
	inputDirectory := path.Join(dir, "input")
	result := []*agents.PipelineRunResult{}

	filepath.WalkDir(inputDirectory, func(filepath string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			codeBytes, err := os.ReadFile(filepath)
			if err != nil {
				return err
			}

			filenameWithoutExtension := strings.Split(d.Name(), ".")[0]

			prompt, err := templateManager.HydrateTemplate(filenameWithoutExtension, string(codeBytes))
			if err != nil {
				return err
			}

			pipelineResult, err := agent.RunPipeline(prompt, filenameWithoutExtension)
			if err != nil {
				return err
			}

			result = append(result, pipelineResult)
		}
		return nil
	})

	// Print output
	formatters.PrintTable(result)
}
