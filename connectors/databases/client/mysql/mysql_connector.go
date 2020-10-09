package mysql

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/raw"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	gautocloud.RegisterConnector(NewMysqlConnector())
}

var allowedParams = []string{
	"allowAllFiles", "allowCleartextPasswords", "allowNativePasswords", "allowOldPasswords", "charset", "checkConnLiveness", "collation",
	"clientFoundRows", "columnsWithAlias", "interpolateParams", "loc", "maxAllowedPacket", "multiStatements", "parseTime", "readTimeout",
	"rejectReadOnly", "serverPubKey", "timeout", "tls", "writeTimeout",
}

type MysqlConnector struct {
	mysqlRawConn connectors.Connector
}

func NewMysqlConnector() connectors.Connector {
	return &MysqlConnector{
		mysqlRawConn: raw.NewMysqlRawConnector(),
	}
}
func (c MysqlConnector) Id() string {
	return "mysql"
}
func (c MysqlConnector) Name() string {
	return c.mysqlRawConn.Name()
}
func (c MysqlConnector) Tags() []string {
	return c.mysqlRawConn.Tags()
}
func (c MysqlConnector) GetConnString(schema dbtype.MysqlDatabase) string {
	connString := schema.User
	if schema.Password != "" {
		connString += ":" + schema.Password
	}
	connString += fmt.Sprintf("@tcp(%s:%d)/%s", schema.Host, schema.Port, schema.Database)
	if schema.Options == "" {
		schema.Options = "parseTime=true"
	} else {
		schema.Options += "&parseTime=true"
	}
	values, _ := url.ParseQuery(schema.Options)
	finalValues := make(url.Values)
	for k, v := range values {
		if c.isAllowedParams(k) {
			finalValues[k] = v
		}
	}
	connString += "?" + finalValues.Encode()
	return connString
}

func (c MysqlConnector) isAllowedParams(param string) bool {
	for _, allowed := range allowedParams {
		if param == allowed {
			return true
		}
	}
	return false
}

func (c MysqlConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := c.mysqlRawConn.Load(schema)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("mysql", c.GetConnString(schema.(dbtype.MysqlDatabase)))
	if err != nil {
		return db, err
	}
	return &dbtype.MysqlDB{db}, nil
}
func (c MysqlConnector) Schema() interface{} {
	return c.mysqlRawConn.Schema()
}
