package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	. "github.com/dirtman/sitepkg"
)

type RefObject struct {
	Ref string `json:"_ref,omitempty"`
}

// Implement the "grid" command.
func commandGrid(invokedAs []string) error {

	commands := Commands{
		"ref":     invokeGetRef,
		"restart": invokeRestart,
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
	// Set any other common options:
	SetBoolOpt("Debug", "", false, false, "Debug mode.")

	// Now that all our options have been specified, configure them, initialize
	// the API, and process user input..
	if args, err = ConfigureOptions(); err != nil {
		return Error("Failure initializing program: %s\n", err)
	} else if err = InitAPI(); err != nil {
		return Error("Failure initializing API: %s", err)
	} else if len(args) != 0 {
		return Error("no operations supported")
	}

	// Run the func for the specified command:
	//    invokedAs = append(invokedAs, command)
	//    err = function(invokedAs)
	//    if err != nil {
	//        err = Error("Failure running \"%s\": %v", strings.Join(invokedAs, " "), err)
	//    }
	//    return err

	// Run the func for the specified command:
	invokedAs = append(invokedAs, command)
	return function(invokedAs)
}

func invokeGetRef(invokedAs []string) error {

	var ref string
	var err error

	if ref, err = gridGetRef(); err != nil {
		return Error("Failure getting grid reference: %s", err)
	}
	Print("Grid Reference ID: %s\n", ref)
	return nil
}

func invokeRestart(invokedAs []string) error {

	return gridRestartServices(!Quiet)
}

// Get the Grid's reference.
func gridGetRef() (string, error) {

	gridRefernce, _ := GetStringOpt("gridReference")
	if gridRefernce != "" {
		return gridRefernce, nil
	}
	body, api_err := IBAPIRequest("GET", "/grid", nil)
	if api_err != nil {
		return "", Error("Failure getting Grid ref: %s", api_err)
	}
	var refObjects []RefObject
	if err := json.Unmarshal(body, &refObjects); err != nil {
		return "", Error("failure unmarshing body (%s): %s", string(body), err)
	}
	return refObjects[0].Ref, nil
}

// Restart grid services if neccessary.
func gridRestartServices(verbose bool) error {

	gridref, err := gridGetRef()
	if err != nil {
		return Error("failure getting grid reference id: %s", err)
	}
	url := "/" + gridref + "?_function=restartservices"
	url += "&restart_option=RESTART_IF_NEEDED"
	url += "&service_option=ALL"
	url += "&member_order=SEQUENTIALLY"
	url += "&sequential_delay=1"
	_, api_err := IBAPIRequest("POST", url, nil)

	if api_err != nil {
		return Error("failure restarting services: %s", api_err)
	}
	if verbose {
		Show("Successfully instructed Infoblox to restart services if needed")
	}
	return nil
}
