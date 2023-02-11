// Package faas runs a health check on the application to let the
// hosting service detect if it is up and running.
package faas

import (
	"net/http"
)

// Translate runs as an FaaS on GCP.
func Translate(w http.ResponseWriter, r *http.Request) {
	// rest.TranslateHandler(w, r)
}
