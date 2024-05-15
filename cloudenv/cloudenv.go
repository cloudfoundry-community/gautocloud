// Package cloudenv It manages the detection of the environment but also the detections of services asked by a Loader.
package cloudenv

import "regexp"

// CloudEnv You must implement this interface if you want to add another cloud environment
// See the cf_cloudenv or heroku_cloudenv or local_cloudenv to have an example of implementation
type CloudEnv interface {
	// Name The name of your cloud environment (e.g.: cloud foundry, heroku, local ...)
	Name() string
	// GetServicesFromTags The loader will call this function and pass a list of tags
	// You will need to give services which match with those tags
	// Note: tag can be an regex, better to take this in consideration
	GetServicesFromTags(tags []string) []Service
	// GetServicesFromName The loader will call this function and pass a service name as a regex
	// You will need to give services which match with this name
	GetServicesFromName(name string) []Service
	// IsInCloudEnv The loader will call this function to see if this cloud environment can be use
	// This function should detect the targeted environment
	IsInCloudEnv() bool
	// Load The loader will call this method to load the environment
	Load() error
	// GetAppInfo This need to return information about application instance information
	GetAppInfo() AppInfo
}
type AppInfo struct {
	Id         string
	Name       string
	Port       int
	Properties map[string]interface{}
}
type Service struct {
	Credentials map[string]interface{}
}

func match(matcher, content string) bool {
	regex, err := regexp.Compile("(?i)^" + matcher + "$")
	if err != nil {
		return false
	}
	return regex.MatchString(content)
}
