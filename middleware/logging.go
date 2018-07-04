package middleware

import (
	"github.com/op/go-logging"
	"net/http"
)

//Logger
var log = logging.MustGetLogger("logging")

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Debug(r.RequestURI)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
