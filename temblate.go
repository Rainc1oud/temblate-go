package temblate_go

import (
	"bytes"
	"fmt"
	"text/template"
)

//go:generate go run ./cmd/gen.go

var templates = make(MessageTemplates, 0)

func GetMessage(lang, key string, data interface{}) string {
	if mts, ok := templates[key]; ok {
		if msg, ok := mts[lang]; ok {
			return newTemplateWrapErr(key, msg, data)
		} else if msg, ok := mts["en"]; ok { // try defaulting to English
			return newTemplateWrapErr(key, msg, data)
		}
	}
	return fmt.Sprintf("Error: no message defined for %q", key)
}

func newTemplateWrapErr(key, message string, data interface{}) string {
	var buf bytes.Buffer
	tmpl, err := template.New(key).Parse(message)
	if err != nil {
		return err.Error() // temporary direct return, probably not so good to expose internal error as template result...
	}
	if err := tmpl.Execute(&buf, data); err != nil {
		return err.Error() // temporary direct return, probably not so good to expose internal error as template result...
	}
	return string(buf.Bytes())
}
