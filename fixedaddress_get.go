package main

import (
	"fmt"
	"strings"

	. "github.com/dirtman/sitepkg"
)

// Implement the "get" command.

func getFixedAddress(invokedAs []string) error {

	var input *UserInput
	var states StatesFixedAddress = make(StatesFixedAddress)
	var err, result error
	var duo, ref, inList bool

	SetStringOpt("view", "V", true, "default", "Specify the the network view to which the record belongs")
	SetStringOpt("fields", "F", true, "", "Specify fields to be used in the search")
	SetStringOpt("rFields", "R", true, "", "Specify additional fields to show in verbose mode")
	SetStringOpt("filename", "f", true, "", "Specify an input file")
	SetBoolOpt("ref", "r", true, false, "Show only the object \"reference\" of each fetched object")

	if input, err = subCommandInit(invokedAs[1], invokedAs[2], duo); err != nil {
		return Error("failure initializing program and getting user input: %v", err)
	}

	// The MAC address is not returned by default, so let's force it, as well as a few others.
	if inList, _ = InList(input.rFields, "name"); !inList {
		input.rFields = append(input.rFields, "name")
	}
	if inList, _ = InList(input.rFields, "mac"); !inList {
		input.rFields = append(input.rFields, "mac")
	}
	if inList, _ = InList(input.rFields, "bootfile"); !inList {
		input.rFields = append(input.rFields, "bootfile")
	}
	if inList, _ = InList(input.rFields, "nextserver"); !inList {
		input.rFields = append(input.rFields, "nextserver")
	}
	if inList, _ = InList(input.rFields, "network"); !inList {
		input.rFields = append(input.rFields, "network")
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

		// I prefer to keep the output short, only showing the user-specified name/data
		// fields.  But if the user provided no name or data, let's show the fields.
		if request == "" {
			request = strings.Join(input.fields, ",")
		}

		if err := states[nameData].err; err != nil {
			Print("%-*s FAILED: %v\n", space, "FixedAddress("+request+")", err)
			result = Error("one or more errors occurred")
		} else if len(records) == 0 {
			if !Quiet {
				Print("%-*s NOTFOUND\n", space, "FixedAddress("+request+")")
			}
			numNotFound++
		}

		for _, record := range records {
			if err = states[nameData].err; err != nil {
				Print("Failure getting %s: %v\n", request, err)
			} else if ref {
				Print("%-*s %s\n", space, "FixedAddress("+request+"): ", record.Ref)
			} else {
				data := record.Mac
				sep := " ("
				end := ""
				if input.view != "default" {
					data += fmt.Sprintf("%s%s mac", sep, record.NetviewName)
					sep = ", "
					end = ")"
				}
				if record.Disable {
					data += fmt.Sprintf("%s%s", sep, "DISABLED")
					end = ")"
				}
				data += end
				Print("%-*s %s %s\n", space, "FixedAddress("+request+"): ", record.IPv4Address, data)
			}
		}
	}
	if result == nil && numNotFound != 0 {
		return Error("One or more records not found.")
	}
	return result
}
