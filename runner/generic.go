package runner

import "fmt"

type Step interface {
	Name() string
	Exec() error
}

type Chain struct {
	steps []Step
}

func (c *Chain) Add(step Step) {

}

func (c *Chain) Exec() error {
	for _, step := range c.steps {
		err := step.Exec()
		if err != nil {
			return fmt.Errorf("error executing step '%s'", step.Name)
		}
	}

	return nil
}
