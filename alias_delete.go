package main

import (
	. "github.com/dirtman/sitepkg"
	"strings"
)

// Implement the "delete" command.

func deleteAlias(invokedAs []string) error {

	var fields []string
	var input *UserInput
	var states StatesAlias = make(StatesAlias)
	var record *RecordAlias
	var err error
	duo := false

	SetStringOpt("view", "V", true, "default", "Specify the the record's view")
	SetStringOpt("filename", "f", true, "", "Specify an input file")
	SetStringOpt("targetType", "T", false, "A", "Specify the target type of the alias")

	if input, err = subCommandInit(invokedAs[1], invokedAs[2], duo); err != nil {
		return Error("failure initializing program and getting user input: %v", err)
	} else if err = getStates(states, input.ndList, input.fields, nil, false, false); err != nil {
		return Error("failure getting states: %v", err)
	}

	// First check if any errors occurred getting the host records. If so, abort.
	if errors := checkStateErrors(states, duo, true); len(errors) != 0 {
		return Error("Aborting process; no records deleted.")
	}

	// Loop through the user provided input (name/data) list.
	space := input.maxNameLength + 8
	var numNotFound, numFailed uint

	for _, nameData := range input.ndList {
		records := states[nameData].records
		request := strings.TrimLeft(nameData, nameDataSep)
		request = strings.TrimRight(request, nameDataSep)

		if len(records) == 0 {
			Print("%-*s NOTFOUND\n", space, "Alias("+request+")")
			numNotFound++
			continue
		}
		record = records[0]
		ref := record.Ref
		_, err := deleteRecord(record.Ref, fields)
		if err != nil {
			Print("%-*s FAILED to delete: %v\n", space, "Alias("+request+")", err)
			numFailed++
			continue
		} else if ref != record.Ref {
			Print("%-*s FAILED to delete: ref mismatch\n", space, "Alias("+request+")")
			numFailed++
		} else {
			Print("%-*s Deleted\n", space, "Alias("+request+")")
		}
	}

	if numFailed != 0 {
		return Error("One or more updates failed")
	} else if numNotFound != 0 {
		return Error("One or more records not found")
	}

	return nil
}
