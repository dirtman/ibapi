package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Implement the "host" command.

func commandHost(invokedAs []string) error {

	commands := Commands{
		"get":    getHost,
		"delete": deleteHost,
		"add":    addHost,
		"update": upateHost,
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

func (s StatesHost) AddRecords(nameData string, body []byte) error {

	if len(s[nameData].records) == 0 {
		s[nameData].err = json.Unmarshal(body, &s[nameData].records)
	} else {
		var records []*RecordHost
		if s[nameData].err = json.Unmarshal(body, &records); s[nameData].err == nil {
			for _, record := range records {
				if !findRefHost(s[nameData].records, record.Ref) {
					s[nameData].records = append(s[nameData].records, record)
				}
			}
		}
	}
	return nil
}

func (s StatesHost) NewState(nameData string) {
	s[nameData] = new(StateHost)
}
func (s StatesHost) GetObjectType() string {
	return "record:host"
}
func (s StatesHost) GetNDKeys() (string, string) {
	return "name", "ipv4addr"
}
func (s StatesHost) GetNDPairs() (nds []string) {
	return keys(s)
}
func (s StatesHost) SetError(nameData string, err error) {
	s[nameData].err = err
}
func (s StatesHost) GetError(nd string) error {
	return s[nd].err
}
func (s StatesHost) GetRecordCount(nd string) int {
	return len(s[nd].records)
}

// Check if one or more records in the specified list have the specified object reference.

func findRefHost(records []*RecordHost, ref string) bool {

	for _, record := range records {
		if record.Ref == ref {
			return true
		}
	}
	return false
}

// Grab the IP addresses from a Host record.

func getHostIPs(record *RecordHost) ([]string, error) {

	var ips []string
	hrIpv4Addrs := record.Ipv4Addrs
	for _, hrIpv4Addr := range hrIpv4Addrs {
		ips = append(ips, hrIpv4Addr.Ipv4Addr)
	}
	hrIpv6Addrs := record.Ipv6Addrs
	for _, hrIpv6Addr := range hrIpv6Addrs {
		ips = append(ips, hrIpv6Addr.Ipv6Addr)
	}
	return ips, nil
}
