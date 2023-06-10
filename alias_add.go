package main

import (
	"strings"

	. "github.com/dirtman/sitepkg"
)

// Implement the "add" command.

func addAlias(invokedAs []string) error {

	var input *UserInput
	var states = make(StatesAlias)
	var statesHost = make(StatesHost)
	var statesCNAME = make(StatesCNAME)
	var statesA = make(StatesA)
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
	SetStringOpt("targetType", "T", false, "A", "Specify the target type of the alias")

	if input, err = subCommandInit(invokedAs[1], invokedAs[2], duo); err != nil {
		return Error("failure initializing program and getting user input: %v", err)
	} else if check, err = GetBoolOpt("checkRecords"); err != nil {
		return Error("failure getting checkRecords option: %v", err)
	}

	// Check if any of the requested records to be added already exist.
	// If --checkRecords, also check for any "related" records.
	// Note we only care about existing records in the same view as ours
	// and with the same target_type as ours.
	targetType := input.targetType
	f := []string{"view=" + input.view, "target_type=" + targetType}
	ndList := input.ndList
	if err = getStates(states, ndList, f, nil, true, false); err != nil {
		return Error("failure getting states: %v", err)
	}

	// Check if any errors occurred getting the above records. If so, abort.
	if errors := checkStateErrors(states, false, true); len(errors) != 0 {
		return Error("Aborting process; no records added.")
	}
	f = []string{"view=" + input.view} // CNAME and A records do not have a target_type.

	// Check for any existing CNAME records with the same name.
	if err = getStates(statesCNAME, ndList, f, nil, true, false); err != nil {
		return Error("failure getting statesCNAME: %v", err)
	} else if errors := checkStateErrors(statesCNAME, true, true); len(errors) != 0 {
		return Error("Aborting process; no records added.")
	}

	if targetType == "A" {
		// Check for any existing A records with the same name.
		if err = getStates(statesA, ndList, f, nil, true, false); err != nil {
			return Error("failure getting statesA: %v", err)
		} else if errors := checkStateErrors(statesA, true, true); len(errors) != 0 {
			return Error("Aborting process; no records added.")
		}
	}
	if check && (targetType == "A" || targetType == "AAAA") {
		// Check for any existing Host records with the same name.
		if err = getStates(statesHost, ndList, f, nil, true, false); err != nil {
			return Error("failure getting statesHost: %v", err)
		} else if errors := checkStateErrors(statesHost, false, true); len(errors) != 0 {
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
				conflict += sep + "Alias record with same name"
			} else {
				conflict += sep + "Alias record with same target"
			}
			sep = ", "
		}

		if len(statesCNAME[nameData].records) != 0 {
			conflict += sep + "CNAME with same name"
			sep = ", "
		}
		if len(statesA) != 0 && len(statesA[nameData].records) != 0 {
			conflict += sep + "A record with same name"
			sep = ", "
		}
		if len(statesHost) != 0 && len(statesHost[nameData].records) != 0 {
			conflict += sep + "Host record with same name"
			sep = ", "
		}

		if conflict != "" {
			Print("%-*s NOT added: %s\n", space, targetType+" Alias("+nameData+")", conflict)
			numConflicts++
			continue
		}

		if _, err := addRecord(object, nKey, dKey, name, data, input.fields); err != nil {
			return Error("aborting! failure adding %s Alias %s: %v", targetType, nameData, err)
		} else {
			Print("%-*s: Added\n", space, targetType+" Alias("+nameData+")")
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
