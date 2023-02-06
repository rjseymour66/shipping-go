package faas

import (
	"hello-api/handlers/rest"
	"net/http"
)

// Translate runs as an FaaS on GCp.
func Translate(w http.ResponseWriter, r *http.Request) {
	rest.TranslateHandler(w, r)
}
