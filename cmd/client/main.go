package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	backendEndpoint = flag.String("uri", "http://localhost:8080", "Backend Server URI")
	help            = flag.Bool("help", false, "Display helpful message")
)

func main() {
	flag.Parse()
	adaptor := newAdaptor(*backendEndpoint)

	if *help || len(os.Args) == 1 {
		adaptor.help()
		return
	}

	err := adaptor.convertCmd(os.Args[1])
	if err != nil {
		fmt.Printf("Execution Error!!! %s\n", err)
		os.Exit(2)
	}
}
