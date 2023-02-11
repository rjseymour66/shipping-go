// Package translation implements the business logic for translating
// the word "hello" into another language.
package translation

import "strings"

// StaticService has data that does not change.
type StaticService struct{}

// NewStaticService creates a new instance of a service that uses static data
func NewStaticService() *StaticService {
	return &StaticService{}
}

// Translate accepts a word and translates it to the given language.
func (s *StaticService) Translate(word string, language string) string {
	word = sanitizeInput(word)
	language = sanitizeInput(language)

	if word != "hello" {
		return ""
	}
	switch language {
	case "english":
		return "hello"
	case "finnish":
		return "hei"
	case "german":
		return "hallo"
	case "french":
		return "bonjour"
	default:
		return ""
	}
}

func sanitizeInput(w string) string {
	w = strings.ToLower(w)
	return strings.TrimSpace(w)
}
