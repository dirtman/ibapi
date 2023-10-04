package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"

	. "github.com/dirtman/sitepkg"
)

/*****************************************************************************\

  This file deals with user input for the Infoblox object(s) to be processed.

  A DNS record has a "name" and resource-type specific "data" value, as well
  as a "TTL" and a few other "fields".  Infoblox WAPI objects have introduced
  many additional fields, and sometimes the concept of a name and data pair
  does not fit well with Infoblox objects - everything is just a field.
  However, this CLI does separate out a "name" and "data" field from the other
  fields.  While to Infoblox everything is a field, this CLI expects the name
  and data fields, if present, to be specified as command line arguments or
  in a file.  The only exception is Get requests, since these support search
  modifiers which can only be expressed via the --fields option, such as in
  "--fields name~=.seci.rice.edu".

  For Add, Update and Delete requests the user must provide a name/data pair
  for each object to be added, updated or deleted.  Depending on the object
  type and operation, the name or data may be empty.  For Get requests, as
  mentioned, name/data pairs are optional.

  The user specifies the name/data pairs via command line arguments, or via a
  file (-f/--file) containing a list of name/data value pairs.

  Fields are specified with the -F/--fields option, or in a few cases a
  separate option is defined for a specific field, such as --view.

\*****************************************************************************/

const nameDataSep = "/"
const requestTypeCreate = 1
const requestTypeGet = 2
const requestTypeDelete = 3
const requestTypeUpdate = 4
const objectTypeHost = 1
const objectTypeA = 2
const objectTypePTR = 3
const objectTypeCNAME = 4
const objectTypeAlias = 5
const objectTypeFixedAddress = 6
const objectTypeMX = 7
const objectTypeTXT = 8
const viewAny = "any"
const targetAny = "any"

type UserInput struct {
	objectType      int      // Type of object specified by user
	operation       int      // Operation requested by user, such as get or add
	ndList          []string // The list of name/data pairs provided by the user
	maxNameLength   int      // For prettier print-out formatting later.
	view            string   // --view value
	comment         string   // --comment value
	disable         string   // --disable: handled differently for Add vs Update
	ttl             uint32   // --ttl value
	fields          []string // --fields
	rFields         []string // --rFields, for Get requests
	enableDNS       string   // --enableDNS, for Host
	enableDHCP      string   // --enableDCHP, for Host
	mac             string   // --mac, for Host
	bootfile        string   // --bootfile, for Host
	nextserver      string   // --nextserver, for Host
	bootserver      string   // --bootserver, for Host
	ipFields        []string // --ipFields, for Host
	restartServices bool     // --restart_if_needed, for Host
	targetType      string   // --targetType, for Alias and CNAME
	txtData			map[string]string	// original, pre-sanitized TXT data.
}

// Process and store the user arguments that define the objects on which to
// operate.  Some requests require both a "name" and a "data" value as input,
// others require one and/or the other, and "get" requests can get by with
// neither, as long as one or more "field" option/value pairs are provided.
// Except for get requests, "duo" determines whether both a name and a data
// value are required.
// In addition to handling user arguments, some commonly used options are
// also handled here.

func getUserInput(objectType, operation string, duo bool, args []string) (*UserInput, error) {

	input := new(UserInput)
	input.ndList = make([]string, 0)
	var err error

	if operation == "add" || operation == "create" {
		input.operation = requestTypeCreate
	} else if operation == "get" || operation == "fetch" {
		input.operation = requestTypeGet
	} else if operation == "delete" {
		input.operation = requestTypeDelete
	} else if operation == "update" {
		input.operation = requestTypeUpdate
	}

	if objectType == "host" {
		input.objectType = objectTypeHost
	} else if objectType == "address" || objectType == "a" {
		input.objectType = objectTypeA
	} else if objectType == "ptr" {
		input.objectType = objectTypePTR
	} else if objectType == "cname" {
		input.objectType = objectTypeCNAME
	} else if objectType == "alias" {
		input.objectType = objectTypeAlias
	} else if objectType == "fixedaddress" {
		input.objectType = objectTypeFixedAddress
	} else if objectType == "mx" {
		input.objectType = objectTypeMX
	} else if objectType == "txt" {
		input.objectType = objectTypeTXT
		input.txtData = make(map[string]string, 0)
	}

	if err = GetFieldOptions(input); err != nil {
		return nil, Error("failure processing fields options: %v", err)
	}
	ShowDebug("getUserInput: input.fields: %#v", input.fields)
	ShowDebug("getUserInput: input.rFields: %#v", input.rFields)

	if filename, _ := GetStringOpt("filename"); filename != "" {
		err = getUserInputFromFile(filename, input, duo, args)
	} else if len(args) > 0 {
		err = getUserInputFromArgs(input, duo, args)
	} else if input.operation != requestTypeGet {
		err = Error("a name and/or data value is required")
	} else if input.fields == nil || len(input.fields) == 0 {
		err = Error("no name, data or fields provided")
	} else {
		input.ndList = append(input.ndList, "/")
	}
	ShowDebug("getUserInput: input.ndList: %#v", input.ndList)
	return input, err
}

// Process the various options that will be part of the WAPI request.
// Most operations support a --fields option, and this string is split
// into a string array and saved as input.fields.  Several other
// options, such as --view, are appended to input.fields.

func GetFieldOptions(input *UserInput) error {

	var ttl uint
	var err error

	// Each operation except Delete has a field option.
	if input.operation != requestTypeDelete {
		if fields, _ := GetStringOpt("fields"); fields != "" {
			input.fields = strings.Split(fields, ",")
		}
	}

	// We are transitioning to multiple views, and it seems it will be safer
	// to always require the view to be specified.  May need to change this...
	// Note the view of an object cannot be updated.  For Update, the view
	// option specifies the object to be updated, not a field to be updated.
	if input.view, _ = GetStringOpt("view"); input.view == "" {
		return Error("the \"view\" option cannot be empty")
	} else if (!(input.view == viewAny && input.operation == requestTypeGet)) &&
		input.operation != requestTypeUpdate {
		if input.objectType == objectTypeFixedAddress {
			input.fields = append(input.fields, "network_view="+input.view)
		} else {
			input.fields = append(input.fields, "view="+input.view)
		}
	}

	// A "target_type" is required for Alias records.  Note that for Update, the
	// target_type specifies the object to be updated, not a field to be updated.
	if input.objectType == objectTypeAlias {
		if input.targetType, _ = GetStringOpt("targetType"); input.targetType == "" {
			return Error("the \"targetType\" option cannot be empty")
		} else if (!(input.targetType == targetAny && input.operation == requestTypeGet)) &&
			input.operation != requestTypeUpdate {
			input.fields = append(input.fields, "target_type="+input.targetType)
		}
	}

	// A host record has a "restart_if_needed" field, and host add, update and
	// delete have the --resetartService option to reflect this.  But note this
	// field is not searchable; my guess is that it is not saved with the host
	// record object, but is more of a one-time flag.  We'll see...
	if (input.objectType == objectTypeHost || input.objectType == objectTypeFixedAddress) &&
		input.operation != requestTypeGet {
		input.restartServices, _ = GetBoolOpt("restartServices")
	}

	// For Deletes we're done.
	if input.operation == requestTypeDelete {
		return nil
	}

	// Get supports a return fields option.
	if input.operation == requestTypeGet {
		if rFields, _ := GetStringOpt("rFields"); rFields != "" {
			input.rFields = strings.Split(rFields, ",")
		}
		// Let's force the "disable" fields, since we may want to let
		// the user know if the fetched record is disabled or not.
		if inList, _ := InList(input.rFields, "disable"); !inList {
			input.rFields = append(input.rFields, "disable")
		}
		// Get MX supports an mx and preference fields.
		if input.objectType == objectTypeMX {
			var mx, preference string
			if mx, err = GetStringOpt("mx"); err != nil {
				return Error("failure getting MX option: %v", err)
			} else if preference, err = GetStringOpt("preference"); err != nil {
				return Error("failure getting preference option: %v", err)
			}
			if mx != "" {
				input.fields = append(input.fields, "mail_exchanger="+mx)
			}
			if preference != "" {
				input.fields = append(input.fields, "preference="+preference)
			}
		}
		// Get TXT supports a txt field.
		if input.objectType == objectTypeTXT {
			var txt string
			if txt, err = GetStringOpt("txt"); err != nil {
				return Error("failure getting TXT option: %v", err)
			}
			if txt != "" {
				input.fields = append(input.fields, "text="+sanitizeRecordData(txt))
			}
		}
	}

	// Handle the comment option.
	if input.comment, _ = GetStringOpt("comment"); input.comment != "" {
		// Replace each space with its hex code (%20).
		comment := strings.ReplaceAll(input.comment, " ", "%20")
		input.fields = append(input.fields, "comment="+comment)
	}

	// The infoblox-go-client package defines a TTL as a uint32.
	if ttl, err = GetUintOpt("ttl"); err != nil &&
		!strings.Contains(err.Error(), ConfErrNoSuchOption) {
		return Error("failure parsing ttl option: %v", err)
	} else if ttl != 0 {
		input.ttl = uint32(ttl)
		input.fields = append(input.fields, "ttl="+fmt.Sprintf("%d", int(ttl)))
	}

	// The disable option is type boolean for Add, but type string for Update.
	if input.disable, err = getStringBool("disable", "", input, &input.fields); err != nil {
		return Error("%v", err)
	}

	// ipFields is only for host records.
	if input.objectType == objectTypeHost {
		if ipFields, _ := GetStringOpt("ipFields"); ipFields != "" {
			input.ipFields = strings.Split(ipFields, ",")
		}
	}

	// The rest only applies to host and fixedaddress records
	if input.objectType != objectTypeHost && input.objectType != objectTypeFixedAddress {
		return nil
	}

	// Handle the mac option.
	if input.mac, _ = GetStringOpt("mac"); input.mac != "" {
		if input.objectType == objectTypeHost {
			input.ipFields = append(input.ipFields, "mac="+input.mac)
		} else if input.operation != requestTypeCreate {
			input.fields = append(input.fields, "mac="+input.mac)
		}
	}

	// Handle the bootfile option.
	if input.bootfile, _ = GetStringOpt("bootfile"); input.bootfile != "" {
		if input.objectType == objectTypeHost {
			input.ipFields = append(input.ipFields, "bootfile="+input.bootfile)
		} else {
			input.fields = append(input.fields, "bootfile="+input.bootfile)
		}
	}

	// Handle the nextserver option.
	if input.nextserver, _ = GetStringOpt("nextserver"); input.nextserver != "" {
		if input.objectType == objectTypeHost {
			input.ipFields = append(input.ipFields, "nextserver="+input.nextserver)
		} else {
			input.fields = append(input.fields, "nextserver="+input.nextserver)
		}
	}

	// Handle the bootserver option.
	if input.bootserver, _ = GetStringOpt("bootserver"); input.bootserver != "" {
		if input.objectType == objectTypeHost {
			input.ipFields = append(input.ipFields, "bootserver="+input.bootserver)
		} else {
			input.fields = append(input.fields, "bootserver="+input.bootserver)
		}
	}

	// The rest only apply to host records
	if input.objectType != objectTypeHost {
		return nil
	}

	// enableDNS is type boolean for Add, but type string for Update.
	if input.enableDNS, err =
		getStringBool("enableDNS", "configure_for_dns", input, &input.fields); err != nil {
		return Error("%v", err)
	}

	// enableDHCP is type boolean for Add, but type string for Update.
	if input.enableDHCP, err =
		getStringBool("enableDHCP", "configure_for_dhcp", input, &input.ipFields); err != nil {
		return Error("%v", err)
	}

	return nil
}

// The "update" and "add" commands all support the "disable" option. Host records
// additionally support the "configure_for_dns" option, and a Host's address records
// support the "configure_for_dhcp" option.  For the add commands, these are defined
// as boolean options, but for the update commands they are string options since 3
// possible values are needed: update to true, update to false, or don't update.
// userOpt is the option presented to the user.
// ibOpt is the corresponding Infoblox setting.
// fields specifies a fields slice to which the option should be added.

func getStringBool(userOpt, ibOpt string, input *UserInput, fields *[]string) (string, error) {

	var boolString string
	var boolean bool
	var err error

	// If the Infloblox option name is not provided, it defaults to the user option name.
	if ibOpt == "" {
		ibOpt = userOpt
	}

	if input.operation == requestTypeUpdate {
		if boolString, _ = GetStringOpt(userOpt); boolString != "" {
			if boolean, err = StringToBool(boolString); err != nil {
				return "", Error("invalid value \"%s\" for %s option", boolString, userOpt)
			} else if boolean {
				boolString = "true"
			} else {
				boolString = "false"
			}
		}
	} else if input.operation == requestTypeCreate {
		if boolean, _ = GetBoolOpt(userOpt); boolean {
			boolString = "true"
		} else {
			boolString = "false"
		}
	}
	if boolString != "" && fields != nil {
		*fields = append(*fields, ibOpt+"="+boolString)
	}

	return boolString, nil
}

// Get name/data input from the user's command line arguments.

func getUserInputFromArgs(input *UserInput, duo bool, args []string) error {

	var name, data string
	var err error

	if name, data, err = getND(input, args); err != nil {
		return err
	} else if duo && (name == "" || data == "") {
		return Error("both a name and a data value must be specified")
	} else if input.objectType == objectTypeHost && name == "" &&
		input.operation != requestTypeGet {
		return Error("a name must be specified")
	} else if name == "" && data == "" {
		return Error("a name and/or a data value must be specified")
	}

	input.ndList = append(input.ndList, name+nameDataSep+data)
	return nil
}

// Get name/data inputs from a user specified input file.

func getUserInputFromFile(filename string, input *UserInput, duo bool, args []string) error {

	var name, data string
	var file *os.File
	var lineNo, fl int
	var err error

	if len(args) > 0 {
		return Error("no arguments allowed when an input file is specified.")
	} else if filename == "" {
		return Error("bad call: filename not defined")
	} else if filename == "-" {
		file = os.Stdin
	} else {
		file, err = os.Open(filename)
		if err != nil {
			return Error("failure opening file \"%s\": %v", filename, err)
		}
		defer file.Close()
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineNo++
		// Remove leading and trailing spaces and tabs:
		line := strings.TrimLeft(scanner.Text(), " \t")
		line = strings.TrimRight(line, " \t")
		// Skip comment (#) lines:
		if strings.HasPrefix(line, "#") {
			continue
		} else if line == "" {
			continue
		}

		// Shave off a trailing comment (must be separated from option value by at least on space):
		comment := regexp.MustCompile("[ \t]+#.*$")
		fields := comment.Split(line, 2)
		line = fields[0]
		space := regexp.MustCompile("[ \t]+")
		fields = space.Split(line, -1)

		if fl = len(fields); fl == 0 {
			return Error("no fields, line %d, file %s", lineNo, filename)
		} else if duo && fl != 2 {
			return Error("wrong field count (%d), line %d, file %s", fl, lineNo, filename)
		} else if name, data, err = getND(input, fields); err != nil {
			return Error("failure parsing line %d, file %s: %v", lineNo, filename, err)
		} else if input.objectType == objectTypeHost && name == "" &&
			input.operation != requestTypeGet {
			return Error("name missing, line %d, file %s", lineNo, filename)
		}

		nameData := name + nameDataSep + data
		input.ndList = append(input.ndList, nameData)
		if len(nameData) > input.maxNameLength {
			input.maxNameLength = len(nameData)
		}
	}
	return nil
}

// getND is used to preprocess input provided by the user via either command line
// arguments or via an input file.  Either 1 or 2 arguments must be provided, and
// one is taken as the "name" of the record, and the other is taken to be the
// "data", or content, of the record.  Some objects allow the name and data values
// to be in either order, and for these object types getND determines which is
// which.  Some objects may require the data to be sanitized.

func getND(input *UserInput, args []string) (string, string, error) {

	ShowDebug("getND: args: %#v", args)

	var name, data string
	numArgs := len(args)

	if numArgs < 1 {
		return "", "", Error("no name or data value specified")
	} else if numArgs > 2 {
		return "", "", Error("extra arguments not allowed (only a single name and/or data value)")
	} else if name = args[0]; len(args) == 2 {
		data = args[1]
	}
	if input.objectType == objectTypeHost || input.objectType == objectTypeA ||
		input.objectType == objectTypePTR {
		return getNameIP(args)
	}
	if input.objectType == objectTypeFixedAddress {
		return getIPMac(args)
	}
	ShowDebug("getND: name: \"%s\";  data: \"%s\".", name, data)

	// Hmmm, trying to get TXT records workings...
	if input.objectType == objectTypeTXT {
		// Save the original raw data provided by the user.
		if data == "" {
			input.txtData[name+nameDataSep] = data
		} else {
			// Sanitize the data provided by the user.
			sanitizedData := sanitizeRecordData(data)
			input.txtData[name+nameDataSep+sanitizedData] = data
			//if input.operation != requestTypeGet {
			data = sanitizedData
			//}
		}
	}
	return name, data, nil
}

// For old-times sake (i.e., pdns_utils), allow either "name IP" or "IP name".

func getNameIP(args []string) (string, string, error) {

	ShowDebug("getNameIP: args: %#v", args)
	var name, data string
	numArgs := len(args)

	arg := args[0]
	if net.ParseIP(arg) != nil {
		data = arg
	} else if validHost(arg) {
		name = arg
	} else {
		return "", "", Error("argument \"%s\" is neither a valid IP or name", arg)
	}

	if numArgs == 2 {
		arg = args[1]
		if data == "" {
			if net.ParseIP(arg) != nil {
				data = arg
			} else {
				return "", "", Error("neither argument is a valid IP address")
			}
		} else if validHost(arg) {
			name = arg
		} else {
			return "", "", Error("neither argument is a valid name")
		}
	}
	return name, data, nil
}

func getIPMac(args []string) (string, string, error) {

	ShowDebug("getIPMac: args: %#v", args)
	var name, data string
	numArgs := len(args)

	arg := args[0]
	if net.ParseIP(arg) != nil {
		name = arg
	} else if validMac(arg) {
		data = arg
	} else {
		return "", "", Error("argument \"%s\" is neither a valid IP or MAC", arg)
	}

	if numArgs == 2 {
		arg = args[1]
		if name == "" {
			if net.ParseIP(arg) != nil {
				name = arg
			} else {
				return "", "", Error("neither argument is a valid IP address")
			}
		} else if validMac(arg) {
			data = arg
		} else {
			return "", "", Error("neither argument is a valid MAC")
		}
	}
	return name, data, nil
}

// Split the name and data parts from the specified nameData pair.

func splitND(nameData string) (string, string, error) {

	// Some data fields may contain the nameDataSep; don't split the data.
	s := strings.SplitN(nameData, nameDataSep, 2)
	if len(s) != 2 {
		return "", "", Error("failure getting name/data from \"%s\"", nameData)
	}
	return s[0], s[1], nil
}

// https://www.socketloop.com/tutorials/golang-validate-hostname
// I doubt if this is totally correct, but good enough for now - sandmant

func validHost(host string) bool {
	host = strings.Trim(host, " ")

	re, _ := regexp.Compile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`)
	return re.MatchString(host)
}

func validMac(mac string) bool {
	mac = strings.Trim(mac, " ")

	re, _ := regexp.Compile(`^([a-zA-Z0-9]{2}:){5}[a-zA-Z0-9]`)
	return re.MatchString(mac)
}
