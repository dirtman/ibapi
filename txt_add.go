package main

import (
	"strings"

	. "github.com/dirtman/sitepkg"
)

// Implement the "add" command.

func addTXT(invokedAs []string) error {

	var input *UserInput
	var states = make(StatesTXT)
	var err error
	duo := true
	invokedString := strings.Join(invokedAs, ":")

	SetStringOpt("view", "V", true, "default", "Specify the the view for the record")
	SetStringOpt("comment", "c", true, invokedString, "Specify the comment for the record")
	SetBoolOpt("disable", "D", true, false, "Disable the new record")
	SetUintOpt("ttl", "", true, 0, "Specify the TTL for the record")
	SetStringOpt("fields", "F", true, "", "Specify additional fields for the record")
	SetStringOpt("filename", "f", true, "", "Specify a name/data input file")

	if input, err = subCommandInit(invokedAs[1], invokedAs[2], duo); err != nil {
		return Error("failure initializing program and getting user input: %v", err)
	}

	// Check if any of the requested records already exist or conflict with existing
	// records in our view.
	ndList := input.ndList
	f := []string{"view=" + input.view}
	if err = getStates(states, ndList, f, nil, true, true); err != nil {
		return Error("failure getting states: %v", err)
	}

	// Check if any errors occurred getting the above records. If so, abort.
	if errors := checkStateErrors(states, false, true); len(errors) != 0 {
		return Error("Aborting process; no records added.")
	}

	// Loop through the user provided input (name/data) list.

	space := input.maxNameLength + 8
	nKey, dKey := states.GetNDKeys()
	object := states.GetObjectType()
	var numConflicts uint

	for nameData, state := range states {

		var name, data, conflict string
		name, data, _ = splitND(nameData)
		ShowDebug("addTXT1: data: %s", data)
		//data = unEscapeURLText(data)
		sep := "Conflicts found: "

		if len(states[nameData].records) != 0 {
			if state.records[0].Name == name && state.records[0].Text == data {
				conflict += sep + "TXT record with same name and data"
			}
			sep = ", "
		}
		// From here on, we are just showing "nameData".
		nameData = unEscapeURLText(nameData)

		if conflict != "" {
			Print("%-*s NOT added: %s\n", space, "TXT("+nameData+")", conflict)
			numConflicts++
			continue
		}

		if _, err := addRecord(object, nKey, dKey, name, data, input.fields); err != nil {
			return Error("aborting! failure adding TXT record %s: %v", nameData, err)
		} else {
			Print("%-*s: Added\n", space, "TXT("+nameData+")")
		}
	}

	if numConflicts > 0 {
		if len(states) > 1 {
			return Error("One or more records not added due to conflict.")
		} else {
			return Error("Record not added due to conflict.")
		}
	} else if input.restartServices {
		if err = gridRestartServices(Verbose); err != nil {
			return Error("failure restarting services: %s", err)
		}
	}

	return nil
}
