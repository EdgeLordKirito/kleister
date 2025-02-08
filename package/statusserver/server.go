package statusserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// jsonResponse represents the JSON structure for /running
type jsonResponse struct {
	RunningSince time.Time `json:"running_since"`
	Status       bool      `json:"status"`
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

		resp := jsonResponse{
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
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		state.running = false
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Server shutting down"}`))

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
