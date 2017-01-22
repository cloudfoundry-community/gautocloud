// It has the responsibility to find the *CloudEnv* where your program run, store *Connector*s and retrieve
// services from *CloudEnv* which corresponds to one or many *Connector* and finally it will pass to *Connector* the service
// and store the result from connector.
package loader

import (
	"errors"
	"reflect"
	"github.com/cloudfoundry-community/gautocloud/decoder"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	"github.com/cloudfoundry-community/gautocloud/cloudenv"
	"strings"
	"fmt"
	ldlogger "github.com/cloudfoundry-community/gautocloud/logger"
	"log"
	"encoding/json"
)

type Loader struct {
	cloudEnvs  []cloudenv.CloudEnv
	connectors map[string]connectors.Connector
	logger     *ldlogger.LoggerLoader
	store      map[string][]StoredService
}
type StoredService struct {
	Data        interface{}
	ReflectType reflect.Type
}

// Create a new loader with cloud environment given
func NewLoader(cloudEnvs []cloudenv.CloudEnv) *Loader {
	loader := &Loader{
		cloudEnvs: cloudEnvs,
		logger: ldlogger.NewLoggerLoader(),
	}
	loader.connectors = make(map[string]connectors.Connector)
	loader.store = make(map[string][]StoredService)
	loader.LoadCloudEnvs()
	return loader
}

// Create a new loader with cloud environment given and a logger
func NewLoaderWithLogger(cloudEnvs []cloudenv.CloudEnv, logger *log.Logger, lvl ldlogger.Level) *Loader {
	loader := &Loader{
		cloudEnvs: cloudEnvs,
		logger: ldlogger.NewLoggerLoader(),
	}
	loader.connectors = make(map[string]connectors.Connector)
	loader.store = make(map[string][]StoredService)
	loader.SetLogger(logger, lvl)
	loader.LoadCloudEnvs()
	return loader
}

// Return all cloud environments loaded
func (l Loader) CloudEnvs() []cloudenv.CloudEnv {
	return l.cloudEnvs
}

// Remove all registered connectors
func (l *Loader) CleanConnectors() {
	l.connectors = make(map[string]connectors.Connector)
}

// Return all services loaded
func (l *Loader) Store() map[string][]StoredService {
	return l.store
}

// Pass a logger to the loader to let you have the possibility to see logs
// the parameter lvl is the level of verbosity which can be
//  - logger.Lall
//  - logger.Loff
//  - logger.Ldebug
//  - logger.Linfo
//  - logger.Lwarning
//  - logger.Lerror
//  - logger.Lsevere
func (l *Loader) SetLogger(logger *log.Logger, lvl ldlogger.Level) {
	l.logger.SetLevel(lvl)
	l.logger.SetLogger(logger)
}

// Register a connector in the loader
// This is mainly use for connectors creators
func (l *Loader) RegisterConnector(connector connectors.Connector) {
	err := l.checkInCloudEnv()
	if err != nil {
		l.logger.Info("Skipping registering connector '%s': %s", connector.Id(), err.Error())
		return
	}
	if _, ok := l.connectors[connector.Id()]; ok {
		l.logger.Error("During registering connector: A connector with id '%s' already exists.", connector.Id())
		return
	}
	l.logger.Debug("Loading connector '%s' ...", connector.Id())
	l.connectors[connector.Id()] = connector
	storedServices := l.load(connector)
	if len(storedServices) == 0 {
		return
	}
	l.store[connector.Id()] = storedServices
	l.logger.Debug("Finished loading connector '%s' .", connector.Id())
}

// Return all registered connectors
func (l Loader) Connectors() map[string]connectors.Connector {
	return l.connectors
}
func (l Loader) LoadCloudEnvs() {
	for _, cloudEnv := range l.cloudEnvs {
		if !cloudEnv.IsInCloudEnv() {
			l.logger.Debug("You are not in a '%s' environment", cloudEnv.Name())
			continue
		}
		err := cloudEnv.Load()
		if err != nil {
			l.logger.Error(
				"Error during loading cloud environment %s: %s",
				cloudEnv.Name(),
				err.Error(),
			)
		}
		l.logger.Info("Environment '%s' detected and loaded.", cloudEnv.Name())
	}
}

// Reload connectors to find services
func (l *Loader) ReloadConnectors() {
	l.LoadCloudEnvs()
	err := l.checkInCloudEnv()
	if err != nil {
		l.logger.Info("Skipping reloading connectors: " + err.Error())
		return
	}
	l.logger.Info("Reloading connectors ...")
	for _, connector := range l.connectors {
		storedServices := l.load(connector)
		l.store[connector.Id()] = storedServices
	}
	l.logger.Info("Finished reloading connectors ...")
}

// Inject service(s) found by connectors with given type
// Example:
//  var svc *dbtype.MysqlDB
//  err = loader.Inject(&svc)
//  // svc will have the value of the first service found with type *dbtype.MysqlDB
// If service parameter is not a slice it will give the first service found
// If you pass a slice of a type in service parameter, it will inject in the slice all services found with this type
// It returns an error if parameter is not a pointer or if no service(s) can be found
func (l Loader) Inject(service interface{}) error {
	err := l.checkInCloudEnv()
	if err != nil {
		return err
	}
	notFound := true
	for id, _ := range l.connectors {
		err = l.InjectFromId(id, service)
		if err == nil && service != nil {
			notFound = false
		}
	}
	if !notFound {
		return nil
	}
	if reflect.TypeOf(service).Kind() != reflect.Ptr {
		return errors.New("You must pass a pointer.")
	}
	reflectType := reflect.TypeOf(service).Elem()

	if reflectType.Kind() == reflect.Slice {
		reflectType = reflectType.Elem()
	}
	return errors.New("Service with the type " + reflectType.String() + " cannot be found. (perhaps no services match any connectors)")
}
// Return the current cloud env detected
func (l Loader) CurrentCloudEnv() cloudenv.CloudEnv {
	return l.getFirstValidCloudEnv()
}
// Return informations about instance of the running application
func (l Loader) GetAppInfo() cloudenv.AppInfo {
	return l.getFirstValidCloudEnv().GetAppInfo()
}
func (l Loader) checkInCloudEnv() error {
	if l.IsInACloudEnv() {
		return nil
	}
	return errors.New(fmt.Sprintf(
		"You are not in any cloud environments (available environments are: [ %s ]).",
		strings.Join(l.getCloudEnvNames(), ", "),
	))
}
func (l Loader) getCloudEnvNames() []string {
	names := make([]string, 0)
	for _, cloudEnv := range l.cloudEnvs {
		names = append(names, cloudEnv.Name())
	}
	return names
}
// Return true if you are in a cloud environment
func (l Loader) IsInACloudEnv() bool {
	for _, cloudEnv := range l.cloudEnvs {
		if !cloudEnv.IsInCloudEnv() {
			continue
		}
		return true
	}
	return false
}
func (l Loader) getFirstValidCloudEnv() cloudenv.CloudEnv {
	var finalCloudEnv cloudenv.CloudEnv
	for _, cloudEnv := range l.cloudEnvs {
		finalCloudEnv = cloudEnv
		if cloudEnv.IsInCloudEnv() {
			break
		}
	}
	return finalCloudEnv
}

// Inject service(s) found by a connector with given type
// id is the id of a connector
// Example:
//  var svc *dbtype.MysqlDB
//  err = gautocloud.InjectFromId("mysql", &svc)
//  // svc will have the value of the first service found with type *dbtype.MysqlDB in this case
// If service parameter is not a slice it will give the first service found
// If you pass a slice of a type in service parameter, it will inject in the slice all services found with this type
// It returns an error if service parameter is not a pointer, if no service(s) can be found and if connector with given id doesn't exist

func (l Loader) InjectFromId(id string, service interface{}) error {
	err := l.checkInCloudEnv()
	if err != nil {
		return err
	}
	err = l.checkConnectorIdExist(id)
	if err != nil {
		return err
	}
	if reflect.TypeOf(service).Kind() != reflect.Ptr {
		return errors.New("You must pass a pointer.")
	}
	reflectType := reflect.TypeOf(service).Elem()

	vService := reflect.ValueOf(service).Elem()
	isArray := false
	if reflectType.Kind() == reflect.Slice {
		isArray = true
		reflectType = reflectType.Elem()
	}
	dataSlice := make([]interface{}, 0)
	for _, store := range l.store[id] {
		if store.ReflectType == reflectType {
			dataSlice = append(dataSlice, store.Data)
		}
	}

	if len(dataSlice) == 0 {
		return errors.New(
			fmt.Sprintf(
				"Connector with id '%s' doesn't give a service with the type '%s'. (perhaps no services match the connector)",
				id,
				reflectType.String(),
			),
		)
	}
	if !isArray {
		vService.Set(reflect.ValueOf(dataSlice[0]))
		return nil
	}
	loadSchemas := reflect.MakeSlice(reflect.SliceOf(reflectType), 0, 0)
	for _, data := range dataSlice {
		loadSchemas = reflect.Append(loadSchemas, reflect.ValueOf(data))
	}
	if service == nil {
		vService.Set(loadSchemas)
		return nil
	}
	for i := 0; i < vService.Len(); i++ {
		loadSchemas = reflect.Append(loadSchemas, vService.Index(i))
	}
	vService.Set(loadSchemas)

	return nil
}

// Return the first service found by a connector
// id is the id of a connector
// Example:
//  var svc *dbtype.MysqlDB
//  data, err = gautocloud.GetFirst("mysql")
//  svc = data.(*dbtype.MysqlDB)
// It returns the first service found or an error if no service can be found or if the connector doesn't exists
func (l Loader) GetFirst(id string) (interface{}, error) {
	err := l.checkInCloudEnv()
	if err != nil {
		return nil, err
	}
	err = l.checkConnectorIdExist(id)
	if err != nil {
		return nil, err
	}
	if len(l.store[id]) == 0 {
		return nil, errors.New("No content have been given by connector with id '" + id + "' (no services match the connector).")
	}
	return l.store[id][0].Data, nil
}
func (l Loader) checkConnectorIdExist(id string) error {
	if _, ok := l.connectors[id]; !ok {
		return errors.New("Connector with id '" + id + "' not found.")
	}
	return nil
}

// Return all services found by a connector
// id is the id of a connector
// Example:
//  var svc []interface{}
//  data, err = gautocloud.GetAll("mysql")
//  svc = data[0].(*dbtype.MysqlDB)
// warning: a connector may give you different types that's why GetAll return a slice of interface{}
// It returns the first service found or an error if no service can be found or if the connector doesn't exists
func (l Loader) GetAll(id string) ([]interface{}, error) {
	err := l.checkInCloudEnv()
	if err != nil {
		return nil, err
	}
	err = l.checkConnectorIdExist(id)
	if err != nil {
		return nil, err
	}

	dataSlice := make([]interface{}, 0)
	for _, store := range l.store[id] {
		dataSlice = append(dataSlice, store.Data)
	}
	return dataSlice, nil
}

func (l *Loader) load(connector connectors.Connector) []StoredService {
	services := make([]cloudenv.Service, 0)
	storedServices := make([]StoredService, 0)
	cloudEnv := l.getFirstValidCloudEnv()
	services = append(services, cloudEnv.GetServicesFromTags(connector.Tags())...)
	l.addService(services, cloudEnv.GetServicesFromName(connector.Name())...)
	if len(services) == 0 {
		l.logger.Debug(
			"No service found for connector '%s' \n\twith name: '%s' \n\tor tags: [ %s ]",
			connector.Id(),
			connector.Name(),
			strings.Join(connector.Tags(), ", "),
		)
		return storedServices
	}
	serviceType := reflect.TypeOf(connector.Schema())
	for _, service := range services {
		element := reflect.New(serviceType)
		decoder.UnmarshalToValue(service.Credentials, element)
		eltInterface := element.Elem().Interface()
		loadedService, err := connector.Load(eltInterface)
		if err != nil {
			l.logger.Error("Error occured during loading connector '%s': %s\n", connector.Id(), err.Error())
			continue
		}
		reflectType := reflect.TypeOf(loadedService)
		b, _ := json.MarshalIndent(service.Credentials, "\t", "\t")
		l.logger.Debug("Connector '%s' load a service which give type '%s' from credentials:\n\t%s\n",
			connector.Id(),
			reflectType.String(),
			string(b),
		)
		storedServices = append(storedServices, StoredService{
			ReflectType: reflectType,
			Data: loadedService,
		})
	}
	return storedServices
}
func (l Loader) addService(services []cloudenv.Service, toAdd ...cloudenv.Service) {
	for _, service := range toAdd {
		if l.serviceAlreadyExists(services, service) {
			continue
		}
		services = append(services, service)
	}
}
func (l Loader) serviceAlreadyExists(services []cloudenv.Service, toFind cloudenv.Service) bool {
	for _, service := range services {
		if reflect.DeepEqual(service, toFind) {
			return true
		}
	}
	return false
}

