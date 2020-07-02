package templates

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"text/template"
)

var TemplateHelperFuncs = map[string]interface{}{
	"title": strings.Title,
	"plural": func(in string) string {
		return in + "s"
	},
	"lower": strings.ToLower,
	"last": func(x int, a interface{}) bool {
		return x == reflect.ValueOf(a).Len()-1
	},
	"incr": func(in int) int {
		return in + 1
	},
	"backticks": func(in string) string {
		return "`" + in + "`"
	},
	"quotes": func(in string) string {
		return "\"" + in + "\""
	},
	"sprintf": func(template string, args ...interface{}) string {
		return fmt.Sprintf(template, args...)
	},
}

func ToStr(model interface{}, name, inputTemplate string) (string, error) {
	tmpl, err := template.New(name).
		Funcs(TemplateHelperFuncs).
		Parse(inputTemplate)
	if err != nil {
		return "", err
	}

	buff := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(buff, model)
	if err != nil {
		return "", err
	}

	return string(buff.Bytes()), nil
}
