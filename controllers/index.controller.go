package controllers

import (
	"ecommerce/models"
	"ecommerce/services/engine"
	"github.com/gorilla/csrf"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	page := models.NewPage(r.Context())
	engine.RenderTemplate(w, "index", map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(r),
		"page": page,
	})
}
