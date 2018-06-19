package gopherpc_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dradtke/gopherpc"
)

func TestGen(t *testing.T) {
	var buf bytes.Buffer
	if err := gopherpc.Gen(gopherpc.GenArgs{
		Verbose:     testing.Verbose(),
		Out:         &buf,
		SrcPackage:  "github.com/dradtke/gopherpc/testdata/server",
		PackageName: "rpc",
	}); err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(buf.String(), "type EchoService struct") {
		t.Error("generated RPC file does not contain EchoService")
	}

	if !strings.Contains(buf.String(), "func (s EchoService) Ping() (string, error)") {
		t.Error("generated RPC file does not contain EchoService.Ping")
	}
}
