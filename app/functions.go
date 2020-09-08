package app

import "github.com/cosmotek/_jenn/ir"

type ClientType string
type ContentType string

const (
	JavascriptClient ClientType = "Javascript Client"
	GoClient         ClientType = "Go Client"
	LinuxCLI         ClientType = "Linux CLI"
	WindowsCLI       ClientType = "Windows CLI"
	MacosCLI         ClientType = "MacOS CLI"

	MarkdownDocs     ContentType = "markdown_docs"
	ExecutableAPI    ContentType = "executable_api"
	ContainerizedAPI ContentType = "containerized_api"

	ClientScript     ContentType = "client_script"
	ClientExecutable ContentType = "client_executable"
)

type Artifact struct {
	Name     string `json:"name"`
	MimeType string `json:"mimetype"`

	ContentType ContentType `json:"content_type"`
	Contents    []byte      `json:"contents"`
}

type ArtifactBundle struct {
	JennVersion string `json:"jenn_version"`

	AppName    string `json:"app_name"`
	AppVersion string `json:"app_version"`
	SchemaHash string `json:"schema_hash"`

	Artifacts []Artifact `json:"artifacts"`
}

func Version() string {
	return "0.0.0"
}

func GenerateAPI(schemaIR ir.ModelIR) (ArtifactBundle, error) {
	return ArtifactBundle{}, nil
}

func GenerateClient(schemaIR ir.ModelIR, outputType ClientType) (ArtifactBundle, error) {
	return ArtifactBundle{}, nil
}
