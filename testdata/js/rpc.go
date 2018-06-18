// +build js

// This file was autogenerated by gopherpc.

package js

import (
	"io"
	"strings"

	"honnef.co/go/js/xhr"
)

// Client represents a handle to an RPC endpoint.
type Client struct {
	// URL is the endpoint to send RPC requests to.
	URL string

	// CSRFToken is the CSRF token to include with each request.
	CSRFToken string

	// EncodeClientRequest encodes the client request. To use the
	// included JSON implementation, set this to json.EncodeClientRequest.
	EncodeClientRequest func(string, interface{}) ([]byte, error)

	// DecodeClientResponse decodes the response. To use the included
	// JSON implementation, set this to json.DecodeClientResponse.
	DecodeClientResponse func(r io.Reader, reply interface{}) error
}

func (c Client) call(serviceMethod string, args, ret interface{}) error {
	message, err := (c.EncodeClientRequest)(serviceMethod, args)
	if err != nil {
		return err
	}

	req := xhr.NewRequest("POST", c.URL)
	req.SetRequestHeader("X-CSRF-Token", c.CSRFToken)
	req.SetRequestHeader("Content-Type", "application/json")
	if err := req.Send(message); err != nil {
		return err
	}

	return (c.DecodeClientResponse)(strings.NewReader(req.ResponseText), &ret)
}

type TestService struct{ Client }

func (c Client) TestService() TestService {
	return TestService{c}
}

func (s TestService) Ping() (string, error) {
	var reply string
	err := s.call("TestService.Ping", nil, &reply)
	return reply, err
}
