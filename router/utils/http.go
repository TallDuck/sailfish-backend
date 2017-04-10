package utils

import (
	"net/http"
)

// UnauthorizedHTTP returns 403 status to request
func UnauthorizedHTTP(w http.ResponseWriter) {
	http.Error(w, http.StatusText(403), 403)
}
