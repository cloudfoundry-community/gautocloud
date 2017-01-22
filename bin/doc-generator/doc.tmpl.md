## Connectors

**Tip**: To load all default connectors import: `_ "github.com/cloudfoundry-community/gautocloud/connectors/all"`

{{ range $conSummary := .Summaries }}- [{{.Name}}](#{{.Slug}}){{range $subSummary := .Sub}}
  - [{{.Name}}](#{{.Slug}}){{end}}
{{end}}
{{template "doc-connector.tmpl.md" .DocsConnector}}
