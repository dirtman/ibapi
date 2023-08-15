package main

import (
	"strconv"
	"strings"

	. "github.com/dirtman/sitepkg"
)

// Implement the "add" command.

func addMX(invokedAs []string) error {

	var input *UserInput
	var states = make(StatesMX)
	var err error
	duo := true
	invokedString := strings.Join(invokedAs, ":")
	var preference string

	SetStringOpt("view", "V", true, "default", "Specify the the view for the record")
	SetStringOpt("preference", "p", true, "10", "Specify the preference value.")
	SetStringOpt("comment", "c", true, invokedString, "Specify the comment for the record")
	SetBoolOpt("disable", "D", true, false, "Disable the new record")
	SetUintOpt("ttl", "", true, 0, "Specify the TTL for the record")
	SetStringOpt("fields", "F", true, "", "Specify additional fields for the record")
	SetStringOpt("filename", "f", true, "", "Specify a name/data input file")

	if input, err = subCommandInit(invokedAs[1], invokedAs[2], duo); err != nil {
		return Error("failure initializing program and getting user input: %v", err)
	} else if preference, err = GetStringOpt("preference"); err != nil {
		return Error("failure getting preference option: %v", err)
	} else if _, err = strconv.Atoi(preference); err != nil {
	    return Error("preference value must be a positive integer")
	}
	// Append preference to the list of field/values
	input.fields = append(input.fields, "preference="+preference)

	// Check if any of the requested records already exist or conflict with existing
	// records in our view.
	ndList := input.ndList
	f := []string{"view=" + input.view}
	f = append(f, "preference=" + preference)
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
		sep := "Conflicts found: "

		if len(states[nameData].records) != 0 {
			if state.records[0].Fqdn == name && state.records[0].MX == data {
				conflict += sep + "MX record with same name and Mail Exchanger"
			}
			sep = ", "
		}

		if conflict != "" {
			Print("%-*s NOT added: %s\n", space, "MX("+nameData+")", conflict)
			numConflicts++
			continue
		}

		if _, err := addRecord(object, nKey, dKey, name, data, input.fields); err != nil {
			return Error("aborting! failure adding MX record %s: %v", nameData, err)
		} else {
			Print("%-*s: Added\n", space, "MX("+nameData+"/"+preference+")")
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
