package prompts

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

func NewPromptTemplate(basePromptFileName string) (*Template, error) {
	basePromptBytes, err := os.ReadFile(basePromptFileName)
	if err != nil {
		return nil, fmt.Errorf("unable to read the base prompt provided %s", err)
	}

	pt := promptTemplate{}
	err = yaml.Unmarshal(basePromptBytes, &pt)
	if err != nil {
		return nil, fmt.Errorf("unable to parse prompt provided %s", err)
	}

	tmpl, err := template.New("prompt").Parse(string(pt.Prompt))
	if err != nil {
		return nil, fmt.Errorf("unable to parse prompt template %s", err)
	}

	return &Template{
		template: tmpl,
	}, nil

}

func (t *Template) HydrateTemplate(packageName, code string) (string, error) {
	w := bytes.NewBufferString("")
	err := t.template.Execute(w, map[string]string{
		"code":     code,
		"basename": packageName,
	})

	if err != nil {
		return "", fmt.Errorf("unable to hydrate the template with code %s", err)
	}

	return w.String(), nil

}

type Template struct {
	template *template.Template
}

type promptTemplate struct {
	Prompt string `yaml:"prompt"`
}
