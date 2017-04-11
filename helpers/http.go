package helpers

import (
	"net/http"
)

// UnauthorizedHTTP returns 403 status to request
func UnauthorizedHTTP(w http.ResponseWriter) {
	http.Error(w, http.StatusText(403), 403)
}

// ErrorHTTP return 500 status to request
func ErrorHTTP(w http.ResponseWriter) {
	http.Error(w, http.StatusText(500), 500)
}
