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
