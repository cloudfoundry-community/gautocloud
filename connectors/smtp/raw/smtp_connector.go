package raw

import (
	"github.com/cloudfoundry-community/gautocloud/connectors"
	. "github.com/cloudfoundry-community/gautocloud/connectors/smtp/schema"
	"github.com/cloudfoundry-community/gautocloud/connectors/smtp/smtptype"
)

type SmtpRawConnector struct{}

func NewSmtpRawConnector() connectors.Connector {
	return &SmtpRawConnector{}
}
func (c SmtpRawConnector) Id() string {
	return "raw:smtp"
}
func (c SmtpRawConnector) Name() string {
	return ".*smtp.*"
}
func (c SmtpRawConnector) Tags() []string {
	return []string{"smtp", "e?mail"}
}
func (c SmtpRawConnector) Load(schema interface{}) (interface{}, error) {
	fSchema := schema.(SmtpSchema)
	if fSchema.Uri.Host == "" {
		return smtptype.Smtp{
			User: fSchema.User,
			Password: fSchema.Password,
			Host: fSchema.Host,
			Port: fSchema.Port,
		}, nil
	}
	port := fSchema.Uri.Port
	if port == 0 {
		port = fSchema.Port
	}
	return smtptype.Smtp{
		User: fSchema.Uri.Username,
		Password: fSchema.Uri.Password,
		Host: fSchema.Uri.Host,
		Port: port,
	}, nil
}
func (c SmtpRawConnector) Schema() interface{} {
	return SmtpSchema{}
}

