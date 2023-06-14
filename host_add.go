package main

import (
	"strings"

	. "github.com/dirtman/sitepkg"
)

// Implement the "add" command.

func addHost(invokedAs []string) error {

	var input *UserInput
	var states = make(StatesHost)
	var statesA = make(StatesA)
	var statesCNAME = make(StatesCNAME)
	var statesAliasA = make(StatesAlias)
	var check bool
	var err error
	duo := true
	invokedString := strings.Join(invokedAs, ":")

	SetStringOpt("view", "V", true, "default", "Specify the view for the record")
	SetStringOpt("comment", "c", true, invokedString, "Specify the comment for the record")
	SetBoolOpt("disable", "D", true, false, "Disable the new record")
	SetUintOpt("ttl", "", true, 0, "Specify the TTL for the record")
	SetStringOpt("fields", "F", true, "", "specify additional fields for the record")
	SetStringOpt("filename", "f", true, "", "Specify a name/data input file")
	SetBoolOpt("checkRecords", "C", true, false, "Check for existing related records")
	SetBoolOpt("enableDNS", "e", true, true, "Configure host record for DNS")
	SetBoolOpt("restartServices", "R", true, false, "Restart Grid services if needed")

	// These all pertain to fields of the Host's IPv4 address.
	SetBoolOpt("enableDHCP", "d", false, false, "Configure the IP for DHCP")
	SetStringOpt("ipFields", "I", false, "", "IP address fields to be updated.")
	SetStringOpt("mac", "m", false, "", "Specify the MAC of the IP address.")
	SetStringOpt("bootfile", "b", false, "", "Specify the bootfile of the IP address.")
	SetStringOpt("nextserver", "N", false, "", "Specify the nextserver of the IP address.")
	SetStringOpt("bootserver", "B", false, "", "Specify the bootserver of the IP address.")

	if input, err = subCommandInit(invokedAs[1], invokedAs[2], duo); err != nil {
		return Error("failure initializing program and getting user input: %v", err)
	} else if check, err = GetBoolOpt("checkRecords"); err != nil {
		return Error("failure getting checkRecords option: %v", err)
	}

	// Check if any of the requested records to be added already exist.
	// If --checkRecords, also check for any "related" records.
	// Note we only care about existing records in the same view as ours.
	f := []string{"view=" + input.view}
	ndList := input.ndList
	if check { // Check for both host with same name and host with same data.
		err = getStates(states, ndList, f, nil, true, true)
	} else { // Check for host with the same name (avoid 400 WAPI error)
		err = getStates(states, ndList, f, nil, true, false)
	}
	if err != nil {
		return Error("failure getting states: %v", err)
	}

	// Check if any errors occurred getting the above records. If so, abort.
	if errors := checkStateErrors(states, false, true); len(errors) != 0 {
		return Error("Aborting process; no records added.")
	}

	// Check for any existing CNAME records with the same name.
	if err = getStates(statesCNAME, ndList, f, nil, true, false); err != nil {
		return Error("failure getting statesCNAME: %v", err)
	} else if errors := checkStateErrors(statesCNAME, true, true); len(errors) != 0 {
		return Error("Aborting process; no records added.")
	}

	// Check for any existing Alias records with the same name and target_type A.
	// To add support for AAAA records, will need to add a "target_type=AAAA" section.
	f = []string{"view=" + input.view, "target_type=A"}
	if err = getStates(statesAliasA, ndList, f, nil, true, false); err != nil {
		return Error("failure getting statesAliasA: %v", err)
	} else if errors := checkStateErrors(statesAliasA, true, true); len(errors) != 0 {
		return Error("Aborting process; no records added.")
	}

	// If check, check for other existing "related" records.
	if check {
		f = []string{"view=" + input.view}
		if err = getStates(statesA, ndList, f, nil, true, true); err != nil {
			return Error("failure getting statesA: %v", err)
		} else if errors := checkStateErrors(statesA, false, true); len(errors) != 0 {
			return Error("Aborting process; no records added.")
		}
	}

	// Loop through the user provided input (name/data) list.

	space := input.maxNameLength + 8
	var numConflicts uint

	for nameData, state := range states {

		var name, data, conflict string
		sep := "Conflicts found: "
		name, data, _ = splitND(nameData)

		if len(states[nameData].records) != 0 {
			if state.records[0].Name == name {
				conflict += sep + "Host record with same name"
			} else {
				conflict += sep + "Host record with same IP"
			}
			sep = ", "
		}

		if len(statesCNAME[nameData].records) != 0 {
			conflict += sep + "CNAME with same name"
			sep = ", "
		}
		if len(statesAliasA[nameData].records) != 0 {
			conflict += sep + "Alias with same name"
			sep = ", "
		}
		if check && len(statesA[nameData].records) != 0 {
			conflict += sep + "related A record"
			sep = ", "
		}

		if conflict != "" {
			Print("%-*s NOT added: %s\n", space, "Host("+nameData+")", conflict)
			numConflicts++
			continue
		}

		if _, err := addRecordHost(name, data, input); err != nil {
			return Error("aborting! failure adding host record %s: %v", nameData, err)
		} else {
			Print("%-*s: Added\n", space, "Host("+nameData+")")
		}
	}

	if numConflicts > 0 {
		if len(states) > 1 {
			return Error("One or more records not added due to conflict.")
		} else {
			return Error("Record not added due to conflict.")
		}
	} else if input.restartServices {
		if err = restartGridServices(Verbose); err != nil {
			return Error("failure restarting services: %s", err)
		}
	}
	return nil

}

func addRecordHost(name, data string, input *UserInput) ([]byte, error) {

	var url, inputBody string

	url = "/" + "record:host"
	inputBody = `{"name":"` + name + `"`

	inputBody = appendFieldsJSON(inputBody, input.fields)

	if strings.Count(data, ":") > 0 {
		inputBody += `,"ipv6addrs":[{"ipv6addr":"` + data + `"`
	} else {
		inputBody += `,"ipv4addrs":[{"ipv4addr":"` + data + `"`
	}

	inputBody = appendFieldsJSON(inputBody, input.ipFields)
	inputBody += "}]}"
	//	inputBody += `,"configure_for_dhcp":` + input.enableDHCP + "}]}"
	ShowDebug("inputBody: %s", inputBody)
	ShowDebug("inputBody: %v", inputBody)
	ShowDebug("inputBody: %#v", inputBody)

	body, ibapiErr := IBAPIPost(url, inputBody)
	if ibapiErr != nil {
		return nil, Error("%s", ibapiErr)
	}
	ShowDebug("body: %s", body)
	return body, nil
}
