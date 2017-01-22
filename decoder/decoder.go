// Package decoder provide a way to decode credentials from a service to a structure
// It provide a cloud tag to help user match the correct credentials
//
// This is what you can pass as a structure:
//  // Name is key of a service credentials, decoder will look at any matching credentials which have the key name and will pass the value of this credentials
//  	Name    string `cloud:"name"`           // note: by default if you don't provide a cloud tag the key will be the field name in snake_case
//  	Uri     decoder.ServiceUri              // ServiceUri is a special type. Decoder will expect an uri as a value and will give a ServiceUri
//  	User    string `cloud:".*user.*,regex"` // by passing `regex` in cloud tag it will say to decoder that the expected key must be match the regex
//  	Pasword string `cloud:".*user.*,regex,default=apassword"` // by passing `default=avalue` decoder will understand that if the key is not found it must fill the field with this value
//  }
package decoder

import (
	"reflect"
	"strings"
	"fmt"
	"regexp"
	"strconv"
	"net/url"
	"errors"
	"github.com/azer/snakecase"
)

const (
	identifier = "cloud"
	regexTag = "regex"
	defaultTag = "default"
	skipTag = "-"
)

type Tag struct {
	Name         string
	Skip         bool
	IsRegex      bool
	DefaultValue string
}
type ServiceUri struct {
	Username string
	Password string
	Scheme   string
	Host     string
	Name     string
	Query    []QueryUri
	RawQuery string
	Port     int
}
type QueryUri struct {
	Key   string
	Value string
}

// Decode a map of credentials into a reflected Value
func UnmarshalToValue(serviceCredentials map[string]interface{}, ps reflect.Value) error {
	v := ps.Elem()
	t := v.Type()
	var err error
	for index := 0; index < v.NumField(); index++ {
		vField := v.Field(index)
		tField := t.Field(index)
		if !vField.CanAddr() || !vField.CanSet() {
			continue
		}
		tag := parseInTag(tField.Tag.Get(identifier), tField.Name)
		if tag.Skip {
			continue
		}
		key := tag.Name
		if tag.IsRegex {
			key = getKeyFromRegex(serviceCredentials, tag.Name)
		}
		if !isValueExists(serviceCredentials, key) && tag.DefaultValue == "" {
			continue
		}
		var data interface{}
		if !isValueExists(serviceCredentials, key) && tag.DefaultValue != "" {
			data = tag.DefaultValue
		} else {
			data = serviceCredentials[key]
		}
		dataKind := reflect.TypeOf(data).Kind()
		if dataKind == reflect.String {
			data, err = convertStringValue(data.(string), vField)
			if err != nil {
				return errors.New(fmt.Sprintf(
					"Error on field '%s' when trying to convert value '%s' in '%s': %s",
					tField.Name,
					tag.DefaultValue,
					vField.Kind().String(),
					err.Error(),
				))
			}
		}
		err = affect(data, vField)
		if err != nil {
			return errors.New(fmt.Sprintf("Error on field '%s': %s", tField.Name, err.Error()))
		}
	}
	return nil
}
// Decode a map of credentials into a structure
func Unmarshal(serviceCredentials map[string]interface{}, obj interface{}) error {
	ps := reflect.ValueOf(obj)
	return UnmarshalToValue(serviceCredentials, ps)
}
func affect(data interface{}, vField reflect.Value) error {
	switch vField.Kind() {
	case reflect.String:
		vField.SetString(data.(string))
		break
	case reflect.Int:
		vField.SetInt(int64(data.(int)))
		break
	case reflect.Int8:
		vField.SetInt(int64(data.(int8)))
		break
	case reflect.Int16:
		vField.SetInt(int64(data.(int16)))
		break
	case reflect.Int32:
		vField.SetInt(int64(data.(int32)))
		break
	case reflect.Int64:
		vField.SetInt(data.(int64))
		break
	case reflect.Uint:
		vField.SetUint(uint64(data.(uint)))
		break
	case reflect.Uint8:
		vField.SetUint(uint64(data.(uint8)))
		break
	case reflect.Uint16:
		vField.SetUint(uint64(data.(uint16)))
		break
	case reflect.Uint32:
		vField.SetUint(uint64(data.(uint32)))
		break
	case reflect.Uint64:
		vField.SetUint(data.(uint64))
		break
	case reflect.Interface:
		vField.Set(reflect.ValueOf(data))
		break
	case reflect.Bool:
		vField.SetBool(data.(bool))
		break
	case reflect.Float32:
		vField.SetFloat(float64(data.(float32)))
		break
	case reflect.Float64:
		vField.SetFloat(data.(float64))
		break
	case reflect.Ptr:
		if vField.IsNil() {
			vField.Set(reflect.New(vField.Type().Elem()))
		}
		err := affect(data, vField.Elem())
		if err != nil {
			return err
		}
		break
	default:
		servUriType := reflect.TypeOf(ServiceUri{})
		if vField.Type() != servUriType {
			return errors.New(fmt.Sprintf("Type '%s' is not supported", vField.Type().String()))
		}
		serviceUrl, err := url.Parse(data.(string))
		if err != nil {
			return err
		}
		serviceUri := urlToServiceUri(serviceUrl)
		vField.Set(reflect.ValueOf(serviceUri))
		break
	}
	return nil
}
func parseInTag(tag, fieldName string) Tag {
	if tag == "" {
		return Tag{
			Name: snakecase.SnakeCase(fieldName),
		}
	}
	tag = strings.TrimSpace(tag)
	splitedTag := strings.Split(tag, ",")
	name := splitedTag[0]
	skipped := false
	if name == skipTag {
		skipped = true
	}
	if name == "" {
		name = snakecase.SnakeCase(fieldName)
	}

	return Tag{
		Name: name,
		Skip: skipped,
		IsRegex: hasRegexTag(splitedTag[1:]),
		DefaultValue: getDefaultTagValue(splitedTag[1:]),
	}
}
func hasRegexTag(tags []string) bool {
	for _, tag := range tags {
		if tag == regexTag {
			return true
		}
	}
	return false
}
func getDefaultTagValue(tags []string) string {

	for _, tag := range tags {
		splitedDefTag := strings.Split(tag, "=")
		if len(splitedDefTag) < 2 || splitedDefTag[0] != defaultTag {
			continue
		}
		return strings.TrimSpace(strings.Join(splitedDefTag[1:], "="))
	}
	return ""
}
func isValueExists(serviceCredentials map[string]interface{}, key string) bool {
	if key == "" {
		return false
	}
	_, ok := serviceCredentials[key]
	return ok
}
func match(matcher, content string) bool {
	regex, err := regexp.Compile("(?i)^" + matcher + "$")
	if err != nil {
		return false
	}
	return regex.MatchString(content)
}
func getKeyFromRegex(serviceCredentials map[string]interface{}, regexKey string) string {
	for key, _ := range serviceCredentials {
		if match(regexKey, key) {
			return key
		}
	}
	return ""
}
func urlToServiceUri(url *url.URL) ServiceUri {
	username := ""
	password := ""
	if url.User != nil {
		if url.User.Username() != "" {
			username = url.User.Username()
		}
		_, hasPassword := url.User.Password()
		if hasPassword {
			password, _ = url.User.Password()
		}
	}
	queries := make([]QueryUri, 0)
	for key, value := range url.Query() {
		queries = append(queries, QueryUri{
			Key: key,
			Value: value[0],
		})
	}
	host := url.Host
	port := 0
	splitedHost := strings.Split(host, ":")
	if len(splitedHost) == 2 {
		host = splitedHost[0]
		port, _ = strconv.Atoi(splitedHost[1])
	}
	return ServiceUri{
		Scheme: url.Scheme,
		Username: username,
		Password: password,
		Host: host,
		Port: port,
		Name: strings.TrimPrefix(url.Path, "/"),
		Query: queries,
		RawQuery: url.RawQuery,
	}
}

func convertStringValue(defVal string, vField reflect.Value) (interface{}, error) {
	switch vField.Kind() {
	case reflect.String:
		return defVal, nil
	case reflect.Interface:
		return defVal, nil
	case reflect.Int:
		return strconv.Atoi(defVal)
	case reflect.Int8:
		val, err := strconv.ParseInt(defVal, 10, 8)
		if err != nil {
			return "", err
		}
		return int8(val), nil
	case reflect.Int16:
		val, err := strconv.ParseInt(defVal, 10, 16)
		if err != nil {
			return "", err
		}
		return int16(val), nil
	case reflect.Int32:
		val, err := strconv.ParseInt(defVal, 10, 32)
		if err != nil {
			return "", err
		}
		return int32(val), nil
	case reflect.Int64:
		val, err := strconv.ParseInt(defVal, 10, 64)
		if err != nil {
			return "", err
		}
		return int64(val), nil
	case reflect.Uint:
		val, err := strconv.ParseUint(defVal, 10, int(strconv.IntSize))
		if err != nil {
			return "", err
		}
		return uint(val), nil
	case reflect.Uint8:
		val, err := strconv.ParseUint(defVal, 10, 8)
		if err != nil {
			return "", err
		}
		return uint8(val), nil
	case reflect.Uint16:
		val, err := strconv.ParseUint(defVal, 10, 16)
		if err != nil {
			return "", err
		}
		return uint16(val), nil
	case reflect.Uint32:
		val, err := strconv.ParseUint(defVal, 10, 32)
		if err != nil {
			return "", err
		}
		return uint32(val), nil
	case reflect.Uint64:
		val, err := strconv.ParseUint(defVal, 10, 64)
		if err != nil {
			return "", err
		}
		return uint64(val), nil
	case reflect.Bool:
		return strconv.ParseBool(defVal)
	case reflect.Float32:
		val, err := strconv.ParseFloat(defVal, 32)
		if err != nil {
			return "", err
		}
		return float32(val), nil
	case reflect.Float64:
		val, err := strconv.ParseFloat(defVal, 64)
		if err != nil {
			return "", err
		}
		return float64(val), nil
	case reflect.Ptr:
		if vField.IsNil() {
			vField.Set(reflect.New(vField.Type().Elem()))
		}
		return convertStringValue(defVal, vField.Elem())
	default:
		servUriType := reflect.TypeOf(ServiceUri{})
		if vField.Type() != servUriType {
			return "", errors.New(fmt.Sprintf("Type %s is not supported", vField.Type().String()))
		}
		return defVal, nil
	}
	return "", errors.New(fmt.Sprintf("Type %s is not supported", vField.Type().String()))
}