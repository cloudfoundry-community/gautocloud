package client

import (
	"crypto/tls"
	"errors"
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/connectors/smtp/raw"
	"github.com/cloudfoundry-community/gautocloud/connectors/smtp/smtptype"
	"net"
	"net/smtp"
	"strconv"
	"time"
)

func init() {
	gautocloud.RegisterConnector(NewSmtpConnector())
}

type SmtpConnector struct {
	wrapConn connectors.Connector
}

func NewSmtpConnector() connectors.Connector {
	return &SmtpConnector{
		wrapConn: raw.NewSmtpRawConnector(),
	}
}
func (c SmtpConnector) Id() string {
	return "smtp"
}
func (c SmtpConnector) Name() string {
	return c.wrapConn.Name()
}
func (c SmtpConnector) Tags() []string {
	return c.wrapConn.Tags()
}
func (c SmtpConnector) GetAuth(schema smtptype.Smtp) smtp.Auth {
	return smtp.PlainAuth("", schema.User, schema.Password, schema.Host+":"+strconv.Itoa(schema.Port))
}
func (c SmtpConnector) GetSmtp(schema smtptype.Smtp, withAuth bool, isTls bool, startTls bool) (*smtp.Client, error) {
	var conn net.Conn
	var err error
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         schema.Host,
	}
	host := schema.Host + ":" + strconv.Itoa(schema.Port)
	conn, err = net.DialTimeout("tcp", host, time.Millisecond*500)
	if err != nil {
		return nil, err
	}
	if isTls {
		conn = tls.Client(conn, tlsconfig)
	}
	client, err := smtp.NewClient(conn, host)
	if err != nil {
		conn.Close()
		return client, err
	}
	if isTls && startTls && withAuth {
		err = client.StartTLS(tlsconfig)
		if err != nil {
			client.Close()
			client.Quit()
			conn.Close()
			return client, err
		}
	}
	if withAuth {
		err = client.Auth(c.GetAuth(schema))
		if err != nil {
			client.Close()
			client.Quit()
			conn.Close()
			return client, err
		}
	}
	return client, nil
}
func (c SmtpConnector) Load(schema interface{}) (interface{}, error) {
	schema, err := c.wrapConn.Load(schema)
	if err != nil {
		return nil, err
	}
	fSchema := schema.(smtptype.Smtp)
	var client *smtp.Client
	errorMessage := ""
	client, err = c.GetSmtp(fSchema, true, true, true)
	if err != nil {
		errorMessage += "\t- tls with startls: " + err.Error() + "\n"
		client, err = c.GetSmtp(fSchema, true, true, false)
	}
	if err != nil {
		errorMessage += "\t- tls: " + err.Error() + "\n"
		client, err = c.GetSmtp(fSchema, true, false, false)
	}
	if err != nil {
		errorMessage += "\t- plain auth: " + err.Error() + "\n"
		client, err = c.GetSmtp(fSchema, false, false, false)
	}
	if err != nil {
		errorMessage += "\t- no auth: " + err.Error() + "\n"
		return nil, errors.New("No smtp are reachable (trying: tls with starttls, tls, plain auth and no auth):\n" + errorMessage)
	}
	return client, nil

}
func (c SmtpConnector) Schema() interface{} {
	return c.wrapConn.Schema()
}
