package main

import (
	"strings"

	. "github.com/dirtman/sitepkg"
)

// Implement the "add" command.

func addFixedAddress(invokedAs []string) error {

	var input *UserInput
	var states = make(StatesFixedAddress)
	var err error
	duo := true
	invokedString := strings.Join(invokedAs, ":")
	var name string

	SetStringOpt("view", "V", true, "default", "Specify the the network view for the record")
	SetStringOpt("name", "n", false, "", "The new record's name")
	SetStringOpt("comment", "c", true, invokedString, "Specify the comment for the record")
	SetBoolOpt("disable", "D", true, false, "Disable the new record")
	SetUintOpt("ttl", "", true, 0, "Specify the TTL for the record")
	SetStringOpt("fields", "F", true, "", "Specify additional fields for the record")
	SetStringOpt("filename", "f", true, "", "Specify a name/data input file")
	SetStringOpt("bootfile", "b", false, "", "Specify the bootfile of the IP address.")
	SetStringOpt("nextserver", "N", false, "", "Specify the nextserver of the IP address.")
	SetStringOpt("bootserver", "B", false, "", "Specify the bootserver of the IP address.")
	SetBoolOpt("restartServices", "R", true, false, "Restart Grid services if needed")

	if input, err = subCommandInit(invokedAs[1], invokedAs[2], duo); err != nil {
		return Error("failure initializing program and getting user input: %v", err)
	} else if name, err = GetStringOpt("name"); err != nil {
		return Error("failure getting name option: %v", err)
	} else if name != "" { // Append it to the list of field/values
		input.fields = append(input.fields, "ptrdname="+name)
	}

	// Check if any of the requested records already exist or conflict with existing
	// records in our view.
	f := []string{"network_view=" + input.view}
	ndList := input.ndList
	if err = getStates(states, ndList, f, nil, true, false); err != nil {
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
		sep := "Conflicts found: "
		name, data, _ = splitND(nameData)

		if len(states[nameData].records) != 0 {
			if state.records[0].Name == name {
				conflict += sep + "FixedAddress record with same name"
			} else {
				conflict += sep + "FixedAddress record with same target"
			}
			sep = ", "
		}

		if conflict != "" {
			Print("%-*s NOT added: %s\n", space, "FixedAddress("+nameData+")", conflict)
			numConflicts++
			continue
		}

		if _, err := addRecord(object, nKey, dKey, name, data, input.fields); err != nil {
			return Error("aborting! failure adding fixedaddress record %s: %v", nameData, err)
		} else {
			Print("%-*s: Added\n", space, "FixedAddress("+nameData+")")
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
