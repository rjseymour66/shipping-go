// Package rest handles requests to the /hello endpoint and
// returns a translated word.
package rest

import (
	"encoding/json"
	"hello-api/translation"
	"net/http"
	"strings"
)

// Resp represents a response from the TranslationHandler.
type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

// TranslateHandler writes a JSON-encoded translation response.
func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	language := r.URL.Query().Get("language")
	if language == "" {
		language = "english"
	}

	word := strings.ReplaceAll(r.URL.Path, "/", "")
	translation := translation.Translate(word, language)
	if translation == "" {
		w.WriteHeader(404)
		return
	}
	resp := Resp{
		Language:    language,
		Translation: translation,
	}

	if err := enc.Encode(resp); err != nil {
		panic("unable to encode response")
	}
}
