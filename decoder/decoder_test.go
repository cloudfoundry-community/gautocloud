package decoder_test

import (
	. "github.com/cloudfoundry-community/gautocloud/decoder"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type TestCompleteStruct struct {
	Uri               ServiceUri
	Name              string `cloud:".*name.*,regex"`
	Nint              int
	Nint8             int8
	Nint16            int16
	Nint32            int32
	Nint64            int64
	Nuint             uint
	Nuint8            uint8
	Nuint16           uint16
	Nuint32           uint32
	Nuint64           uint64
	Ainterface        interface{}
	Aslice            []string
	Abool             bool
	Nfloat32          float32
	Nfloat64          float64
	Npint             *int
	UriDefault        ServiceUri `cloud:"uri_default,default=srv://user:pass@host.com:12/data?options=1"`
	NintDefault       int `cloud:",default=1"`
	Nint8Default      int8 `cloud:",default=2"`
	Nint16Default     int16 `cloud:",default=3"`
	Nint32Default     int32 `cloud:",default=4"`
	Nint64Default     int64 `cloud:",default=5"`
	NuintDefault      uint `cloud:",default=6"`
	Nuint8Default     uint8 `cloud:",default=7"`
	Nuint16Default    uint16 `cloud:",default=8"`
	Nuint32Default    uint32 `cloud:",default=9"`
	Nuint64Default    uint64 `cloud:",default=10"`
	AinterfaceDefault interface{} `cloud:",default=myinterface"`
	AboolDefault      bool `cloud:",default=true"`
	Nfloat32Default   float32 `cloud:",default=1.1"`
	Nfloat64Default   float64 `cloud:",default=1.2"`
	NpintDefault      *int `cloud:",default=11"`
}
type InvalidStruct struct{}
type TestInvalidStruct struct {
	MyStruct InvalidStruct
}
type TestDefaultInvalidStruct struct {
	MyStruct InvalidStruct `cloud:",default=1"`
}

var _ = Describe("Decoder", func() {
	var expectedStruct TestCompleteStruct
	expectDefaultUri := ServiceUri{
		Host: "host.com",
		Name: "data",
		Password: "pass",
		Port: 12,
		Query: []QueryUri{QueryUri{
			Key: "options",
			Value: "1",
		}},
		RawQuery: "options=1",
		Username: "user",
		Scheme: "srv",
	}
	pint := 11
	BeforeEach(func() {
		expectedStruct = TestCompleteStruct{
			UriDefault: expectDefaultUri,
			NintDefault: 1,
			Nint8Default: int8(2),
			Nint16Default: int16(3),
			Nint32Default: int32(4),
			Nint64Default: int64(5),
			NuintDefault: uint(6),
			Aslice: []string{"titi", "toto"},
			Nuint8Default: uint8(7),
			Nuint16Default: uint16(8),
			Nuint32Default: uint32(9),
			Nuint64Default: uint64(10),
			AinterfaceDefault: "myinterface",
			AboolDefault: true,
			Nfloat32Default: float32(1.1),
			Nfloat64Default: float64(1.2),
			NpintDefault: &pint,
		}
	})
	It("should decode struct if credentials map type match structure type", func() {
		test := TestCompleteStruct{}
		data := map[string]interface{}{
			"uri": "srv://user:pass@host.com:12/data?options=1",
			"myname": "myservice",
			"nint": 1,
			"nint8": int8(2),
			"nint16": int16(3),
			"nint32": int32(4),
			"nint64": int64(5),
			"nuint": uint(6),
			"aslice": []string{"titi", "toto"},
			"nuint8": uint8(7),
			"nuint16": uint16(8),
			"nuint32": uint32(9),
			"nuint64": uint64(10),
			"ainterface": "myinterface",
			"abool": true,
			"nfloat32": float32(1.1),
			"nfloat64": float64(1.2),
			"npint": 11,
		}
		expectedStruct.Uri = expectDefaultUri
		expectedStruct.Name = "myservice"
		expectedStruct.Nint = 1
		expectedStruct.Nint8 = int8(2)
		expectedStruct.Nint16 = int16(3)
		expectedStruct.Nint32 = int32(4)
		expectedStruct.Nint64 = int64(5)
		expectedStruct.Aslice = []string{"titi", "toto"}
		expectedStruct.Nuint = uint(6)
		expectedStruct.Nuint8 = uint8(7)
		expectedStruct.Nuint16 = uint16(8)
		expectedStruct.Nuint32 = uint32(9)
		expectedStruct.Nuint64 = uint64(10)
		expectedStruct.Ainterface = "myinterface"
		expectedStruct.Abool = true
		expectedStruct.Nfloat32 = float32(1.1)
		expectedStruct.Nfloat64 = float64(1.2)
		expectedStruct.Npint = &pint
		err := Unmarshal(data, &test)
		Expect(err).NotTo(HaveOccurred())
		Expect(test).Should(BeEquivalentTo(expectedStruct))
	})
	It("should decode even if credential value is a string but structure expect other type", func() {
		test := TestCompleteStruct{}
		data := map[string]interface{}{
			"uri": "srv://user:pass@host.com:12/data?options=1",
			"myname": "myservice",
			"nint": "1",
			"nint8": "2",
			"nint16": "3",
			"nint32": "4",
			"nint64": "5",
			"nuint": "6",
			"aslice": "titi, toto",
			"nuint8": "7",
			"nuint16": "8",
			"nuint32": "9",
			"nuint64": "10",
			"ainterface": "myinterface",
			"abool": "true",
			"nfloat32": "1.1",
			"nfloat64": "1.2",
			"npint": "11",
		}
		expectedStruct.Uri = expectDefaultUri
		expectedStruct.Name = "myservice"
		expectedStruct.Nint = 1
		expectedStruct.Nint8 = int8(2)
		expectedStruct.Nint16 = int16(3)
		expectedStruct.Nint32 = int32(4)
		expectedStruct.Nint64 = int64(5)
		expectedStruct.Nuint = uint(6)
		expectedStruct.Nuint8 = uint8(7)
		expectedStruct.Aslice = []string{"titi", "toto"}
		expectedStruct.Nuint16 = uint16(8)
		expectedStruct.Nuint32 = uint32(9)
		expectedStruct.Nuint64 = uint64(10)
		expectedStruct.Ainterface = "myinterface"
		expectedStruct.Abool = true
		expectedStruct.Nfloat32 = float32(1.1)
		expectedStruct.Nfloat64 = float64(1.2)
		expectedStruct.Npint = &pint
		err := Unmarshal(data, &test)
		Expect(err).NotTo(HaveOccurred())
		Expect(test).Should(BeEquivalentTo(expectedStruct))
	})
	It("should give an error if credentials map type don't match structure type and type is not a string", func() {
		test := TestCompleteStruct{}
		data := map[string]interface{}{
			"myname": 1,
		}
		Expect(func() {
			Unmarshal(data, &test)
		}).Should(Panic())
	})
	It("should give an error if structure have invalid type", func() {
		test := TestInvalidStruct{}
		data := map[string]interface{}{
			"my_struct": 1,
		}
		err := Unmarshal(data, &test)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(ContainSubstring("Error on field 'MyStruct'"))
		Expect(err.Error()).Should(ContainSubstring("is not supported"))
	})
	It("should give an error if structure have invalid type and try to set default value", func() {
		test := TestDefaultInvalidStruct{}
		data := map[string]interface{}{}
		err := Unmarshal(data, &test)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(ContainSubstring("Error on field 'MyStruct' when trying to convert value"))
		Expect(err.Error()).Should(ContainSubstring("is not supported"))
	})

})
