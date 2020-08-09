package templates

import "io/ioutil"

const (
	GoControllerTemplate      = "templates/goController.go.tmpl"
	GoServiceTemplate         = "templates/goService.go.tmpl"
	ProvisionDatabaseTemplate = "templates/provisionDatabase.sql.tmpl"
)

func LoadFile(filepath string) (string, error) {
	fileBytes, err := ioutil.ReadFile(filepath)
	return string(fileBytes), err
}
