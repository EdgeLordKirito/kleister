package statusserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// runningResponse represents the JSON structure for /running
type runningResponse struct {
	RunningSince time.Time `json:"running_since"`
	Status       bool      `json:"status"`
}

type stopResponse struct {
	Message    string `json:"message"`
	Uptime     string `json:"uptime"`
	UptimeSecs int    `json:"uptime_secs"`
}

type errorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// SetupStatusServer starts the HTTP server with context handling
func SetupStatusServer(ctxBundle ContextBundle, settings ServerSettings) {

	defer ctxBundle.Waiter.Done()
	state := &serverState{
		running:      true,
		runningSince: time.Now(),
	}

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    settings.Adress,
		Handler: mux,
	}

	// /running endpoint
	mux.HandleFunc("/running", func(w http.ResponseWriter, r *http.Request) {
		state.mu.Lock()
		defer state.mu.Unlock()

		resp := runningResponse{
			RunningSince: state.runningSince,
			Status:       state.running,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	// /stop endpoint
	mux.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		state.mu.Lock()
		defer state.mu.Unlock()

		if !settings.Auth.Authenticate(r) {
			response := errorResponse{
				Error:   "unauthorized",
				Message: "Authentication failed",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response) // Send JSON error response
			return
		}

		state.running = false
		uptime := time.Since(state.runningSince) // Calculate the uptime

		// Create a response struct
		response := stopResponse{
			Message:    "Server shutting down",
			Uptime:     uptime.String(),
			UptimeSecs: int(uptime.Seconds()),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response) // Encode the response struct as JSON

		fmt.Println("Shutdown triggered via /stop endpoint")
		ctxBundle.CancelFunc()
	})

	// Run the server in a goroutine
	go func() {
		fmt.Println("Status server running on http://" + settings.Adress)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for context cancellation
	<-ctxBundle.Context.Done()
	fmt.Println("Shutting down status server...")
	server.Close()
	fmt.Println("Status server stopped")
	//Maybe not needed
}
