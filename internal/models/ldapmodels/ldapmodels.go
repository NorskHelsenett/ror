package ldapmodels

type LdapConfigs struct {
	Domains []LdapConfig `json:"domains"`
}

type LdapConfig struct {
	Domain       string       `json:"domain"`
	BindUser     string       `json:"bindUser"`
	BindPassword string       `json:"bindPassword"`
	BaseDN       string       `json:"basedn"`
	Development  bool         `json:"development,omitempty"`
	Servers      []LdapServer `json:"servers"`
}

type LdapServer struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
