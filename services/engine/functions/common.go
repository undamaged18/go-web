package functions

import (
	"ecommerce/services/config"
	"html/template"
)

// Assets is used in the HTML templates to return the CSS/JS assets commonly used
func Assets(item string) string {
	return config.Assets(item)
}

// Links is used in the HTML templates to return the URL link
func Links(item string) string {
	return config.Routes(item)
}

func Unescape(text string) template.HTML {
	return template.HTML(text)
}
