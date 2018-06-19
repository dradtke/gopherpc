package server

import "net/http"

// EchoService is a simple RPC service that implements Ping.
// gopherpc:generate
type EchoService struct{}

// Ping always responds with "pong".
func (s EchoService) Ping(r *http.Request, _ *struct{}, reply *string) error {
	*reply = "pong"
	return nil
}
