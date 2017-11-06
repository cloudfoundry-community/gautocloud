package configfile

import (
	"errors"
	"fmt"
	"github.com/cloudfoundry-community/gautocloud/decoder"
	"github.com/cloudfoundry-community/gautocloud/interceptor"
	"github.com/cloudfoundry-community/gautocloud/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"reflect"
)

type ConfigFileInterceptor struct {
	configPath string
}

func NewConfigFile() *ConfigFileInterceptor {
	return &ConfigFileInterceptor{}
}
func (i ConfigFileInterceptor) Intercept(current, found interface{}) (interface{}, error) {
	confPath := i.configPath
	if confPath == "" {
		log.Warn("ConfigFileInterceptor: Skipping loading config file, config file path not set.")
		return found, nil
	}
	_, err := os.Stat(confPath)
	if err != nil {
		log.Warnf(
			"ConfigFileInterceptor: Skipping loading config file, can't load config file '%s', see details: %s",
			confPath,
			err.Error(),
		)
		return found, nil
	}

	schema := current
	if schema == nil {
		schema = found
	}
	viper.SetConfigType(filepath.Ext(confPath)[1:])
	viper.SetConfigFile(confPath)
	err = viper.ReadInConfig()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Fatal error on reading config file: %s \n", err.Error()))
	}
	var creds map[interface{}]interface{}
	err = viper.Unmarshal(&creds)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Fatal error when unmarshaling config file: %s \n", err.Error()))
	}
	finalCreds := utils.ConvertMapInterface(creds).(map[string]interface{})
	schemaPtr := interceptor.InterfaceAsPtr(schema)

	err = decoder.Unmarshal(finalCreds, schemaPtr)
	if err != nil {
		return nil, err
	}
	if reflect.TypeOf(schema).Kind() == reflect.Ptr {
		return interceptor.NewOverwrite().Intercept(schemaPtr, found)
	}
	return interceptor.NewOverwrite().Intercept(
		reflect.ValueOf(schemaPtr).Elem().Interface(),
		found,
	)
}

func (i *ConfigFileInterceptor) SetConfigPath(configPath string) {
	i.configPath = configPath
}
