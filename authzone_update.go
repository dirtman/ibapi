package main

import (
	"strings"

	. "github.com/dirtman/sitepkg"
)

// Implement the "update" command.

func upateZoneAuth(invokedAs []string) error {

	var input *UserInput
	var nsgroup, message string
	var err error
	duo := false

	SetStringOpt("view", "V", true, "default", "Specify the view of the record to update")
	SetStringOpt("nsgroup", "n", false, "", "Update the name server group.")
	SetStringOpt("comment", "c", true, "", "Update the record's comment")
	SetStringOpt("disable", "D", true, "", "Disable the specified record")
	SetStringOpt("fields", "F", false, "", "Additional fields to be updated")
	SetStringOpt("filename", "f", true, "", "Specify a name/data input file")
	SetBoolOpt("restartServices", "R", true, false, "Restart Grid services if needed")

	if input, err = subCommandInit(invokedAs[1], invokedAs[2], duo); err != nil {
		return Error("failure initializing program and getting user input: %v", err)
	} else if nsgroup, err = GetStringOpt("nsgroup"); err != nil {
		return Error("failure getting nsgroup option: %v", err)
	}

	if nsgroup != "" { // Append it to the list of field/values to be updated.
		input.fields = append(input.fields, "ns_group="+nsgroup)
	}

	// Query the record being updated, and check for errors.
	states := make(StatesZoneAuth)
	f := []string{"view=" + input.view}
	if err = getStates(states, input.ndList, f, nil, false, false); err != nil {
		return Error("failure getting states: %v", err)
	} else if errors := checkStateErrors(states, duo, true); len(errors) != 0 {
		return Error("Aborting process; no records updated.")
	}
	space := input.maxNameLength + 8

	// Loop through the user provided input (name/data) list.
	var numNotFound, numFailed uint

	for _, nameData := range input.ndList {
		records := states[nameData].records
		request := strings.TrimLeft(nameData, nameDataSep)
		request = strings.TrimRight(request, nameDataSep)

		if len(records) == 0 {
			Print("%-*s NOTFOUND\n", space, "ZoneAuth("+request+")")
			numNotFound++
			continue
		}
		_, err = updateRecord(records[0].Ref, input.fields)
		message = "(fields: " + strings.Join(input.fields, ",") + ")"

		if err != nil {
			Print("%-*s FAILED to update: %v\n", space, "ZoneAuth("+request+")", err)
			numFailed++
			continue
		} else {
			Print("%-*s Updated %s\n", space, "ZoneAuth("+request+")", message)
		}
	}

	if numFailed != 0 {
		return Error("One or more updates failed")
	} else if numNotFound != 0 {
		return Error("One or more records not found")
	} else if input.restartServices {
		if err := gridRestartServices(Verbose); err != nil {
			return Error("failure restarting services: %s", err)
		}
	}
	return nil
}
