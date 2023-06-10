package main

import (
	"strings"

	. "github.com/dirtman/sitepkg"
)

// Implement the "add" command.

func addA(invokedAs []string) error {

	var input *UserInput
	var states = make(StatesA)
	var statesPTR = make(StatesPTR)
	var statesHost = make(StatesHost)
	var statesCNAME = make(StatesCNAME)
	var statesAlias = make(StatesAlias)
	var check bool
	var err error
	duo := true
	invokedString := strings.Join(invokedAs, ":")

	SetStringOpt("view", "V", true, "default", "Specify the view for the record")
	SetStringOpt("comment", "c", true, invokedString, "Specify the comment for the record")
	SetBoolOpt("disable", "D", true, false, "Disable the new record")
	SetUintOpt("ttl", "", true, 0, "Specify the TTL for the record")
	SetStringOpt("fields", "F", true, "", "Specify additional fields for the record")
	SetStringOpt("filename", "f", true, "", "Specify a name/data input file")
	SetBoolOpt("checkRecords", "C", true, false, "Check for existing related records")

	if input, err = subCommandInit(invokedAs[1], invokedAs[2], duo); err != nil {
		return Error("failure initializing program and getting user input: %v", err)
	} else if check, err = GetBoolOpt("checkRecords"); err != nil {
		return Error("failure getting checkRecords option: %v", err)
	}

	// Check if any of the requested records to be added already exist.
	// If --checkRecords, also check for any "related" records.
	// Note we only care about existing records in the same view as ours.
	f := []string{"view=" + input.view}
	ndList := input.ndList
	if check {
		err = getStates(states, ndList, f, nil, true, true)
	} else {
		err = getStates(states, ndList, f, nil, false, false)
	}
	if err != nil {
		return Error("failure getting states: %v", err)
	}

	// Check if any errors occurred getting the above records. If so, abort.
	if errors := checkStateErrors(states, false, true); len(errors) != 0 {
		return Error("Aborting process; no records added.")
	}

	// Check for any existing CNAME records with the same name.
	if err = getStates(statesCNAME, ndList, f, nil, true, false); err != nil {
		return Error("failure getting statesCNAME: %v", err)
	} else if errors := checkStateErrors(statesCNAME, true, true); len(errors) != 0 {
		return Error("Aborting process; no records added.")
	}

	// Check for any existing Alias records with the same name.
	f = []string{"view=" + input.view, "target_type=A"}
	if err = getStates(statesAlias, ndList, f, nil, true, false); err != nil {
		return Error("failure getting statesAlias: %v", err)
	} else if errors := checkStateErrors(statesAlias, true, true); len(errors) != 0 {
		return Error("Aborting process; no records added.")
	}

	// If check, check for other existing "related" records.
	if check {
		f := []string{"view=" + input.view}
		if err = getStates(statesHost, ndList, f, nil, true, true); err != nil {
			return Error("failure getting statesHost: %v", err)
		} else if errors := checkStateErrors(statesHost, false, true); len(errors) != 0 {
			return Error("Aborting process; no records added.")
		}
		if err = getStates(statesPTR, ndList, f, nil, true, true); err != nil {
			return Error("failure getting statesPTR: %v", err)
		} else if errors := checkStateErrors(statesPTR, true, true); len(errors) != 0 {
			return Error("Aborting process; no records added.")
		}
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
				conflict += sep + "A record with same name"
			} else {
				conflict += sep + "A record with same IP"
			}
			sep = ", "
		}

		if len(statesCNAME[nameData].records) != 0 {
			conflict += sep + "CNAME with same name"
			sep = ", "
		}
		if len(statesAlias[nameData].records) != 0 {
			conflict += sep + "Alias with same name"
			sep = ", "
		}
		if check && len(statesHost[nameData].records) != 0 {
			conflict += sep + "related Host record"
			sep = ", "
		}
		if check && len(statesPTR[nameData].records) != 0 {
			conflict += sep + "related PTR record"
			sep = ", "
		}

		if conflict != "" {
			Print("%-*s NOT added: %s\n", space, "A("+nameData+")", conflict)
			numConflicts++
			continue
		}

		if _, err := addRecord(object, nKey, dKey, name, data, input.fields); err != nil {
			return Error("aborting! failure adding host record %s: %v", nameData, err)
		} else {
			Print("%-*s: Added\n", space, "A("+nameData+")")
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
