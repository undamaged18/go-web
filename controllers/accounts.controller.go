package controllers

import (
	"ecommerce/models"
	"ecommerce/services/engine"
	"ecommerce/services/utils/errors"
	"ecommerce/services/utils/forms"
	"github.com/gorilla/csrf"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	user := models.NewUser()
	if r.Method == http.MethodPost {
		user.Authenticate()
	}
	page := models.NewPage(r.Context())
	engine.RenderTemplate(w, "account:login", map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(r),
		"page": page,
	})
}

func register(w http.ResponseWriter, r *http.Request) {
	user := models.NewUser()
	var formErrors models.FormErrors
	if r.Method == http.MethodPost {
		if err := forms.Parse(r); err != nil {
			fn, line := errors.FuncTrace()
			errors.Panic(http.StatusInternalServerError, fn, line, err)
		}
		if err := forms.Decoder(user, r.PostForm, true); err != nil {
			fn, line := errors.FuncTrace()
			errors.Panic(http.StatusInternalServerError, fn, line, err)
		}
		formErrors = user.Create()
	}
	page := models.NewPage(r.Context())
	engine.RenderTemplate(w, "account:register", map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(r),
		"page": page,
		"user": user,
		"errors": formErrors,
	})
}

func validate() {

}