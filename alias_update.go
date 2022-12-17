package main

import (
	. "github.com/dirtman/sitepkg"
	"strings"
)

// Implement the "update" command.

func upateAlias(invokedAs []string) error {

	var input *UserInput
	var name, target, message string
	var check bool
	var err error
	duo := false

	SetStringOpt("view", "V", true, "default", "Specify the view of the record to update")
	SetStringOpt("name", "n", false, "", "Update the record's name")
	SetStringOpt("comment", "c", true, "", "Update the record's comment")
	SetUintOpt("ttl", "", true, 0, "Update the the record's TTL")
	SetStringOpt("disable", "D", true, "", "Disable the specified record")
	SetStringOpt("target", "t", false, "", "Update the record's target (Alias)")
	SetStringOpt("fields", "F", false, "", "Additional fields to be updated")
	SetStringOpt("filename", "f", true, "", "Specify a name/data input file")
	SetBoolOpt("checkRecords", "C", true, false, "Check for existing related records")
	SetStringOpt("targetType", "T", false, "A", "Specify the target type of the alias to update")

	if input, err = subCommandInit(invokedAs[1], invokedAs[2], duo); err != nil {
		return Error("failure initializing program and getting user input: %v", err)
	} else if target, err = GetStringOpt("target"); err != nil {
		return Error("failure getting target option: %v", err)
	} else if name, err = GetStringOpt("name"); err != nil {
		return Error("failure getting name option: %v", err)
	} else if check, err = GetBoolOpt("checkRecords"); err != nil {
		return Error("failure getting checkcheckRecords option: %v", err)
	}

	if name != "" { // Append it to the list of field/values to be updated.
		input.fields = append(input.fields, "name="+name)
	}
	if target != "" { // Append it to the list of field/values to be updated.
		input.fields = append(input.fields, "target_name="+target)
	}

	// Query the record being updated, and check for errors.
	states := make(StatesAlias)
	f := []string{"view=" + input.view, "target_type=" + input.targetType}
	if err = getStates(states, input.ndList, f, nil, false, false); err != nil {
		return Error("failure getting states: %v", err)
	} else if errors := checkStateErrors(states, duo, true); len(errors) != 0 {
		return Error("Aborting process; no records updated.")
	}
	space := input.maxNameLength + 8

	// If we are updating the record's name itself, check for existing name conflicts
	// once, before going into the loop.  And if we are updating the record's target, check
	// for existing target conflicts (if check).  If a conflict is found, there is no need
	// to go into the loop, since the conflict will exist for each user-specified host.
	request := strings.TrimLeft(input.ndList[0], nameDataSep)
	request = strings.TrimRight(request, nameDataSep)
	var conflict string
	tt := input.targetType
	if name != "" {
		f := []string{"view=" + input.view, "name=" + name}
		if conflict, err = checkConflict(f, check, true, true, true, true, tt); err != nil {
			return Error("failure checking host conflicts: %v", err)
		}
	}
	if conflict != "" {
		return Error("%-*s NOT updated: %s\n", space, "Alias("+request+")", conflict)
	}

	// Loop through the user provided input (name/data) list.
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
		_, err = updateRecord(records[0].Ref, input.fields)
		message = "(fields: " + strings.Join(input.fields, ",") + ")"

		if err != nil {
			Print("%-*s FAILED to update: %v\n", space, "Alias("+request+")", err)
			numFailed++
			continue
		} else {
			Print("%-*s Updated %s\n", space, "Alias("+request+")", message)
		}
	}

	if numFailed != 0 {
		return Error("One or more updates failed")
	} else if numNotFound != 0 {
		return Error("One or more records not found")
	}
	return nil
}
