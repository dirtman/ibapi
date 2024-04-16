package main

import (
	"fmt"
	"strings"

	. "github.com/dirtman/sitepkg"
)

// Implement the "delete" command.

func deleteZoneAuth(invokedAs []string) error {

	var fields []string
	var input *UserInput
	var states StatesZoneAuth = make(StatesZoneAuth)
	var record *ZoneAuth
	var err error
	var duo, assumeyes bool

	SetStringOpt("view", "V", true, "default", "Specify the network view to which the record belongs")
	SetStringOpt("filename", "f", true, "", "Specify an input file")
	SetBoolOpt("assumeyes", "", false, false, "Do not ask for deletion confirmation (assume \"yes\"")
	SetBoolOpt("restartServices", "R", true, false, "Restart Grid services if needed")

	if input, err = subCommandInit(invokedAs[1], invokedAs[2], duo); err != nil {
		return Error("failure initializing program and getting user input: %v", err)
	} else if err = getStates(states, input.ndList, input.fields, nil, false, false); err != nil {
		return Error("failure getting states: %v", err)
	}
	if assumeyes, err = GetBoolOpt("assumeyes"); err != nil {
		return Error("failure getting assumeyes option: %v", err)
	}

	// First check if any errors occurred getting the host records. If so, abort.
	if errors := checkStateErrors(states, duo, true); len(errors) != 0 {
		return Error("Aborting process; no records deleted.")
	}

	// Loop through the user provided input (name/data) list.
	space := input.maxNameLength + 8
	var numNotFound, numFailed uint

	for _, nameData := range input.ndList {
		records := states[nameData].records
		request := strings.TrimLeft(nameData, nameDataSep)
		request = strings.TrimRight(request, nameDataSep)

		if len(records) == 0 {
			Print("%-*s NOTFOUND\n", space, "ZoneAuth("+request+")")
			numNotFound++
			continue
		}
		record = records[0]
		ref := record.Ref
		name := record.Fqdn
		var verify string
		var err error
		var protected bool

		if protected, err = CheckProtectedZoneAuth(name); err != nil {
			Print("%-*s FAILED to delete: %v\n", space, "ZoneAuth("+request+")", err)
			numFailed++
			continue
		}

		if protected {
			Print("%-*s FAILED to delete: Zone is protected\n", space, "ZoneAuth("+request+")")
			numFailed++
			continue
		}

		if !assumeyes {
			Print("Are you SURE you want to delete the WHOLE %s ZONE? Enter yes if so: [no]: ", name)
			fmt.Scanln(&verify)
			if verify != "yes" {
				Print("%-*s Skipping delete at your request\n", space, "ZoneAuth("+request+")")
				numFailed++
				continue
			}
		}

		// OK, we're sure we want to delete a zone:
		_, err = deleteRecord(record.Ref, fields)

		if err != nil {
			Print("%-*s FAILED to delete: %v\n", space, "ZoneAuth("+request+")", err)
			numFailed++
			continue
		} else if ref != record.Ref {
			Print("%-*s FAILED to delete: ref mismatch\n", space, "ZoneAuth("+request+")")
			numFailed++
		} else {
			Print("%-*s Deleted\n", space, "ZoneAuth("+request+")")
		}
	}

	if numFailed != 0 {
		return Error("One or more updates failed")
	} else if numNotFound != 0 {
		return Error("One or more records not found")
	} else if input.restartServices {
		if err = gridRestartServices(Verbose); err != nil {
			return Error("failure restarting services: %s", err)
		}
	}

	return nil
}

func CheckProtectedZoneAuth(name string) (inList bool, err error) {

	// Out of paranoia, I am adding all domains as of 04/10/2024:
	var ProtectedZones = []string{
		"10.0.0.0/8",
		"10.134.196.0/24",
		"10.134.96.0/24",
		"10.225.96.0/24",
		"10.226.96.0/24",
		"128.42.0.0/16",
		"128.42.1.0/24",
		"128.42.130.0/24",
		"128.42.16.0/24",
		"128.42.167.0/24",
		"128.42.17.0/24",
		"128.42.173.0/24",
		"128.42.223.0/24",
		"128.42.60.0/24",
		"128.42.61.0/24",
		"128.42.63.0/24",
		"168.2.0.0/16",
		"168.3.0.0/16",
		"168.4.0.0/16",
		"168.5.0.0/16",
		"168.6.0.0/16",
		"168.7.0.0/16",
		"172.16.0.0/16",
		"172.17.0.0/16",
		"172.30.0.0/16",
		"192.136.144.0/24",
		"192.136.153.0/24",
		"192.168.0.0/16",
		"192.168.110.0/24",
		"192.168.111.0/24",
		"192.168.120.0/24",
		"192.168.137.0/24",
		"192.168.172.0/24",
		"192.168.177.0/24",
		"192.168.71.0/24",
		"192.168.72.0/24",
		"192.168.75.0/24",
		"192.168.77.0/24",
		"192.168.80.0/24",
		"192.168.92.0/24",
		"192.168.94.0/24",
		"206.223.129.0/24",
		"239.255.0.0/16",
		"2604:5e80:20:199f::/64",
		"3dml.rice.edu",
		"acs.rice.edu",
		"appleid.rice.edu",
		"arch.rice.edu",
		"arc.rice.edu",
		"avsys.rice.edu",
		"bakerinstitute.org",
		"baker.rice.edu",
		"bao.rice.edu",
		"bioraft.rice.edu",
		"biosci.rice.edu",
		"business.rice.edu",
		"byod.rice.edu",
		"caam.rice.edu",
		"cachetest.rice.edu",
		"can.rice.edu",
		"carbonhubmail.rice.edu",
		"ceve.rice.edu",
		"chenlab.rice.edu",
		"civicleadership.rice.edu",
		"civil.rice.edu",
		"clear.rice.edu",
		"cloudmail.rice.edu",
		"cnx.rice.edu",
		"coi.rice.edu",
		"comms.rice.edu",
		"continue.rice.edu",
		"crc.rice.edu",
		"cs.rice.edu",
		"ctbp.rice.edu",
		"dev.iam.rice.edu",
		"drive.rice.edu",
		"dyndns.rice.edu",
		"ece.rice.edu",
		"education.rice.edu",
		"eeps.rice.edu",
		"eilab.rice.edu",
		"ei.rice.edu",
		"entrepreneurship.rice.edu",
		"experience.rice.edu",
		"fom.rice.edu",
		"give.rice.edu",
		"giving.rice.edu",
		"glasscock.rice.edu",
		"gulfcoastconsortia.com",
		"gulfcoastconsortia.org",
		"guscus.rice.edu",
		"gworkspacetest.rice.edu",
		"healthhumanitiessyllabi.rice.edu",
		"help.rice.edu",
		"houstoncommunitysustainability.org",
		"houstonfriendsofchambermusic.com",
		"houstonfriendsofchambermusic.org",
		"houstonfriendsofmusic.com",
		"houstonfriendsofmusic.org",
		"iam.rice.edu",
		"info.rice.edu",
		"iondistrict.com",
		"ise.rice.edu",
		"iso.rice.edu",
		"is.rice.edu",
		"itso.rice.edu",
		"kinder.rice.edu",
		"logs.rice.edu",
		"lom.rice.edu",
		"lovett.rice.edu",
		"lync.rice.edu",
		"mail.healthsense.rice.edu",
		"mail.rice.edu",
		"maksud.rice.edu",
		"martel.rice.edu",
		"mat.rice.edu",
		"mems.rice.edu",
		"monitor.rice.edu",
		"moody.rice.edu",
		"msapps.rice.edu",
		"music.rice.edu",
		"navigate.rice.edu",
		"nei.rice.edu",
		"net.rice.edu",
		"ni.rice.edu",
		"oedk.rice.edu",
		"og-hpc.rice.edu",
		"online.rice.edu",
		"osi.rice.edu",
		"owlnet.rice.edu",
		"owlpc.rice.edu",
		"panw.rice.edu",
		"pr.rice.edu",
		"puentes-consortium.org",
		"rac.rice.edu",
		"rcsg.rice.edu",
		"rcs.rice.edu",
		"rdf.rice.edu",
		"rhaptos.com",
		"rhaptos.org",
		"rice.edu",
		"riceoffice365.rice.edu",
		"riceworks.rice.edu",
		"ruee.rice.edu",
		"ruf.rice.edu",
		"scholar.rice.edu",
		"seci.rice.edu",
		"sheelab.rice.edu",
		"sid.rice.edu",
		"spatialstudieslab.rice.edu",
		"spf.rice.edu",
		"splunk.rice.edu",
		"stat.rice.edu",
		"test.iam.rice.edu",
		"testipa.iam.rice.edu",
		"test.rice.edu",
		"touchnet.rice.edu",
		"uss.rice.edu",
		"wrc.rice.edu",
		"xdspam1.mail.rice.edu",
	}

	// Do allow deletion of second level domains, like rice.edu or joe.com:
	numSubZones := strings.Count(name, ".")
	if numSubZones <= 1 {
		return true, nil
	}
	// Don't allow the deletion of Purdue domains:
	if strings.Contains(name, "purdue.edu") {
		return true, nil
	}
	if inList, err = InList(ProtectedZones, name); err != nil {
		return false, Error("Failure checking if zone is protected: %v", err)
	}
	return inList, nil
}
