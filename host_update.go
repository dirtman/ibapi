package main

import (
	. "github.com/dirtman/sitepkg"
	"strings"
)

// Implement the "update" command.

// According to the Infoblox WAPI docs:
//
// * Only PUT and POST methods can have a Body on input
// * The data to be updated can be given as arguments in the URL or the body of
//   the request (but not both).
// * You can specify only atomic values as arguments (i.e. booleans, integers, or
//   strings). You must use a method that contains a body if lists or structures
//   are needed.

func upateHost(invokedAs []string) error {

	var input *UserInput
	var name, ip, alias, message string
	var check bool
	var err error

	// Until we process the user's input, we can't tell what duo should be.  Chicken
	// or egg issue.  For instance, if one of the IP addresses in the Host's IP list
	// is being modified, we certainly need that IP address to be specified.  But if
	// you only want to update a simple field of the Host record, such as to disable
	// it, the user certainly does not need to specify an IP address.
	// Here we assume duo is false, but we may need to flip it to true before calling
	// checkStateErrors().
	duo := false

	SetStringOpt("view", "V", true, "default", "Specify the view of the record to update.")
	SetStringOpt("name", "n", false, "", "Update the record's name.")
	SetStringOpt("comment", "c", true, "", "Update the record's comment.")
	SetUintOpt("ttl", "", true, 0, "Update the the record's TTL.")
	SetStringOpt("disable", "D", true, "", "Disable the specified record")
	SetStringOpt("enableDNS", "e", false, "", "Configure host record for DNS.")
	SetStringOpt("fields", "F", false, "", "Additional fields to be updated.")
	SetStringOpt("filename", "f", true, "", "Specify a name/data input file.")
	SetBoolOpt("checkRecords", "C", true, false, "Check for existing related records.")

	// These involve lists as opposed to simple atomic values.
	SetStringOpt("ip", "i", false, "", "Update, add or delete an IP address.")
	SetStringOpt("alias", "a", false, "", "Update, add or delete a Host alias.")

	// These all pertain to updating fields of an IP address in the Host's IP list.
	SetStringOpt("enableDHCP", "d", false, "", "Configure the IP for DHCP")
	SetStringOpt("ipFields", "I", false, "", "IP address fields to be updated.")
	SetStringOpt("mac", "m", false, "", "Update the MAC of the IP address.")
	SetStringOpt("bootfile", "b", false, "", "Update the bootfile of the IP address.")
	SetStringOpt("nextserver", "N", false, "", "Update the nextserver of the IP address.")
	SetStringOpt("bootserver", "B", false, "", "Update the bootserver of the IP address.")

	if input, err = subCommandInit(invokedAs[1], invokedAs[2], duo); err != nil {
		return Error("failure initializing program and getting user input: %v", err)
	} else if ip, err = GetStringOpt("ip"); err != nil {
		return Error("failure getting IP option: %v", err)
	} else if name, err = GetStringOpt("name"); err != nil {
		return Error("failure getting name option: %v", err)
	} else if alias, err = GetStringOpt("alias"); err != nil {
		return Error("failure getting Alias option: %v", err)
	} else if check, err = GetBoolOpt("checkRecords"); err != nil {
		return Error("failure getting checkcheckRecords option: %v", err)
	}

	// If we are updating the Host's name itself, we'll use fields for the update.
	if name != "" { // Append the name to the list of field/values to be updated.
		input.fields = append(input.fields, "name="+name)
	}

	// If you want to update a field of a particular IP, you must specify that
	// IP.  Ditto if you want to update the IP itself to a different IP.
	if (ip != "" && !(strings.HasPrefix(ip, "-") || strings.HasPrefix(ip, "+"))) ||
		len(input.ipFields) != 0 {
		duo = true
	}

	// For sanity's sake, limit the types of updates that can be requested simultaneously.
	if len(input.fields) != 0 {
		if ip != "" {
			return Error("cannot update Host fields and an IP address at the same time")
		} else if alias != "" {
			return Error("cannot update Host fields and an Alias at the same time")
		} else if len(input.ipFields) != 0 {
			return Error("cannot update both the Host fields and IP fields together")
		}
	} else if len(input.ipFields) != 0 {
		if ip != "" {
			return Error("cannot update both an IP address and it's fields at the same time")
		} else if alias != "" {
			return Error("cannot update both IP address fields and an alias at the same time")
		}
	} else if ip != "" && alias != "" {
		return Error("cannot update both an IP address an alias together")
	}

	// The user can specify multiple name/data pairs to be updated only if:
	// * the Host's name is not being updated
	// * the Hosts's aliases or IPs are not being updated
	// * no IP fields are being updated.
	if len(input.ndList) != 1 {
		if name != "" || alias != "" || ip != "" || len(input.ipFields) != 0 {
			return Error("only one name/data pair can be specified for your update")
		}
	}

	// In order to update a specific field we need to be sure to fetch that field.
	// Note: this actually depends on how I make the update. But it won't hurt.
	var rFields []string
	for _, f := range input.fields {
		rFields = append(rFields, strings.SplitN(f, "=", 2)[0])
	}
	if alias != "" {
		rFields = append(rFields, "aliases")
	}

	// Query the record being updated, and check for errors.
	states := make(StatesHost)
	f := []string{"view=" + input.view}
	if err = getStates(states, input.ndList, f, rFields, false, false); err != nil {
		return Error("failure getting states: %v", err)
	} else if errors := checkStateErrors(states, duo, true); len(errors) != 0 {
		return Error("Aborting process; no records updated.")
	}
	space := input.maxNameLength + 8

	// If we are updating only simple fields, other than the Host's name itself,
	// we do not need to check for any possible conflicts.  Furthermore, this type
	// of update is the only type that allows multiple name/data pairs (i.e., more
	// than one Host to be updated at a time).  Let's get this case out the way.

	if name == "" && len(input.fields) > 0 {
		return updateHostFields(input, states, space)
	}

	// At this point, only one Host record can be updated at a time.
	if len(input.ndList) > 1 { // Should never happen
		return Error("multiple Host updates not allowed at this point")
	}

	nameData := input.ndList[0]
	records := states[nameData].records
	request := strings.TrimLeft(nameData, nameDataSep)
	request = strings.TrimRight(request, nameDataSep)
	if len(records) > 1 {
		return Error("multiple records found for single Host update")
	} else if len(records) == 0 {
		Print("%-*s NOTFOUND\n", space, "HOST("+request+")")
		return Error("Host not updated")
	}
	record := records[0]
	var conflict string

	// If we are updating the records's name itself, check for existing conflicts.
	if name != "" {
		f := []string{"view=" + input.view, "name=" + name}
		if conflict, err = checkConflict(f, true, check, check, true, true, "A"); err != nil {
			return Error("failure checking host conflicts: %v", err)
		}
	}

	// Ditto if we are adding/updating an alias.
	if alias != "" && !strings.HasPrefix(alias, "-") && conflict == "" {
		strippedAlias := strings.TrimPrefix(alias, "+")
		f := []string{"view=" + input.view, "name=" + strippedAlias}
		if conflict, err = checkConflict(f, true, true, check, true, true, "A"); err != nil {
			return Error("failure checking host conflicts: %v", err)
		}
	}

	// And if check is true, ditto if we are adding/updating an IP.
	if check && ip != "" && !strings.HasPrefix(ip, "-") && conflict == "" {
		strippedIP := strings.TrimPrefix(ip, "+")
		f := []string{"view=" + input.view, "ipv4addr=" + strippedIP}
		if conflict, err = checkConflict(f, true, true, false, false, false, ""); err != nil {
			return Error("failure checking host conflicts: %v", err)
		}
	}
	if conflict != "" {
		return Error("%-*s NOT updated: %s\n", space, "Host("+request+")", conflict)
	}

	// Make the appropriate update.
	_, data, _ := splitND(nameData)

	if len(input.fields) != 0 {
		return updateHostFields(input, states, space)
	} else if ip != "" {
		err = updateHostIP(record, data, ip)
		message = "(IP: " + ip + ")"
	} else if alias != "" {
		err = updateHostAlias(record, alias)
		message = "(Alias: " + alias + ")"
	} else if len(input.ipFields) != 0 {
		err = updateHostIPFields(record, data, input.ipFields)
		message = "(fields: " + strings.Join(input.ipFields, ",") + ")"
	} else {
		err = Error("don't know what to update")
	}

	if err != nil {
		Print("%-*s NOT updated: %v\n", space, "HOST("+request+")", err)
	} else {
		Print("%-*s Updated %s\n", space, "HOST("+request+")", message)
	}

	return nil
}

// Update the Hosts's fields.  This is the only update type that allows multiple
// hosts (name/data pairs) to be updated sumultaneously.

func updateHostFields(input *UserInput, states StatesHost, space int) error {

	// Loop through the user provided input (name/data) list.
	var numNotFound, numFailed uint
	var message string

	for _, nameData := range input.ndList {
		records := states[nameData].records
		request := strings.TrimLeft(nameData, nameDataSep)
		request = strings.TrimRight(request, nameDataSep)

		if len(records) == 0 {
			Print("%-*s NOTFOUND\n", space, "Host("+request+")")
			numNotFound++
			continue
		} else if len(records) > 1 { // Should NOT happen.
			Print("%-*s MULTIFOUND\n", space, "Host("+request+")")
			numNotFound++
			continue
		}

		_, err := updateRecord(records[0].Ref, input.fields)
		message = "(fields: " + strings.Join(input.fields, ",") + ")"

		if err != nil {
			Print("%-*s NOT updated: %v\n", space, "Host("+request+")", err)
			numFailed++
			continue
		} else {
			Print("%-*s Updated %s\n", space, "Host("+request+")", message)
		}
	}

	if numFailed != 0 {
		return Error("One or more updates failed")
	} else if numNotFound != 0 {
		return Error("One or more records not found")
	}
	return nil
}

// Change one of the Host's IP addresses, or add a new IP address, or remove
// an exising IP address.  For simplicity, we won't allow any Host or IP fields
// to be updated at the same time.
//
// To update a particular IP to a new IP, the particular IP must be specified
// via "currentIP", and the new IP is specified by "ip".
// To add a new IP, preceed ip with a '+' (+10.10.10.100).
// To remove an existing IP, preceed ip with a '-' (+10.10.10.100).
// In the latter 2 cases, "currentIP" is optional (it is ignored).
//
// When changing one of a Host's IP addresses, care must be taken not to lose
// any of the field settings of that IP address.

func updateHostIP(record *RecordHost, currentIP, ip string) error {

	var ipUpdateType string
	var hostRef = record.Ref
	var recordIpv4Addr HostRecordIpv4Addr
	var newHRIP HostRecordIpv4Addr
	var newHR struct {
		Ipv4Addrs []HostRecordIpv4Addr `json:"ipv4addrs"`
	}
	newHR.Ipv4Addrs = make([]HostRecordIpv4Addr, 0)

	if strings.HasPrefix(ip, "-") { // Delete the IP
		var found bool
		ipUpdateType = "deletion"
		ip = strings.TrimPrefix(ip, "-")
		for _, recordIpv4Addr = range record.Ipv4Addrs {
			if recordIpv4Addr.Ipv4Addr == ip {
				found = true
			} else {
				newHR.Ipv4Addrs = append(newHR.Ipv4Addrs, recordIpv4Addr)
			}
		}
		if !found {
			return Error("Host does not have IP %s", ip)
		}
	} else if strings.HasPrefix(ip, "+") { // Add the IP
		ip = strings.TrimPrefix(ip, "+")
		ipUpdateType = "addition"
		for _, recordIpv4Addr = range record.Ipv4Addrs {
			if recordIpv4Addr.Ipv4Addr == ip {
				return Error("Host already has IP %s", ip)
			}
		}
		newHRIP.Ipv4Addr = ip
		newHR.Ipv4Addrs = append(record.Ipv4Addrs, newHRIP)
	} else { // Update the IP address
		ipFields := []string{"ipv4addr=" + ip}
		return updateHostIPFields(record, currentIP, ipFields)
	}

	_ = ipUpdateType // We may want this later...
	url := "/" + hostRef
	body, ibapiErr := IBAPIPut(url, newHR)
	if ibapiErr != nil {
		return Error("%s", ibapiErr)
	}
	ShowDebug("body: %s", body)
	return nil
}

// Update one or more fields of one of a Host's IP addresses (without affecting
// all the other fields).

func updateHostIPFields(record *RecordHost, ip string, ipFields []string) error {

	var hripRef string

	// Find the IP address we want to update, and get its object ref.
	for _, recordIpv4Addr := range record.Ipv4Addrs {
		if recordIpv4Addr.Ipv4Addr == ip {
			hripRef = recordIpv4Addr.Ref
			break
		}
	}
	if hripRef == "" {
		return Error("Host does not have IP %s", ip)
	}
	_, err := updateRecord(hripRef, ipFields)
	return err
}

// Add or remove a Host alias.

func updateHostAlias(record *RecordHost, alias string) error {

	var hostRef = record.Ref
	var newHR struct {
		Aliases []string `json:"aliases"`
	}
	newHR.Aliases = make([]string, 0)

	if strings.HasPrefix(alias, "-") { // Delete the Alias
		var found bool
		alias = strings.TrimPrefix(alias, "-")
		for _, a := range record.Aliases {
			if a == alias {
				found = true
			} else {
				newHR.Aliases = append(newHR.Aliases, a)
			}
		}
		if !found {
			return Error("Host does not have Alias %s", alias)
		}
	} else { // Add the Alias
		if strings.HasPrefix(alias, "+") {
			alias = strings.TrimPrefix(alias, "+")
		}
		for _, a := range record.Aliases {
			if a == alias {
				return Error("Host already has Alias %s", alias)
			}
		}
		newHR.Aliases = append(record.Aliases, alias)
	}

	url := "/" + hostRef
	body, ibapiErr := IBAPIPut(url, newHR)
	if ibapiErr != nil {
		return Error("%s", ibapiErr)
	}
	ShowDebug("body: %s", body)
	return nil
}
