package middlewares

import (
	. "github.com/ashwinipatankar/tempalte-go-microservice-ssl/services/common/functions"
	"net/http"
)

// MIDDLEWARES
// Sets Content Type header to application/json
func SetContentTypeJson(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// our middleware logic
		SetResponseContentType(w, "application/json")
		next(w, r)
	}
}

// Checks if valid json body is provided
func ParseJsonBody(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// our middleware logic
		if IsNilReqBody(w, r) {
			return
		}
		if NotJsonReqBody(w, r) {
			return
		}
		next(w, r)
	}
}
