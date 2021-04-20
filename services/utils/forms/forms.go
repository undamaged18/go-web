package forms

import (
	"github.com/gorilla/schema"
	"net/http"
)

// Transfer data from PostForm to user struct
func Decoder(dest interface{}, source map[string][]string, ignoreUnknownKeys bool) error {
	// Use Gorilla/Schema to parse the PostForm into a Struct
	reader := schema.NewDecoder()
	// Ignore any additional fields passed in the PostForm
	// This is done to avoid unnecessary panics
	reader.IgnoreUnknownKeys(ignoreUnknownKeys)
	if err := reader.Decode(dest, source); err != nil {
		return err
	}
	return nil
}

// If form has been submitted (HTTP POST) process the form
// Parse the HTTP form to ensure is valid
func Parse(r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	return nil
}
