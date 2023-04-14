package middlewares

import "net/http"

func validateRoute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {

		wr.WriteHeader(http.StatusOK)
	})
}
