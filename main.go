// Infoblox WAPI CLI to add/delete/get/update basic Infoblox objects.

package main

import (
	"fmt"
	"os"
	"strings"

	. "github.com/dirtman/sitepkg"
)

// Mapping of command names to corresponding functions.
type Commands map[string]func([]string) error

// Root ibapi command that invokes the specified subcommand.

func main() {

	// Let's point DefaultShow and DefaultDebug to Stderr.
	DefaultShow = os.Stderr
	DefaultDebug = os.Stderr

	// Define our command/func mapping:
	commands := Commands{
		"a":     commandA,
		"alias": commandAlias,
		"host":  commandHost,
		"cname": commandCNAME,
		"ptr":   commandPTR,
		"url":   commandURL,
	}

	err := PackageInit("ibapi", "0.0")
	if err != nil {
		Warn("Failure initializing program: %v", err)
		os.Exit(1)
	}

	if err = makePodMap(); err != nil {
		Warn("Failure making PodMap; may be OK")
	}

	// Set options common to all commands.
	auth_types := strings.Join([]string{AuthMethodBasic, AuthMethodBearer}, ", ")
	auth_help := fmt.Sprintf("API authentication type (%s)", auth_types)
	//	SetStringOpt("APIServer", "", true, "infoblox.rice.edu", "API base URL")
	//	SetStringOpt("APIVersion", "", true, "2.11", "API Version")
	//	SetStringOpt("APIPort", "", true, "443", "API Port")
	SetStringOpt("APIBaseURL", "", true, "", "API base URL")
	SetStringOpt("APIAuthMethod", "", true, AuthMethodBasic, auth_help)
	//	SetStringOpt("APIAuthToken", "", true, "", "API auth token")
	SetStringOpt("APIAuthTokenID", "u", true, "", "username or auth token ID")
	SetStringOpt("Password", "p", true, "", "password for Basic authentication")
	SetStringOpt("SecretsDir", "", true, "", "Location of \"secrets files\"")
	SetIntOpt("HTTPTimeout", "", true, 60, "Timeout in seconds of the HTTP connection")
	SetBoolOpt("Debug", "", false, false, "Debug mode.")

	// Verify that at least one command has been specified:
	invokedAs := []string{ProgramName}
	if len(os.Args) < 2 {
		commandHelp(invokedAs, commands, "need one or more arguments", 1)
	}

	if strings.HasPrefix(os.Args[1], "-") || strings.EqualFold(os.Args[1], "help") {
		commandHelp(invokedAs, commands, "", 0)
	}

	// Grab the specified command and remove it from os.Args:
	command := strings.ToLower(os.Args[1])
	os.Args[0] += " " + os.Args[1]
	os.Args = append(os.Args[:1], os.Args[2:]...)

	// Get the func for the specified command:
	function, ok := commands[command]
	if !ok {
		commandHelp(invokedAs, commands, fmt.Sprintf("unrecognized command \"%s\"", command), 1)
	}

	// Run the func for the specified command:
	err = function(append(invokedAs, command))
	if err != nil {
		Warn("%v", err)
		os.Exit(1)
	}
}

// Print out short usage information when a command is invoked improperly. This is
// generally only needed for the root command and the top level command since full
// configuration has not yet been completed.

func commandHelp(invokedAs []string, commands Commands, message string, exitCode int) {

	// POD is more useful info; if the invoker is asking for help, show the POD doc.
	for _, option := range os.Args[1:] {
		opt := strings.ToLower(option)
		if opt == "-h" || opt == "help" || opt == "--help" {
			if err := ShowPod(); err == nil {
				Exit(exitCode)
			}
		}
	}

	if message != "" {
		Warn(message)
	}

	Println("Usage:  %s <command> <options>\nSupported commands:", strings.Join(invokedAs, " "))
	for name := range commands {
		Println("  %s", name)
	}
	Exit(exitCode)
}

// Common initialization sequence needed by each object type subcommand.

func subCommandInit(objectType, operator string, duo bool) (*UserInput, error) {

	var input *UserInput
	var args []string
	var err error

	if args, err = ConfigureOptions(); err != nil {
		return nil, Error("Failure initializing program: %s\n", err)
	} else if err = InitAPI(); err != nil {
		return nil, Error("Failure initializing API: %s", err)
	} else if input, err = getUserInput(objectType, operator, duo, args); err != nil {
		return nil, Error("failure getting user input: %v", err)
	}
	return input, nil
}
