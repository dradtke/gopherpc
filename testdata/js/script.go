// +build js

// This is an example of GopherJS code imports and uses a generated RPC client.

package main

import (
	"github.com/dradtke/gopherpc/json"
	"github.com/dradtke/gopherpc/testdata/js/rpc"
)

func main() {
	client := rpc.Client{
		URL:       "https://your-site.com/rpc",
		CSRFToken: "abcd",
		Encoding:  json.Encoding{},
	}

	result, err := client.EchoService().Ping()
	if err != nil {
		println("failed to call EchoService.Ping: " + err.Error())
	} else {
		println(result)
	}
}
