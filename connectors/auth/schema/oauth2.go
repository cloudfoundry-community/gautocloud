package schema

type Oauth2Schema struct {
	AuthorizationUri string `cloud:".*auth.*,regex"`
	UserInfoUri      string `cloud:".*info.*,regex"`
	TokenUri         string `cloud:".*token.*,regex"`
	ClientId         string `cloud:".*id.*,regex"`
	ClientSecret     string `cloud:".*secret.*,regex"`
	GrantTypes       []string `cloud:".*grant.*,regex"`
	Scopes           []string `cloud:".*scope.*,regex"`
}
