package steps

import (
	"fmt"
	"os/exec"
)

type Pylint struct {
}

func NewPylint() Step {
	return &Pylint{}
}

func (s *Pylint) Execute(filepath string) error {
	mainFile := fmt.Sprintf("%s.py", filepath)
	res, err := exec.Command("pylint", "--fail-under=6", mainFile).CombinedOutput()
	if err != nil {
		return fmt.Errorf("lint output: %s \n %s", err, res)
	}
	return nil
}
