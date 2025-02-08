package statusserver

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type ContextBundle struct {
	Context    context.Context
	CancelFunc context.CancelFunc
	Waiter     *sync.WaitGroup
}

type ServerSettings struct {
	Adress string
	Auth   Authenticator
}

// serverState holds the running state of the server
type serverState struct {
	mu           sync.Mutex
	running      bool
	runningSince time.Time
}

type Authenticator interface {
	Authenticate(r *http.Request) bool
}
