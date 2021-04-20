package engine

import (
	"ecommerce/services/engine/functions"
)

func funcMap() map[string]interface{} {
	return map[string]interface{}{
		"asset":    functions.Assets,
		"link":     functions.Links,
		"unescape": functions.Unescape,
		"error":    functions.ErrorMessage,
	}
}

