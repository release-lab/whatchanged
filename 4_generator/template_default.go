package generator

const defaultTemplate = `# {{ .Version }}

{{- define "body" -}}
{{range . -}}
- {{if .Field.Header.Scope }}**{{ .Field.Header.Scope }}**: {{ end }}{{ .Field.Header.Subject }}({{.Short}}) (thanks @{{ .Author.Name }}){{if .Field.Footer }} {{if .Field.Footer.Closes }}, Closes: {{ StringsJoin .Field.Footer.Closes "," }} {{- end }}  {{- end}}
{{ end }}
{{- end -}}
{{if .Feat}}
### New feature:
{{ template "body" .Feat }}
{{- end -}}
{{if .Fix}}
### Bugs fixed:
{{ template "body" .Fix }}
{{- end -}}
{{if .Refactor}}
### Code Refactoring:
{{ template "body" .Refactor }}
{{- end -}}
{{if .Test}}
### Testing:
{{ template "body" .Test }}
{{- end -}}
{{if .Perf}}
### Performance improves:
{{ template "body" .Perf }}
{{- end -}}
{{if .Docs}}
### Documentation:
{{ template "body" .Docs }}
{{- end -}}
{{if .BreakingChanges}}
### BREAKING CHANGES:
{{ range .BreakingChanges -}}
- {{if .Field.Footer.BreakingChange.Title}}{{ .Field.Footer.BreakingChange.Title }}{{ else }}{{ .Field.Title }}{{ end }}
{{ .Field.Footer.BreakingChange.Content }}
{{- end -}}
{{- end}}

### Commits({{ len .Commits }}):
{{range .Commits -}}
- **{{ .Short }}** {{ .Field.Title }}
{{ end }}`
