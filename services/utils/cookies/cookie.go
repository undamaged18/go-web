package cookies

import (
	"github.com/gorilla/securecookie"
	"net/http"
	"os"
	"time"
)

// Set - Creates a new HTTP cookie with the provided Name, Value and Expiration time
func Set(w http.ResponseWriter, name string, value string, expires time.Time, maxAge int, secure bool, httpOnly bool) error {
	var s = securecookie.New([]byte(os.Getenv("COOKIE_HASH_KEY")), []byte(os.Getenv("COOKIE_BLOCK_KEY")))
	if encoded, err := s.Encode(name, value); err == nil {
		cookie := &http.Cookie{
			Name:     name,
			Value:    encoded,
			Expires:  expires,
			MaxAge:   maxAge,
			Path:     "/",
			Secure:   secure,
			HttpOnly: httpOnly,
			SameSite: http.SameSiteStrictMode,
		}
		http.SetCookie(w, cookie)
	} else {
		return err
	}
	return nil
}

// Expire any cookie with the provided name
func Expire(w http.ResponseWriter, name string) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    "",
		Expires:  time.Now(),
		MaxAge:   -1,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
}

// MaxAge - Gets the number of seconds until Midnight
// This forces a new login to be required after each day
func MaxAge() int {
	t := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, 0, 0, 0, 0, time.Local)
	currentTime := time.Now()
	difference := t.Sub(currentTime)
	return int(difference.Seconds())
}
