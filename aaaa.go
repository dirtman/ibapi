package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Implement the "aaaa" command.

func commandAAAA(invokedAs []string) error {

	commands := Commands{
		"get":    getAAAA,
		"delete": deleteAAAA,
		"add":    addAAAA,
		"update": upateAAAA,
	}

	if len(os.Args) < 2 {
		commandHelp(invokedAs, commands, "need a subcommand", 1)
	}
	if strings.HasPrefix(os.Args[1], "-") || strings.EqualFold(os.Args[1], "help") {
		commandHelp(invokedAs, commands, "", 0)
	}

	command := strings.ToLower(os.Args[1])
	os.Args[0] += " " + os.Args[1]
	os.Args = append(os.Args[:1], os.Args[2:]...)

	function, ok := commands[command]
	if !ok {
		commandHelp(invokedAs, commands, fmt.Sprintf("unrecognized command \"%s\"", command), 1)
	}

	// Run the func for the specified command:
	invokedAs = append(invokedAs, command)
	return function(invokedAs)
}

// Define a set of methods for managing record "states".  See also the states.go file.

func (s StatesAAAA) AddRecords(nameData string, body []byte) error {

	if len(s[nameData].records) == 0 {
		s[nameData].err = json.Unmarshal(body, &s[nameData].records)
	} else {
		var records []*RecordAAAA
		if s[nameData].err = json.Unmarshal(body, &records); s[nameData].err == nil {
			for _, record := range records {
				if !findRefAAAA(s[nameData].records, record.Ref) {
					s[nameData].records = append(s[nameData].records, record)
				}
			}
		}
	}
	return nil
}

func (s StatesAAAA) NewState(nameData string) {
	s[nameData] = new(StateAAAA)
}
func (s StatesAAAA) GetObjectType() string {
	return "record:aaaa"
}
func (s StatesAAAA) GetNDKeys(ndValues ...string) (string, string) {
	return "name", "ipv6addr"
}
func (s StatesAAAA) GetNDPairs() (nds []string) {
	return keys(s)
}
func (s StatesAAAA) SetError(nameData string, err error) {
	s[nameData].err = err
}
func (s StatesAAAA) GetError(nd string) error {
	return s[nd].err
}
func (s StatesAAAA) GetRecordCount(nd string) int {
	return len(s[nd].records)
}

// Check if one or more records in the specified list have the specified object reference.

func findRefAAAA(records []*RecordAAAA, ref string) bool {

	for _, record := range records {
		if record.Ref == ref {
			return true
		}
	}
	return false
}
