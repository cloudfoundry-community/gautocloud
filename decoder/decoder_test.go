package decoder_test

import (
	. "github.com/cloudfoundry-community/gautocloud/decoder"

	"encoding/json"
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
	FloatJsonNumber   float64
	IntJsonNumber     int
	Ainterface        interface{}
	Amap              map[string]interface{}
	Aslice            []string
	Abool             bool
	Nfloat32          float32
	Nfloat64          float64
	Asubstruct        SubStruct
	Slicesubstruct    []SubStruct
	Npint             *int
	UriDefault        ServiceUri      `cloud:"uri_default" cloud-default:"srv://user:pass@host.com:12/data?options=1"`
	NintDefault       int             `cloud-default:"1"`
	Nint8Default      int8            `cloud-default:"2"`
	Nint16Default     int16           `cloud-default:"3"`
	Nint32Default     int32           `cloud-default:"4"`
	Nint64Default     int64           `cloud-default:"5"`
	NuintDefault      uint            `cloud-default:"6"`
	Nuint8Default     uint8           `cloud-default:"7"`
	Nuint16Default    uint16          `cloud-default:"8"`
	Nuint32Default    uint32          `cloud-default:"9"`
	Nuint64Default    uint64          `cloud-default:"10"`
	AinterfaceDefault interface{}     `cloud-default:"myinterface"`
	AboolDefault      bool            `cloud-default:"true"`
	Nfloat32Default   float32         `cloud-default:"1.1"`
	Nfloat64Default   float64         `cloud-default:"1.2"`
	NpintDefault      *int            `cloud-default:"11"`
	PtrAsString       *SubStructPtr   `cloud-default:"foo"`
	SlicePtrAsString  []*SubStructPtr `cloud-default:"foo"`
}

type SubStructPtr struct {
	Name string
}

func (p *SubStructPtr) UnmarshalCloud(data interface{}) error {
	p.Name = data.(string)
	return nil
}

type SubStruct struct {
	Name        string
	NameDefault string `cloud-default:"myname"`
}

type WrapStruct struct {
	Nint int
	SubStruct
}

type InvalidStruct struct{}

type TestInvalidStruct struct {
	MyStruct InvalidStruct
}
type TestDefaultInvalidStruct struct {
	MyStruct InvalidStruct `cloud-default:"1"`
}

var _ = Describe("Decoder", func() {
	var expectedStruct TestCompleteStruct
	expectDefaultUri := ServiceUri{
		Host:     "host.com",
		Name:     "data",
		Password: "pass",
		Port:     12,
		Query: []QueryUri{{
			Key:   "options",
			Value: "1",
		}},
		RawQuery: "options=1",
		Username: "user",
		Scheme:   "srv",
	}
	pint := 11
	BeforeEach(func() {
		expectedStruct = TestCompleteStruct{
			UriDefault:      expectDefaultUri,
			NintDefault:     1,
			Nint8Default:    int8(2),
			Nint16Default:   int16(3),
			Nint32Default:   int32(4),
			Nint64Default:   int64(5),
			NuintDefault:    uint(6),
			FloatJsonNumber: float64(1.2),
			IntJsonNumber:   2,
			Asubstruct: SubStruct{
				Name:        "subname",
				NameDefault: "myname",
			},
			Slicesubstruct: []SubStruct{
				{
					Name:        "name",
					NameDefault: "myname",
				},
			},
			Amap: map[string]interface{}{
				"name": "name",
			},
			Aslice:            []string{"titi", "toto"},
			Nuint8Default:     uint8(7),
			Nuint16Default:    uint16(8),
			Nuint32Default:    uint32(9),
			Nuint64Default:    uint64(10),
			AinterfaceDefault: "myinterface",
			AboolDefault:      true,
			Nfloat32Default:   float32(1.1),
			Nfloat64Default:   float64(1.2),
			NpintDefault:      &pint,
			PtrAsString: &SubStructPtr{
				Name: "subname",
			},
			SlicePtrAsString: []*SubStructPtr{{
				Name: "subname",
			}},
		}
	})
	It("should decode struct if credentials map type match structure type", func() {
		test := TestCompleteStruct{}
		data := map[string]interface{}{
			"uri":                 "srv://user:pass@host.com:12/data?options=1",
			"myname":              "myservice",
			"nint":                1,
			"nint8":               int8(2),
			"nint16":              int16(3),
			"nint32":              int32(4),
			"nint64":              int64(5),
			"nuint":               uint(6),
			"amap":                map[string]interface{}{"name": "name"},
			"asubstruct":          map[string]interface{}{"name": "subname"},
			"slicesubstruct":      []map[string]interface{}{{"name": "name"}},
			"aslice":              []string{"titi", "toto"},
			"nuint8":              uint8(7),
			"nuint16":             uint16(8),
			"nuint32":             uint32(9),
			"nuint64":             uint64(10),
			"ainterface":          "myinterface",
			"float_json_number":   json.Number("0.12e+1"),
			"int_json_number":     json.Number("2"),
			"abool":               true,
			"nfloat32":            float32(1.1),
			"nfloat64":            float64(1.2),
			"npint":               11,
			"ptr_as_string":       "subname",
			"slice_ptr_as_string": []string{"subname"},
		}
		expectedStruct.Uri = expectDefaultUri
		expectedStruct.Name = "myservice"
		expectedStruct.Nint = 1
		expectedStruct.Nint8 = int8(2)
		expectedStruct.Nint16 = int16(3)
		expectedStruct.Nint32 = int32(4)
		expectedStruct.Nint64 = int64(5)
		expectedStruct.Asubstruct = SubStruct{
			Name:        "subname",
			NameDefault: "myname",
		}
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
		expectedStruct.PtrAsString = &SubStructPtr{
			Name: "subname",
		}
		expectedStruct.SlicePtrAsString = []*SubStructPtr{{
			Name: "subname",
		}}
		err := Unmarshal(data, &test)
		Expect(err).NotTo(HaveOccurred())
		Expect(test).Should(BeEquivalentTo(expectedStruct))
	})
	It("should decode even if credential value is a string but structure expect other type", func() {
		test := TestCompleteStruct{}
		data := map[string]interface{}{
			"uri":                 "srv://user:pass@host.com:12/data?options=1",
			"myname":              "myservice",
			"nint":                "1",
			"nint8":               "2",
			"nint16":              "3",
			"nint32":              "4",
			"nint64":              "5",
			"nuint":               "6",
			"amap":                map[string]interface{}{"name": "name"},
			"slicesubstruct":      []map[string]interface{}{{"name": "name"}},
			"float_json_number":   json.Number("0.12e+1"),
			"int_json_number":     json.Number("2"),
			"aslice":              "titi, toto",
			"nuint8":              "7",
			"nuint16":             "8",
			"nuint32":             "9",
			"nuint64":             "10",
			"ainterface":          "myinterface",
			"abool":               "true",
			"nfloat32":            "1.1",
			"nfloat64":            "1.2",
			"npint":               "11",
			"ptr_as_string":       "subname",
			"slice_ptr_as_string": []string{"subname"},
		}
		expectedStruct.Uri = expectDefaultUri
		expectedStruct.Name = "myservice"
		expectedStruct.Nint = 1
		expectedStruct.Nint8 = int8(2)
		expectedStruct.Nint16 = int16(3)
		expectedStruct.Nint32 = int32(4)
		expectedStruct.Nint64 = int64(5)
		expectedStruct.Asubstruct = SubStruct{
			NameDefault: "myname",
		}
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
		expectedStruct.PtrAsString = &SubStructPtr{
			Name: "subname",
		}
		expectedStruct.SlicePtrAsString = []*SubStructPtr{{
			Name: "subname",
		}}
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
		Expect(err).Should(BeAssignableToTypeOf(ErrDecode{}))
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
	It("should give corresponding value on a number if float is given", func() {
		test := TestCompleteStruct{}
		data := map[string]interface{}{
			"uri":                 "srv://user:pass@host.com:12/data?options=1",
			"myname":              "myservice",
			"nint":                1,
			"nint8":               float32(2),
			"nint16":              float32(3),
			"nint32":              float32(4),
			"nint64":              float32(5),
			"nuint":               float32(6),
			"float_json_number":   json.Number("0.12e+1"),
			"int_json_number":     json.Number("2"),
			"amap":                map[string]interface{}{"name": "name"},
			"asubstruct":          map[string]interface{}{"name": "subname"},
			"slicesubstruct":      []map[string]interface{}{{"name": "name"}},
			"aslice":              []string{"titi", "toto"},
			"nuint8":              float64(7),
			"nuint16":             float32(8),
			"nuint32":             float32(9),
			"nuint64":             float32(10),
			"ainterface":          "myinterface",
			"abool":               true,
			"nfloat32":            float32(1.1),
			"nfloat64":            float64(1.2),
			"npint":               11,
			"ptr_as_string":       "subname",
			"slice_ptr_as_string": []string{"subname"},
		}
		expectedStruct.Uri = expectDefaultUri
		expectedStruct.Name = "myservice"
		expectedStruct.Nint = 1
		expectedStruct.Nint8 = int8(2)
		expectedStruct.Nint16 = int16(3)
		expectedStruct.Nint32 = int32(4)
		expectedStruct.Nint64 = int64(5)
		expectedStruct.Asubstruct = SubStruct{
			Name:        "subname",
			NameDefault: "myname",
		}
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
		expectedStruct.PtrAsString = &SubStructPtr{
			Name: "subname",
		}
		expectedStruct.SlicePtrAsString = []*SubStructPtr{{
			Name: "subname",
		}}
		err := Unmarshal(data, &test)
		Expect(err).NotTo(HaveOccurred())
		Expect(test).Should(BeEquivalentTo(expectedStruct))
	})
	It("should decode struct with no default value when set", func() {
		test := struct {
			Default string `cloud-default:"adefault"`
		}{}
		data := map[string]interface{}{}
		err := UnmarshalNoDefault(data, &test)
		Expect(err).NotTo(HaveOccurred())
		Expect(test.Default).Should(BeEmpty())
	})
	It("should decode map inside sub struct", func() {
		test := struct {
			SubStructs []struct {
				Amap        map[string]string
				AComplexMap map[string]struct {
					Toto string
				}
				AComplexPtrMap map[string]*struct {
					Titi string
				}
			}
		}{}
		data := map[string]interface{}{
			"sub_structs": []map[string]interface{}{
				{
					"amap": map[string]interface{}{
						"name": "name",
					},
					"a_complex_map": map[string]interface{}{
						"avalue": map[string]interface{}{"toto": "toto"},
					},
					"a_complex_ptr_map": map[string]interface{}{
						"avalue": map[string]interface{}{"titi": "titi"},
					},
				},
			},
		}
		err := Unmarshal(data, &test)
		Expect(err).NotTo(HaveOccurred())
		Expect(test.SubStructs[0].Amap).Should(HaveKey("name"))
		Expect(test.SubStructs[0].Amap["name"]).Should(Equal("name"))

		Expect(test.SubStructs[0].AComplexMap).Should(HaveKey("avalue"))
		Expect(test.SubStructs[0].AComplexMap["avalue"].Toto).Should(Equal("toto"))

		Expect(test.SubStructs[0].AComplexPtrMap).Should(HaveKey("avalue"))
		Expect(test.SubStructs[0].AComplexPtrMap["avalue"].Titi).Should(Equal("titi"))
	})
	It("should decode wrapped struct", func() {
		test := WrapStruct{}
		data := map[string]interface{}{
			"name": "wrapname",
			"nint": 1,
		}

		err := Unmarshal(data, &test)
		Expect(err).NotTo(HaveOccurred())
		Expect(test.Name).Should(Equal("wrapname"))
		Expect(test.Nint).Should(Equal(1))
	})
})
