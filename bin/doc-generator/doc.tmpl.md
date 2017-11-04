## Connectors

A connector can be also a connector intercepter. An interceptor work like a http middleware. 
This permit to intercept data which will be given back by gautocloud and modified it before giving back to user.

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

This connector is also connector intercepter, it use interceptor [schema](https://godoc.org/github.com/cloudfoundry-community/gautocloud/interceptor#NewSchema).

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

// this show how to intercept data which will be injected to modify it.
// Here it will get interface found by gautocloud and add `intercepted`, after calling Inject, struct receive will have this modification.
//  func (s *MySchema) Intercept(found interface{}) error{
//      f := found.(MySchema)
//      s.MyData = f.MyData + " intercepted"
//      return nil
//  }

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
        err = gautocloud.InjectFromId("id-my-connector", &svc)
        // or
        data, err := gautocloud.GetFirst("id-my-connector")
        svc = data.(MySchema)
        // ----------------------
        // as slice of elements
        var svcSlice []MySchema
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("id-my-connector", &svcSlice)
        // or
        data, err := gautocloud.GetAll("id-my-connector")
        svcSlice = make([]MySchema,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(MySchema))
        }
}
```

#### Config

This is a schema based connectors but `id`, `name` and `tags` are already set (can be registered multiple times).

This connector is also connector intercepter, it use 2 interceptors:
- [schema](https://godoc.org/github.com/cloudfoundry-community/gautocloud/interceptor#NewSchema)
- [overwrite](https://godoc.org/github.com/cloudfoundry-community/gautocloud/interceptor#NewOverwrite) 
(this will be use if struc from user does not implement [SchemaIntercepter](https://godoc.org/github.com/cloudfoundry-community/gautocloud/interceptor#SchemaIntercepter))

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
        // you can set values before inject:
        //  svc.ConfigParam = "my data"
        // this is handle by overwrite interceptor
        err = gautocloud.Inject(&svc)
        // ----------------------
        // as slice of elements
        var svcSlice []MyConfig
        err = gautocloud.Inject(&svcSlice)
}
```

{{template "doc-connector.tmpl.md" .DocsConnector}}
