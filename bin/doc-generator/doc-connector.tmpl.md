{{range $doc := .}}
### {{.Name}}

All of these connectors responds on:
- Regex name: `{{ .RespondName }}`
- Regex tags:{{ range $tag := .RespondTags }}
  - `{{ $tag }}`{{ end }}

{{ range $connector := .Connectors }}
#### {{$connector.Name}}

- **Id**: `{{$connector.Id}}`
- **Given type**: `{{$connector.TypeName}}`

**Tip**: You can load all based *{{$connector.GlobalType}} {{.SimpleName}}* by importing: `_ "{{$connector.Tip}}"`

##### Type documentation
The type `{{$connector.TypeName}}` can be found in package: `{{$connector.TypePkg}}`.
{{if $connector.TypeWrapped}}
The type `{{$connector.TypeName}}` is a wrapper on the real package `{{$connector.TypeWrapped}}`, 
you can find doc on real type here: [{{$connector.DocUrl}}]({{$connector.DocUrl}}). 
{{else if $connector.DocUrl}}
You can find documentation related to package `{{$connector.TypePkg}}` here: [{{$connector.DocUrl}}]({{$connector.DocUrl}}).
{{else}}
This type refers to this structure:
```go
type {{$connector.StructGiven.Name}} struct { {{ range $field := $connector.StructGiven.Fields}}
        {{$field.Name}} {{$field.Type}} {{if $field.Comment}}// {{$field.Comment}}{{end}}{{end}}
}
```
{{end}}

##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "{{$connector.Pkg}}"
        "{{$connector.TypePkg}}"
)
func main() {
        var err error
        // As single element
        var svc {{$connector.TypeName}}
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("{{$connector.Id}}", &svc)
        // or
        data, err := gautocloud.GetFirst("{{$connector.Id}}")
        svc = data.({{$connector.TypeName}}){{if .Closeable}}
        defer svc.Close(){{end}}
        // ----------------------
        // as slice of elements
        var svcSlice []{{$connector.TypeName}}
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("{{$connector.Id}}", &svcSlice)
        // or
        data, err := gautocloud.GetAll("{{$connector.Id}}")
        svcSlice = make([]{{$connector.TypeName}},0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.({{$connector.TypeName}}))
        }
}
```
{{ end }}
{{end}}