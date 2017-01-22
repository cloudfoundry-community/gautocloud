package fake

import (
	"github.com/cloudfoundry-community/gautocloud/cloudenv"
)

type FakeCloudEnv struct {
	services   []cloudenv.Service
	inCloudEnv bool
	callLoad   int
	appInfo    cloudenv.AppInfo
}

func NewFakeCloudEnv() cloudenv.CloudEnv {
	return &FakeCloudEnv{
		services: make([]cloudenv.Service, 0),
		inCloudEnv: true,
		callLoad: 0,
	}
}
func (c FakeCloudEnv) Name() string {
	return "fakecloudenv"
}
func (c FakeCloudEnv) GetServicesFromTags(tags []string) ([]cloudenv.Service) {
	return c.services
}
func (c FakeCloudEnv) GetServicesFromName(name string) ([]cloudenv.Service) {
	return c.services
}
func (c *FakeCloudEnv) SetAppInfo(appInfo cloudenv.AppInfo) {
	c.appInfo = appInfo
}
func (c *FakeCloudEnv) SetInCloudEnv(inCloudEnv bool) {
	c.inCloudEnv = inCloudEnv
}
func (c FakeCloudEnv) IsInCloudEnv() bool {
	return c.inCloudEnv
}
func (c *FakeCloudEnv) SetServices(services []cloudenv.Service) {
	c.services = services
}
func (c FakeCloudEnv) Services() []cloudenv.Service {
	return c.services
}
func (c FakeCloudEnv) CallLoad() int {
	return c.callLoad
}
func (c *FakeCloudEnv) Load() error {
	c.callLoad++
	return nil
}
func (c FakeCloudEnv) GetAppInfo() cloudenv.AppInfo {
	return c.appInfo
}