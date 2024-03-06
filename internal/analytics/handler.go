package analytics

import "net/http"

func CountHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	event := r.PathValue("event")
	IncrementEventCount(event, "metadata")
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(http.StatusText(http.StatusNoContent)))
}
