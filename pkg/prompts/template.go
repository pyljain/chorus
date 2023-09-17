package prompts

import (
	"bytes"
	"fmt"
	"text/template"
)

func NewPromptTemplate(basePrompt string) (*Template, error) {
	tmpl, err := template.New("prompt").Parse(basePrompt)
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
