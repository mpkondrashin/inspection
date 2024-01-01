package main

import (
	"os"
	"text/template"
)

var versionGo = `package main

var MajorMinorRevision = "{{.MajorMinorRevision}}"
`

func main() {
	if len(os.Args) != 3 {
		panic("correct options are file path and version value")
	}
	filePath := os.Args[1]
	ver := os.Args[2]
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	tmpl, err := template.New("version_go").Parse(versionGo)
	if err != nil {
		panic(err)
	}
	data := struct{ MajorMinorRevision string }{ver}
	if err = tmpl.Execute(f, &data); err != nil {
		panic(err)
	}
}
