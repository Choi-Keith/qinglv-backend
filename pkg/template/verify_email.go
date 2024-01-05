package template

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
	"runtime"
)

func GenerateVerifyBody(registerUrl string) (string, error) {
	_, filename, _, _ := runtime.Caller(0)
	file := fmt.Sprintf("%s/%s", filepath.Dir(filename), "verify_email.html")
	tmpl, err := template.ParseFiles(file)
	if err != nil {
		return "", err
	}
	bodyBuf := &bytes.Buffer{}
	templateData := struct {
		RegisterUrl string
	}{
		RegisterUrl: registerUrl,
	}
	if err := tmpl.Execute(bodyBuf, templateData); err != nil {
		return "", err
	}
	return bodyBuf.String(), nil
}
