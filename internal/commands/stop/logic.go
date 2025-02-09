package stop

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/EdgeLordKirito/wallpapersetter/package/statusserver"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) error {
	url := "http://" + statusserver.DefaultAdress + "/stop"

	// Create a context with a 3-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Create an HTTP request with the context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Use an HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		// Check for timeout errors
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, os.ErrDeadlineExceeded) {
			fmt.Println("Request timed out. No kleister instance running.")
		} else if netErr, ok := err.(*net.OpError); ok && netErr.Op == "dial" {
			// Check if it's a connection refused error
			fmt.Println("No kleister instance running.")
		} else if strings.Contains(err.Error(), "connectex: No connection could be made") {
			// Handle Windows-specific error messages
			fmt.Println("No kleister instance running.")
		} else {
			return fmt.Errorf("request failed: %w", err)
		}
		return nil // Don't return an error; just print a message
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Print response
	fmt.Println("Response:", string(body))
	return nil
}
