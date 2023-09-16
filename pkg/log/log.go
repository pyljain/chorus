package log

import (
	"codeagent/pkg/llm"
	"strings"

	"github.com/fatih/color"
)

func Write(memory []llm.Message) {

	latestMessage := memory[len(memory)-1]

	actorFormatting := color.New(color.Bold, color.FgHiYellow).PrintfFunc()
	actorFormatting("%s says: ", strings.ToUpper(latestMessage.Role))
	color.White("%s", latestMessage.Content)
}
