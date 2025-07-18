package main

import (
	. "github.com/dirtman/sitepkg"
	"strings"
)

// Implement the "delete" command.

func deleteAAAA(invokedAs []string) error {

	var fields []string
	var input *UserInput
	var states StatesAAAA = make(StatesAAAA)
	var record *RecordAAAA
	var err error
	duo := true

	SetStringOpt("view", "V", true, "default", "Specify the the record's view")
	SetStringOpt("filename", "f", true, "", "Specify an input file")

	if input, err = subCommandInit(invokedAs[1], invokedAs[2], duo); err != nil {
		return Error("failure initializing program and getting user input: %v", err)
	} else if err = getStates(states, input.ndList, input.fields, nil, false, false); err != nil {
		return Error("failure getting states: %v", err)
	}

	// First check if any errors occurred getting the host records. If so, abort.
	if errors := checkStateErrors(states, duo, true); len(errors) != 0 {
		return Error("Aborting process; no records deleted.")
	}
	// This must first be enabled via Grid DNS Properties:
	// Enable PTR record removal for A/AAAA records:
	// var removePtr bool;
	// SetBoolOpt("removePtr", "P", true, false, "Also remove the associated PTR record")
	// if removePtr, err = GetBoolOpt("removePtr"); err != nil {
	//    return Error("failure getting removePtr option: %v", err)
	// } else if removePtr {
	//	fields = []string{"remove_associated_ptr=true"}
	// }

	// Loop through the user provided input (name/data) list.
	space := input.maxNameLength + 8
	var numNotFound, numFailed uint

	for _, nameData := range input.ndList {
		records := states[nameData].records
		request := strings.TrimLeft(nameData, nameDataSep)
		request = strings.TrimRight(request, nameDataSep)

		if len(records) == 0 {
			Print("%-*s NOTFOUND\n", space, "AAAA("+request+")")
			numNotFound++
			continue
		}
		record = records[0]
		ref := record.Ref
		_, err := deleteRecord(record.Ref, fields)
		if err != nil {
			Print("%-*s FAILED to delete: %v\n", space, "AAAA("+request+")", err)
			numFailed++
			continue
		} else if ref != record.Ref {
			Print("%-*s FAILED to delete: ref mismatch\n", space, "AAAA("+request+")")
			numFailed++
		} else {
			Print("%-*s Deleted\n", space, "AAAA("+request+")")
		}
	}

	if numFailed != 0 {
		return Error("One or more updates failed")
	} else if numNotFound != 0 {
		return Error("One or more records not found")
	}

	return nil
}
