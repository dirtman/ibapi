package main

import (
	"encoding/json"
	"fmt"

	. "github.com/dirtman/sitepkg"
)

// On failure Infoblox WAPI requests return with a code of 400 or higher and
// include a body with additional detail.  An object of type ibError can be
// be used for un-marshalling the body.  Most of the time.  For auth and perhaps
// other failures, an HTML string is returned.

type ibError struct {
	Error string `json:"Error,omitempty"`
	Code  string `json:"code,omitempty"`
	Text  string `json:"text,omitempty"`
}

// IBAPIError implements the Error interface.  It includes additional fields
// to ibError, and the "Error" field has been renamed to accommodate an
// "Error" method.

type IBAPIError struct {
	ibError    string // Infoblox "Error"
	code       string // Infoblox "code"
	text       string // Infoblox "text"
	apiError   string // err returned by lower layer.
	requestURL string // The url request
	method     string // The request method
}

func (err *IBAPIError) Error() string {
	if err.apiError == "" {
		return ""
	} else if err.text != "" {
		if Verbose {
			return fmt.Sprintf("%s: %s", err.apiError, err.ibError)
		} else {
			return err.text
		}
	} else if err.ibError != "" {
		return fmt.Sprintf("%s: %s", err.apiError, err.ibError)
	} else {
		return err.apiError
	}
}

// Infoblox-specific wrapper for API calls.  Since Infoblox always returns a well
// defined error result in the body, we can provide the caller with more detail.

func IBAPIGet(url string) (body []byte, err *IBAPIError) {
	return IBAPIRequest("GET", url, nil)
}
func IBAPIPost(url string, data interface{}) (body []byte, err *IBAPIError) {
	return IBAPIRequest("POST", url, data)
}
func IBAPIPut(url string, data interface{}) (body []byte, err *IBAPIError) {
	return IBAPIRequest("PUT", url, data)
}
func IBAPIDelete(url string, data interface{}) (body []byte, err *IBAPIError) {
	return IBAPIRequest("DELETE", url, data)
}

func IBAPIRequest(method string, url string, data interface{}) ([]byte, *IBAPIError) {

	body, apiErr := APIRequest(method, url, data)
	if apiErr == nil {
		return body, nil
	}
	ibErr := new(ibError)
	ibapiErr := new(IBAPIError)
	ibapiErr.ibError = "empty body returned for failed request"
	ibapiErr.apiError = fmt.Sprintf("%v", apiErr)
	ibapiErr.method = method
	ibapiErr.requestURL = url

	// If body is nil, let's just populate the error from upstream.
	if body == nil {
		return body, ibapiErr
	} else if len(body) < 1 { // Hmmm, not sure what's going on here.
		ibapiErr.ibError = "almost empty body returned for failed request"
		return body, ibapiErr
	} else if body[0] != 123 { // check if it looks like json;  123 == `{`
		ibapiErr.ibError = string(body)
		ibapiErr.text = fmt.Sprintf("%s", apiErr)
		return body, ibapiErr
	} else if err := json.Unmarshal(body, ibErr); err != nil {
		ibapiErr.ibError = fmt.Sprintf("failure unmarshing body: %s", err)
		return body, ibapiErr
	}
	ibapiErr.ibError = ibErr.Error
	ibapiErr.text = ibErr.Text
	ibapiErr.code = ibErr.Code
	return body, ibapiErr
}

/*****************************************************************************\

  Make a WAPI get request and return the raw JSON body.

  object: the Infoblox "wapitype", such as "record:a" or "record:cname".
  nKey: the name key for the type of object, most often "name".
  dKey: the data key, such as "ipv4addr", "canonical", or "target_name".
  sf: the set of fields to be included in the request, such as "zone=external".
  rf: the desired return fields (_return_fields).

  While the Infoblox WAPI does not point out a specific "data" field, ibapi
  distinguishes a specific "data" field for each record type, such as "ipv4addr"
  for A records, "target_name" for Alias records, "canonical" for CNAME records,
  and so on.

  "name" is most often the "name" field specified by Infoblox, but for
  fixedaddresses, for instance, the ibapi "name" refers to the Infoblox
  FixedAddress "ipv4addr" field.  And for authzone "name" refers to the
  ZoneAuth "fqdn" field.  And, more confusingly, for PTR records, "name"
  refers to the Infoblox record:ptr "ptrdname" field, such as "rb4.rice.edu",
  and "data" refers to either the Infoblox record:ptr "name", ipv4addr or
  ipv6addr field, such as 236.182.42.128.in-addr.arpa or 128.42.182.236.

  Separating out a data field and the name field from the rest of the fields
  improves the user experience, especially when operationg on multiple records.
  Likewise for "_return_fields".

\*****************************************************************************/

func getRecords(object, nKey, dKey, name, data string, sf, rf []string) ([]byte, error) {

	ShowDebug("getRecords: object \"%s\".", object)
	ShowDebug("getRecords: name \"%s\".", name)
	ShowDebug("getRecords: data \"%s\".", data)
	ShowDebug("getRecords: sf \"%#v\".", sf)
	ShowDebug("getRecords: rf \"%#v\".", rf)

	if name == "" && data == "" && (len(sf) == 0) {
		return nil, Error("no name, data or fields specified for GET request")
	}
	url := "/" + object
	sep := "?"

	if name != "" {
		url += sep + nKey + "=" + name
		sep = "&"
	}
	if data != "" {
		url += sep + dKey + "=" + data
		sep = "&"
	}
	if len(sf) > 0 {
		for _, field := range sf {
			url += sep + field
			sep = "&"
		}
	}
	if len(rf) > 0 {
		url += sep + "_return_fields%2b" // %2b == '+'
		sep = "="
		for _, field := range rf {
			url += sep + field
			sep = ","
		}
	}
	body, ibapiErr := IBAPIGet(url)
	if ibapiErr != nil {
		//		return nil, Error("failure fetching \"%s\": %v", url, ibapiErr)
		return nil, ibapiErr
	} else if Verbose && !Debug {
		Show("GET result for %s:\n%s\n", url, body)
	}
	return body, nil
}

// Add the specified record.

func addRecord(object, nKey, dKey, name, data string, f []string) ([]byte, error) {

	// For Adds, a name is always required:
	if name == "" {
		return nil, Error("a name value must be provided")
	}
	// And unless dKey == "", a data value is also required:
	if dKey != "" && data == "" {
		return nil, Error("a data value must be provided for %s", dKey)
	}
	url := ""
	sep := "&"

	// The ZoneAuth object does not have a "data" field.
	if dKey == "" {
		url = "/" + object + "?" + nKey + "=" + name
	} else {
		url = "/" + object + "?" + nKey + "=" + name + sep + dKey + "=" + data
	}

	if len(f) > 0 {
		for _, field := range f {
			url += sep + field
		}
	}
	body, ibapiErr := IBAPIPost(url, nil)
	if ibapiErr != nil {
		//		return nil, Error("POST failure for \"%s\": %v", url, ibapiErr)
		return nil, ibapiErr
	}
	ShowDebug("body: %s", body)
	return body, nil
}

func updateRecord(ref string, f []string) ([]byte, error) {

	url := "/" + ref
	sep := "?"

	if f == nil || len(f) < 1 {
		return nil, Error("no fields to update")
	}
	for _, field := range f {
		url += sep + field
		sep = "&"
	}
	body, ibapiErr := IBAPIPut(url, nil)
	if ibapiErr != nil {
		//		return nil, Error("PUT failure for \"%s\": %v", url, ibapiErr)
		return nil, ibapiErr
	}
	ShowDebug("body: %s", body)
	return body, nil
}

// Delete the specified record.

func deleteRecord(ref string, f []string) ([]byte, error) {

	url := "/" + ref
	sep := "?"
	for _, field := range f {
		url += sep + field
		sep = "&"
	}
	body, ibapiErr := IBAPIDelete(url, nil)
	if ibapiErr != nil {
		//		return nil, Error("DELETE failure for \"%s\": %v", url, ibapiErr)
		return nil, ibapiErr
	}
	ShowDebug("body: %s", body)
	return body, nil
}
