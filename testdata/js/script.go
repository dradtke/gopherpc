// +build js

// This is an example of GopherJS code imports and uses a generated RPC client.

package main

import (
	"github.com/dradtke/gopherpc/json"
	"github.com/dradtke/gopherpc/testdata/js/rpc"
)

func main() {
	client := rpc.Client{
		URL:                  "https://your-site.com/rpc",
		CSRFToken:            "abcd",
		EncodeClientRequest:  json.EncodeClientRequest,
		DecodeClientResponse: json.DecodeClientResponse,
	}

	result, err := client.TestService().Ping()
	if err != nil {
		println("failed to call TestService.Ping: " + err.Error())
	} else {
		println(result)
	}
}
