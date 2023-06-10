package main

import (
	"strings"

	. "github.com/dirtman/sitepkg"
)

// Implement the "add" command.

func addCNAME(invokedAs []string) error {

	var input *UserInput
	var states = make(StatesCNAME)
	var statesHost = make(StatesHost)
	var statesA = make(StatesA)
	var statesAlias = make(StatesAlias)
	var err error
	duo := true
	invokedString := strings.Join(invokedAs, ":")

	SetStringOpt("view", "V", true, "default", "Specify the view for the record")
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
	f := []string{"view=" + input.view}
	ndList := input.ndList
	if err = getStates(states, ndList, f, nil, true, false); err != nil {
		return Error("failure getting states: %v", err)
	}

	// Check if any errors occurred getting the above records. If so, abort.
	if errors := checkStateErrors(states, false, true); len(errors) != 0 {
		return Error("Aborting process; no records added.")
	}

	// Check for any existing A records with the same name.
	if err = getStates(statesA, ndList, f, nil, true, false); err != nil {
		return Error("failure getting statesA: %v", err)
	} else if errors := checkStateErrors(statesA, true, true); len(errors) != 0 {
		return Error("Aborting process; no records added.")
	}

	// Check for any existing Alias records with the same name.
	// This check will only check for the default target type, A.
	// For other types, we'll settle for a 4xx WAPI error.
	if err = getStates(statesAlias, ndList, f, nil, true, false); err != nil {
		return Error("failure getting statesAlias: %v", err)
	} else if errors := checkStateErrors(statesAlias, true, true); len(errors) != 0 {
		return Error("Aborting process; no records added.")
	}

	// Check for any existing Host records with the same name.
	if err = getStates(statesHost, ndList, f, nil, true, false); err != nil {
		return Error("failure getting statesHost: %v", err)
	} else if errors := checkStateErrors(statesHost, false, true); len(errors) != 0 {
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
				conflict += sep + "CNAME record with same name"
			} else {
				conflict += sep + "CNAME record with same target"
			}
			sep = ", "
		}

		if len(statesA[nameData].records) != 0 {
			conflict += sep + "A record with same name"
			sep = ", "
		}
		if len(statesAlias[nameData].records) != 0 {
			conflict += sep + "Alias with same name"
			sep = ", "
		}
		if len(statesHost) != 0 && len(statesHost[nameData].records) != 0 {
			conflict += sep + "Host record with same name"
			sep = ", "
		}

		if conflict != "" {
			Print("%-*s NOT added: %s\n", space, "CNAME("+nameData+")", conflict)
			numConflicts++
			continue
		}

		if _, err := addRecord(object, nKey, dKey, name, data, input.fields); err != nil {
			return Error("aborting! failure adding cname record %s: %v", nameData, err)
		} else {
			Print("%-*s: Added\n", space, "CNAME("+nameData+")")
		}
	}

	if numConflicts == 0 {
		return nil
	} else if len(states) > 1 {
		return Error("One or more records not added due to conflict.")
	} else {
		return Error("Record not added due to conflict.")
	}
}
