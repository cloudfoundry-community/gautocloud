package raw

import (
	"github.com/cloudfoundry-community/gautocloud/connectors"
	. "github.com/cloudfoundry-community/gautocloud/connectors/auth/schema"
)

type Oauth2RawConnector struct{}

func NewOauth2RawConnector() connectors.Connector {
	return &Oauth2RawConnector{}
}
func (c Oauth2RawConnector) Id() string {
	return "raw:oauth2"
}
func (c Oauth2RawConnector) Name() string {
	return ".*oauth.*"
}
func (c Oauth2RawConnector) Tags() []string {
	return []string{"oauth.*", "sso"}
}
func (c Oauth2RawConnector) Load(schema interface{}) (interface{}, error) {
	return schema.(Oauth2Schema), nil
}
func (c Oauth2RawConnector) Schema() interface{} {
	return Oauth2Schema{}
}
