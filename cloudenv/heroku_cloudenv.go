package cloudenv

import (
	"os"
	"strings"
	"strconv"
	"net"
)

type HerokuCloudEnv struct {
	envVars []EnvVar
}
type EnvVar struct {
	Key   string
	Value string
}

const ENV_KEY_APP_NAME string = "GAUTOCLOUD_APP_NAME"

func NewHerokuCloudEnv() CloudEnv {
	return &HerokuCloudEnv{}

}
func NewHerokuCloudEnvEnvironment(environ []string) CloudEnv {
	cloudEnv := &HerokuCloudEnv{}
	cloudEnv.InitEnv(environ)
	return cloudEnv

}
func (c *HerokuCloudEnv) Load() error {
	c.InitEnv(os.Environ())
	return nil
}
func (c *HerokuCloudEnv) InitEnv(environ []string) {
	envVars := make([]EnvVar, 0)
	for _, envVar := range environ {
		splitEnv := strings.Split(envVar, "=")
		envVars = append(envVars, EnvVar{
			Key: strings.ToLower(splitEnv[0]),
			Value: strings.Join(splitEnv[1:], "="),
		})
	}
	c.envVars = envVars
}
func (c HerokuCloudEnv) GetServicesFromTags(tags []string) ([]Service) {
	services := make([]Service, 0)
	for _, tag := range tags {
		services = append(services, c.getServicesFromPrefix(tag)...)
	}
	return services
}
func (c HerokuCloudEnv) GetServicesFromName(name string) ([]Service) {
	return c.getServicesFromPrefix(name)
}
func (c HerokuCloudEnv) getServicesFromPrefix(prefix string) []Service {
	services := make(map[string]Service)
	for _, envVar := range c.envVars {
		splitKey := c.splitKey(envVar.Key)
		posKey := c.findPosInKey(strings.ToLower(prefix), splitKey)
		if posKey == -1 {
			continue
		}
		toSplitPos := posKey + 1
		name := splitKey[0]
		if len(splitKey) > 1 {
			name = strings.Join(splitKey[0:toSplitPos], "_")
		}
		if _, ok := services[name]; !ok {
			services[name] = Service{
				Credentials: make(map[string]interface{}),
			}
		}
		if len(splitKey) == 1 {
			services[name].Credentials[splitKey[0]] = c.extractCredValue(splitKey, envVar.Value)
			services[name].Credentials["uri"] = c.extractCredValue(splitKey, envVar.Value)
			continue
		}
		services[name].Credentials[splitKey[len(splitKey) - 1]] = c.extractCredValue(splitKey[toSplitPos:], envVar.Value)

	}
	sliceServices := make([]Service, 0)
	for _, service := range services {
		sliceServices = append(sliceServices, service)
	}
	return sliceServices
}
func (c HerokuCloudEnv) extractCredValue(splitKey []string, value string) interface{} {
	if len(splitKey) > 1 {
		return c.extractCredValue(splitKey[1:], value)
	}
	return value
}
func (c HerokuCloudEnv) splitKey(key string) []string {
	return strings.Split(key, "_")
}
func (c HerokuCloudEnv) findPosInKey(matcher string, splitKey []string) int {
	splitMatcher := c.splitKey(matcher)
	for index, key := range splitKey {
		if len(splitMatcher) == 1 && match(splitMatcher[0], key) {
			return index
		}
		nextIndex := index + 1
		if len(splitKey) > nextIndex && match(splitMatcher[0], key) {
			pos := c.findPosInKey(strings.Join(splitMatcher[1:], "_"), splitKey[nextIndex:])
			if pos == -1 {
				return -1
			}
			return pos + 1
		}
	}
	return -1
}
func (c HerokuCloudEnv) Name() string {
	return "heroku"
}

func (c HerokuCloudEnv) IsInCloudEnv() bool {
	_, isIn := os.LookupEnv("DYNO")
	return isIn
}
func (c HerokuCloudEnv) externalIP() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, iface := range ifaces {
		if iface.Flags & net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags & net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return ""
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String()
		}
	}
	return ""
}
func (c HerokuCloudEnv) getEnvValueName(key string) string {
	for _, envVar := range c.envVars {
		if envVar.Key == strings.ToLower(key) {
			return envVar.Value
		}
	}
	return ""
}
func (c HerokuCloudEnv) GetAppInfo() AppInfo {
	name := c.getEnvValueName(ENV_KEY_APP_NAME)
	if name == "" {
		name = "<unknown>"
	}
	port, _ := strconv.Atoi(c.getEnvValueName("PORT"))
	return AppInfo{
		Id: c.getEnvValueName("DYNO"),
		Name: name,
		Properties: map[string]interface{}{
			"port": port,
			"host": c.externalIP(),
		},
	}
}