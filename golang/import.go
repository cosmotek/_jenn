package golang

type Import struct {
	Package         string
	Alias           string
	SideEffectsOnly bool
}

var (
	TimePackage = Import{
		Package:         "time",
		Alias:           "",
		SideEffectsOnly: false,
	}
)
