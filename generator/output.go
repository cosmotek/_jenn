package generator

type OutputConfig struct {
	OutputDirectory string
	Outputs         map[OutputType]OutputOpts

	TransportsEnabled []TransportType
}

type OutputOpts struct {
	Namespaces   []string
	UIDocEnabled bool
}
