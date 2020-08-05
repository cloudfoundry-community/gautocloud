## Connectors

A connector can be also a connector intercepter. An interceptor work like a http middleware. 
This permit to intercept data which will be given back by gautocloud and modified it before giving back to user.

**Tip**: To load all default connectors import: `_ "github.com/cloudfoundry-community/gautocloud/connectors/all"`

- [Generic](#generic)
  - [Schema based](#schema-based)
  - [Config](#config)
- [Amqp](#amqp)
  - [Amqp - Client](#amqp---client)
  - [Amqp - Raw](#amqp---raw)
- [Mongodb](#mongodb)
  - [Mongodb - Client](#mongodb---client)
  - [Mongodb - Raw](#mongodb---raw)
- [Mssql](#mssql)
  - [Mssql - Client](#mssql---client)
  - [Mssql - Gorm](#mssql---gorm)
  - [Mssql - Raw](#mssql---raw)
- [Mysql](#mysql)
  - [Mysql - Client](#mysql---client)
  - [Mysql - Gorm](#mysql---gorm)
  - [Mysql - Raw](#mysql---raw)
- [Oauth2](#oauth2)
  - [Oauth2 - Config](#oauth2---config)
  - [Oauth2 - Raw](#oauth2---raw)
- [Oracle](#oracle)
  - [Oracle - Raw](#oracle---raw)
- [Postgresql](#postgresql)
  - [Postgresql - Client](#postgresql---client)
  - [Postgresql - Gorm](#postgresql---gorm)
  - [Postgresql - Raw](#postgresql---raw)
- [Redis](#redis)
  - [Redis - Client](#redis---client)
  - [Redis - Raw](#redis---raw)
- [S3](#s3)
  - [S3 - Goamz](#s3---goamz)
  - [S3 - Minio](#s3---minio)
  - [S3 - Raw](#s3---raw)
- [Smtp](#smtp)
  - [Smtp - Client](#smtp---client)
  - [Smtp - Raw](#smtp---raw)

### Generic

These connectors are specials connectors that final users need to register manually when needed.

One of usecase is to be able to retrieve configuration from services or simply add your own connector easily.

#### Schema based

Add a straight forward connector which give back schema fed by loader.

This connector is also connector intercepter, it use interceptor [schema](https://godoc.org/github.com/cloudfoundry-community/gautocloud/interceptor#NewSchema) 
as default interceptor. 

You can also set your interceptors to connector, they will be used only if struct does not implement interface [SchemaIntercepter](https://godoc.org/github.com/cloudfoundry-community/gautocloud/interceptor#SchemaIntercepter).

You can find interceptors already created at https://godoc.org/github.com/cloudfoundry-community/gautocloud/interceptor .

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
        // MyFirstInterceptor,
        // MySecondInterceptor,
        // ....
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

This connector is a connector intercepter, it use the default interceptor [overwrite](https://godoc.org/github.com/cloudfoundry-community/gautocloud/interceptor#NewOverwrite)  
which only be used if struct does not implement interface [SchemaIntercepter](https://godoc.org/github.com/cloudfoundry-community/gautocloud/interceptor#SchemaIntercepter) 
and if no interceptors has been given to connector.

You can find interceptors already created at https://godoc.org/github.com/cloudfoundry-community/gautocloud/interceptor .

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
        gautocloud.RegisterConnector(
            generic.NewConfigGenericConnector(
            MySecondConfig{},
            // MyFirstInterceptor,
            // MySecondInterceptor,
            // ....
            ),
        )
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

### Amqp

All of these connectors responds on:
- Regex name: `.*amqp.*`
- Regex tags:
  - `amqp`
  - `rabbitmq`


#### Amqp - Client

- **Id**: `amqp`
- **Given type**: `*amqp.Connection`

**Tip**: You can load all based *Amqp Client* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/amqp"`

##### Type documentation
The type `*amqp.Connection` can be found in package: `github.com/streadway/amqp`.

You can find documentation related to package `github.com/streadway/amqp` here: [https://github.com/streadway/amqp](https://github.com/streadway/amqp).


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/amqp/client"
        "github.com/streadway/amqp"
)
func main() {
        var err error
        // As single element
        var svc *amqp.Connection
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("amqp", &svc)
        // or
        data, err := gautocloud.GetFirst("amqp")
        svc = data.(*amqp.Connection)
        defer svc.Close()
        // ----------------------
        // as slice of elements
        var svcSlice []*amqp.Connection
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("amqp", &svcSlice)
        // or
        data, err := gautocloud.GetAll("amqp")
        svcSlice = make([]*amqp.Connection,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(*amqp.Connection))
        }
}
```

#### Amqp - Raw

- **Id**: `raw:amqp`
- **Given type**: `amqptype.Amqp`

**Tip**: You can load all based *Amqp Raw* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/amqp"`

##### Type documentation
The type `amqptype.Amqp` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/amqp/amqptype`.

This type refers to this structure:
```go
type Amqp struct { 
        User string 
        Password string 
        Host string 
        Port int 
}
```


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/amqp"
        "github.com/cloudfoundry-community/gautocloud/connectors/amqp/amqptype"
)
func main() {
        var err error
        // As single element
        var svc amqptype.Amqp
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("raw:amqp", &svc)
        // or
        data, err := gautocloud.GetFirst("raw:amqp")
        svc = data.(amqptype.Amqp)
        // ----------------------
        // as slice of elements
        var svcSlice []amqptype.Amqp
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("raw:amqp", &svcSlice)
        // or
        data, err := gautocloud.GetAll("raw:amqp")
        svcSlice = make([]amqptype.Amqp,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(amqptype.Amqp))
        }
}
```


### Mongodb

All of these connectors responds on:
- Regex name: `.*mongo.*`
- Regex tags:
  - `mongo.*`


#### Mongodb - Client

- **Id**: `mongodb`
- **Given type**: `*mgo.Session`

**Tip**: You can load all based *Databases Client* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/databases/client"`

##### Type documentation
The type `*mgo.Session` can be found in package: `gopkg.in/mgo.v2`.

You can find documentation related to package `gopkg.in/mgo.v2` here: [https://gopkg.in/mgo.v2](https://gopkg.in/mgo.v2).


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/databases/client/mongodb"
        "gopkg.in/mgo.v2"
)
func main() {
        var err error
        // As single element
        var svc *mgo.Session
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("mongodb", &svc)
        // or
        data, err := gautocloud.GetFirst("mongodb")
        svc = data.(*mgo.Session)
        defer svc.Close()
        // ----------------------
        // as slice of elements
        var svcSlice []*mgo.Session
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("mongodb", &svcSlice)
        // or
        data, err := gautocloud.GetAll("mongodb")
        svcSlice = make([]*mgo.Session,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(*mgo.Session))
        }
}
```

#### Mongodb - Raw

- **Id**: `raw:mongodb`
- **Given type**: `dbtype.MongodbDatabase`

**Tip**: You can load all based *Databases Raw* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/databases"`

##### Type documentation
The type `dbtype.MongodbDatabase` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype`.

This type refers to this structure:
```go
type MongodbDatabase struct { 
        User string 
        Password string 
        Host string 
        Port int 
        Database string 
        Options string 
}
```


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/databases"
        "github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)
func main() {
        var err error
        // As single element
        var svc dbtype.MongodbDatabase
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("raw:mongodb", &svc)
        // or
        data, err := gautocloud.GetFirst("raw:mongodb")
        svc = data.(dbtype.MongodbDatabase)
        // ----------------------
        // as slice of elements
        var svcSlice []dbtype.MongodbDatabase
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("raw:mongodb", &svcSlice)
        // or
        data, err := gautocloud.GetAll("raw:mongodb")
        svcSlice = make([]dbtype.MongodbDatabase,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(dbtype.MongodbDatabase))
        }
}
```


### Mssql

All of these connectors responds on:
- Regex name: `.*mssql.*`
- Regex tags:
  - `mssql.*`
  - `sqlserver`


#### Mssql - Client

- **Id**: `mssql`
- **Given type**: `*dbtype.MssqlDB`

**Tip**: You can load all based *Databases Client* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/databases/client"`

##### Type documentation
The type `*dbtype.MssqlDB` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype`.

The type `*dbtype.MssqlDB` is a wrapper on the real package `*sql.DB`, 
you can find doc on real type here: [https://golang.org/pkg/database/sql](https://golang.org/pkg/database/sql). 


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/databases/client/mssql"
        "github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)
func main() {
        var err error
        // As single element
        var svc *dbtype.MssqlDB
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("mssql", &svc)
        // or
        data, err := gautocloud.GetFirst("mssql")
        svc = data.(*dbtype.MssqlDB)
        defer svc.Close()
        // ----------------------
        // as slice of elements
        var svcSlice []*dbtype.MssqlDB
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("mssql", &svcSlice)
        // or
        data, err := gautocloud.GetAll("mssql")
        svcSlice = make([]*dbtype.MssqlDB,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(*dbtype.MssqlDB))
        }
}
```

#### Mssql - Gorm

- **Id**: `gorm:mssql`
- **Given type**: `*gorm.DB`

**Tip**: You can load all based *Databases Gorm* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/databases/gorm"`

##### Type documentation
The type `*gorm.DB` can be found in package: `github.com/jinzhu/gorm`.

You can find documentation related to package `github.com/jinzhu/gorm` here: [https://github.com/jinzhu/gorm](https://github.com/jinzhu/gorm).


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/databases/gorm/mssql"
        "github.com/jinzhu/gorm"
)
func main() {
        var err error
        // As single element
        var svc *gorm.DB
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("gorm:mssql", &svc)
        // or
        data, err := gautocloud.GetFirst("gorm:mssql")
        svc = data.(*gorm.DB)
        defer svc.Close()
        // ----------------------
        // as slice of elements
        var svcSlice []*gorm.DB
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("gorm:mssql", &svcSlice)
        // or
        data, err := gautocloud.GetAll("gorm:mssql")
        svcSlice = make([]*gorm.DB,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(*gorm.DB))
        }
}
```

#### Mssql - Raw

- **Id**: `raw:mssql`
- **Given type**: `dbtype.MssqlDatabase`

**Tip**: You can load all based *Databases Raw* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/databases"`

##### Type documentation
The type `dbtype.MssqlDatabase` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype`.

This type refers to this structure:
```go
type MssqlDatabase struct { 
        User string 
        Password string 
        Host string 
        Port int 
        Database string 
        Options string 
}
```


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/databases"
        "github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)
func main() {
        var err error
        // As single element
        var svc dbtype.MssqlDatabase
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("raw:mssql", &svc)
        // or
        data, err := gautocloud.GetFirst("raw:mssql")
        svc = data.(dbtype.MssqlDatabase)
        // ----------------------
        // as slice of elements
        var svcSlice []dbtype.MssqlDatabase
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("raw:mssql", &svcSlice)
        // or
        data, err := gautocloud.GetAll("raw:mssql")
        svcSlice = make([]dbtype.MssqlDatabase,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(dbtype.MssqlDatabase))
        }
}
```


### Mysql

All of these connectors responds on:
- Regex name: `.*(mysql|maria).*`
- Regex tags:
  - `mysql`
  - `maria.*`


#### Mysql - Client

- **Id**: `mysql`
- **Given type**: `*dbtype.MysqlDB`

**Tip**: You can load all based *Databases Client* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/databases/client"`

##### Type documentation
The type `*dbtype.MysqlDB` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype`.

The type `*dbtype.MysqlDB` is a wrapper on the real package `*sql.DB`, 
you can find doc on real type here: [https://golang.org/pkg/database/sql](https://golang.org/pkg/database/sql). 


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/databases/client/mysql"
        "github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)
func main() {
        var err error
        // As single element
        var svc *dbtype.MysqlDB
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("mysql", &svc)
        // or
        data, err := gautocloud.GetFirst("mysql")
        svc = data.(*dbtype.MysqlDB)
        defer svc.Close()
        // ----------------------
        // as slice of elements
        var svcSlice []*dbtype.MysqlDB
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("mysql", &svcSlice)
        // or
        data, err := gautocloud.GetAll("mysql")
        svcSlice = make([]*dbtype.MysqlDB,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(*dbtype.MysqlDB))
        }
}
```

#### Mysql - Gorm

- **Id**: `gorm:mysql`
- **Given type**: `*gorm.DB`

**Tip**: You can load all based *Databases Gorm* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/databases/gorm"`

##### Type documentation
The type `*gorm.DB` can be found in package: `github.com/jinzhu/gorm`.

You can find documentation related to package `github.com/jinzhu/gorm` here: [https://github.com/jinzhu/gorm](https://github.com/jinzhu/gorm).


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/databases/gorm/mysql"
        "github.com/jinzhu/gorm"
)
func main() {
        var err error
        // As single element
        var svc *gorm.DB
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("gorm:mysql", &svc)
        // or
        data, err := gautocloud.GetFirst("gorm:mysql")
        svc = data.(*gorm.DB)
        defer svc.Close()
        // ----------------------
        // as slice of elements
        var svcSlice []*gorm.DB
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("gorm:mysql", &svcSlice)
        // or
        data, err := gautocloud.GetAll("gorm:mysql")
        svcSlice = make([]*gorm.DB,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(*gorm.DB))
        }
}
```

#### Mysql - Raw

- **Id**: `raw:mysql`
- **Given type**: `dbtype.MysqlDatabase`

**Tip**: You can load all based *Databases Raw* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/databases"`

##### Type documentation
The type `dbtype.MysqlDatabase` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype`.

This type refers to this structure:
```go
type MysqlDatabase struct { 
        User string 
        Password string 
        Host string 
        Port int 
        Database string 
        Options string 
}
```


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/databases"
        "github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)
func main() {
        var err error
        // As single element
        var svc dbtype.MysqlDatabase
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("raw:mysql", &svc)
        // or
        data, err := gautocloud.GetFirst("raw:mysql")
        svc = data.(dbtype.MysqlDatabase)
        // ----------------------
        // as slice of elements
        var svcSlice []dbtype.MysqlDatabase
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("raw:mysql", &svcSlice)
        // or
        data, err := gautocloud.GetAll("raw:mysql")
        svcSlice = make([]dbtype.MysqlDatabase,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(dbtype.MysqlDatabase))
        }
}
```


### Oauth2

All of these connectors responds on:
- Regex name: `.*oauth.*`
- Regex tags:
  - `oauth.*`
  - `sso`


#### Oauth2 - Config

- **Id**: `config:oauth2`
- **Given type**: `*oauth2.Config`

**Tip**: You can load all based *Auth Config* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/auth/config"`

##### Type documentation
The type `*oauth2.Config` can be found in package: `golang.org/x/oauth2`.

You can find documentation related to package `golang.org/x/oauth2` here: [https://golang.org/x/oauth2](https://golang.org/x/oauth2).


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/auth/config/oauth2"
        "golang.org/x/oauth2"
)
func main() {
        var err error
        // As single element
        var svc *oauth2.Config
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("config:oauth2", &svc)
        // or
        data, err := gautocloud.GetFirst("config:oauth2")
        svc = data.(*oauth2.Config)
        // ----------------------
        // as slice of elements
        var svcSlice []*oauth2.Config
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("config:oauth2", &svcSlice)
        // or
        data, err := gautocloud.GetAll("config:oauth2")
        svcSlice = make([]*oauth2.Config,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(*oauth2.Config))
        }
}
```

#### Oauth2 - Raw

- **Id**: `raw:oauth2`
- **Given type**: `schema.Oauth2Schema`

**Tip**: You can load all based *Auth Raw* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/auth"`

##### Type documentation
The type `schema.Oauth2Schema` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/auth/schema`.

This type refers to this structure:
```go
type Oauth2Schema struct { 
        AuthorizationUri string 
        UserInfoUri string 
        TokenUri string 
        ClientId string 
        ClientSecret string 
        GrantTypes []string 
        Scopes []string 
}
```


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/auth"
        "github.com/cloudfoundry-community/gautocloud/connectors/auth/schema"
)
func main() {
        var err error
        // As single element
        var svc schema.Oauth2Schema
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("raw:oauth2", &svc)
        // or
        data, err := gautocloud.GetFirst("raw:oauth2")
        svc = data.(schema.Oauth2Schema)
        // ----------------------
        // as slice of elements
        var svcSlice []schema.Oauth2Schema
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("raw:oauth2", &svcSlice)
        // or
        data, err := gautocloud.GetAll("raw:oauth2")
        svcSlice = make([]schema.Oauth2Schema,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(schema.Oauth2Schema))
        }
}
```


### Oracle

All of these connectors responds on:
- Regex name: `.*oracle.*`
- Regex tags:
  - `oracle`
  - `oci.*`


#### Oracle - Raw

- **Id**: `raw:oracle`
- **Given type**: `dbtype.OracleDatabase`

**Tip**: You can load all based *Databases Raw* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/databases"`

##### Type documentation
The type `dbtype.OracleDatabase` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype`.

This type refers to this structure:
```go
type OracleDatabase struct { 
        User string 
        Password string 
        Host string 
        Port int 
        Database string 
        Options string 
}
```


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/databases"
        "github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)
func main() {
        var err error
        // As single element
        var svc dbtype.OracleDatabase
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("raw:oracle", &svc)
        // or
        data, err := gautocloud.GetFirst("raw:oracle")
        svc = data.(dbtype.OracleDatabase)
        // ----------------------
        // as slice of elements
        var svcSlice []dbtype.OracleDatabase
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("raw:oracle", &svcSlice)
        // or
        data, err := gautocloud.GetAll("raw:oracle")
        svcSlice = make([]dbtype.OracleDatabase,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(dbtype.OracleDatabase))
        }
}
```


### Postgresql

All of these connectors responds on:
- Regex name: `.*postgres.*`
- Regex tags:
  - `postgres.*`


#### Postgresql - Client

- **Id**: `postgresql`
- **Given type**: `*dbtype.PostgresqlDB`

**Tip**: You can load all based *Databases Client* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/databases/client"`

##### Type documentation
The type `*dbtype.PostgresqlDB` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype`.

The type `*dbtype.PostgresqlDB` is a wrapper on the real package `*sql.DB`, 
you can find doc on real type here: [https://golang.org/pkg/database/sql](https://golang.org/pkg/database/sql). 


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/databases/client/postgresql"
        "github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)
func main() {
        var err error
        // As single element
        var svc *dbtype.PostgresqlDB
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("postgresql", &svc)
        // or
        data, err := gautocloud.GetFirst("postgresql")
        svc = data.(*dbtype.PostgresqlDB)
        defer svc.Close()
        // ----------------------
        // as slice of elements
        var svcSlice []*dbtype.PostgresqlDB
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("postgresql", &svcSlice)
        // or
        data, err := gautocloud.GetAll("postgresql")
        svcSlice = make([]*dbtype.PostgresqlDB,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(*dbtype.PostgresqlDB))
        }
}
```

#### Postgresql - Gorm

- **Id**: `gorm:postgresql`
- **Given type**: `*gorm.DB`

**Tip**: You can load all based *Databases Gorm* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/databases/gorm"`

##### Type documentation
The type `*gorm.DB` can be found in package: `github.com/jinzhu/gorm`.

You can find documentation related to package `github.com/jinzhu/gorm` here: [https://github.com/jinzhu/gorm](https://github.com/jinzhu/gorm).


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/databases/gorm/postgresql"
        "github.com/jinzhu/gorm"
)
func main() {
        var err error
        // As single element
        var svc *gorm.DB
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("gorm:postgresql", &svc)
        // or
        data, err := gautocloud.GetFirst("gorm:postgresql")
        svc = data.(*gorm.DB)
        defer svc.Close()
        // ----------------------
        // as slice of elements
        var svcSlice []*gorm.DB
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("gorm:postgresql", &svcSlice)
        // or
        data, err := gautocloud.GetAll("gorm:postgresql")
        svcSlice = make([]*gorm.DB,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(*gorm.DB))
        }
}
```

#### Postgresql - Raw

- **Id**: `raw:postgresql`
- **Given type**: `dbtype.PostgresqlDatabase`

**Tip**: You can load all based *Databases Raw* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/databases"`

##### Type documentation
The type `dbtype.PostgresqlDatabase` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype`.

This type refers to this structure:
```go
type PostgresqlDatabase struct { 
        User string 
        Password string 
        Host string 
        Port int 
        Database string 
        Options string 
}
```


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/databases"
        "github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)
func main() {
        var err error
        // As single element
        var svc dbtype.PostgresqlDatabase
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("raw:postgresql", &svc)
        // or
        data, err := gautocloud.GetFirst("raw:postgresql")
        svc = data.(dbtype.PostgresqlDatabase)
        // ----------------------
        // as slice of elements
        var svcSlice []dbtype.PostgresqlDatabase
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("raw:postgresql", &svcSlice)
        // or
        data, err := gautocloud.GetAll("raw:postgresql")
        svcSlice = make([]dbtype.PostgresqlDatabase,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(dbtype.PostgresqlDatabase))
        }
}
```


### Redis

All of these connectors responds on:
- Regex name: `.*redis.*`
- Regex tags:
  - `redis`


#### Redis - Client

- **Id**: `redis`
- **Given type**: `*redis.Client`

**Tip**: You can load all based *Databases Client* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/databases/client"`

##### Type documentation
The type `*redis.Client` can be found in package: `github.com/go-redis/redis/v7`.

You can find documentation related to package `github.com/go-redis/redis/v7` here: [https://pkg.go.dev/github.com/go-redis/redis?tab=doc](https://pkg.go.dev/github.com/go-redis/redis?tab=doc).


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/databases/client/redis"
        "github.com/go-redis/redis/v7"
)
func main() {
        var err error
        // As single element
        var svc *redis.Client
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("redis", &svc)
        // or
        data, err := gautocloud.GetFirst("redis")
        svc = data.(*redis.Client)
        defer svc.Close()
        // ----------------------
        // as slice of elements
        var svcSlice []*redis.Client
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("redis", &svcSlice)
        // or
        data, err := gautocloud.GetAll("redis")
        svcSlice = make([]*redis.Client,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(*redis.Client))
        }
}
```

#### Redis - Raw

- **Id**: `raw:redis`
- **Given type**: `dbtype.RedisDatabase`

**Tip**: You can load all based *Databases Raw* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/databases"`

##### Type documentation
The type `dbtype.RedisDatabase` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype`.

This type refers to this structure:
```go
type RedisDatabase struct { 
        Password string 
        Host string 
        Port int 
}
```


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/databases"
        "github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)
func main() {
        var err error
        // As single element
        var svc dbtype.RedisDatabase
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("raw:redis", &svc)
        // or
        data, err := gautocloud.GetFirst("raw:redis")
        svc = data.(dbtype.RedisDatabase)
        // ----------------------
        // as slice of elements
        var svcSlice []dbtype.RedisDatabase
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("raw:redis", &svcSlice)
        // or
        data, err := gautocloud.GetAll("raw:redis")
        svcSlice = make([]dbtype.RedisDatabase,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(dbtype.RedisDatabase))
        }
}
```


### S3

All of these connectors responds on:
- Regex name: `.*s3.*`
- Regex tags:
  - `s3`
  - `riak.*`


#### S3 - Goamz

- **Id**: `s3`
- **Given type**: `*s3.Bucket`

**Tip**: You can load all based *Objstorage Goamz* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/objstorage/client/s3"`

##### Type documentation
The type `*s3.Bucket` can be found in package: `github.com/goamz/goamz/s3`.

You can find documentation related to package `github.com/goamz/goamz/s3` here: [https://github.com/goamz/goamz](https://github.com/goamz/goamz).


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/objstorage/client/s3/goamz"
        "github.com/goamz/goamz/s3"
)
func main() {
        var err error
        // As single element
        var svc *s3.Bucket
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("s3", &svc)
        // or
        data, err := gautocloud.GetFirst("s3")
        svc = data.(*s3.Bucket)
        // ----------------------
        // as slice of elements
        var svcSlice []*s3.Bucket
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("s3", &svcSlice)
        // or
        data, err := gautocloud.GetAll("s3")
        svcSlice = make([]*s3.Bucket,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(*s3.Bucket))
        }
}
```

#### S3 - Minio

- **Id**: `minio:s3`
- **Given type**: `*miniotype.MinioClient`

**Tip**: You can load all based *Objstorage Minio* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/objstorage/client/s3"`

##### Type documentation
The type `*miniotype.MinioClient` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/objstorage/objstoretype/miniotype`.

This type refers to this structure:
```go
type MinioClient struct { 
        Client *minio.Client // See doc: https://github.com/minio/minio-go
        Bucket string 
}
```


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/objstorage/client/s3/minio"
        "github.com/cloudfoundry-community/gautocloud/connectors/objstorage/objstoretype/miniotype"
)
func main() {
        var err error
        // As single element
        var svc *miniotype.MinioClient
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("minio:s3", &svc)
        // or
        data, err := gautocloud.GetFirst("minio:s3")
        svc = data.(*miniotype.MinioClient)
        // ----------------------
        // as slice of elements
        var svcSlice []*miniotype.MinioClient
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("minio:s3", &svcSlice)
        // or
        data, err := gautocloud.GetAll("minio:s3")
        svcSlice = make([]*miniotype.MinioClient,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(*miniotype.MinioClient))
        }
}
```

#### S3 - Raw

- **Id**: `raw:s3`
- **Given type**: `objstoretype.S3`

**Tip**: You can load all based *Objstorage Raw* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/objstorage"`

##### Type documentation
The type `objstoretype.S3` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/objstorage/objstoretype`.

This type refers to this structure:
```go
type S3 struct { 
        Host string 
        AccessKeyID string 
        SecretAccessKey string 
        Bucket string 
        Port int 
        UseSsl bool 
}
```


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/objstorage"
        "github.com/cloudfoundry-community/gautocloud/connectors/objstorage/objstoretype"
)
func main() {
        var err error
        // As single element
        var svc objstoretype.S3
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("raw:s3", &svc)
        // or
        data, err := gautocloud.GetFirst("raw:s3")
        svc = data.(objstoretype.S3)
        // ----------------------
        // as slice of elements
        var svcSlice []objstoretype.S3
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("raw:s3", &svcSlice)
        // or
        data, err := gautocloud.GetAll("raw:s3")
        svcSlice = make([]objstoretype.S3,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(objstoretype.S3))
        }
}
```


### Smtp

All of these connectors responds on:
- Regex name: `.*smtp.*`
- Regex tags:
  - `smtp`
  - `e?mail`


#### Smtp - Client

- **Id**: `smtp`
- **Given type**: `*smtp.Client`

**Tip**: You can load all based *Smtp Client* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/smtp"`

##### Type documentation
The type `*smtp.Client` can be found in package: `net/smtp`.

You can find documentation related to package `net/smtp` here: [https://golang.org/pkg/net/smtp](https://golang.org/pkg/net/smtp).


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/smtp/client"
        "net/smtp"
)
func main() {
        var err error
        // As single element
        var svc *smtp.Client
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("smtp", &svc)
        // or
        data, err := gautocloud.GetFirst("smtp")
        svc = data.(*smtp.Client)
        defer svc.Close()
        // ----------------------
        // as slice of elements
        var svcSlice []*smtp.Client
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("smtp", &svcSlice)
        // or
        data, err := gautocloud.GetAll("smtp")
        svcSlice = make([]*smtp.Client,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(*smtp.Client))
        }
}
```

#### Smtp - Raw

- **Id**: `raw:smtp`
- **Given type**: `smtptype.Smtp`

**Tip**: You can load all based *Smtp Raw* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/smtp"`

##### Type documentation
The type `smtptype.Smtp` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/smtp/smtptype`.

This type refers to this structure:
```go
type Smtp struct { 
        User string 
        Password string 
        Host string 
        Port int 
}
```


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/smtp"
        "github.com/cloudfoundry-community/gautocloud/connectors/smtp/smtptype"
)
func main() {
        var err error
        // As single element
        var svc smtptype.Smtp
        err = gautocloud.Inject(&svc)
        // or
        err = gautocloud.InjectFromId("raw:smtp", &svc)
        // or
        data, err := gautocloud.GetFirst("raw:smtp")
        svc = data.(smtptype.Smtp)
        // ----------------------
        // as slice of elements
        var svcSlice []smtptype.Smtp
        err = gautocloud.Inject(&svcSlice)
        // or
        err = gautocloud.InjectFromId("raw:smtp", &svcSlice)
        // or
        data, err := gautocloud.GetAll("raw:smtp")
        svcSlice = make([]smtptype.Smtp,0)
        for _, elt := range data {
                svcSlice = append(svcSlice, elt.(smtptype.Smtp))
        }
}
```


