package main

import (
	"fmt"
	. "github.com/dirtman/sitepkg"
	"strings"
)

// Implement the "get" command.

func getHost(invokedAs []string) error {

	var input *UserInput
	var states StatesHost = make(StatesHost)
	var err, result error
	var duo, ref bool

	SetStringOpt("view", "V", true, "any", "Specify the the record's view")
	SetStringOpt("fields", "F", true, "", "specify fields to be used in the search")
	SetStringOpt("rfields", "R", true, "", "specify additional fields to show in verbose mode")
	SetStringOpt("filename", "f", true, "", "Specify an input file")
	SetBoolOpt("ref", "r", true, false, "show only the object \"reference\" of each fetched object")
	SetStringOpt("ipFields", "I", false, "", "IP address fields to show in verbose mode.")

	if input, err = subCommandInit(invokedAs[1], invokedAs[2], duo); err != nil {
		return Error("failure initializing program and getting user input: %v", err)
	} else if ref, err = GetBoolOpt("ref"); err != nil {
		return Error("failure getting ref option: %v", err)
	}

	// GET allows a "parent.subject" argument syntax for return fields.
	// Let's take advantage, and append our ipFields to our rFields.
	if input.ipFields != nil && len(input.ipFields) > 0 {
		for _, field := range input.ipFields {
			input.rFields = append(input.rFields, "ipv4addrs."+field)
		}
	}
	err = getStates(states, input.ndList, input.fields, input.rFields, false, false)
	if err != nil {
		return Error("failure getting states: %v", err)
	} else if errors := checkStateErrors(states, false, true); errors != nil && len(errors) > 0 {
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
			Print("%-*s FAILED: %v\n", space, "Host("+request+")", err)
			result = Error("one or more errors occurred")
		} else if len(records) == 0 {
			if !Quiet {
				Print("%-*s NOTFOUND\n", space, "Host("+request+")")
			}
			numNotFound++
		}

		for _, record := range records {
			if err = states[nameData].err; err != nil {
				Print("Failure getting %s: %v\n", request, err)
			} else if ref {
				Print("%-*s %s\n", space, "Host("+request+"): ", record.Ref)
			} else {
				ips, _ := getHostIPs(record)
				if len(ips) == 0 {
					return Error("Failure getting IPs from Host record for %s", request)
				}
				data := strings.Join(ips, ", ")
                sep := " ("
                end := ""
                if input.view == "any" {
                    data += fmt.Sprintf("%s%s view", sep, record.View)
                    sep = ", "
                    end = ")"
                }
                if record.Disable {
                    data += fmt.Sprintf("%s%s", sep, "DISABLED")
                    end = ")"
                }
                data += fmt.Sprintf("%s", end)
				Print("%-*s %s %s\n", space, "Host("+request+"): ", record.Name, data)
			}
		}
	}
	if result == nil && numNotFound != 0 {
		return Error("One or more records not found.")
	}
	return result
}
