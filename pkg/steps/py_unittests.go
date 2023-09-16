package steps

import (
	"fmt"
	"os/exec"
)

type PyUnitTest struct {
}

func NewPyUnitTest() Step {
	return &PyUnitTest{}
}

func (s *PyUnitTest) Execute(filepath string) error {
	testFile := fmt.Sprintf("%s_test.py", filepath)
	res, err := exec.Command("python3", testFile).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s . %s", err, res)
	}
	return nil
}
