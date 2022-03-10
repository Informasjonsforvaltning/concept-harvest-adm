package env

import "os"

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func MongoHost() string {
	return getEnv("MONGO_HOST", "localhost:27017")
}

func MongoPassword() string {
	return getEnv("MONGO_PASSWORD", "admin")
}

func MongoUsername() string {
	return getEnv("MONGO_USERNAME", "admin")
}

func KeycloakHost() string {
	return getEnv("KEYCLOAK_HOST", "https://sso.staging.fellesdatakatalog.digdir.no")
}

type Constants struct {
	MongoAuthParams string
	MongoCollection string
	MongoDatabase   string
}

type Paths struct {
	Datasource     string
	Datasources    string
	OrgDatasources string
	Organizations  string
	Ping           string
	Ready          string
}

type Security struct {
	TokenAudience   string
	SysAdminAuth    string
	OrgType         string
	AdminPermission string
}

var ConstantValues = Constants{
	MongoAuthParams: "authSource=admin&authMechanism=SCRAM-SHA-1",
	MongoCollection: "datasources",
	MongoDatabase:   "fdk-harvest-admin",
}

var PathValues = Paths{
	Datasource:     "organizations/:org/datasources/:id",
	Datasources:    "datasources",
	OrgDatasources: "organizations/:org/datasources",
	Organizations:  "organizations",
	Ping:           "ping",
	Ready:          "ready",
}

var SecurityValues = Security{
	TokenAudience:   "fdk-harvest-admin",
	SysAdminAuth:    "system:root:admin",
	OrgType:         "organization",
	AdminPermission: "admin",
}
