package main

import (
	"fmt"
	"strings"

	. "github.com/dirtman/sitepkg"
)

// Implement the "get" command.

func getTXT(invokedAs []string) error {

	var input *UserInput
	var states StatesTXT = make(StatesTXT)
	var err, result error
	var duo, ref bool
	var txt string

	SetStringOpt("view", "V", true, "default", "Specify the the view to which the record belongs")
	SetStringOpt("fields", "F", true, "", "Specify fields to be used in the search")
	SetStringOpt("rfields", "R", true, "", "Specify additional fields to show in verbose mode")
	SetStringOpt("filename", "f", true, "", "Specify an input file")
	SetStringOpt("txt", "t", false, "", "Specify the TXT of the record to fetch")
	SetBoolOpt("ref", "r", true, false, "Show only the object \"reference\" of each fetched object")

	if input, err = subCommandInit(invokedAs[1], invokedAs[2], duo); err != nil {
		return Error("failure initializing program and getting user input: %v", err)
	} else if txt, err = GetStringOpt("txt"); err != nil {
		return Error("failure getting TXT option: %v", err)
	}

	if txt != "" {
		input.fields = append(input.fields, "text="+sanitizeTXT(txt))
	}

	if err = getStates(states, input.ndList, input.fields, input.rFields, false, false); err != nil {
		return Error("failure getting states: %v", err)
	} else if ref, err = GetBoolOpt("ref"); err != nil {
		return Error("failure getting ref option: %v", err)
	}

	if errors := checkStateErrors(states, false, true); len(errors) > 0 {
		return Error("Aborting process; no records fetched.")
	}

	// Loop through the user provided input (name/data) list.
	space := input.maxNameLength + 8
	var numNotFound uint

	for _, nameData := range input.ndList {
		records := states[nameData].records
		request := strings.TrimLeft(nameData, nameDataSep)
		request = strings.TrimRight(request, nameDataSep)
		request = unEscapeURLText(request)

		// I prefer to keep the output short, only showing the user-specified name/data
		// fields.  But if the user provided no name or data, let's show the fields.
		if request == "" {
			request = unEscapeURLText(strings.Join(input.fields, ","))
		} else {
			request = unEscapeURLText(request)
		}

		if err := states[nameData].err; err != nil {
			Print("%-*s FAILED: %v\n", space, "TXT("+request+")", err)
			result = Error("one or more errors occurred")
		} else if len(records) == 0 {
			if !Quiet {
				Print("%-*s NOTFOUND\n", space, "TXT("+request+")")
			}
			numNotFound++
		}

		for _, record := range records {
			if err = states[nameData].err; err != nil {
				Print("Failure getting %s: %v\n", request, err)
			} else if ref {
				Print("%-*s %s\n", space, "TXT("+request+"): ", record.Ref)
			} else {
				data := record.Text
				sep := " ("
				end := ""
				if input.view != "default" {
					data += fmt.Sprintf("%s%s view", sep, record.View)
					sep = ", "
					end = ")"
				}
				if record.Disable {
					data += fmt.Sprintf("%s%s", sep, "DISABLED")
					end = ")"
				}
				data += end
				Print("%-*s %s %s\n", space, "TXT("+request+"): ", record.Name, data)
			}
		}
	}
	if result == nil && numNotFound != 0 {
		return Error("One or more records not found.")
	}
	return result
}
