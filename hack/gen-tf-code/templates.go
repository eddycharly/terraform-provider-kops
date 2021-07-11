package main

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"text/template"
)

const schemasTemplate = "hack/gen-tf-code/templates/schema.tpl"
const testsTemplate = "hack/gen-tf-code/templates/test.tpl"
const docsTemplate = "hack/gen-tf-code/templates/doc.tpl"

func executeTemplate(t reflect.Type, tplFile, folder, file string, funcMaps ...template.FuncMap) {
	tmpl := template.New(tplFile)
	for _, funcMap := range funcMaps {
		tmpl = tmpl.Funcs(funcMap)
	}
	if tmpl, err := tmpl.Parse(readFile(tplFile)); err != nil {
		panic(err)
	} else {
		if err := os.MkdirAll(folder, 0755); err != nil {
			panic(fmt.Sprintf("Failed to create directories for %s", folder))
		}
		f, err := os.Create(path.Join(folder, file))
		if err != nil {
			panic(fmt.Sprintf("Failed to create file %s", path.Join(folder, file)))
		}
		defer f.Close()
		if err := tmpl.Execute(f, t); err != nil {
			panic(err)
		}
	}
}
