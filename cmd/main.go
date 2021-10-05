package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s'search/replace' file ...\n", path.Base(os.Args[0]))
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {

}
