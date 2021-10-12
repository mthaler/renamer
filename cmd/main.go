package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"renamer/modifier"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s'search/replace' file ...\n", path.Base(os.Args[0]))
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	dryrun := flag.Bool("d", false, "dryrun only")
	verbose := flag.Bool("v", false, "verbose mode")
	flag.Usage = usage
	flag.Parse()

	if *dryrun {
		fmt.Printf("Dryrun mode\n")
	}

	if len(flag.Args()) < 2 {
		usage()
	}

	cmd := flag.Args()[0]
	files := flag.Args()[1]
	modfier, err := modifier.Mkmodifier()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid command: %s\n", cmd)
	}

	for _, file := range files {

	}
}
