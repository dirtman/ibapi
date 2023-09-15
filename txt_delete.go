package main

import (
	"strings"

	. "github.com/dirtman/sitepkg"
)

// Implement the "delete" command.

func deleteTXT(invokedAs []string) error {

	var fields []string
	var input *UserInput
	var states StatesTXT = make(StatesTXT)
	var record *RecordTXT
	var err error
	duo := true

	SetStringOpt("view", "V", true, "default", "Specify the view to which the record belongs")
	SetStringOpt("filename", "f", true, "", "Specify an input file")

	if input, err = subCommandInit(invokedAs[1], invokedAs[2], duo); err != nil {
		return Error("failure initializing program and getting user input: %v", err)
	}
	if err = getStates(states, input.ndList, input.fields, nil, false, false); err != nil {
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
		request = unEscapeURLText(request)

		if len(records) == 0 {
			Print("%-*s NOTFOUND\n", space, "TXT("+request+")")
			numNotFound++
			continue
		}
		record = records[0]
		ref := record.Ref
		_, err := deleteRecord(record.Ref, fields)
		if err != nil {
			Print("%-*s FAILED to delete: %v\n", space, "TXT("+request+")", err)
			numFailed++
			continue
		} else if ref != record.Ref {
			Print("%-*s FAILED to delete: ref mismatch\n", space, "TXT("+request+")")
			numFailed++
		} else {
			Print("%-*s Deleted\n", space, "TXT("+request+")")
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
