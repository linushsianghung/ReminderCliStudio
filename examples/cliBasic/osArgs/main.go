package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	/*
		CLI format: CMD Subcommand Args...
		os.Args[0]: CMD
		os.Args[1]: Command
		os.Args[2]: Command Flag
	*/
	if len(os.Args) < 2 {
		fmt.Println("Waring!!! No commands found...")
		os.Exit(2)
	}

	cmd := os.Args[1]
	switch cmd {
	case "greet":
		msg := "REMINDER CLI - CLI BASICS"
		if len(os.Args) > 2 {
			f := strings.Split(os.Args[2], "=")
			if len(f) == 2 && f[0] == "--msg" {
				msg = f[1]
			}
		}
		fmt.Printf("Hello & Welcome => %s\n", msg)
	case "help":
		fmt.Println("Some help over here~~~")
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
	}

	fmt.Println(os.Args)
}
