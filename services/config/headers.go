package config

type policy struct {
	Enabled bool `yaml:"enabled"`
	Value string `yaml:"value"`
}

type headers struct {
	CSP policy `yaml:"content-security-policy"`
	HSTS policy `yaml:"hsts"`
	ReferrerPolicy policy `yaml:"referrer-policy"`
	XSSProtection policy `yaml:"x-xss-protection"`
	ContentOptions policy `yaml:"x-content-type-options"`
	FrameOptions policy `yaml:"x-frame-options"`
	PermissionsPolicy policy `yaml:"permissions-policy"`
}
var h headers
const headerFile = "app.headers.yaml"

func Headers() *headers {
	return &h
}