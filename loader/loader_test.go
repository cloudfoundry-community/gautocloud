package loader_test

import (
	. "github.com/cloudfoundry-community/gautocloud/loader"

	"bytes"
	"github.com/cloudfoundry-community/gautocloud/cloudenv"
	fakecloud "github.com/cloudfoundry-community/gautocloud/cloudenv/fake"
	"github.com/cloudfoundry-community/gautocloud/connectors"
	fakecon "github.com/cloudfoundry-community/gautocloud/connectors/fake"
	"github.com/cloudfoundry-community/gautocloud/decoder"
	ldlogger "github.com/cloudfoundry-community/gautocloud/logger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"reflect"
)

type FakeSchema struct {
	Uri      decoder.ServiceUri `cloud:"ur(i|l),regex"`
	Host     string             `cloud:"host.*,regex"`
	Username string             `cloud:"user.*,regex"`
	Password string             `cloud:"pass.*,regex"`
	Port     int
}
type SecondFakeSchema struct {
	Uri      decoder.ServiceUri `cloud:"ur(i|l),regex"`
	Host     string             `cloud:"host.*,regex"`
	Username string             `cloud:"user.*,regex"`
	Password string             `cloud:"pass.*,regex"`
	Port     int
}

var defaultServices []cloudenv.Service = []cloudenv.Service{
	{
		Credentials: map[string]interface{}{
			"uri": "postgres://seilbmbd:PHxTPJSbkcDakfK4cYwXHiIX9Q8p5Bxn@babar.elephantsql.com:5432/seilbmbd",
		},
	},
	{
		Credentials: map[string]interface{}{
			"hostname": "smtp.sendgrid.net",
			"port":     25,
			"username": "QvsXMbJ3rK",
			"password": "HCHMOYluTv",
		},
	},
}
var srv1Expected FakeSchema = FakeSchema{
	Uri: decoder.ServiceUri{
		Scheme:   "postgres",
		Username: "seilbmbd",
		Password: "PHxTPJSbkcDakfK4cYwXHiIX9Q8p5Bxn",
		Host:     "babar.elephantsql.com",
		Port:     5432,
		Query:    make([]decoder.QueryUri, 0),
		Name:     "seilbmbd",
	},
}
var srv2Expected FakeSchema = FakeSchema{
	Host:     "smtp.sendgrid.net",
	Port:     25,
	Username: "QvsXMbJ3rK",
	Password: "HCHMOYluTv",
}
var _ = Describe("Loader", func() {
	var fakeCloudEnv cloudenv.CloudEnv
	var loader Loader
	logBuf := new(bytes.Buffer)
	logger := log.New(logBuf, "", log.Ldate|log.Ltime)
	BeforeEach(func() {
		fakeCloudEnv = fakecloud.NewFakeCloudEnv()
		fakeCloudEnv.(*fakecloud.FakeCloudEnv).SetServices(defaultServices)
		loader = NewLoader([]cloudenv.CloudEnv{fakeCloudEnv})
		loader.SetLogger(logger, ldlogger.Lall)
	})
	AfterEach(func() {
		logBuf.Reset()
	})
	Context("RegisterConnector", func() {
		It("should log an info and not load connector if not in a cloud environment", func() {
			fakeCloudEnv.(*fakecloud.FakeCloudEnv).SetInCloudEnv(false)
			loader.RegisterConnector(fakecon.NewFakeConnector(FakeSchema{}))
			Expect(logBuf.String()).Should(ContainSubstring("Skipping loading connector"))
			Expect(loader.Connectors()).Should(HaveLen(0))
		})
		It("should register and load connector if in a cloud environment", func() {
			loader.RegisterConnector(fakecon.NewFakeConnector(FakeSchema{}))
			Expect(loader.Connectors()).Should(HaveLen(1))
		})
	})
	Context("CleanConnectors", func() {
		It("should remove all registered connectors", func() {
			loader.RegisterConnector(fakecon.NewFakeConnector(FakeSchema{}))
			Expect(loader.Connectors()).To(HaveLen(1))
			loader.CleanConnectors()
			Expect(loader.Connectors()).Should(HaveLen(0))
		})
	})
	Context("ReloadConnectors", func() {
		It("should log an info and not reload connectors if not in a cloud environment", func() {
			fakeCloudEnv.(*fakecloud.FakeCloudEnv).SetInCloudEnv(false)
			loader.ReloadConnectors()
			Expect(logBuf.String()).Should(ContainSubstring("Skipping reloading connector"))
		})
	})
	Context("GetFirst", func() {
		var connector connectors.Connector
		BeforeEach(func() {
			connector = fakecon.NewFakeConnector(FakeSchema{})
			loader.RegisterConnector(connector)
		})
		Context("connector give a structure", func() {
			It("should return a correct content given by connector", func() {
				data, err := loader.GetFirst(connector.Id())
				Expect(err).ToNot(HaveOccurred())
				Expect(data).Should(BeEquivalentTo(srv1Expected))
			})
		})
		Context("connector give a pointer to a structure", func() {
			BeforeEach(func() {
				loader.CleanConnectors()
				connector = fakecon.NewFakePtrConnector(FakeSchema{})
				loader.RegisterConnector(connector)
			})
			It("should return a correct content given by connector", func() {
				data, err := loader.GetFirst(connector.Id())
				Expect(err).ToNot(HaveOccurred())
				Expect(data).Should(BeEquivalentTo(&srv1Expected))
				Expect(reflect.TypeOf(data).Kind()).Should(Equal(reflect.Ptr))
			})
		})
		It("should return an error if no content cannot be given", func() {
			fakeCloudEnv.(*fakecloud.FakeCloudEnv).SetServices(make([]cloudenv.Service, 0))
			loader.ReloadConnectors()
			data, err := loader.GetFirst(connector.Id())
			Expect(data).To(BeNil())
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).Should(ContainSubstring("No content have been given by connector with id"))
		})
		It("should return an error if no connector exists with this id", func() {
			loader.ReloadConnectors()
			data, err := loader.GetFirst("nonvalidconnector")
			Expect(data).To(BeNil())
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).Should(ContainSubstring("not found"))
		})

	})
	Context("GetAll", func() {
		var connector connectors.Connector
		BeforeEach(func() {
			loader.CleanConnectors()
			connector = fakecon.NewFakeConnector(FakeSchema{})
			loader.RegisterConnector(connector)
		})
		Context("connector give a structure", func() {
			It("should return a correct content given by connector", func() {
				data, err := loader.GetAll(connector.Id())
				Expect(err).ToNot(HaveOccurred())
				Expect(data).Should(BeEquivalentTo([]interface{}{srv1Expected, srv2Expected}))
			})
		})
		Context("connector give a pointer to a structure", func() {
			BeforeEach(func() {
				loader.CleanConnectors()
				connector = fakecon.NewFakePtrConnector(FakeSchema{})
				loader.RegisterConnector(connector)
			})
			It("should return a correct content given by connector", func() {
				data, err := loader.GetAll(connector.Id())
				Expect(err).ToNot(HaveOccurred())
				Expect(data).Should(BeEquivalentTo([]interface{}{&srv1Expected, &srv2Expected}))
				Expect(reflect.TypeOf(data).Kind()).Should(Equal(reflect.Slice))
				Expect(reflect.TypeOf(data[0]).Kind()).Should(Equal(reflect.Ptr))
				Expect(reflect.TypeOf(data[1]).Kind()).Should(Equal(reflect.Ptr))
			})
		})
		It("should return an empty slice if no content cannot be given", func() {
			fakeCloudEnv.(*fakecloud.FakeCloudEnv).SetServices(make([]cloudenv.Service, 0))
			loader.ReloadConnectors()
			data, err := loader.GetAll(connector.Id())
			Expect(err).ToNot(HaveOccurred())
			Expect(data).Should(HaveLen(0))
			Expect(reflect.TypeOf(data).Kind()).Should(Equal(reflect.Slice))
		})
		It("should return an error if no connector exists with this id", func() {
			loader.ReloadConnectors()
			data, err := loader.GetAll("nonvalidconnector")
			Expect(data).To(BeNil())
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).Should(ContainSubstring("not found"))
		})

	})

	Context("Inject", func() {
		var connector connectors.Connector
		BeforeEach(func() {
			connector = fakecon.NewFakeConnector(FakeSchema{})
			loader.RegisterConnector(connector)
		})
		Context("connector give a structure", func() {
			It("should inject the correct content given by connector when asking structure", func() {
				var data FakeSchema
				err := loader.Inject(&data)
				Expect(err).ToNot(HaveOccurred())
				Expect(data).Should(BeEquivalentTo(srv1Expected))
			})
			It("should inject the correct content given by connector when asking slice of structure", func() {
				var data []FakeSchema
				err := loader.Inject(&data)
				Expect(err).ToNot(HaveOccurred())
				Expect(data).Should(BeEquivalentTo([]FakeSchema{srv1Expected, srv2Expected}))
			})
		})
		Context("connector give a pointer to a structure", func() {
			BeforeEach(func() {
				loader.CleanConnectors()
				connector = fakecon.NewFakePtrConnector(FakeSchema{})
				loader.RegisterConnector(connector)
			})
			It("should inject the correct content given by connector when asking structure", func() {
				var data *FakeSchema
				err := loader.Inject(&data)
				Expect(err).ToNot(HaveOccurred())
				Expect(data).Should(BeEquivalentTo(&srv1Expected))
			})
			It("should inject the correct content given by connector when asking slice of structure", func() {
				var data []*FakeSchema
				err := loader.Inject(&data)
				Expect(err).ToNot(HaveOccurred())
				Expect(data).Should(BeEquivalentTo([]*FakeSchema{&srv1Expected, &srv2Expected}))
			})
		})
		It("should return an error if content to inject is not a pointer", func() {
			var data FakeSchema
			err := loader.Inject(data)
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).Should(ContainSubstring("You must pass a pointer"))
		})
		It("should return an error if no content with given type can be found", func() {
			var data SecondFakeSchema
			err := loader.Inject(&data)
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).Should(ContainSubstring("Service with the type"))
			Expect(err.Error()).Should(ContainSubstring("cannot be found"))
		})
		It("should return an error if no content can be given", func() {
			fakeCloudEnv.(*fakecloud.FakeCloudEnv).SetServices(make([]cloudenv.Service, 0))
			loader.ReloadConnectors()
			var data FakeSchema
			err := loader.Inject(&data)
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).Should(ContainSubstring("Service with the type"))
			Expect(err.Error()).Should(ContainSubstring("cannot be found"))
		})

	})
	Context("InjectFromId", func() {
		var connector connectors.Connector
		BeforeEach(func() {
			connector = fakecon.NewFakeConnector(FakeSchema{})
			loader.RegisterConnector(connector)
		})
		Context("connector give a structure", func() {
			It("should inject the correct content given by connector when asking structure", func() {
				var data FakeSchema
				err := loader.InjectFromId(connector.Id(), &data)
				Expect(err).ToNot(HaveOccurred())
				Expect(data).Should(BeEquivalentTo(srv1Expected))
			})
			It("should inject the correct content given by connector when asking slice of structure", func() {
				var data []FakeSchema
				err := loader.Inject(&data)
				Expect(err).ToNot(HaveOccurred())
				Expect(data).Should(BeEquivalentTo([]FakeSchema{srv1Expected, srv2Expected}))
			})
		})
		Context("connector give a pointer to a structure", func() {
			BeforeEach(func() {
				loader.CleanConnectors()
				connector = fakecon.NewFakePtrConnector(FakeSchema{})
				loader.RegisterConnector(connector)
			})
			It("should inject the correct content given by connector when asking structure", func() {
				var data *FakeSchema
				err := loader.InjectFromId(connector.Id(), &data)
				Expect(err).ToNot(HaveOccurred())
				Expect(data).Should(BeEquivalentTo(&srv1Expected))
			})
			It("should inject the correct content given by connector when asking slice of structure", func() {
				var data []*FakeSchema
				err := loader.InjectFromId(connector.Id(), &data)
				Expect(err).ToNot(HaveOccurred())
				Expect(data).Should(BeEquivalentTo([]*FakeSchema{&srv1Expected, &srv2Expected}))
			})
		})
		It("should return an error if content to inject is not a pointer", func() {
			var data FakeSchema
			err := loader.InjectFromId(connector.Id(), data)
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).Should(ContainSubstring("You must pass a pointer"))
		})
		It("should return an error if no connector exists with this id", func() {
			var data FakeSchema
			err := loader.InjectFromId("notavalidconnector", data)
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).Should(ContainSubstring("not found"))
		})
		It("should return an error if no content with given type can be found", func() {
			var data SecondFakeSchema
			err := loader.InjectFromId(connector.Id(), &data)
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).Should(ContainSubstring("Connector with id"))
			Expect(err.Error()).Should(ContainSubstring("doesn't give a service with the type"))
		})
		It("should return an error if no content can be given", func() {
			fakeCloudEnv.(*fakecloud.FakeCloudEnv).SetServices(make([]cloudenv.Service, 0))
			loader.ReloadConnectors()
			var data FakeSchema
			err := loader.InjectFromId(connector.Id(), &data)
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).Should(ContainSubstring("Connector with id"))
			Expect(err.Error()).Should(ContainSubstring("doesn't give a service with the type"))
		})

	})
	Context("LoadCloudEnvs", func() {
		It("should call load on cloud env", func() {
			fakeCloudEnv = fakecloud.NewFakeCloudEnv()
			fakeCloudEnv.(*fakecloud.FakeCloudEnv).SetServices(defaultServices)
			loader = NewLoader([]cloudenv.CloudEnv{fakeCloudEnv})
			Expect(fakeCloudEnv.(*fakecloud.FakeCloudEnv).CallLoad()).Should(Equal(1))
		})
	})
	Context("IsInACloudEnv", func() {
		It("should return true if one of CloudEnv detected its environment", func() {
			fakeCloudEnv = fakecloud.NewFakeCloudEnv()
			fakeCloudEnv.(*fakecloud.FakeCloudEnv).SetInCloudEnv(false)
			fakeCloudEnv1 := fakecloud.NewFakeCloudEnv()
			fakeCloudEnv1.(*fakecloud.FakeCloudEnv).SetInCloudEnv(true)
			loader = NewLoader([]cloudenv.CloudEnv{fakeCloudEnv, fakeCloudEnv1})
			Expect(loader.IsInACloudEnv()).Should(BeTrue())
		})
		It("should return false if no CloudEnv detected its environment", func() {
			fakeCloudEnv = fakecloud.NewFakeCloudEnv()
			fakeCloudEnv.(*fakecloud.FakeCloudEnv).SetInCloudEnv(false)
			fakeCloudEnv1 := fakecloud.NewFakeCloudEnv()
			fakeCloudEnv1.(*fakecloud.FakeCloudEnv).SetInCloudEnv(false)
			loader = NewLoader([]cloudenv.CloudEnv{fakeCloudEnv, fakeCloudEnv1})
			Expect(loader.IsInACloudEnv()).Should(BeFalse())
		})
	})
})
