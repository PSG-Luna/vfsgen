package main

const generateTemplateText = `package main

import (
	"log"

	"github.com/shurcooL/vfsgen"

	{{.ImportPath | quote}}
)

func main() {
	err := vfsgen.Generate({{.PackageName}}.{{.VariableName}}, vfsgen.Options{
		PackageName:  {{.PackageName | quote}},{{with .BuildTags}}
		BuildTags:    {{. | quote}},{{end}}
		VariableName: {{.VariableName | quote}},
	})
	if err != nil {
		log.Fatalln(err)
	}
}
`
