package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/brucedominguez/web-api-healthcheck/driver"
)

// sends payload to JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

type healthcheckResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
	Time    string `json:"time"`
	DBinfo  string `json:"db_information"`
	DBerr   string `json:"error"`
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Add Context
	ctx := context.Background()
	// Adding context to timeout DB connection within 3 second
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)

	currentTime := time.Now().Add(time.Hour * 8) // display UTC + 8 hours (perth time)
	var version string

	defer cancel()

	err := driver.DB.QueryRowContext(ctx, "SELECT version ()").Scan(&version)
	if err != nil {
		switch {
		case err.Error() == "pq: canceling statement due to user request":
			respondWithJSON(w, http.StatusOK, healthcheckResponse{
				http.StatusText(500),
				os.Getenv("VERSION"),
				currentTime.Format("2006-01-02 15:04:05"),
				"DB Unavailable",
				fmt.Sprintf("Failed healthcheck timeout: %v", ctx.Err()),
			})
			return
		default:
			respondWithJSON(w, http.StatusOK, healthcheckResponse{
				http.StatusText(500),
				os.Getenv("VERSION"),
				currentTime.Format("2006-01-02 15:04:05"),
				"DB Unavailable",
				fmt.Sprintf("Failed healthcheck timeout: %v", err),
			})
			return
		}
	}
	healthcheck := healthcheckResponse{
		http.StatusText(200),
		os.Getenv("VERSION"),
		currentTime.Format("2006-01-02 15:04:05"),
		version,
		"",
	}

	respondWithJSON(w, http.StatusOK, healthcheck)
}

