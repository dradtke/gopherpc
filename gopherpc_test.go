package gopherpc_test

import (
	"bytes"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
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
	if err := gopherpc.Gen(pkgs[0], "rpc", &buf, gopherpc.Default); err != nil {
		t.Fatal(err)
	}

	// Log the result so that it's visible if -v is specified.
	t.Log(buf.String())

	fileset := token.NewFileSet()
	f, err := parser.ParseFile(fileset, "", buf.String(), parser.AllErrors)
	if err != nil {
		t.Fatalf("generated code failed to parse: %s", err)
	}

	config := types.Config{Importer: importer.Default()}
	if _, err = config.Check("", fileset, []*ast.File{f}, nil); err != nil {
		t.Fatalf("generated code failed to typecheck: %s", err)
	}

	if !strings.Contains(buf.String(), "type EchoService struct") {
		t.Error("generated RPC file does not contain EchoService")
	}

	if !strings.Contains(buf.String(), "func (s EchoService) Ping() (string, error)") {
		t.Error("generated RPC file does not contain EchoService.Ping")
	}
}
