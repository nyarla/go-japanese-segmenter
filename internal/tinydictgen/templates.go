package tinydictgen

import (
	"io"
	"text/template"
)

const codeTemplate = `
{{ define "calc" }}
	if {{ .Var }} == 0x{{ .Rune | printf "%x" }} {
		{{ if (ne .Bias 0) }}n += {{ .Bias }}{{ end }}
		{{ if (ne (len .List) 0) }}
		{{ range $item := .List }}{{ template "calc" $item }}{{ end }}
		{{ end }} }
{{ end }}
// this code is auto-generated. DO NOT EDIT.
package {{ .Package }}

const initialBias = {{ .InitialBias }}

func CalculateBias(p1, p2, p3, r1, r2, r3, r4, r5, r6, t1, t2, t3, t4, t5, t6 rune) int64 {
  n := int64(initialBias);

{{ range $item := .List }}{{ template "calc" $item }}{{ end }}

	return n
}
`

var instance = template.Must(template.New("gen").Parse(codeTemplate))

type TemplateParams struct {
	Package     string
	InitialBias string
	List        []*Item
}

func Render(w io.Writer, pkg string, bias string, json JSONData) error {
	return instance.Execute(w, &TemplateParams{
		Package:     pkg,
		InitialBias: bias,
		List:        json.Items(),
	})
}
