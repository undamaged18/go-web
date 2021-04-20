package config

import (
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/acme/autocert"
	"html/template"
)

const configFile = "app.config.yaml"

type config struct {
	Environment string `yaml:"environment"`
	ServerName  string `yaml:"server_name"`
	Ports       struct {
		HTTP  string `yaml:"http"`
		HTTPS string `yaml:"https"`
	} `yaml:"ports"`
	TLS struct {
		Cert string `yaml:"cert"`
		Key  string `yaml:"key"`
	} `yaml:"tls"`
	Keys struct {
		JWT string `yaml:"jwt"`
		CSRF string `yaml:"csrf"`
		Session string `yaml:"session"`
	} `yaml:"keys"`
	Paths struct {
		Root      string `yaml:"dist"`
		Templates string `yaml:"templates"`
		Logs      string `yaml:"logs"`
	} `yaml:"paths"`
	CertManager *autocert.Manager             `yaml:"-"`
	Templates   map[string]*template.Template `yaml:"-"`
	Session     *sessions.Session             `yaml:"-"`
}

var conf *config

func New() *config {
	return conf
}
