package middleware

import (
	"github.com/gorilla/sessions"
	"net/http"
	"time"
)

func session(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var sess *sessions.Session
		var store = sessions.NewCookieStore([]byte(conf.Keys.Session))

		sess, _ = store.Get(r, "sess")

		sess.Options.Secure = true
		sess.Options.HttpOnly = true
		sess.Options.SameSite = http.SameSiteStrictMode
		sess.Options.MaxAge = int(time.Second * 3600)
		err := sess.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		conf.Session = sess
		next.ServeHTTP(w, r)
	})
}
