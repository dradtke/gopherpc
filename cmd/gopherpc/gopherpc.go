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
		verbose = flag.Bool("v", false, "enable verbose output")
		scan    = flag.String("scan", "", "path of the package to scan")
		output  = flag.String("o", "", "output file to generate")
		pkg     = flag.String("pkg", "", "name of the package to generate; if not specified, will be inferred from -o")
	)
	flag.Parse()

	if *scan == "" {
		log.Fatal("-scan must be specified")
	}

	args := gopherpc.GenArgs{
		Verbose:     *verbose,
		SrcPackage:  *scan,
		PackageName: *pkg,
	}

	if *output == "" {
		args.Out = os.Stdout

		if *pkg == "" {
			log.Fatal("-pkg must be specified if -o is not")
		}
	} else {
		var err error
		if args.Out, err = os.Create(*output); err != nil {
			log.Fatal(err)
		}

		if *pkg == "" {
			abs, err := filepath.Abs(*output)
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
