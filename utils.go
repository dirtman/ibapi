package main

import (
	"strconv"
	"strings"

	. "github.com/dirtman/sitepkg"
)

// Very crude routine to add a slice of key=value pairs to a JSON "string".
// "value" is assumed to be a JSON string, unless it looks like a boolean
// or a number.  "body" is assumed to be an existing JSON string to which
// each key/value pair is appended, which each pair preceeded with a ','.

func appendFieldsJSON(body string, fields []string) string {

	for _, field := range fields {
		f := strings.SplitN(field, "=", 2)
		setting := f[0]
		value := f[1]
		body += `,"` + setting + `":`
		_, err := strconv.Atoi(value)
		if err == nil || value == "true" || value == "false" {
			body += value
		} else {
			value = strings.ReplaceAll(value, `"`, `\"`)
			body += `"` + value + `"`
		}
	}
	return body
}

// checkConflict is called by some of the Update commands to check for already
// existing records that conflict with the updates being requested for a record.
// The supported record types for which to check are Host, A, AAAA, CNAME, and Alias.
// The f slice defines the query for all the record types being checked.
// The boolean flags determine which of the supported record types to check.
// tt specifies an Alias target_type and is only applicable to Alias checks.
// Return when the first conflict or error is encountered.
// Note that we do not bother unmarshalling the raw records, and an "empty" body
// should be 2 bytes: "[]".

func checkConflict(f []string, host, a, aaaa, cname, alias bool, tt string) (string, error) {

	var conflict string
	ShowDebug("checkConflict: fields: %#v", f)

	if a {
		ShowDebug("  checkConflict: checking A record")
		if records, err := getRecords("record:a", "", "", "", "", f, nil); err != nil {
			return conflict, Error("A record query failed: %v", err)
		} else if !(records == nil || len(records) <= 2) {
			return "Address record with same name or value already exists", nil
		}
	}
	if host {
		ShowDebug("  checkConflict: checking Host record")
		if records, err := getRecords("record:host", "", "", "", "", f, nil); err != nil {
			return conflict, Error("host query failed: %v", err)
		} else {
			ShowDebug("   checkConflict: Host record: %#v", records)
			ShowDebug("   checkConflict: Host record: %s", records)
			ShowDebug("   checkConflict: len(record): %d", len(records))
			if !(records == nil || len(records) <= 2) {
				return "Host with same name or value already exists", nil
			}
			ShowDebug("   checkConflict: Host: no return")
		}
	}
	if cname {
		ShowDebug("  checkConflict: checking CNAME record")
		if records, err := getRecords("record:cname", "", "", "", "", f, nil); err != nil {
			return conflict, Error("cname query failed: %v", err)
		} else if !(records == nil || len(records) <= 2) {
			return "CNAME record with same name or value already exists", nil
		}
	}
	if alias {
		ShowDebug("  checkConflict: checking Alias record")
		if tt != "" {
			f = append(f, "target_type="+tt)
		}
		if records, err := getRecords("record:alias", "", "", "", "", f, nil); err != nil {
			return conflict, Error("alias query failed: %v", err)
		} else if !(records == nil || len(records) <= 2) {
			return "Alias record with same name or value already exists", nil
		}
	}
	return conflict, nil
}

// Return the keys of a map (https://gosamples.dev/generics-map-keys/).
func keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

/*****************************************************************************\

  sanitizeRecordData is used to sanitize the record data provided by the user,
  generally for TXT record requests, in 2 ways:

  1) special URL chars in the data string, such as "&", need to be escaped.
  2) strings > 255 chars need to be split into multiple sub-strings.

  The max size of a string in a TXT record is 255 chars, but a record can have
  multiple strings (the client joins the strings back into one).  In the
  Infloblox GUI, this could be done by splitting a long string into sub-strings
  with doule quotes: "bigString1" "bigString2".  Here I do the same.

  I noticed via trial and error that if I split a string with `""` like this:
    "thisistoolong"  ->  "this""isto""olon""g",
  an add works, and the resultant record is valid.  But I cannot get the txt
  record with the data ("thisistoolong"). I then noticed that Infoblox shows
  the data as  "this" "isto" "olon" "g".  So I tried spitting with `" "`, but
  then the data gets quietly truncated at the first space (so Infoblox has
  "this".  I then somehow stumbled upon using `"+"`, and this seems to work
  best.  Go figure.  A `+` seems to get converted to a space.

\*****************************************************************************/

func sanitizeRecordData(dataString string) string {

	if dataString == "" {
		return dataString
	}

	var splitString, splitter string
	var currentLen, currentStart int

	// If the TXT data contains spaces, it must be quoted; else Infoblox will
	// convert it into individual quoted words, and the spaces will be lost.
	// It does not seem to hurt if the data does not contain spaces and I quote
	// it anyway.  So I will always quote the data if it is not already quoted.
	if dataString[0] != 34 {	// 34 == `"`
		dataString = `"` + dataString + `"`
		ShowDebug("Added double quotes to TXT data: \"%v\"", dataString)
	}
	if len(dataString) < maxDataStringSize {
		return escapeURLText(dataString)
	}

	for i := range dataString {
		if currentLen == maxDataStringSize {
			splitString += splitter + escapeURLText(dataString[currentStart:i])
			splitter = `"+"`
			currentLen = 0
			currentStart = i
			ShowDebug("sanitizeRecordData: splitString: %s", splitString)
		}
		currentLen++
	}

	// We may have some remaining string (< maxDataStringSize) to append.
	if len(dataString[currentStart:]) > 0 {
		// Corner case: only the trailing double quote is left.  Let's not
		// append two additional double quotes in this case.
		if dataString[currentStart:] == `"` {
			splitString += `"`
		} else {
			splitString += splitter + escapeURLText(dataString[currentStart:])
		}
	}
	ShowDebug("sanitizeRecordData: splitString: %s", splitString)

	if Debug && splitString != dataString {
		ShowDebug("sanitizeRecordData: %s", dataString)
		ShowDebug("                    %s", splitString)
	}
	return splitString
}

// joinDataStrings joins multiple strings into one long string.  This is the reverse
// of sanitizeRecordData, except joinDataStrings does not need to "un-escape" anything.

func joinDataStrings(dataString string) string {

	var joinedString string
	var subStrings []string
	quote := "%22"
	splitter := `"+"`

	ShowDebug("dataString[0:3]: %s", dataString[0:3])
	ShowDebug("dataString[len(dataString)-3:]: %s", dataString[len(dataString)-3:])

	if dataString[0:3] == quote && dataString[len(dataString)-3:] == quote {
		Print("**********************************")
		subStrings = strings.Split(dataString[3:len(dataString)-3], splitter)
	} else {
		subStrings = strings.Split(dataString, splitter)
	}

	for _, subString := range subStrings {
		joinedString += subString
	}

	if Debug && joinedString != dataString {
		ShowDebug("joinDataStrings: %s", dataString)
		ShowDebug("                 %s", joinedString)
	}
	return joinedString
}

// escapeURLText escapes chars in a TXT record that cause issues in a URL.  These
// seem to get converted back by the WAPI.

func escapeURLText(urlText string) string {

	replacer := strings.NewReplacer(
		`+`, "%2B",
		`=`, "%3D",
		`;`, "%3B",
		` `, "%20",
	)
	escaped := replacer.Replace(urlText)
	//	if Debug && urlText != escaped {
	//		ShowDebug("escapeURLText: %s", urlText)
	//		ShowDebug("               %s", escaped)
	//	}
	return escaped
}
