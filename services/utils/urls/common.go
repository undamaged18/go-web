package urls

import (
	"net/http"
	"net/url"
)

// urls.Next checks the "next" URL query parameter to ensure that it doesn't redirect
//
// away from the website and if attempting to do so results with a redirect to the homepage
func Next(r *http.Request, defaultPath string) string {
	if defaultPath == "" {
		defaultPath = "/"
	}
	// Check if ?next exists and is not set to blank
	if r.URL.Query().Get("next") != "" {
		// Parse the value of next to ensure it is in the correct format
		u, err := url.ParseRequestURI(r.URL.Query().Get("next"))
		if err != nil {
			// If the format is incorrect redirect to home page
			return defaultPath
		} else {
			// If u.Host is set, then an external link has been provided
			// force the redirect to the home page
			if u.Host != "" {
				return defaultPath
			} else {
				// If only a relative path has been provided redirect to the value of next
				return u.Path
			}
		}
	} else {
		return defaultPath
	}
}
