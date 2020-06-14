package generator

type OutputType string

const (
	GoClient        OutputType = "goClient"
	GoServer        OutputType = "goServer"
	JSBrowserClient OutputType = "jsBrowserClient"
	CSharpClient    OutputType = "csClient"
	RubyClient      OutputType = "rbClient"
)
