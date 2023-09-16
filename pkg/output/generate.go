package output

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func Generate(basePath string, filename string, openAIResponse string) error {
	lr := llmResponse{}
	err := json.Unmarshal([]byte(openAIResponse), &lr)
	if err != nil {
		return err
	}
	fullFileName := path.Join(basePath, fmt.Sprintf("%s%s", filename, lr.Extension))
	err = os.WriteFile(fullFileName, []byte(lr.Code), 0644)
	if err != nil {
		return err
	}
	testFileName := path.Join(basePath, fmt.Sprintf("%s_test%s", filename, lr.Extension))
	err = os.WriteFile(testFileName, []byte(lr.UnitTest), 0644)
	if err != nil {
		return err
	}

	return nil
}

type llmResponse struct {
	Explanation     string `json:"explanation"`
	Extension       string `json:"extension"`
	Code            string `json:"code"`
	UnitTest        string `json:"unitTest"`
	ConfidenceScore int32  `json:"confidenceScore"`
}
