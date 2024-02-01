// timeapi/timeapi.go
package timeapi

import (
	"encoding/json"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time  string `json:"time,omitempty"`
	Error string `json:"error,omitempty"`
}

func GetTime(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	if location == "" {
		writeErrorResponse(w, "Missing 'location' query parameter", http.StatusBadRequest)
		return
	}

	loc, err := time.LoadLocation(location)
	if err != nil {
		writeErrorResponse(w, "Invalid timezone", http.StatusBadRequest)
		return
	}

	now := time.Now().In(loc)
	response := TimeResponse{Time: now.Format(time.RFC3339)}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func writeErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	response := TimeResponse{Error: message}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
