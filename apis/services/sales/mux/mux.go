// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"encoding/json"
	"net/http"

	"github.com/sralapati/service/foundation/logger"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(log *logger.Logger) *http.ServeMux {
	mux := http.NewServeMux()

	h := func(w http.ResponseWriter, r *http.Request) {
		status := struct {
			Status string `json:"status"`
		}{
			Status: "OK",
		}

		json.NewEncoder(w).Encode(status)
		log.Info(r.Context(), "request", "method", r.Method, "path", r.URL.Path)
	}

	mux.HandleFunc("GET /test", h)

	// Define your routes here
	return mux
}
