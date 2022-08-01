package middleware

import "net/http"

func InternalServerError(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("Server error"))
			}
		}()

		handler.ServeHTTP(w, r)
	}
}
