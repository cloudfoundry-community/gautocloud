
# Gautocloud [![Build Status](https://travis-ci.org/cloudfoundry-community/gautocloud.svg?branch=master)](https://travis-ci.org/cloudfoundry-community/gautocloud) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![GoDoc](https://godoc.org/github.com/cloudfoundry-community/gautocloud?status.svg)](https://godoc.org/github.com/cloudfoundry-community/gautocloud)

Gautocloud provides a simple abstraction that golang based applications can use 
to discover information about the cloud environment on which they are running, 
to connect to services automatically with ease of use in mind. It provides out-of-the-box support 
for discovering common services on Heroku and Cloud Foundry cloud platforms, 
and it supports custom automatic connectors.

This project can be assimilated to the [spring-cloud-connector](https://github.com/spring-cloud/spring-cloud-connectors) project
 but for golang (and with is own concepts).

## Summary

- [Usage by example](#usage-by-example)
- [Default connectors](/docs/connectors.md)
- [Cloud Environments](#cloud-environments)
  - [Cloud Foundry](#cloud-foundry)
  - [Heroku](#heroku)
  - [Local](#local)
- [Concept](#concept)
  - [Architecture](#architecture)
  - [Connector registration sequence](#connector-registration-sequence)
  - [Usage by injection sequence](#usage-by-injection-sequence)
- [Create your own connector](#create-your-own-connector)
- [Create your own Cloud Environment](#create-your-own-cloud-environment)
- [Use it without the facade](#use-it-without-the-facade)
- [Use mocked facade for your test](#use-mocked-facade-for-your-test)
- [Run tests](#run-tests)
- [Contributing](#contributing)
- [FAQ](#faq)

## Usage by example

Let's define a context: We are in a Cloud Foundry environment where we connect a MySql service on our 
application.

We now wants to use this service without parsing a json or anything else to have a MySql client to use our service.
Gautocloud is here to help for this kind of use case.

This software will retrieve all services found in your environment and will pass informations from service to what we call
a connector. A connector is responsible to create, in our context, a MySql client which make it available in your program.

You only needs to import a connector to make it useable by gautocloud. 
This system of import let you have only what you need to run your app (=do not create a huge binary) and let the possibility to create custom connector.

Example (provide a `*net/sql.DB` struct):
```go
package main
import (
        "fmt"
        "github.com/cloudfoundry-community/gautocloud"
        "github.com/cloudfoundry-community/gautocloud/logger"
        "log"
        "os"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/databases/client/mysql" // this register the connector mysql to gautocloud
        "github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)
func main() {
        // always attach a logger to see what happens
        gautocloud.SetLogger(log.New(os.Stdout, "", log.Ldate | log.Ltime), logger.Linfo) // set to level Ldebug to see services found
        appInfo := gautocloud.GetAppInfo() // retrieve all informations about your application instance
        fmt.Println(appInfo.Name) // give the app name
        // by injection 
        var c *dbtype.MysqlDB // this is just a wrapper of *net/sql.DB you can use as normal sql.DB client
        err := gautocloud.Inject(&c) // you can also use gautocloud.InjectFromId("mysql", &c) where "mysql" is the id of the connector to use
        if err != nil {
                panic(err)
        }
        defer c.Close()
        // c is now useable as a *sql.DB
        // e.g.: err = c.Ping()
        
        // or you can also do by return
        // data, err := gautocloud.GetFirst("mysql")
        // if err != nil {
        //         panic(err)
        // }
        // c = data.(*dbtype.MysqlDB)
        
}
```

Imagine now that we have multiple MySql services connected to your Cloud Foundry app, you can also have multiple client:
```go
package main
import (
        "github.com/cloudfoundry-community/gautocloud"
        _ "github.com/cloudfoundry-community/gautocloud/connectors/databases/client/mysql" // this register the connector mysql to gautocloud
        "github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)
func main() {
        // by injection 
        var cs []*dbtype.MysqlDB // this is just a wrapper of *net/sql.DB you can use as normal sql.DB client
        err := gautocloud.Inject(&cs) // you can also use gautocloud.InjectFromId("mysql", &cs) where "mysql" is the id of the connector to use
        if err != nil {
                panic(err)
        }
        // you have now a slice containing all mysql client you can have
        
        // or you can also do by return
        // data, err := gautocloud.GetAll("mysql") // a connector may give you different types that's why GetAll return a slice of interface{}
        // cs = make([]*dbtype.MysqlDB,0)
        // for _, elt := range data {
        //        svcSlice = append(cs, elt.(*dbtype.MysqlDB))
        // }
        
}
```

**Tip**: You can either do the same thing without the gautocloud facade, see: [use gautocloud without facade](#use-it-without-the-facade)

## Connectors

Doc for default connectors can be found here: [/docs/connectors.md](/docs/connectors.md).

You can see connectors made by the community on the dedicated wiki page: https://github.com/cloudfoundry-community/gautocloud/wiki/Connectors

## Cloud Environments

### Cloud Foundry

- **Cloud Detection**: if the `VCAP_APPLICATION` env var exists and not empty
- **Service detection by name**: Look if a service in `VCAP_SERVICES` match the name required by a connector.
- **Service detection by tags**: Look if a service in `VCAP_SERVICES` match one of tag required by a connector.
- **App information id**: id of the app given by Cloud Foundry
- **App information name**: name of the app given during `cf push`
- **App information properties**:
  - `uris`: (type: *[]string*) list of routes associated to the apps.
  - `host`: (type: *string*) host of the app.
  - `home`: (type: *string*) root folder for the deployed app.
  - `index`: (type: *int*) index of the app.
  - `memory_limit`: (type: *string*) maximum amount of memory that each instance of the application can consume.
  - `port`: (type: *int*) port of the app.
  - `space_id`: (type: *string*) id of the space.
  - `space_name_id`: (type: *string*) name of the space.
  - `temp_dir`: (type: *string*) directory location where temporary and staging files are stored.
  - `user`: (type: *string*) user account under which the container runs.
  - `version`: (type: *string*) version of the app.
  - `working_dir`: (type: *string*) present working directory, where the buildpack that processed the application ran.

### Heroku

**Tip**: you can also use in local but settings the env var `DYNO` and create env var corresponding 
to a service you want to connect. (see: [/test-integration/test_integration_test.go](/test-integration/test_integration_test.go) as an example)

- **Cloud Detection**: if the `DYNO` env var exists
- **Service detection by name**: Look all env var which contains the name required by a connector. Env var key are after parsed to create credentials.
Example:
```
you have env var:
- `MY_SVC_NAME=myname`
- `MY_SVC_HOST=localhost`

Connector required name `SVC`.
CloudEnv decode `MY_SVC_NAME` to [MY, SVC, NAME] and [MY, SVC, VALUE] 
and detect that there is SVC in those two env var.
It returns a service with credentials:
{
  "name": "myname",
  "host": "localhost"
}
```
**Note**: if a env var key doesn't contains `_` (e.g.: `SVC=localhost`) it will give those credentials: `{"svc": "localhost", "uri": "localhost"}`.

- **Service detection by tags**: each tag work like by name.
- **App information id**: id of the app given by the env var `DYNO`
- **App information name**: Set the env var `GAUTOCLOUD_APP_NAME` to give a name to your app instead it will be `<unknown>`
- **App information properties**:
  - `host`: (type: *string*) host of the app.
  - `port`: (type: *int*) port of the app.

### Local

This is a special *CloudEnv* and can be considered as a fake one.

You need to set the env var `CLOUD_FILE` which contains the path of a configuration files containing services. 
This config file can be a `yml`, `json`, `toml` or `hcl` file. It only requires to follow this pattern (example in yml):

```yml
app_name: "myapp" # set the app name you want (it can be not set)
services:
- name: myelephantsql
  tags: [postgresql, service] # ...
  credentials:
    uri: postgres://seilbmbd:PHxTPJSn@babar.elephantsql.com:5432/seilbmbd
    host: babar.elephantsql.com
    # ... you can have other credentials
```

You can see how to follow the same pattern with other format here: [/cloudenv/local_cloudenv_test.go#L13-L86](/cloudenv/local_cloudenv_test.go#L13-L86).

- **Cloud Detection**: if the `CLOUD_FILE` env var exists and not empty.
- **Service detection by name**: Look if a service in the config file match the name required by a connector.
- **Service detection by tags**: Look if a service in the config file match one of tag required by a connector.
- **App information id**: random uuid
- **App information name**: The name given in the config file, if not set it will be `<unknown>`
- **App information properties**: *None*

## Concept

Gautocloud have a lot of black magics but in fact the concept is quite simple.

### Architecture

![Architecture](/docs/arch.png)

- **Loader**: It has the responsibility to find the *CloudEnv* where your program run, store *Connector*s and retrieve 
services from *CloudEnv* which corresponds to one or many *Connector* and finally it will pass to *Connector* the service
and store the result from connector.
- **Gautocloud *facade***: This facade was made to make things easier for users. It store one instance of a *Loader*
 and give the ability to make lazy loading (this is why to register a *Connector* you only need to do `import _ "a/connector"`) 
- **CloudEnv**: Each *CloudEnv* correspond to a real cloud. It manages the detection of the environment but 
also the detections of services asked by the *Loader*.
- **Connector**: A connector register itself on the loader when using *Gautocloud Facade*. It handles the conversion of 
a service to a real client or structure which can be manipulated after by user.
- **CloudDecoder**: This decoder do the conversion of a service to an expected schema. 
In *Gautocloud* context this decoder is used to convert a given service to an expected schema given by a connector.
This decoder can be used in other context. (see: [/decoder/decoder.go](/decoder/decoder.go) to know about it)

### Connector registration sequence

![Connector registration sequence](/docs/reg.png)

### Usage by injection sequence

![Usage injection sequence](/docs/inject.png)

## Create your own connector

The best way is to look at an example here: [/connectors/example_test.go](/connectors/example_test.go).

**Note**: Add your connector on the dedicated wiki page: https://github.com/cloudfoundry-community/gautocloud/wiki/Connectors

## Create your own Cloud Environment

The best way to implement yourself a cloud environment is too look at interface here [/cloudenv/cloudenv.go](/cloudenv/cloudenv.go).

You will need to load you cloud env after by [use gautocloud without facade](#use-it-without-the-facade), 
you can either do a pull request to had your cloud environment as builtin by doing a pull request.

## Use it without the facade

We will take the same example as we see in [Usage by example](#usage-by-example) but we will not use the facade this time:

```go
package main
import (
        "fmt"
        "github.com/cloudfoundry-community/gautocloud/loader"
        "github.com/cloudfoundry-community/gautocloud/cloudenv"
        "github.com/cloudfoundry-community/gautocloud/logger"
        "log"
        "os"
        "github.com/cloudfoundry-community/gautocloud/connectors/databases/client/mysql" // this register the connector mysql to gautocloud
        "github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
)
func main() {
        ld := loader.NewLoaderWithLogger(
            []cloudenv.CloudEnv{
                cloudenv.NewCfCloudEnv(),
                cloudenv.NewHerokuCloudEnv(),
                cloudenv.NewLocalCloudEnv(),
            },
            log.New(os.Stdout, "", log.Ldate | log.Ltime), 
            logger.Linfo,
        )
        ld.RegisterConnector(mysql.NewMysqlConnector()) // you need to manually register connectors
        
        appInfo := ld.GetAppInfo() // retrieve all informations about your application instance
        fmt.Println(appInfo.Name) // give the app name
        // by injection 
        var c *dbtype.MysqlDB // this is just a wrapper of *net/sql.DB you can use as normal sql.DB client
        err := ld.Inject(&c) // you can also use gautocloud.InjectFromId("mysql", &c) where "mysql" is the id of the connector to use
        if err != nil {
                panic(err)
        }
        defer c.Close()
        
}
```

## Use mocked facade for your test

If you need to write your tests with a mocked gautocloud (and use the facade), you can ask to go compiler to get a mocked version.

To perform this, simply run your test with the tags `gautocloud_mock` (e.g.: `go test -tags gautocloud_mock`).

The facade will load a [gomock](https://github.com/golang/mock) version of the loader. 
You can find, for example, how to perform injection with this mocked version here: 
[/test-mock/test_mock_test.go](/test-mock/test_mock_test.go) (see also [gomock documentation](https://github.com/golang/mock) to learn more)

## Run tests

Requirements:
- [ginkgo](https://onsi.github.io/ginkgo/)

Requirements for integration tests:
- [docker](https://docs.docker.com/engine/installation/)
- [docker-compose](https://docs.docker.com/compose/install/)

**Note**: We need docker and docker-compose for integrations to run services and ensure clients works.

Simply run in a terminal `bin/test.sh`.

## Contributing

Any PR or/and issues are welcomes.

Don't be shy to send a PR to add another cloud environment as a builtin one.

## FAQ

**Why do I need to import a connector even if it's a builtin one ?**

You need to import it because if you didn't have to it will requires to load all default connectors with 
associated dependencies to the connector which can make a huge binary. 

In our case, it will compile only what you need by importing the connector.
 
**Why the *CloudEnv* interface works with a concept of tags and name?**
 
This concept comes directly from Cloud Foundry and the way it gives services. Cloud Foundry has a api called 
[service brokers](https://docs.cloudfoundry.org/services/api.html) this api return services with tags and name. 

This concept will now be used in the future by a lot of cloud environment (PaaS and CaaS) because this api have now 
a dedicated governance managed by people from Google (Kubernetes in mind), Pivotal (Cloud Foundry), Red Hat (Openshift) ... 
You can found their website here: https://www.openservicebrokerapi.org/


