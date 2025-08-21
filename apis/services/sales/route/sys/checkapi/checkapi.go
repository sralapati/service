package checkapi

import (
	"context"
	"encoding/json"
	"net/http"
)

func liveness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}

	json.NewEncoder(w).Encode(status)

	return nil
}

func readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}

	json.NewEncoder(w).Encode(status)

	return nil
}
