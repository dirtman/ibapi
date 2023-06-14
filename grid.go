package main

import (
	"encoding/json"

	. "github.com/dirtman/sitepkg"
)

type RefObject struct {
	Ref string `json:"_ref,omitempty"`
}

// The grid's reference ID can be obtained with a simple fetch, but since
// it rarely changes, I will define it here to save the one fetch.
var GridRef = "grid/b25lLmNsdXN0ZXIkMA:Infoblox"

// Get the Grid's reference.
func getGridRef() (string, error) {

	if GridRef == "" {
		return GridRef, nil
	}
	body, api_err := IBAPIRequest("GET", "/grid", nil)
	if api_err != nil {
		return "", Error("Failure getting Grid ref: %s", api_err)
	}
	var refObjects []RefObject
	if err := json.Unmarshal(body, &refObjects); err != nil {
		return "", Error("failure unmarshing body (%s): %s", string(body), err)
	}
	GridRef = refObjects[0].Ref
	return refObjects[0].Ref, nil
}

// Restart grid services if neccessary.
func restartGridServices(verbose bool) error {

	gridref, err := getGridRef()
	if err != nil {
		return Error("failure getting grid reference id: %s", err)
	}
	url := "/" + gridref + "?_function=restartservices"
	url += "&restart_option=RESTART_IF_NEEDED"
	url += "&service_option=ALL"
	url += "&member_order=SEQUENTIALLY"
	url += "&sequential_delay=1"
	_, api_err := IBAPIRequest("POST", url, nil)

	if api_err != nil {
		return Error("failure restarting services: %s", api_err)
	}
	if verbose {
		Show("Successfully instructed Infoblox to restart services if needed")
	}
	return nil
}
