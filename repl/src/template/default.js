const TEMPLATE_DEFAULT = `# {{ .Version }}

{{- define "body" -}}
{{range . -}}
- {{if .Field.Header.Scope }}**{{ unescape .Field.Header.Scope }}**: {{ end }}{{ unescape .Field.Header.Subject }}({{.Short}}) (thanks @{{ unescape .Author.Name }}){{if .Field.Footer }} {{if .Field.Footer.Closes }}, Closes: {{ stringsJoin .Field.Footer.Closes "," }} {{- end }}  {{- end}}
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

{{if .Build}}
### Build system:
{{ template "body" .Build }}
{{- end -}}

{{if .Ci}}
### CI:
{{ template "body" .Ci }}
{{- end -}}

{{if .Chore}}
### Chore:
{{ template "body" .Chore }}
{{- end -}}

{{if .Docs}}
### Documentation:
{{ template "body" .Docs }}
{{- end -}}

{{if .Revert}}
### 🔙 Revert:
{{range .Revert -}}
- {{if .RevertCommitHash }}revert {{ hashURL .RevertCommitHash }}, {{ end }}{{ unescape .Field.Header.Subject }}({{ hashURL .Hash}})
{{ end }}
{{ end }}

{{if .BreakingChanges}}
### BREAKING CHANGES:
{{ range .BreakingChanges -}}
- {{if .Field.Footer.BreakingChange.Title}}{{ unescape .Field.Footer.BreakingChange.Title }}{{ else }}{{ unescape .Field.Title }}{{ end }}
{{ unescape .Field.Footer.BreakingChange.Content }}
{{- end -}}
{{- end}}

### Commits({{ len .Commits }}):
{{range .Commits -}}
- **{{ .Short }}** {{ unescape .Field.Title }}
{{ end }}
`;

export default TEMPLATE_DEFAULT;
