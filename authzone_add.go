package main

import (
	"strings"

	. "github.com/dirtman/sitepkg"
)

// Implement the "add" command.

func addZoneAuth(invokedAs []string) error {

	var input *UserInput
	var states = make(StatesZoneAuth)
	var err error
	duo := false
	invokedString := strings.Join(invokedAs, ":")
	var nsgroup, zoneFormat string

	SetStringOpt("view", "V", true, "default", "Specify the the network view for the zone")
	SetStringOpt("nsgroup", "n", true, "", "Specify the name server group for the zone")
	SetStringOpt("zoneFormat", "z", true, "FORWARD", "Specify the zone_format for the zone")
	SetStringOpt("comment", "c", true, invokedString, "Specify the comment for the zone")
	SetBoolOpt("disable", "D", true, false, "Disable the new zone")
	SetStringOpt("fields", "F", true, "", "Specify additional fields for the zone")
	SetStringOpt("filename", "f", true, "", "Specify a name/data input file")
	SetBoolOpt("restartServices", "R", true, false, "Restart Grid services if needed")

	if input, err = subCommandInit(invokedAs[1], invokedAs[2], duo); err != nil {
		return Error("failure initializing program and getting user input: %v", err)
	} else if nsgroup, err = GetStringOpt("nsgroup"); err != nil {
		return Error("failure getting nsgroup option: %v", err)
	} else if zoneFormat, err = GetStringOpt("zoneFormat"); err != nil {
		return Error("failure getting zoneFormat option: %v", err)
	}
	if nsgroup == "" {
		if input.view == "external" {
			nsgroup = "external_Rice"
		}
	}
	if nsgroup != "" { // Append it to the list of field/values
		input.fields = append(input.fields, "ns_group="+nsgroup)
	}
	if zoneFormat != "" { // Append it to the list of field/values
		input.fields = append(input.fields, "zone_format="+zoneFormat)
	}

	// Check if any of the requested records already exist or conflict with existing
	// records in our view.
	f := []string{"view=" + input.view}
	ndList := input.ndList
	if err = getStates(states, ndList, f, nil, true, false); err != nil {
		return Error("failure getting states: %v", err)
	}

	// Check if any errors occurred getting the above records. If so, abort.
	if errors := checkStateErrors(states, false, true); len(errors) != 0 {
		return Error("Aborting process; no records added.")
	}

	// Loop through the user provided input (name/data) list.

	space := input.maxNameLength + 8
	nKey, dKey := states.GetNDKeys() // nKey == "" for ZoneAuth object
	object := states.GetObjectType()
	var numConflicts uint

	for nameData, state := range states {

		var name, data, conflict string
		sep := "Conflicts found: "
		name, data, _ = splitND(nameData)

		if len(states[nameData].records) != 0 {
			if state.records[0].Fqdn == name {
				conflict += sep + "ZoneAuth record with same name"
			} else { // Should never happen
				conflict += sep + "ZoneAuth record with same target"
			}
			sep = ", "
		}

		if conflict != "" {
			Print("%-*s NOT added: %s\n", space, "ZoneAuth("+nameData+")", conflict)
			numConflicts++
			continue
		}

		if _, err := addRecord(object, nKey, dKey, name, data, input.fields); err != nil {
			return Error("aborting! failure adding authoritative zone %s: %v", nameData, err)
		} else {
			Print("%-*s: Added\n", space, "ZoneAuth("+nameData+")")
		}
	}

	if numConflicts > 0 {
		if len(states) > 1 {
			return Error("One or more records not added due to conflict.")
		} else {
			return Error("Record not added due to conflict.")
		}
	} else if input.restartServices {
		if err = gridRestartServices(Verbose); err != nil {
			return Error("failure restarting services: %s", err)
		}
	}

	return nil
}
