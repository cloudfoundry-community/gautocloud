package test_utils

import (
	"os"
	"strconv"
)

type ServiceUrl struct {
	Type     string
	User     string
	Password string
	Port     int
	Target   string
	Options  string
}

func CreateEnvValue(serviceUrl ServiceUrl) string {
	host := os.Getenv("GAUTOCLOUD_HOST_SERVICES")
	envValue := serviceUrl.Type + "://"
	if serviceUrl.User != "" {
		envValue += serviceUrl.User
	}
	if serviceUrl.User != "" && serviceUrl.Password != "" {
		envValue += ":" + serviceUrl.Password
	}
	if serviceUrl.User != "" {
		envValue += "@"
	}
	envValue += host
	if serviceUrl.Port != 0 {
		envValue += ":" + strconv.Itoa(serviceUrl.Port)
	}
	if serviceUrl.Target != "" {
		envValue += "/" + serviceUrl.Target
	}
	if serviceUrl.Options != "" {
		envValue += "?" + serviceUrl.Options
	}
	return envValue
}
