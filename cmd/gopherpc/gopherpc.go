package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/dradtke/gopherpc"
)

func main() {
	var (
		verbose           = flag.Bool("v", false, "enable verbose output")
		scan              = flag.String("scan", "", "path of the package to scan")
		outputFile        = flag.String("o", "", "output file to generate")
		outputPackageName = flag.String("pkg", "", "name of the package to generate; if not specified, will be inferred from -o")
	)
	flag.Parse()

	if *scan == "" {
		log.Fatal("-scan must be specified")
	}

	args := gopherpc.GenArgs{
		Verbose:     *verbose,
		SrcPackage:  *scan,
		PackageName: *outputPackageName,
	}

	if *outputFile == "" {
		args.Out = os.Stdout

		if *outputPackageName == "" {
			log.Fatal("-pkg must be specified if -o is not")
		}
	} else {
		var err error
		if args.Out, err = os.Create(*outputFile); err != nil {
			log.Fatal(err)
		}

		if *outputPackageName == "" {
			abs, err := filepath.Abs(*outputFile)
			if err != nil {
				log.Fatal(err)
			}
			args.PackageName = filepath.Base(filepath.Dir(abs))
		}
	}

	if err := gopherpc.Gen(args); err != nil {
		log.Fatal(err)
	}
}
