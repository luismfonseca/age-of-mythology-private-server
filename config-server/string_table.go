package main

type StringTable struct {
	Version    string       `xml:"Version,attr"`
	TableEntry []TableEntry `xml:"TableEntry"`
}

type TableEntry struct {
	ID    string `xml:"ID,attr"`
	Value string `xml:"Value,attr"`
}

var BaseXMLStringTable = StringTable{
	Version: "1",
	TableEntry: []TableEntry{
		{ID: "0", Value: "Unable to connect with ESO (on UDP port %d) due to a network error or lack of administrator rights."},
		{ID: "1", Value: "ESO detected that your computer date is incorrect by more than 30 days. Adjust the date and log in again."},
		{ID: "2", Value: "ESO cannot connect to your computer. This may be caused by an incorrect language setting in Internet Explorer, an invalid CD key, or your Secret Question/Answer is too short."},
		{ID: "3", Value: "Your connection to ESO has been lost. Please reconnect to ESO."},
	},
}
