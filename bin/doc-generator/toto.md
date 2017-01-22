## Connectors

**Tip**: To load all default connectors import: `_ "github.com/cloudfoundry-community/gautocloud/connectors/all"`

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


### Amqp

This connector responds on:
- Regex name: `amqp`
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
        ////
        // as single element
        ////
        var svc *amqp.Connection
        err = gautocloud.Get(&svc)
        // or
        err = gautocloud.GetWithId("amqp", &svc)
        // or
        data, err := gautocloud.GetFirst("amqp")
        svc = data.(*amqp.Connection)
        ////
        // as slice of elements
        ////
        var svcSlice []*amqp.Connection
        err = gautocloud.Get(&svcSlice)
        // or
        err = gautocloud.GetWithId("amqp", &svcSlice)
        // or
        data, err := gautocloud.GetAll("amqp")
        svcSlice = data.([]*amqp.Connection)
        ////
}

```

#### Amqp - Raw

- **Id**: `raw:amqp`
- **Given type**: `amqptype.Amqp`

**Tip**: You can load all based *Amqp Raw* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/amqp"`

##### Type documentation
The type `amqptype.Amqp` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/amqp/amqptype`.

This structure will be given for this type:
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
        _ "github.com/cloudfoundry-community/gautocloud/connectors/amqp/raw"
        "github.com/cloudfoundry-community/gautocloud/connectors/amqp/amqptype"
)
func main() {
        var err error
        ////
        // as single element
        ////
        var svc amqptype.Amqp
        err = gautocloud.Get(&svc)
        // or
        err = gautocloud.GetWithId("raw:amqp", &svc)
        // or
        data, err := gautocloud.GetFirst("raw:amqp")
        svc = data.(amqptype.Amqp)
        ////
        // as slice of elements
        ////
        var svcSlice []amqptype.Amqp
        err = gautocloud.Get(&svcSlice)
        // or
        err = gautocloud.GetWithId("raw:amqp", &svcSlice)
        // or
        data, err := gautocloud.GetAll("raw:amqp")
        svcSlice = data.([]amqptype.Amqp)
        ////
}

```


### Mongodb

This connector responds on:
- Regex name: `mongo`
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
        ////
        // as single element
        ////
        var svc *mgo.Session
        err = gautocloud.Get(&svc)
        // or
        err = gautocloud.GetWithId("mongodb", &svc)
        // or
        data, err := gautocloud.GetFirst("mongodb")
        svc = data.(*mgo.Session)
        ////
        // as slice of elements
        ////
        var svcSlice []*mgo.Session
        err = gautocloud.Get(&svcSlice)
        // or
        err = gautocloud.GetWithId("mongodb", &svcSlice)
        // or
        data, err := gautocloud.GetAll("mongodb")
        svcSlice = data.([]*mgo.Session)
        ////
}

```

#### Mongodb - Raw

- **Id**: `raw:mongodb`
- **Given type**: `dbtype.MongodbDatabase`

**Tip**: You can load all based *Databases Raw* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/databases"`

##### Type documentation
The type `dbtype.MongodbDatabase` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype`.

This structure will be given for this type:
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
        _ "github.com/cloudfoundry-community/gautocloud/connectors/databases/raw"
        "github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)
func main() {
        var err error
        ////
        // as single element
        ////
        var svc dbtype.MongodbDatabase
        err = gautocloud.Get(&svc)
        // or
        err = gautocloud.GetWithId("raw:mongodb", &svc)
        // or
        data, err := gautocloud.GetFirst("raw:mongodb")
        svc = data.(dbtype.MongodbDatabase)
        ////
        // as slice of elements
        ////
        var svcSlice []dbtype.MongodbDatabase
        err = gautocloud.Get(&svcSlice)
        // or
        err = gautocloud.GetWithId("raw:mongodb", &svcSlice)
        // or
        data, err := gautocloud.GetAll("raw:mongodb")
        svcSlice = data.([]dbtype.MongodbDatabase)
        ////
}

```


### Mssql

This connector responds on:
- Regex name: `mssql`
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
        ////
        // as single element
        ////
        var svc *dbtype.MssqlDB
        err = gautocloud.Get(&svc)
        // or
        err = gautocloud.GetWithId("mssql", &svc)
        // or
        data, err := gautocloud.GetFirst("mssql")
        svc = data.(*dbtype.MssqlDB)
        ////
        // as slice of elements
        ////
        var svcSlice []*dbtype.MssqlDB
        err = gautocloud.Get(&svcSlice)
        // or
        err = gautocloud.GetWithId("mssql", &svcSlice)
        // or
        data, err := gautocloud.GetAll("mssql")
        svcSlice = data.([]*dbtype.MssqlDB)
        ////
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
        ////
        // as single element
        ////
        var svc *gorm.DB
        err = gautocloud.Get(&svc)
        // or
        err = gautocloud.GetWithId("gorm:mssql", &svc)
        // or
        data, err := gautocloud.GetFirst("gorm:mssql")
        svc = data.(*gorm.DB)
        ////
        // as slice of elements
        ////
        var svcSlice []*gorm.DB
        err = gautocloud.Get(&svcSlice)
        // or
        err = gautocloud.GetWithId("gorm:mssql", &svcSlice)
        // or
        data, err := gautocloud.GetAll("gorm:mssql")
        svcSlice = data.([]*gorm.DB)
        ////
}

```

#### Mssql - Raw

- **Id**: `raw:mssql`
- **Given type**: `dbtype.MssqlDatabase`

**Tip**: You can load all based *Databases Raw* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/databases"`

##### Type documentation
The type `dbtype.MssqlDatabase` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype`.

This structure will be given for this type:
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
        _ "github.com/cloudfoundry-community/gautocloud/connectors/databases/raw"
        "github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)
func main() {
        var err error
        ////
        // as single element
        ////
        var svc dbtype.MssqlDatabase
        err = gautocloud.Get(&svc)
        // or
        err = gautocloud.GetWithId("raw:mssql", &svc)
        // or
        data, err := gautocloud.GetFirst("raw:mssql")
        svc = data.(dbtype.MssqlDatabase)
        ////
        // as slice of elements
        ////
        var svcSlice []dbtype.MssqlDatabase
        err = gautocloud.Get(&svcSlice)
        // or
        err = gautocloud.GetWithId("raw:mssql", &svcSlice)
        // or
        data, err := gautocloud.GetAll("raw:mssql")
        svcSlice = data.([]dbtype.MssqlDatabase)
        ////
}

```


### Mysql

This connector responds on:
- Regex name: `(mysql|maria)`
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
        ////
        // as single element
        ////
        var svc *dbtype.MysqlDB
        err = gautocloud.Get(&svc)
        // or
        err = gautocloud.GetWithId("mysql", &svc)
        // or
        data, err := gautocloud.GetFirst("mysql")
        svc = data.(*dbtype.MysqlDB)
        ////
        // as slice of elements
        ////
        var svcSlice []*dbtype.MysqlDB
        err = gautocloud.Get(&svcSlice)
        // or
        err = gautocloud.GetWithId("mysql", &svcSlice)
        // or
        data, err := gautocloud.GetAll("mysql")
        svcSlice = data.([]*dbtype.MysqlDB)
        ////
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
        ////
        // as single element
        ////
        var svc *gorm.DB
        err = gautocloud.Get(&svc)
        // or
        err = gautocloud.GetWithId("gorm:mysql", &svc)
        // or
        data, err := gautocloud.GetFirst("gorm:mysql")
        svc = data.(*gorm.DB)
        ////
        // as slice of elements
        ////
        var svcSlice []*gorm.DB
        err = gautocloud.Get(&svcSlice)
        // or
        err = gautocloud.GetWithId("gorm:mysql", &svcSlice)
        // or
        data, err := gautocloud.GetAll("gorm:mysql")
        svcSlice = data.([]*gorm.DB)
        ////
}

```

#### Mysql - Raw

- **Id**: `raw:mysql`
- **Given type**: `dbtype.MysqlDatabase`

**Tip**: You can load all based *Databases Raw* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/databases"`

##### Type documentation
The type `dbtype.MysqlDatabase` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype`.

This structure will be given for this type:
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
        _ "github.com/cloudfoundry-community/gautocloud/connectors/databases/raw"
        "github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)
func main() {
        var err error
        ////
        // as single element
        ////
        var svc dbtype.MysqlDatabase
        err = gautocloud.Get(&svc)
        // or
        err = gautocloud.GetWithId("raw:mysql", &svc)
        // or
        data, err := gautocloud.GetFirst("raw:mysql")
        svc = data.(dbtype.MysqlDatabase)
        ////
        // as slice of elements
        ////
        var svcSlice []dbtype.MysqlDatabase
        err = gautocloud.Get(&svcSlice)
        // or
        err = gautocloud.GetWithId("raw:mysql", &svcSlice)
        // or
        data, err := gautocloud.GetAll("raw:mysql")
        svcSlice = data.([]dbtype.MysqlDatabase)
        ////
}

```


### Oracle

This connector responds on:
- Regex name: `oracle`
- Regex tags:
  - `oracle`
  - `oci.*`


#### Oracle - Raw

- **Id**: `raw:oracle`
- **Given type**: `dbtype.OracleDatabase`

**Tip**: You can load all based *Databases Raw* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/databases"`

##### Type documentation
The type `dbtype.OracleDatabase` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype`.

This structure will be given for this type:
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
        _ "github.com/cloudfoundry-community/gautocloud/connectors/databases/raw"
        "github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)
func main() {
        var err error
        ////
        // as single element
        ////
        var svc dbtype.OracleDatabase
        err = gautocloud.Get(&svc)
        // or
        err = gautocloud.GetWithId("raw:oracle", &svc)
        // or
        data, err := gautocloud.GetFirst("raw:oracle")
        svc = data.(dbtype.OracleDatabase)
        ////
        // as slice of elements
        ////
        var svcSlice []dbtype.OracleDatabase
        err = gautocloud.Get(&svcSlice)
        // or
        err = gautocloud.GetWithId("raw:oracle", &svcSlice)
        // or
        data, err := gautocloud.GetAll("raw:oracle")
        svcSlice = data.([]dbtype.OracleDatabase)
        ////
}

```


### Postgresql

This connector responds on:
- Regex name: `postgres`
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
        ////
        // as single element
        ////
        var svc *dbtype.PostgresqlDB
        err = gautocloud.Get(&svc)
        // or
        err = gautocloud.GetWithId("postgresql", &svc)
        // or
        data, err := gautocloud.GetFirst("postgresql")
        svc = data.(*dbtype.PostgresqlDB)
        ////
        // as slice of elements
        ////
        var svcSlice []*dbtype.PostgresqlDB
        err = gautocloud.Get(&svcSlice)
        // or
        err = gautocloud.GetWithId("postgresql", &svcSlice)
        // or
        data, err := gautocloud.GetAll("postgresql")
        svcSlice = data.([]*dbtype.PostgresqlDB)
        ////
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
        ////
        // as single element
        ////
        var svc *gorm.DB
        err = gautocloud.Get(&svc)
        // or
        err = gautocloud.GetWithId("gorm:postgresql", &svc)
        // or
        data, err := gautocloud.GetFirst("gorm:postgresql")
        svc = data.(*gorm.DB)
        ////
        // as slice of elements
        ////
        var svcSlice []*gorm.DB
        err = gautocloud.Get(&svcSlice)
        // or
        err = gautocloud.GetWithId("gorm:postgresql", &svcSlice)
        // or
        data, err := gautocloud.GetAll("gorm:postgresql")
        svcSlice = data.([]*gorm.DB)
        ////
}

```

#### Postgresql - Raw

- **Id**: `raw:postgresql`
- **Given type**: `dbtype.PostgresqlDatabase`

**Tip**: You can load all based *Databases Raw* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/databases"`

##### Type documentation
The type `dbtype.PostgresqlDatabase` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype`.

This structure will be given for this type:
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
        _ "github.com/cloudfoundry-community/gautocloud/connectors/databases/raw"
        "github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)
func main() {
        var err error
        ////
        // as single element
        ////
        var svc dbtype.PostgresqlDatabase
        err = gautocloud.Get(&svc)
        // or
        err = gautocloud.GetWithId("raw:postgresql", &svc)
        // or
        data, err := gautocloud.GetFirst("raw:postgresql")
        svc = data.(dbtype.PostgresqlDatabase)
        ////
        // as slice of elements
        ////
        var svcSlice []dbtype.PostgresqlDatabase
        err = gautocloud.Get(&svcSlice)
        // or
        err = gautocloud.GetWithId("raw:postgresql", &svcSlice)
        // or
        data, err := gautocloud.GetAll("raw:postgresql")
        svcSlice = data.([]dbtype.PostgresqlDatabase)
        ////
}

```


### Redis

This connector responds on:
- Regex name: `redis`
- Regex tags:
  - `redis`


#### Redis - Client

- **Id**: `redis`
- **Given type**: `*redis.Client`

**Tip**: You can load all based *Databases Client* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/databases/client"`

##### Type documentation
The type `*redis.Client` can be found in package: `gopkg.in/redis.v5`.

You can find documentation related to package `gopkg.in/redis.v5` here: [https://gopkg.in/redis.v5](https://gopkg.in/redis.v5).


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/databases/client/redis"
        "gopkg.in/redis.v5"
)
func main() {
        var err error
        ////
        // as single element
        ////
        var svc *redis.Client
        err = gautocloud.Get(&svc)
        // or
        err = gautocloud.GetWithId("redis", &svc)
        // or
        data, err := gautocloud.GetFirst("redis")
        svc = data.(*redis.Client)
        ////
        // as slice of elements
        ////
        var svcSlice []*redis.Client
        err = gautocloud.Get(&svcSlice)
        // or
        err = gautocloud.GetWithId("redis", &svcSlice)
        // or
        data, err := gautocloud.GetAll("redis")
        svcSlice = data.([]*redis.Client)
        ////
}

```

#### Redis - Raw

- **Id**: `raw:redis`
- **Given type**: `dbtype.RedisDatabase`

**Tip**: You can load all based *Databases Raw* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/databases"`

##### Type documentation
The type `dbtype.RedisDatabase` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype`.

This structure will be given for this type:
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
        _ "github.com/cloudfoundry-community/gautocloud/connectors/databases/raw"
        "github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)
func main() {
        var err error
        ////
        // as single element
        ////
        var svc dbtype.RedisDatabase
        err = gautocloud.Get(&svc)
        // or
        err = gautocloud.GetWithId("raw:redis", &svc)
        // or
        data, err := gautocloud.GetFirst("raw:redis")
        svc = data.(dbtype.RedisDatabase)
        ////
        // as slice of elements
        ////
        var svcSlice []dbtype.RedisDatabase
        err = gautocloud.Get(&svcSlice)
        // or
        err = gautocloud.GetWithId("raw:redis", &svcSlice)
        // or
        data, err := gautocloud.GetAll("raw:redis")
        svcSlice = data.([]dbtype.RedisDatabase)
        ////
}

```


### S3

This connector responds on:
- Regex name: `s3`
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
        ////
        // as single element
        ////
        var svc *s3.Bucket
        err = gautocloud.Get(&svc)
        // or
        err = gautocloud.GetWithId("s3", &svc)
        // or
        data, err := gautocloud.GetFirst("s3")
        svc = data.(*s3.Bucket)
        ////
        // as slice of elements
        ////
        var svcSlice []*s3.Bucket
        err = gautocloud.Get(&svcSlice)
        // or
        err = gautocloud.GetWithId("s3", &svcSlice)
        // or
        data, err := gautocloud.GetAll("s3")
        svcSlice = data.([]*s3.Bucket)
        ////
}

```

#### S3 - Minio

- **Id**: `minio:s3`
- **Given type**: `objstoretype.MinioClient`

**Tip**: You can load all based *Objstorage Minio* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/objstorage/client/s3"`

##### Type documentation
The type `objstoretype.MinioClient` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/objstorage/objstoretype`.

This structure will be given for this type:
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
        "github.com/cloudfoundry-community/gautocloud/connectors/objstorage/objstoretype"
)
func main() {
        var err error
        ////
        // as single element
        ////
        var svc objstoretype.MinioClient
        err = gautocloud.Get(&svc)
        // or
        err = gautocloud.GetWithId("minio:s3", &svc)
        // or
        data, err := gautocloud.GetFirst("minio:s3")
        svc = data.(objstoretype.MinioClient)
        ////
        // as slice of elements
        ////
        var svcSlice []objstoretype.MinioClient
        err = gautocloud.Get(&svcSlice)
        // or
        err = gautocloud.GetWithId("minio:s3", &svcSlice)
        // or
        data, err := gautocloud.GetAll("minio:s3")
        svcSlice = data.([]objstoretype.MinioClient)
        ////
}

```

#### S3 - Raw

- **Id**: `raw:s3`
- **Given type**: `objstoretype.S3`

**Tip**: You can load all based *Objstorage Raw* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/objstorage"`

##### Type documentation
The type `objstoretype.S3` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/objstorage/objstoretype`.

This structure will be given for this type:
```go
type S3 struct { 
        Host string 
        AccessKeyID string 
        SecretAccessKey string 
        Bucket string 
        UseSsl bool 
}
```


##### Example
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/objstorage/raw"
        "github.com/cloudfoundry-community/gautocloud/connectors/objstorage/objstoretype"
)
func main() {
        var err error
        ////
        // as single element
        ////
        var svc objstoretype.S3
        err = gautocloud.Get(&svc)
        // or
        err = gautocloud.GetWithId("raw:s3", &svc)
        // or
        data, err := gautocloud.GetFirst("raw:s3")
        svc = data.(objstoretype.S3)
        ////
        // as slice of elements
        ////
        var svcSlice []objstoretype.S3
        err = gautocloud.Get(&svcSlice)
        // or
        err = gautocloud.GetWithId("raw:s3", &svcSlice)
        // or
        data, err := gautocloud.GetAll("raw:s3")
        svcSlice = data.([]objstoretype.S3)
        ////
}

```


### Smtp

This connector responds on:
- Regex name: `smtp`
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
        ////
        // as single element
        ////
        var svc *smtp.Client
        err = gautocloud.Get(&svc)
        // or
        err = gautocloud.GetWithId("smtp", &svc)
        // or
        data, err := gautocloud.GetFirst("smtp")
        svc = data.(*smtp.Client)
        ////
        // as slice of elements
        ////
        var svcSlice []*smtp.Client
        err = gautocloud.Get(&svcSlice)
        // or
        err = gautocloud.GetWithId("smtp", &svcSlice)
        // or
        data, err := gautocloud.GetAll("smtp")
        svcSlice = data.([]*smtp.Client)
        ////
}

```

#### Smtp - Raw

- **Id**: `raw:smtp`
- **Given type**: `smtptype.Smtp`

**Tip**: You can load all based *Smtp Raw* by importing: `_ "github.com/cloudfoundry-community/gautocloud/connectors/smtp"`

##### Type documentation
The type `smtptype.Smtp` can be found in package: `github.com/cloudfoundry-community/gautocloud/connectors/smtp/smtptype`.

This structure will be given for this type:
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
        _ "github.com/cloudfoundry-community/gautocloud/connectors/smtp/raw"
        "github.com/cloudfoundry-community/gautocloud/connectors/smtp/smtptype"
)
func main() {
        var err error
        ////
        // as single element
        ////
        var svc smtptype.Smtp
        err = gautocloud.Get(&svc)
        // or
        err = gautocloud.GetWithId("raw:smtp", &svc)
        // or
        data, err := gautocloud.GetFirst("raw:smtp")
        svc = data.(smtptype.Smtp)
        ////
        // as slice of elements
        ////
        var svcSlice []smtptype.Smtp
        err = gautocloud.Get(&svcSlice)
        // or
        err = gautocloud.GetWithId("raw:smtp", &svcSlice)
        // or
        data, err := gautocloud.GetAll("raw:smtp")
        svcSlice = data.([]smtptype.Smtp)
        ////
}

```


