package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, `nginx version: md5x/0.1.19
Usage: md5x [-h] -dir dir [-out filename] [-export filename]

Options:
`)
	flag.PrintDefaults()
}
