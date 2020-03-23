package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/dradtke/gopherpc"
	"golang.org/x/tools/go/packages"
)

func main() {
	var (
		verbose           = flag.Bool("v", false, "enable verbose output")
		scan              = flag.String("scan", "", "path of the package to scan")
		outputFile        = flag.String("o", "", "output file to generate")
		outputPackageName = flag.String("pkg", "", "name of the package to generate; if not specified, will be inferred from -o")
		mode              = flag.String("mode", "", "output mode, use 'js' for compatibility with GopherJS")
	)
	flag.Parse()

	if *scan == "" {
		log.Fatal("-scan must be specified")
	}

	genMode, err := gopherpc.NewMode(*mode)
	if err != nil {
		log.Fatal(err)
	}

	pkgs, err := packages.Load(&packages.Config{
		BuildFlags: []string{"-tags", "!js"},
		Mode:       packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax,
	}, *scan)
	if err != nil {
		log.Fatalf("failed to load package %s: %s", *scan, err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("found unexpected number of packages; want 1, got %d", len(pkgs))
	}

	if !*verbose {
		log.SetOutput(ioutil.Discard)
	}

	var w io.Writer

	if *outputFile == "" {
		w = os.Stdout
		if *outputPackageName == "" {
			log.Fatal("-pkg must be specified if -o is not")
		}
	} else {
		var err error
		if w, err = os.Create(*outputFile); err != nil {
			log.Fatal(err)
		}

		if *outputPackageName == "" {
			abs, err := filepath.Abs(*outputFile)
			if err != nil {
				log.Fatal(err)
			}
			*outputPackageName = filepath.Base(filepath.Dir(abs))
		}
	}

	if err := gopherpc.Gen(pkgs[0], *outputPackageName, w, genMode); err != nil {
		log.Fatal(err)
	}
}
