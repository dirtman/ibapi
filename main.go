// Infoblox WAPI CLI to add/delete/get/update basic Infoblox objects.

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	. "github.com/dirtman/sitepkg"
)

// Mapping of command names to corresponding functions.
type Commands map[string]func([]string) error

// Root ibapi command that invokes the specified subcommand.

func main() {

	// Let's point DefaultShow and DefaultDebug to Stderr so that
	// Stdout is consistent and not dependent on --verbose, --debug
	// etc (extra details will go to Stderr).
	DefaultShow = os.Stderr
	DefaultDebug = os.Stderr

	// Define our command/func mapping:
	commands := Commands{
		"a":            commandA,
		"alias":        commandAlias,
		"host":         commandHost,
		"cname":        commandCNAME,
		"ptr":          commandPTR,
		"url":          commandURL,
		"fixedaddress": commandFixedAddress,
		"grid":         commandGrid,
		"mx":           commandMX,
		"txt":          commandTXT,
		"authzone":     commandZoneAuth,
	}

	// Initialize as a SitePkg.
	err := PackageInit("ibapi", "1.0")
	if err != nil {
		Warn("Failure initializing program: %v", err)
		os.Exit(1)
	}

	// Create the in-memory documentation.
	if err = makePodMap(); err != nil {
		Warn("Failure making PodMap; may be OK")
	}

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

	Print("Usage:  %s [-h] <object> [-h] <operation> <args>\nSupported objects:\n",
		strings.Join(invokedAs, " "))
	var commandsSorted []string
	for name := range commands {
		commandsSorted = append(commandsSorted, name)
	}
	sort.Strings(commandsSorted)
	for _, name := range commandsSorted {
		Println("  %s", name)
	}
	Exit(exitCode)
}

// Common initialization sequence needed by each object type subcommand.

func subCommandInit(objectType, operator string, duo bool) (*UserInput, error) {

	var input *UserInput
	var args []string
	var err error

	// Set API-related options:
	if err = SetAPIOptions(); err != nil {
		Warn("Failure setting API options: %v", err)
		os.Exit(1)
	}
	// Set options common to all commands.
	SetBoolOpt("Debug", "", false, false, "Debug mode.")

	// Now that all our options have been specified, configure them, initialize
	// the API, and process user input..
	if args, err = ConfigureOptions(); err != nil {
		return nil, Error("Failure initializing program: %s\n", err)
	} else if err = InitAPI(); err != nil {
		return nil, Error("Failure initializing API: %s", err)
	} else if input, err = getUserInput(objectType, operator, duo, args); err != nil {
		return nil, Error("failure getting user input: %v", err)
	}
	return input, nil
}
