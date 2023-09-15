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

  sanitizeTXT is used to sanitize the TXT in TXT records.

  The max size of a string in a TXT record is 255 chars, but a record can have
  multiple strings (the client joins the strings back into one).  In the
  Infloblox GUI, this could be done by splitting a long string into sub-strings
  with doule quotes: "bigString1" "bigString2".  I can't figure out how to get
  this working from the command line.
  
  I did notice by chance that a "+" in the TXT seems to split the TXT in two:
    ibapi txt add txt.rice.edu onetwo+three
      Infoblox:  onetwo three
    ibapi txt add txt.rice.edu 'one two + three'
      Infoblox:  one two three
  
  No double quotes appear in Inflox, and this will cause issues with spaces:
    ibapi txt add --debug  txt.rice.edu "Plus+Plus-Equal=Equal-Semi;Semi-2Space  2Space"
      Infoblox: Plus+Plus-Equal=Equal-Semi;Semi-2Space 2Space
	  Dig: "Plus+Plus-Equal=Equal-Semi;Semi-2Space" "2Space"
  The client with join the two and the space (was actually 2) will be lost.
  If I manually quote the record in Infoblox, all is well:
    Infoblox: "Plus+Plus-Equal=Equal-Semi;Semi-2Space 2Space"
	Dig: "Plus+Plus-Equal=Equal-Semi;Semi-2Space 2Space"

\*****************************************************************************/

const maxTXTRecordChunkSize = 100

func sanitizeTXT(txt string) string {
    if maxTXTRecordChunkSize >= len(txt) {
        return escapeURLText(txt)
    }
	var splitString, splitter string
    currentLen := 0
    currentStart := 0
    for i := range txt {
        if currentLen == maxTXTRecordChunkSize {
			splitString += splitter + escapeURLText(txt[currentStart:i])
            currentLen = 0
            currentStart = i
			splitter = "+"
        }
        currentLen++
    }
    splitString += splitter + escapeURLText(txt[currentStart:])
	if Debug && splitString != txt {
		ShowDebug("splitTXT: %s", txt)
		ShowDebug("          %s", splitString)
	}
    return splitString
}

// escapeURLText escapes chars in a TXT record that cause issues in a URL.  This
// seem to get converted back by the WAPI.

func escapeURLText(htmlText string) string {

	replacer := strings.NewReplacer(
		`+`, "%2B",
		`=`, "%3D",
		`;`, "%3B",
		` `, "%20",
	)
	escaped := replacer.Replace(htmlText)
	if Debug && htmlText != escaped {
		ShowDebug("escapeURLText: %s", htmlText)
		ShowDebug("               %s", escaped)
	}
	return escaped
}

// Un-Escape URL special chars:
func unEscapeURLText(escaped string) string {

	replacer := strings.NewReplacer(
		"%2B", `+`,
		"%3D", `=`,
		"%3B", `;`,
		"%20", ` `,
	)
	htmlText := replacer.Replace(escaped)
	if Debug && htmlText != escaped {
		ShowDebug("unEscapeURLText: %s", escaped)
		ShowDebug("                 %s", htmlText)
	}
	return htmlText
}


func splitTXT(s string) string {
    if maxTXTRecordChunkSize >= len(s) {
        return s
    }
	var splitString, splitter string
    currentLen := 0
    currentStart := 0
    for i := range s {
        if currentLen == maxTXTRecordChunkSize {
			splitString += splitter + s[currentStart:i]
            currentLen = 0
            currentStart = i
			splitter = "+"
        }
        currentLen++
    }
    splitString += splitter + s[currentStart:]
	if Debug && splitString != s {
		ShowDebug("splitTXT: %s", s)
		ShowDebug("          %s", splitString)
	}
    return splitString
}
