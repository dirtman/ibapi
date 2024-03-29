package main

import (
	"strings"

	. "github.com/dirtman/sitepkg"
)

// Implement the "delete" command.

func deleteHost(invokedAs []string) error {

	var fields []string
	var input *UserInput
	var states StatesHost = make(StatesHost)
	var record *RecordHost
	var err error
	duo := false

	SetStringOpt("view", "V", true, "default", "Specify the the record's view")
	SetStringOpt("filename", "f", true, "", "Specify an input file")
	SetBoolOpt("restartServices", "R", true, false, "Restart Grid services if needed")

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
			Print("%-*s NOTFOUND\n", space, "Host("+request+")")
			numNotFound++
			continue
		}
		record = records[0]
		ref := record.Ref
		_, err := deleteRecord(record.Ref, fields)
		if err != nil {
			Print("%-*s FAILED to delete: %v\n", space, "Host("+request+")", err)
			numFailed++
			continue
		} else if ref != record.Ref {
			Print("%-*s FAILED to delete: ref mismatch\n", space, "Host("+request+")")
			numFailed++
		} else {
			Print("%-*s Deleted\n", space, "Host("+request+")")
		}
	}

	if numFailed != 0 {
		return Error("One or more updates failed")
	} else if numNotFound != 0 {
		return Error("One or more records not found")
	} else if input.restartServices {
		if err = gridRestartServices(Verbose); err != nil {
			return Error("failure restarting services: %s", err)
		}
	}
	return nil
}
