package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
	"text/template"
)

type itemValue struct {
	VarName, Value, Text string
}

type templateArgs struct {
	Map map[string][]itemValue
}

const tplt = `
	{{ range $typeName, $items := .Map }}
	export const {{ $typeName }}Fields = {
		{{- range $items -}}
			{{ .VarName }}: {{ .Value}},
		{{ end -}}
	};
  	export const {{ $typeName }}s = [
		  {{- range $items -}}
		 	 { value: {{ $typeName }}Fields.{{ .VarName }}, text: "{{ .Text }}" },
		  {{ end -}}
	  ];

	{{ end }}
`

var tp = template.Must(template.New("enums").Parse(tplt))

func main() {
	enums := parse()
	if err := tp.Execute(os.Stdout, templateArgs{Map: enums}); err != nil {
		log.Fatal(err)
	}
}

func parse() map[string][]itemValue {
	t := token.NewFileSet()
	f, err := parser.ParseFile(t, "server/models/enums.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	enums := map[string][]itemValue{} // type -> datas
	for _, decl := range f.Decls {
		stm, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range stm.Specs {
			s, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}
			if s.Comment == nil {
				log.Fatal("value as comment expected")
			}
			text := strings.TrimSpace(s.Comment.Text())
			varName := s.Names[0].String()
			typeName := s.Type.(*ast.Ident).String()
			value := s.Values[0].(*ast.BasicLit).Value
			enums[typeName] = append(enums[typeName], itemValue{VarName: varName, Value: value, Text: text})
		}
	}
	return enums
}
