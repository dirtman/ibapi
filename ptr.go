package main

import (
    "encoding/json"
    "fmt"
    "os"
    "strings"
)

// Implement the "ptr" command.

func commandPTR(invokedAs []string) error {

    commands := Commands{
        "get":    getPTR,
        "delete": deletePTR,
        "add":    addPTR,
        "update": upatePTR,
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

func (s StatesPTR) AddRecords(nameData string, body []byte) error {

    if len(s[nameData].records) == 0 {
        s[nameData].err = json.Unmarshal(body, &s[nameData].records)
    } else {
        var records []*RecordPTR
        if s[nameData].err = json.Unmarshal(body, &records); s[nameData].err == nil {
            for _, record := range records {
                if !findRefPTR(s[nameData].records, record.Ref) {
                    s[nameData].records = append(s[nameData].records, record)
                }
            }
        }
    }
    return nil
}

func (s StatesPTR) NewState(nameData string) {
    s[nameData] = new(StatePTR)
}
func (s StatesPTR) GetObjectType() string {
    return "record:ptr"
}

// An Infoblox record:ptr object has fields like this:
//    "name": "236.182.42.128.in-addr.arpa",
//    "ipv4addr": "128.42.182.236",
//    "ptrdname": "help.rice.edu",
// Or for IPv6 addresses:
//    "name": "4.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.f.9.9.1.0.2.0.0.0.8.e.5.4.0.6.2.ip6.arpa",
//    "ipv6addr": "c00:1234:5678:9abc::21",
//    "ptrdname": "ps2.crc.rice.edu",
// It seems the most sensible mapping between the DNS record name/data pair
// and the Infoblox pair would be name to name and data to ptrdname.  However
// this makes checking for conflicting A, AAAA, Host etc records more difficult.
// So I will map name to ptrdname, which more closely matches the mapping for
// the other record types.  And for the data, I will map either the Infoblox
// record:ptr name, ipv4addr, or ipv6addr field.

func (s StatesPTR) GetNDKeys(ndValues ...string) (string, string) {
    var name string
    dataKey := "name"
    if len(ndValues) > 1 {
        name = ndValues[1]
        if !validPtrHost(name) {
			if isIPv4(name) {
				dataKey = "ipv4addr"
			} else {
				dataKey = "ipv6addr"
			}
		}
    }
    return "ptrdname", dataKey
}

func (s StatesPTR) GetNDPairs() (nds []string) {
    return keys(s)
}
func (s StatesPTR) SetError(nameData string, err error) {
    s[nameData].err = err
}
func (s StatesPTR) GetError(nd string) error {
    return s[nd].err
}
func (s StatesPTR) GetRecordCount(nd string) int {
    return len(s[nd].records)
}

// Check if one or more records in the specified list have the specified object reference.

func findRefPTR(records []*RecordPTR, ref string) bool {

    for _, record := range records {
        if record.Ref == ref {
            return true
        }
    }
    return false
}
