package config

const linksFile = "app.links.yaml"

type paths struct {
	Routes map[string]string `yaml:"routes"`
	Assets map[string]string `yaml:"assets"`
}
var r paths

func Routes(route string) string {

	return r.Routes[route]
}

func Assets(item string) string {
	return r.Assets[item]
}