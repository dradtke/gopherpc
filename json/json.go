package json

import (
	"encoding/json"
	"errors"
	"io"
	"sync/atomic"
)

// clientRequest represents a JSON-RPC request sent by a client.
type clientRequest struct {
	// A String containing the name of the method to be invoked.
	Method string `json:"method"`
	// Object to pass as request parameter to the method.
	Params [1]interface{} `json:"params"`
	// The request id. This can be of any type. It is used to match the
	// response with the request that it is replying to.
	Id uint64 `json:"id"`
}

// clientResponse represents a JSON-RPC response returned to a client.
type clientResponse struct {
	Result *json.RawMessage `json:"result"`
	Error  interface{}      `json:"error"`
	Id     uint64           `json:"id"`
}

type Encoding struct {
	id uint64
}

func (e *Encoding) ContentType() string {
	return "application/json"
}

func (e *Encoding) EncodeRequest(serviceMethod string, arg interface{}) ([]byte, error) {
	c := &clientRequest{
		Method: serviceMethod,
		Params: [1]interface{}{arg},
		Id:     atomic.AddUint64(&e.id, 1),
	}
	return json.Marshal(c)
}

func (e *Encoding) DecodeResponse(r io.Reader, reply interface{}) error {
	var c clientResponse
	if err := json.NewDecoder(r).Decode(&c); err != nil {
		return errors.New("failed to decode response: " + err.Error())
	}
	if c.Error != nil {
		switch t := c.Error.(type) {
		case error:
			return t
		case string:
			return errors.New(t)
		default:
			println(c.Error)
			return errors.New("client response returned error, see above")
		}
	}
	if c.Result == nil {
		return errors.New("unexpected null result")
	}
	if err := json.Unmarshal(*c.Result, reply); err != nil {
		return errors.New("failed to unmarshal response: " + err.Error())
	}
	return nil
}
