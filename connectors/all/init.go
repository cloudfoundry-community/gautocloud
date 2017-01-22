package all

import (
	_ "github.com/cloudfoundry-community/gautocloud/connectors/databases/gorm"
	_ "github.com/cloudfoundry-community/gautocloud/connectors/databases"
	_ "github.com/cloudfoundry-community/gautocloud/connectors/databases/client"
	_ "github.com/cloudfoundry-community/gautocloud/connectors/amqp"
	_ "github.com/cloudfoundry-community/gautocloud/connectors/amqp/client"
	_ "github.com/cloudfoundry-community/gautocloud/connectors/smtp"
	_ "github.com/cloudfoundry-community/gautocloud/connectors/smtp/client"
	_ "github.com/cloudfoundry-community/gautocloud/connectors/objstorage"
	_ "github.com/cloudfoundry-community/gautocloud/connectors/objstorage/client"
)

