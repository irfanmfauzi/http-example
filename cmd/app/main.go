package main

import (
	"internal-tools/internal/routes"
	"log"
	"log/slog"
	"os"

	"net/http"
)

func main() {
	handler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	mux := routes.NewRouter()

	logger.Info("Storytale Internal-Tooling")

	log.Fatal(http.ListenAndServe(":8000", mux))
}
