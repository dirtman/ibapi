package main

import (
	"strings"

	. "github.com/dirtman/sitepkg"
)

// Implement the "update" command.

func upateTXT(invokedAs []string) error {

	var input *UserInput
	var name, txt, message string
	var err error
	duo := true

	SetStringOpt("view", "V", true, "default", "Specify the view of the record to update")
	SetStringOpt("name", "n", false, "", "Update the record's name")
	SetStringOpt("comment", "c", true, "", "Update the record's comment")
	SetStringOpt("disable", "D", true, "", "Disable the specified record")
	SetStringOpt("txt", "t", false, "", "Update the record's text (TXT)")
	SetStringOpt("fields", "F", false, "", "Additional fields to be updated")
	SetStringOpt("filename", "f", true, "", "Specify a name/data input file")

	if input, err = subCommandInit(invokedAs[1], invokedAs[2], duo); err != nil {
		return Error("failure initializing program and getting user input: %v", err)
	} else if txt, err = GetStringOpt("txt"); err != nil {
		return Error("failure getting TXT option: %v", err)
	} else if name, err = GetStringOpt("name"); err != nil {
		return Error("failure getting name option: %v", err)
	}

	if name != "" { // Append it to the list of field/values to be updated.
		input.fields = append(input.fields, "name="+name)
	}
	if txt != "" { // Append it to the list of field/values to be updated.
		input.fields = append(input.fields, "text="+sanitizeTXT(txt))
	}

	// Query the record being updated, and check for errors.
	states := make(StatesTXT)
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
		request = unEscapeURLText(request)

		if len(records) == 0 {
			Print("%-*s NOTFOUND\n", space, "TXT("+request+")")
			numNotFound++
			continue
		}
		_, err = updateRecord(records[0].Ref, input.fields)
		message = "(fields: " + strings.Join(input.fields, ",") + ")"

		if err != nil {
			Print("%-*s FAILED to update: %v\n", space, "TXT("+request+")", err)
			numFailed++
			continue
		} else {
			Print("%-*s Updated %s\n", space, "TXT("+request+")", message)
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
