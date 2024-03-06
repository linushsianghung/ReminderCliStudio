package main

import (
	"fmt"
	"net/http"
)

type HTTPClient struct {
	client          *http.Client
	backendEndpoint string
}

type Adaptor struct {
	HTTPClient // Promoted Field; Ref: https://medium.com/golangspec/promoted-fields-and-methods-in-go-4e8d7aefb3e3
	commands   map[string]func() func(string) error
}

func newAdaptor(uri string) Adaptor {
	a := Adaptor{}
	a.client = &http.Client{}
	a.backendEndpoint = uri
	a.commands = map[string]func() func(string) error{
		"create": a.create,
		"edit":   a.edit,
		"fetch":  a.fetch,
		"delete": a.delete,
		"health": a.health,
	}

	return a
}

func (a Adaptor) convertCmd(cmdName string) error {
	cmd, ok := a.commands[cmdName]
	if !ok {
		return fmt.Errorf("Unknown command: %s\n", cmdName)
	}

	return cmd()(cmdName)
}

func (a Adaptor) help() {
	fmt.Println("Operation Commands")
	for cmd, _ := range a.commands {
		fmt.Printf(" %s\t\t--[Description]\n", cmd)
	}
	fmt.Println("Use <command> --help for more information about a given command.")
}

func (a Adaptor) create() func(string) error {
	return func(cmd string) error {
		fmt.Println("Create Reminder")
		return nil
	}
}

func (a Adaptor) edit() func(string) error {
	return func(cmd string) error {
		fmt.Println("Edit Reminder")
		return nil
	}
}

func (a Adaptor) fetch() func(string) error {
	return func(cmd string) error {
		fmt.Println("Fetch Reminder")
		return nil
	}
}

func (a Adaptor) delete() func(string) error {
	return func(cmd string) error {
		fmt.Println("Delete Reminder")
		return nil
	}
}

func (a Adaptor) health() func(string) error {
	return func(cmd string) error {
		fmt.Println("Health Check")
		return nil
	}
}
