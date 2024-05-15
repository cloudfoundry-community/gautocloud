package main

import (
	"github.com/cloudfoundry-community/gautocloud"
	_ "github.com/cloudfoundry-community/gautocloud/connectors/all"
	. "github.com/cloudfoundry-community/gautocloud/test-utils"
	"html/template"
	"os"
	"path"
	"reflect"
	"runtime"
	"sort"
	"strings"
)

type DocStruct struct {
	Name   string
	Fields []DocField
}
type DocField struct {
	Name    string
	Type    string
	Comment string
}
type DocConnector struct {
	Name        string
	Id          string
	Pkg         string
	TypeName    string
	TypePkg     string
	DocUrl      string
	TypeWrapped string
	SimpleName  string
	Tip         string
	GlobalType  string
	Closeable   bool
	StructGiven DocStruct
}
type DocConnectors []DocConnector
type Doc struct {
	Name        string
	RespondName string
	RespondTags []string
	Connectors  DocConnectors
}
type Summary struct {
	Name string
	Slug string
	Sub  []Summary
}
type GlobalDoc struct {
	Summaries     []Summary
	DocsConnector []Doc
}

func before() {
	os.Setenv("MYSQL_URL", CreateEnvValue(ServiceUrl{
		Type:     "mysql",
		User:     "user",
		Password: "password",
		Port:     3406,
		Target:   "mydb",
	}))
	os.Setenv("POSTGRES_URL", CreateEnvValue(ServiceUrl{
		Type:     "postgres",
		User:     "user",
		Password: "password",
		Port:     5532,
		Target:   "mydb",
		Options:  "sslmode=disable",
	}))
	os.Setenv("MSSQL_URL", CreateEnvValue(ServiceUrl{
		Type:     "sqlserver",
		User:     "sa",
		Password: "password",
		Port:     1433,
		Target:   "test",
	}))
	os.Setenv("SSO_TOKEN_URI", "http://localhost/tokenUri")
	os.Setenv("SSO_AUTH_URI", "http://localhost/authUri")
	os.Setenv("SSO_USER_INFO_URI", "http://localhost/userInfo")
	os.Setenv("SSO_CLIENT_ID", "myId")
	os.Setenv("SSO_CLIENT_SECRET", "mySecret")
	os.Setenv("SSO_GRANT_TYPE", "grant1,grant2")
	os.Setenv("SSO_SCOPES", "scope1,scope2")
	os.Setenv("MONGODB_URL", CreateEnvValue(ServiceUrl{
		Type:   "mongo",
		Port:   27017,
		Target: "test",
	}))
	os.Setenv("ORACLE_URL", CreateEnvValue(ServiceUrl{
		Type:   "oci",
		Port:   27017,
		Target: "test",
	}))
	os.Setenv("REDIS_URL", CreateEnvValue(ServiceUrl{
		Type:     "redis",
		User:     "redis",
		Password: "redis",
		Port:     6379,
	}))
	os.Setenv("AMQP_URL", CreateEnvValue(ServiceUrl{
		Type:     "amqp",
		User:     "user",
		Password: "password",
		Port:     5672,
	}))
	os.Setenv("SMTP_URL", CreateEnvValue(ServiceUrl{
		Type: "smtp",
		Port: 587,
	}))
	os.Setenv("S3_URL", CreateEnvValue(ServiceUrl{
		Type:     "s3",
		User:     "accessKey1",
		Password: "verySecretKey1",
		Port:     8090,
		Target:   "bucket",
	}))
	gautocloud.ReloadConnectors()
}
func main() {
	before()
	_, currentPath, _, _ := runtime.Caller(0)
	dir := path.Dir(currentPath)
	tmplPath := path.Join(dir, "doc-connector.tmpl.md")
	tmplDocPath := path.Join(dir, "doc.tmpl.md")
	tmpl, err := template.ParseFiles(tmplDocPath, tmplPath)
	fatalIf(err)
	docs := getDocMap()
	mk := make([]string, len(docs))
	i := 0
	for k := range docs {
		mk[i] = k
		i++
	}
	sort.Strings(mk)
	summaries := make([]Summary, len(docs))
	docSlice := make([]Doc, len(docs))
	for index, key := range mk {
		docSlice[index] = docs[key]
		summaries[index] = Summary{
			Name: docs[key].Name,
			Slug: toSlug(docs[key].Name),
			Sub:  make([]Summary, 0),
		}
		for _, docConn := range docs[key].Connectors {
			summary := summaries[index]
			summary.Sub = append(summary.Sub, Summary{
				Name: docConn.Name,
				Slug: toSlug(docConn.Name),
				Sub:  make([]Summary, 0),
			})
			summaries[index] = summary
		}
	}
	gd := GlobalDoc{
		DocsConnector: docSlice,
		Summaries:     summaries,
	}
	err = tmpl.Execute(os.Stdout, gd)
	fatalIf(err)

}
func toSlug(name string) string {
	return strings.ToLower(strings.Replace(name, " ", "-", -1))
}
func getDocMap() map[string]Doc {
	docs := make(map[string]Doc)
	for id, conn := range gautocloud.Connectors() {
		rootIdSplit := strings.Split(id, ":")
		rootId := rootIdSplit[len(rootIdSplit)-1]
		if _, ok := docs[rootId]; !ok {
			docs[rootId] = Doc{
				Name:        strings.Title(rootId),
				RespondName: conn.Name(),
				RespondTags: conn.Tags(),
				Connectors:  make([]DocConnector, 0),
			}
		}
		connType := reflect.TypeOf(conn).Elem()
		store := gautocloud.Store()[id]
		if len(store) == 0 {
			continue
		}
		givenData := store[0].Data
		givenType := reflect.TypeOf(givenData)
		giventTypePkg := givenType
		if giventTypePkg.Kind() == reflect.Ptr {
			giventTypePkg = giventTypePkg.Elem()
		}
		docUrl := generateDocUrl(giventTypePkg.PkgPath())
		tip := path.Dir(connType.PkgPath())
		connName := path.Base(connType.PkgPath())
		if connName == rootId {
			connName = path.Base(path.Dir(connType.PkgPath()))
		}
		givenValue := reflect.ValueOf(givenData)
		if givenType.Kind() == reflect.Ptr {
			givenValue = givenValue.Elem()
		}
		typeWrapped := ""
		if givenValue.NumField() == 1 {
			typeShortSplit := strings.Split(givenValue.Field(0).Type().String(), ".")
			typeShort := typeShortSplit[len(typeShortSplit)-1]
			if len(typeShortSplit) > 1 && typeShort == givenValue.Type().Field(0).Name {
				typeField := givenValue.Field(0).Type()
				if typeField.Kind() == reflect.Ptr {
					typeField = typeField.Elem()
				}
				if strings.Contains(giventTypePkg.PkgPath(), "gautocloud") {
					docUrl = generateDocUrl(typeField.PkgPath())
					typeWrapped = givenValue.Field(0).Type().String()
				}
			}
		}
		pkgConn := connType.PkgPath()
		if connName == "raw" {
			pkgConn = path.Dir(pkgConn)
		}
		globalType := ""
		pkgSplit := strings.Split(connType.PkgPath(), "/")
		if len(pkgSplit) >= 5 {
			globalType = pkgSplit[4]
		}
		connDoc := DocConnector{
			Name:        strings.Title(rootId + " - " + connName),
			Id:          conn.Id(),
			Pkg:         pkgConn,
			TypeName:    givenType.String(),
			TypePkg:     giventTypePkg.PkgPath(),
			TypeWrapped: typeWrapped,
			DocUrl:      docUrl,
			Tip:         tip,
			GlobalType:  strings.Title(globalType),
			SimpleName:  strings.Title(connName),
			StructGiven: generateDocStruct(givenData),
			Closeable:   isCloseable(givenData),
		}
		doc := docs[rootId]
		connectors := append(doc.Connectors, connDoc)
		sort.Sort(connectors)
		doc.Connectors = connectors
		docs[rootId] = doc
	}
	return docs
}
func generateDocUrl(pkgPath string) string {
	splitGivenPkg := strings.Split(pkgPath, "/")
	if pkgPath == "" {
		return ""
	}
	if !strings.Contains(splitGivenPkg[0], ".") {
		return "https://golang.org/pkg/" + pkgPath
	}
	if len(splitGivenPkg) >= 3 && splitGivenPkg[2] == "gautocloud" {
		return ""
	}
	maxLen := 3
	if len(splitGivenPkg) < 3 {
		maxLen = len(splitGivenPkg)
	}
	return "https://" + strings.Join(splitGivenPkg[0:maxLen], "/")
}
func isCloseable(data interface{}) bool {
	givenType := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	resp := v.MethodByName("Close")
	if resp != (reflect.Value{}) {
		return true
	}
	if givenType.Kind() == reflect.Ptr {
		v = v.Elem()
	} else {
		return false
	}
	resp = v.MethodByName("Close")
	if resp == (reflect.Value{}) {
		return false
	}
	return true
}
func generateDocStruct(data interface{}) DocStruct {
	givenType := reflect.TypeOf(data)
	nameSplit := strings.Split(givenType.String(), ".")

	structGiven := DocStruct{
		Name: nameSplit[len(nameSplit)-1],
	}
	fields := make([]DocField, 0)
	v := reflect.ValueOf(data)
	if givenType.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for index := 0; index < v.NumField(); index++ {
		vField := v.Field(index)
		tField := t.Field(index)
		docType := vField.Type()
		if docType.Kind() == reflect.Ptr || docType.Kind() == reflect.Slice {
			docType = docType.Elem()
		}
		docUrl := generateDocUrl(docType.PkgPath())
		if docUrl != "" {
			docUrl = "See doc: " + docUrl
		}
		fields = append(fields, DocField{
			Name:    tField.Name,
			Type:    vField.Type().String(),
			Comment: docUrl,
		})
	}
	structGiven.Fields = fields
	return structGiven
}
func fatalIf(err error) {
	if err == nil {
		return
	}
	panic(err)
}
func (d DocConnectors) Len() int {
	return len(d)
}
func (d DocConnectors) Less(i, j int) bool {
	return d[i].Name < d[j].Name
}
func (d DocConnectors) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
