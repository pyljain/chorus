package main

import (
	"codeagent/pkg/agents"
	"codeagent/pkg/prompts"
	"codeagent/pkg/steps"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {

	baseprompt := flag.String("baseprompt", "", "Provide the file location for the base prompt YAML")
	flag.Parse()

	// Get directory with incoming files to convert
	dir := flag.Arg(0)
	if dir == "" {
		fmt.Printf("Make sure to provide the root directory for incoming code files to process")
		os.Exit(-1)
	}

	// Get the prompt template from a file location

	if *baseprompt == "" {
		fmt.Printf("Make sure to provide the base prompt YAML location")
		os.Exit(-1)
	}

	templateManager, err := prompts.NewPromptTemplate(*baseprompt)
	if err != nil {
		fmt.Println(err)
	}

	outputDirectory := path.Join(dir, "output")

	openAIAPIKey := os.Getenv("OPENAI_API_KEY")
	agent := agents.NewOpenAICodeAgent(openAIAPIKey, outputDirectory)
	agent.ConfigurePipeline([]steps.Step{
		steps.NewPyUnitTest(),
		steps.NewPylint(),
	})

	// Read files
	inputDirectory := path.Join(dir, "input")

	err = filepath.WalkDir(inputDirectory, func(filepath string, d fs.DirEntry, err error) error {

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

			err = agent.RunPipeline(prompt, filenameWithoutExtension)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Unable to parse input directory %s", err)
		os.Exit(-1)
	}

	// Generate hydrated prompt
	// Populate the output directory
}
