package config

type Ldap struct {
	Url      string `yaml:"url"`
	Dc       string `yaml:"dc"`
	Password string `yaml:"password"`
}
