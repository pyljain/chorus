package steps

type Step interface {
	Execute(filepath string) error
}
