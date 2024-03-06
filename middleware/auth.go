package middleware

import (
	"encoding/json"
	"net/http"
)

type BasicAuthMiddleware struct {
	BasicAuth string
}

func (b *BasicAuthMiddleware) Validate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		basicAuth := r.Header.Get("Authorization")
		if b.BasicAuth != basicAuth {
			unAtuhroizedMap := map[string]string{
				"message": "Unauthroized",
			}

			resp, _ := json.Marshal(unAtuhroizedMap)

			w.WriteHeader(http.StatusUnauthorized)
			w.Write(resp)
			return
		}

		next.ServeHTTP(w, r)
	})
}
