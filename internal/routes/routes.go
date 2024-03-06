package routes

import (
	"encoding/json"
	"fmt"
	"internal-tools/internal/analytics"
	"internal-tools/middleware"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter() http.Handler {

	mux := http.NewServeMux()

	handler := middleware.LoggerMiddleware(mux)

	mux.Handle("/metrics", promhttp.HandlerFor(analytics.Registry, promhttp.HandlerOpts{
		Registry: analytics.Registry,
	}))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.PathValue("/") == r.RequestURI {
			fmt.Printf("r.PathValue : %v \n", r.PathValue("/"))
			bodyByte, _ := json.Marshal(map[string]string{
				"message": "Hello World",
			})

			w.WriteHeader(http.StatusOK)
			w.Write(bodyByte)
			return
		}

		bodyByte, _ := json.Marshal(map[string]string{
			"message": http.StatusText(http.StatusNotFound),
		})

		w.WriteHeader(http.StatusNotFound)
		w.Write(bodyByte)
	})

	mux.HandleFunc("POST /count/{event}", analytics.CountHandler)

	return handler

}
