package main

import (
	. "github.com/dirtman/sitepkg"
	"strings"
)

// Implement the "update" command.

func upatePTR(invokedAs []string) error {

	var input *UserInput
	var name, ip, message string
	var check bool
	var err error
	duo := true

	SetStringOpt("view", "V", true, "default", "Specify the view of the record to update")
	SetStringOpt("name", "n", false, "", "Update the record's name")
	SetStringOpt("comment", "c", true, "", "Update the record's comment")
	SetUintOpt("ttl", "", true, 0, "Update the the record's TTL")
	SetStringOpt("disable", "D", true, "", "Disable the specified record")
	SetStringOpt("ip", "i", false, "", "Update the record's IP address")
	SetStringOpt("fields", "F", false, "", "Additional fields to be updated")
	SetStringOpt("filename", "f", true, "", "Specify a name/data input file")
	SetBoolOpt("checkRecords", "C", true, false, "Check for existing related records")

	if input, err = subCommandInit(invokedAs[1], invokedAs[2], duo); err != nil {
		return Error("failure initializing program and getting user input: %v", err)
	} else if ip, err = GetStringOpt("ip"); err != nil {
		return Error("failure getting IP option: %v", err)
	} else if name, err = GetStringOpt("name"); err != nil {
		return Error("failure getting name option: %v", err)
	} else if check, err = GetBoolOpt("checkRecords"); err != nil {
		return Error("failure getting checkcheckRecords option: %v", err)
	}

	if name != "" { // Append it to the list of field/values to be updated.
		input.fields = append(input.fields, "ptrdname="+name)
	}
	if ip != "" { // Append it to the list of field/values to be updated.
		input.fields = append(input.fields, "ipv4addr="+ip)
	}

	// Query the record being updated, and check for errors.
	states := make(StatesPTR)
	f := []string{"view=" + input.view}
	if err = getStates(states, input.ndList, f, nil, false, false); err != nil {
		return Error("failure getting states: %v", err)
	} else if errors := checkStateErrors(states, duo, true); len(errors) != 0 {
		return Error("Aborting process; no records updated.")
	}
	space := input.maxNameLength + 8

	// If we are updating the record's name itself, check for existing name conflicts
	// once, before going into the loop.  And if we are updating the record's IP, check
	// for existing IP conflicts (if check).  If a conflict is found, there is no need
	// to go into the loop, since the conflict will exist for each user-specified host.
	request := strings.TrimLeft(input.ndList[0], nameDataSep)
	request = strings.TrimRight(request, nameDataSep)
	var conflict string
	if name != "" {
		f := []string{"view=" + input.view, "name=" + name}
		if conflict, err = checkConflict(f, true, check, check, true, true, "A"); err != nil {
			return Error("failure checking host conflicts: %v", err)
		}
	}
	if check && ip != "" && conflict == "" {
		f := []string{"view=" + input.view, "ipv4addr=" + ip}
		if conflict, err = checkConflict(f, true, true, false, false, false, ""); err != nil {
			return Error("failure checking host conflicts: %v", err)
		}
	}
	if conflict != "" {
		return Error("%-*s NOT updated: %s\n", space, "PTR("+request+")", conflict)
	}

	// Loop through the user provided input (name/data) list.
	var numNotFound, numFailed uint

	for _, nameData := range input.ndList {
		records := states[nameData].records
		request := strings.TrimLeft(nameData, nameDataSep)
		request = strings.TrimRight(request, nameDataSep)

		if len(records) == 0 {
			Print("%-*s NOTFOUND\n", space, "PTR("+request+")")
			numNotFound++
			continue
		}
		_, err = updateRecord(records[0].Ref, input.fields)
		message = "(fields: " + strings.Join(input.fields, ",") + ")"

		if err != nil {
			Print("%-*s FAILED to update: %v\n", space, "PTR("+request+")", err)
			numFailed++
			continue
		} else {
			Print("%-*s Updated %s\n", space, "PTR("+request+")", message)
		}
	}

	if numFailed != 0 {
		return Error("One or more updates failed")
	} else if numNotFound != 0 {
		return Error("One or more records not found")
	}
	return nil
}
