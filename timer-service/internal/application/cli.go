package application

import (
	"flag"
	"fmt"
	"os"
)

var (
	help bool
	path string
)

func init() {
	flag.StringVar(&path, "config-file", "", "it will be reading from config file")
	flag.BoolVar(&help, "help", false, "")
}

func parseFlags() string {
	flag.Parse()

	if help {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	return path
}
