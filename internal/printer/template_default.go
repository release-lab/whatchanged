package printer

const defaultTemplate = `{{ .Version }}{{if .Feat}}
New feature:
{{ range .Feat }}
- {{if .Field.Header.Scope }}{{ .Field.Header.Scope }}: {{ end }}{{ .Field.Header.Subject }} (thanks @{{ .Author.Name }}){{if .Field.Footer }} {{if .Field.Footer.Closes }}, Closes: {{ StringsJoin .Field.Footer.Closes "," }} {{- end }}  {{- end }} {{- end }}{{- end}}
{{if .Fix}}
Bugs fixed:
{{ range .Fix }}
- {{if .Field.Header.Scope }}{{ .Field.Header.Scope }}: {{ end }}{{ .Field.Header.Subject }} (thanks @{{ .Author.Name }}){{if .Field.Footer }} {{if .Field.Footer.Closes }}, Closes: {{ StringsJoin .Field.Footer.Closes "," }} {{- end }}  {{- end }} {{- end }}{{- end }}
{{if .Refactor}}
Code Refactoring:
{{ range .Refactor }}
- {{if .Field.Header.Scope }}{{ .Field.Header.Scope }}: {{ end }}{{ .Field.Header.Subject }} (thanks @{{ .Author.Name }}){{if .Field.Footer }} {{if .Field.Footer.Closes }}, Closes: {{ StringsJoin .Field.Footer.Closes "," }} {{- end }}  {{- end }} {{- end }}{{- end }}
{{if .Test}}
Test:
{{ range .Test }}
- {{if .Field.Header.Scope }}{{ .Field.Header.Scope }}: {{ end }}{{ .Field.Header.Subject }} (thanks @{{ .Author.Name }}){{if .Field.Footer }} {{if .Field.Footer.Closes }}, Closes: {{ StringsJoin .Field.Footer.Closes "," }} {{- end }}  {{- end }} {{- end }}{{- end }}
{{if .Perf}}
Performance improves:
{{ range .Perf }}
- {{if .Field.Header.Scope }}{{ .Field.Header.Scope }}: {{ end }}{{ .Field.Header.Subject }} (thanks @{{ .Author.Name }}){{if .Field.Footer }} {{if .Field.Footer.Closes }}, Closes: {{ StringsJoin .Field.Footer.Closes "," }} {{- end }}  {{- end }} {{- end }}{{- end }}
{{if .Docs}}
Documentation:
{{ range .Docs }}
- {{if .Field.Header.Scope }}{{ .Field.Header.Scope }}: {{ end }}{{ .Field.Header.Subject }} (thanks @{{ .Author.Name }}){{if .Field.Footer }} {{if .Field.Footer.Closes }}, Closes: {{ StringsJoin .Field.Footer.Closes "," }} {{- end }}  {{- end }} {{- end }}{{- end }}
{{if .BreakingChanges}}
BREAKING CHANGES:
{{ range .BreakingChanges }}
- {{ .Field.Footer.BreakingChange.Title }} {{ .Field.Title }}
{{ .Field.Footer.BreakingChange.Content }}{{- end }}{{- end }}
Commits({{ len .Commits }}):{{ range .Commits }}
{{ .Short }} {{ .Field.Title }}
{{- end }}`
