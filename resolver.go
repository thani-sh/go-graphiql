package graphiql

import (
	"io"
	"text/template"
)

const tFile = `/*
 * DO NOT EDIT
 * CODE GENERATED AUTOMATICALLY WITH "go run gen.go"
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package {{.Name}}
{{$t:=.T}}
import ({{range .Imports}}
	"{{.}}"{{end}}
)

type T{{.Name}} struct { {{range .Fields}}
	{{index . 0}} {{index . 1}} {{$t}}json:"{{index . 2}}"{{$t}}{{end}}
}

type Resolver struct {
	s *T{{.Name}}
}

func (R *Resolver) Set(s *T{{.Name}}) {
	R.s = s
}

{{range .Fields}}
func (R *Resolver) {{index . 0}}() {{index . 1}} {
	return R.s.{{index . 0}}
}
{{end}}

`

type ResolverOptions struct {
	T       string     `yaml:"t"`
	Name    string     `yaml:"name"`
	Imports []string   `yaml:"imports"`
	Fields  [][]string `yaml:"fields"`
}

func GenResolver(w io.Writer, options ResolverOptions) (err error) {
	options.T = "`"
	tpl, err := template.New("Resolver").Parse(tFile)
	if err != nil {
		return
	}

	return tpl.Execute(w, options)
}
