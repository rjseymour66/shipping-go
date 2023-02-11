// Package rest handles requests to the /hello endpoint and
// returns a translated word.
package rest

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Translator is an interface for a struct that implements the Translate
// method that translates a word the given language.
type Translator interface {
	Translate(word string, language string) string
}

// TranslateHandler translates calls for the caller
type TranslateHandler struct {
	service Translator
}

// NewTranslateHandler creates a new instance of the handler using a
// translation service.
func NewTranslateHandler(service Translator) *TranslateHandler {
	return &TranslateHandler{
		service: service,
	}
}

// Resp represents a response from the TranslationHandler.
type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

// TranslateHandler takes a given request with a path value of the
// word to be translated and a query parameter of the language to translate to.
func (t *TranslateHandler) TranslateHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	language := r.URL.Query().Get("language")
	if language == "" {
		language = "english"
	}

	word := strings.ReplaceAll(r.URL.Path, "/", "")
	translation := t.service.Translate(word, language)
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
