package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type todoFlag []string

func (td todoFlag) String() string {
	return strings.Join(td, " & ")
}

func (td *todoFlag) Set(s string) error {
	*td = append(*td, s)
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Waring!!! No commands found...")
		os.Exit(2)
	}

	// This if block implementation just in order for "flag.Value" example
	if strings.HasPrefix(os.Args[1], "-") {
		var todo todoFlag
		flag.Var(&todo, "td", "Specify the TODO list")
		flag.Parse()
		fmt.Printf("Here is your TODO list: %s\n", todo)
	} else {
		cmd := os.Args[1]
		switch cmd {
		case "greet":
			greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
			msgFlag := greetCmd.String("msg", "REMINDER CLI - CLI BASICS", "Message for 'greet' command")
			err := greetCmd.Parse(os.Args[2:])
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Printf("Hello & Welcome => %s\n", *msgFlag) // Remember to dereference pointer
		case "help":
			fmt.Println("Some help over here~~~")
		default:
			fmt.Printf("Unknown command: %s\n", cmd)
		}
	}
}
