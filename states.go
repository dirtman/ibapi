package main

import (
	"fmt"

	. "github.com/dirtman/sitepkg"
)

/*****************************************************************************\

  Define a few funcs and "wrapper" objects to fetch and hold the results of
  Infoblox WAPI Get requests for the entire set of name/data pairs provided by
  the user.

  Using these funcs/objects is not so useful for user Get requests themselves,
  but rather for user Add requests.  They can be used to check for the
  existence of possibly conflicting records before making any new Add requests.
  While adding an A record with the same name as an existing Host record is
  certainly allowed by Infoblox and may certainly be what is desired, I find it
  beneficial to alert the user just in case.  Furthermore, it may be desirable
  to not add any records if one or more of the requested records will not be
  added due to a record conflict or bad user input.  This is where these record
  "state" objects come into play.  They can be used to hold the current state
  of affairs before taking any actions at all.  While this -drastically-
  increases the overhead and complexity needed to simply add a record, I think
  its worth the effort to possibly avoid DNS related fiascoes.

  StatesXYX is a string map of pointers to StateXYX.  The map keys are the
  name/data pairs provided by the user.  The key values (the StateXYX pointers)
  hold the records retrieved from the WAPI for that name/data key, as well as
  an error to track if an error occurred with a WAPI Get request for the
  name/data pair.

  Update:
  In a attempt to reduce the amount of code required with each new record
  type to be supported, I defined a "Fetcher" interface, and I separated the
  fetching of the raw JSON results from the unmarshalling of those results.  I
  am not very happy with the results - I think there's most likely a better
  approach.  But for now I need to move on.

\*****************************************************************************/

type IBObject interface {
	GetObjectType() string
	GetNDKeys() (string, string)
}
type Fetcher interface {
	IBObject
	GetNDPairs() []string
	GetError(nameData string) error
	GetRecordCount(nameData string) int
	AddRecords(nameData string, body []byte) error
	SetError(nameData string, err error)
	NewState(nameData string)
}

type StatesA map[string]*StateA
type StatesHost map[string]*StateHost
type StatesPTR map[string]*StatePTR
type StatesCNAME map[string]*StateCNAME
type StatesAlias map[string]*StateAlias
type StatesTXT map[string]*StateTXT

type StateBase struct {
	err error
}
type StateA struct {
	StateBase
	records []*RecordA
}
type StateHost struct {
	StateBase
	records []*RecordHost
}
type StatePTR struct {
	StateBase
	records []*RecordPTR
}
type StateCNAME struct {
	StateBase
	records []*RecordCNAME
}
type StateAlias struct {
	StateBase
	records []*RecordAlias
}
type StateTXT struct {
	StateBase
}

/*****************************************************************************\

  Query Infoblox for each name/data pair provided by the user and store the
  results.

  nd is the set of name/data pairs provided as input by the user.
  sf specifies the fields to be used in the request.
  rf specifies the "_return_fields" and determine which fields are returned.

  checkN and checkD are flags that can be used to alter the records search.
  If both are false, a single query is made with both the name/data values as
  specified by the user (i.e., record:RR?nnn=name&ddd=data, assuming both
  name and data were specified by the user).  If only checkN, a query is made
  with only the name value (record:RR?nnn=name).  If only checkD, a query is
  made with only the data value (record:RR?ddd=data). If checkN and checkD,
  both above queries are made, and the results are combined.

  F is Fetcher interface that hold the record object being queried.

\*****************************************************************************/

func getStates(F Fetcher, nd, sf, rf []string, checkN, checkD bool) error {

	var name, data string
	var err error

	for _, nameData := range nd {
		F.NewState(nameData)
		if name, data, err = splitND(nameData); err == nil {
			if name == "" && data == "" && (len(sf) == 0) {
				err = Error("cannot GET with no name, data or search fields")
			} else {
				err = getState(F, nameData, sf, rf, checkN, checkD)
			}
		}
		if err != nil {
			F.SetError(nameData, err)
		}
		ShowDebug("getStates(%T): nameData \"%s\"; err: %v", F, nameData, err)
	}
	return nil
}

// getState is simply a split out from getStates().

func getState(F Fetcher, nameData string, sf, rf []string, checkN, checkD bool) error {

	name, data, _ := splitND(nameData)
	var err error

	ShowDebug("getState(%s); checkN: %v;  checkD: %v;  sf: %#v", nameData, checkN, checkD, sf)
	// If both are false, use whatever input the user provided
	if !checkN && !checkD {
		return getStateRecords(F, nameData, name, data, sf, rf)
	}
	if name != "" && checkN {
		ShowDebug("Calling getStateRecords: name: %s; checkN: %v", name, checkN)
		if err = getStateRecords(F, nameData, name, "", sf, rf); err != nil {
			return err
		}
	}
	if data != "" && checkD {
		ShowDebug("Calling getStateRecords: data: %s; checkD: %v", data, checkD)
		return getStateRecords(F, nameData, "", data, sf, rf)
	}
	return nil
}

// Fetch records and populate the nameData map.

func getStateRecords(F Fetcher, nameData, name, data string, sf, rf []string) error {

	nKey, dKey := F.GetNDKeys()
	object := F.GetObjectType()
	rawRecords, err := getRecords(object, nKey, dKey, name, data, sf, rf)

	ShowDebug("getStateRecords(%s); name: %s;  data: %s;  sf: %v; err: %v",
		nameData, name, data, sf, err)

	if err != nil {
		F.SetError(nameData, err)
	} else {
		F.AddRecords(nameData, rawRecords)
	}

	return nil
}

// Check for and return any errors in the Fetcher object.

func checkStateErrors(F Fetcher, duo, show bool) []error {

	var name, data string
	var errors []error
	var err error
	intro := fmt.Sprintf("Get(%s)", F.GetObjectType())

	for _, nameData := range F.GetNDPairs() {
		if err = F.GetError(nameData); err != nil {
			err = Error("%s \"%s\": %v", intro, nameData, err)
		} else if name, data, err = splitND(nameData); err != nil {
			err = Error("%s \"%s\": failure parsing nameData: %v", intro, nameData, err)
		} else if duo && (name == "" || data == "") {
			err = Error("%s \"%s\": need both a name and a data value", intro, nameData)
		} else if duo && F.GetRecordCount(nameData) > 1 {
			err = Error("%s \"%s\": multiple A records found for unique record", intro, nameData)
		}
		if err != nil {
			errors = append(errors, err)
			if show {
				Warn("%v", err)
			}
		}
	}
	return errors
}
