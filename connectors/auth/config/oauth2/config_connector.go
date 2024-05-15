package oauth2

import (
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/connectors/auth/raw"
	. "github.com/cloudfoundry-community/gautocloud/connectors/auth/schema"
	"golang.org/x/oauth2"
)

func init() {
	gautocloud.RegisterConnector(NewOauth2ConfigConnector())
}

type Oauth2ConfigConnector struct {
	rawConn connectors.Connector
}

func NewOauth2ConfigConnector() connectors.Connector {
	return &Oauth2ConfigConnector{
		rawConn: raw.NewOauth2RawConnector(),
	}
}
func (c Oauth2ConfigConnector) Id() string {
	return "config:oauth2"
}
func (c Oauth2ConfigConnector) Name() string {
	return c.rawConn.Name()
}
func (c Oauth2ConfigConnector) Tags() []string {
	return c.rawConn.Tags()
}
func (c Oauth2ConfigConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := c.rawConn.Load(schema)
	if err != nil {
		return nil, err
	}
	fSchema := schema.(Oauth2Schema)
	config := &oauth2.Config{
		ClientID:     fSchema.ClientId,
		ClientSecret: fSchema.ClientSecret,
		Scopes:       fSchema.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  fSchema.AuthorizationUri,
			TokenURL: fSchema.TokenUri,
		},
	}
	return config, nil
}
func (c Oauth2ConfigConnector) Schema() interface{} {
	return c.rawConn.Schema()
}
