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
	var duo, multiple bool

	SetStringOpt("view", "V", true, "default", "Specify the view of the record to update")
	SetStringOpt("name", "n", false, "", "Update the record's name")
	SetStringOpt("comment", "c", true, "", "Update the record's comment")
	SetStringOpt("disable", "D", true, "", "Disable the specified record")
	SetStringOpt("txt", "t", false, "", "Update the record's text (TXT)")
	SetStringOpt("fields", "F", false, "", "Additional fields to be updated")
	SetStringOpt("filename", "f", true, "", "Specify a name/data input file")
	SetBoolOpt("multiple", "", true, false, "Allow multiple records to be updated")

	if input, err = subCommandInit(invokedAs[1], invokedAs[2], duo); err != nil {
		return Error("failure initializing program and getting user input: %v", err)
	} else if txt, err = GetStringOpt("txt"); err != nil {
		return Error("failure getting TXT option: %v", err)
	} else if name, err = GetStringOpt("name"); err != nil {
		return Error("failure getting name option: %v", err)
	} else if multiple, err = GetBoolOpt("multiple"); err != nil {
	    return Error("failure getting multiple option: %v", err)
	}

	if name != "" { // Append it to the list of field/values to be updated.
		input.fields = append(input.fields, "name="+name)
	}

	var prettyFields []string
	for _, field := range input.fields {
		prettyFields = append(prettyFields, field)
	}

	if txt != "" { // Append it to the list of field/values to be updated.
		input.fields = append(input.fields, "text="+sanitizeRecordData(txt))
		prettyFields = append(prettyFields, "text="+txt)
	}

	// Query the record being updated, and check for errors.
	states := make(StatesTXT)
	f := []string{"view=" + input.view}
	if err = getStates(states, input.ndList, f, nil, false, false); err != nil {
		return Error("failure getting states: %v", err)
	} else if errors := checkStateErrors(states, duo, true); len(errors) != 0 {
		return Error("Aborting process; no records updated.")
	}

    // Unless the --multiple option was specified, let's verify that only
    // a single record was found per input request.
    if ! multiple {
        for _, nameData := range input.ndList {
            if len(states[nameData].records) > 1 {
                return Error("Multiple records found for \"%s\"; see the --multiple option", nameData)
            }
        }
    }

	// Loop through the user provided input (name/data) list.
	space := input.maxNameLength + 8
	var numNotFound, numFailed uint

	for _, nameData := range input.ndList {
		records := states[nameData].records
        name, _, _ = splitND(nameData)
        request := name + nameDataSep + input.txtData[nameData]
        request = strings.TrimLeft(request, nameDataSep)
        request = strings.TrimRight(request, nameDataSep)

		if len(records) == 0 {
			Print("%-*s NOTFOUND\n", space, "TXT("+request+")")
			numNotFound++
			continue
		}
		for _, record := range records {
			_, err = updateRecord(record.Ref, input.fields)
			message = "(fields: " + strings.Join(prettyFields, ",") + ")"

			if err != nil {
				Print("%-*s FAILED to update: %v\n", space, "TXT("+request+")", err)
				numFailed++
				continue
			} else {
				Print("%-*s Updated %s\n", space, "TXT("+request+")", message)
			}
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
