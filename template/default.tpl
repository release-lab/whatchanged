## {{ .Version }} ({{ .Date }})

{{- define "body" -}}
{{range . -}}
- {{if .Field.Header.Scope }}**{{ .Field.Header.Scope | unescape }}**: {{ end }}{{ .Field.Header.Subject | unescape }}({{ .HashURL }}) (by {{ .Author.Name | unescape }}){{if .Closes }}, Closes: {{ range $index, $element := .Closes}}{{if $index}},{{end}}{{$element}}{{- end -}} {{- end }}
{{ end }}
{{- end -}}

{{if .Feat}}
### New feature:
{{ template "body" .Feat }}
{{ end }}

{{if .Fix}}
### Bugs fixed:
{{ template "body" .Fix }}
{{ end }}

{{if .Perf}}
### Performance improves:
{{ template "body" .Perf }}
{{ end }}

{{if .Revert}}
### Revert:
{{range .Revert -}}
- {{if .RevertCommitHash }}revert {{ .RevertCommitHashURL }}, {{ end }}{{ .Field.Header.Subject | unescape }}({{ .HashURL }})
{{ end }}
{{ end }}

{{if .BreakingChanges}}
### BREAKING CHANGES:
{{ range .BreakingChanges -}}

- {{if .Field.Footer.BreakingChange.Title}}{{ .Field.Footer.BreakingChange.Title | unescape }}{{ else }}{{ .Field.Title | unescape }}{{ end }}

{{ .Field.Footer.BreakingChange.Content | indent 2 | unescape }}

{{ end -}}
{{ end }}
