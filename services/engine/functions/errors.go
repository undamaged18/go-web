package functions

import "net/http"

// ErrorMessage is used in the HTML templates to return an error message based on the HTTP status code
func ErrorMessage(code int) string {
	switch code {
	case http.StatusOK:
		return "Hmmm... Nothing seems to be the problem? Why are you here?"
	case http.StatusMethodNotAllowed:
		return "Oops... that method is not allowed"
	case http.StatusNotFound:
		return "Oops... this page could not be found"
	case http.StatusInternalServerError:
		return "Oops... something went wrong..."
	default:
		return "Oops... something went wrong..."
	}
}

