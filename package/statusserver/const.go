package statusserver

import "net/http"

const (
	DefaultAdress string = "127.0.0.1:8080"
)

type FalsyAuth struct{}

func (self FalsyAuth) Authenticate(r *http.Request) bool {
	return false
}

type TruthyAuth struct{}

func (self TruthyAuth) Authenticate(r *http.Request) bool {
	return true
}
