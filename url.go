package main

import (
	"fmt"
	"os"
	"strings"

	. "github.com/dirtman/sitepkg"
)

// Implement the "url" command.

var urlRequest string

func commandURL(invokedAs []string) error {

	commands := Commands{
		"get":    getURL,
		"delete": deleteURL,
		"add":    addURL,
		"update": updateURL,
	}
	var args []string
	var err error

	if len(os.Args) < 2 {
		commandHelp(invokedAs, commands, "need a subcommand", 1)
	}
	if strings.HasPrefix(os.Args[1], "-") || strings.EqualFold(os.Args[1], "help") {
		commandHelp(invokedAs, commands, "", 0)
	}

	command := strings.ToLower(os.Args[1])
	os.Args[0] += " " + os.Args[1]
	os.Args = append(os.Args[:1], os.Args[2:]...)

	function, ok := commands[command]
	if !ok {
		commandHelp(invokedAs, commands, fmt.Sprintf("unrecognized command \"%s\"", command), 1)
	}

	// Set API-related options:
	if err = SetAPIOptions(); err != nil {
		Warn("Failure setting API options: %v", err)
		os.Exit(1)
	}
	// Set any other options:
	SetBoolOpt("Debug", "", false, false, "Debug mode.")

	// Now that all our options have been specified, configure them, initialize
	// the API, and process user input..
	if args, err = ConfigureOptions(); err != nil {
		return Error("Failure initializing program: %s\n", err)
	} else if err = InitAPI(); err != nil {
		return Error("Failure initializing API: %s", err)
	} else if len(args) == 0 {
		return Error("no URL specified")
	} else if len(args) > 1 {
		return Error("only one URL allowed")
	}
	if urlRequest = args[0]; !strings.HasPrefix(urlRequest, "/") {
		urlRequest = "/" + urlRequest
	}

	// Run the func for the specified command:
	invokedAs = append(invokedAs, command)
	err = function(invokedAs)
	if err != nil {
		err = Error("Failure running \"%s\": %v", strings.Join(invokedAs, " "), err)
	}
	return err
}

// Implement the "get" command.

func getURL(invokedAs []string) error {

	body, err := IBAPIRequest("GET", urlRequest, nil)
	if err != nil {
		Print("%s:  %s\n", strings.Join(invokedAs, " "), "Failed")
		return Error("GET request failed: %s", err)
	}
	Print("%s:  %s\n%s\n", strings.Join(invokedAs, " "), "Success", body)
	return nil
}

func deleteURL(invokedAs []string) error {

	if _, err := IBAPIRequest("DELETE", urlRequest, nil); err != nil {
		Print("%s:  %s\n", strings.Join(invokedAs, " "), "Failed")
		return Error("DELETE request failed: %s", err)
	}
	Print("%s:  %s\n", strings.Join(invokedAs, " "), "Success")
	return nil
}

func addURL(invokedAs []string) error {

	if _, err := IBAPIRequest("POST", urlRequest, nil); err != nil {
		Print("%s:  %s\n", strings.Join(invokedAs, " "), "Failed")
		return Error("POST request failed: %s", err)
	}
	Print("%s:  %s\n", strings.Join(invokedAs, " "), "Success")
	return nil
}

func updateURL(invokedAs []string) error {

	if _, err := IBAPIRequest("PUT", urlRequest, nil); err != nil {
		Print("%s:  %s\n", strings.Join(invokedAs, " "), "Failed")
		return Error("PUT request failed: %s", err)
	}
	Print("%s:  %s\n", strings.Join(invokedAs, " "), "Success")
	return nil
}
