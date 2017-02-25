## Connectors

**Tip**: To load all default connectors import: `_ "github.com/cloudfoundry-community/gautocloud/connectors/all"`

- [Generic](#generic)
  - [Schema based](#schema-based)
  - [Config](#config)
{{ range $conSummary := .Summaries }}- [{{.Name}}](#{{.Slug}}){{range $subSummary := .Sub}}
  - [{{.Name}}](#{{.Slug}}){{end}}
{{end}}

### Generic

These connectors are specials connectors that final users need to register manually when needed.

One of usecase is to be able to retrieve configuration from services or simply add your own connector easily.

#### Schema based

Add a straight forward connector which give back schema fed by loader.

##### Example

```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        "github.com/cloudfoundry-community/gautocloud/connectors/generic"
)

type MySchema struct {
        MyData string
}

func init() {
        gautocloud.RegisterConnector(generic.NewSchemaBasedGenericConnector(
        "id-my-connector",
        ".*my_connector.*",
        []string{"my_connector.*"},
        MySchema{},
        ))
}
func main() {
        var err error
        // As single element
        var svc MySchema
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("{{$connector.Id}}", &svc)
        // or
        data, err := gautocloud.GetFirst("{{$connector.Id}}")
        svc = data.(MySchema)
        // ----------------------
        // as slice of elements
        var svcSlice []MySchema
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("{{$connector.Id}}", &svcSlice)
        // or
        data, err := gautocloud.GetAll("{{$connector.Id}}")
        svcSlice = make([]MySchema,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(MySchema))
        }
}
```

#### Config

This is a schema based connectors but `id`, `name` and `tags` are already set (can be registered multiple times).

This generic connector responds on:
- Regex name: `.*config.*`
- Regex tags:
  - `config.*`

##### Example

```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        "github.com/cloudfoundry-community/gautocloud/connectors/generic"
)

type MyConfig struct {
        ConfigParam string
}

type MySecondConfig struct {
        ConfigParam string
}

func init() {
        gautocloud.RegisterConnector(generic.NewConfigGenericConnector(MyConfig{}))
        gautocloud.RegisterConnector(generic.NewConfigGenericConnector(MySecondConfig{}))
}
func main() {
        var err error
        // As single element
        var svc MyConfig
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("{{$connector.Id}}", &svc)
        // or
        data, err := gautocloud.GetFirst("{{$connector.Id}}")
        svc = data.(MyConfig)
        // ----------------------
        // as slice of elements
        var svcSlice []MyConfig
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("{{$connector.Id}}", &svcSlice)
        // or
        data, err := gautocloud.GetAll("{{$connector.Id}}")
        svcSlice = make([]MyConfig,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(MyConfig))
        }
}
```

{{template "doc-connector.tmpl.md" .DocsConnector}}
