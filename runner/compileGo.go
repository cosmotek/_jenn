package runner

type GoCompilationOp struct{}

func (g GoCompilationOp) Name() string {
	return ""
}

func (g GoCompilationOp) Exec() error {
	return nil
}
