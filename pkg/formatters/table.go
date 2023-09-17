package formatters

import (
	"codeagent/pkg/agents"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func PrintTable(result []*agents.PipelineRunResult) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("FILENAME", "CONVERSATION ROUNDS", "COMPLETED SUCCESSFULLY")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, pr := range result {
		tbl.AddRow(pr.Filename, pr.ConversationRounds, pr.Success)
	}

	tbl.Print()
}
