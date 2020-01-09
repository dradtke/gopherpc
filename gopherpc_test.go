package gopherpc_test

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"
	"testing"

	"github.com/dradtke/gopherpc"
	"golang.org/x/tools/go/packages"
)

func TestGen(t *testing.T) {
	if !testing.Verbose() {
		log.SetOutput(ioutil.Discard)
	}

	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax,
	}, "github.com/dradtke/gopherpc/testdata/server")
	if err != nil {
		t.Fatal(err)
	}
	if len(pkgs) != 1 {
		t.Fatalf("unexpected number of packages: %d", len(pkgs))
	}

	var buf bytes.Buffer
	if err := gopherpc.Gen(pkgs[0], "rpc", &buf, gopherpc.Wasm); err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(buf.String(), "type EchoService struct") {
		t.Error("generated RPC file does not contain EchoService")
	}

	if !strings.Contains(buf.String(), "func (s EchoService) Ping() (string, error)") {
		t.Error("generated RPC file does not contain EchoService.Ping")
	}
}
