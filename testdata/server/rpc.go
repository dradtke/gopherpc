package rpc

import "net/http"

// rpc:gen
type TestService struct{}

func (s TestService) Ping(_ *http.Request, _ *struct{}, reply *string) error {
	*reply = "pong"
	return nil
}
