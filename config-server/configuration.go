package main

import (
	"fmt"
	"time"
)

type Configuration struct {
	Property                     []Property                     `xml:"Property"`
	StringTable                  StringTableConfig              `xml:"StringTable"`
	ChatChannel                  []ChatChannel                  `xml:"ChatChannel"`
	MessageServer                MessageServer                  `xml:"MessageServer"`
	ChatChannelChatChannelConfig []ChatChannelChatChannelConfig `xml:"ChatChannelConfig>ChatChannel"`
}

type Property struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:"Value,attr"`
}

type StringTableConfig struct {
	URL     string `xml:"URL,attr"`
	Version string `xml:"Version,attr"`
}

type ChatChannel struct {
	Name                string `xml:"Name,attr"`
	Address             string `xml:"Address,attr"`
	Port                string `xml:"Port,attr"`
	InternalServiceName string `xml:"InternalServiceName,attr"`
	Population          string `xml:"Population,attr"`
	LastUpdate          string `xml:"LastUpdate,attr"`
}

type MessageServer struct {
	Port                string `xml:"Port,attr"`
	InternalServiceName string `xml:"InternalServiceName,attr"`
	Population          string `xml:"Population,attr"`
	LastUpdate          string `xml:"LastUpdate,attr"`
	Name                string `xml:"Name,attr"`
	Address             string `xml:"Address,attr"`
}

type ChatChannelChatChannelConfig struct {
	InternalServiceName string `xml:"InternalServiceName,attr"`
	IsLobby             string `xml:"IsLobby,attr"`
	AppID               string `xml:"AppID,attr"`
}

var BaseXMLConfiguration = Configuration{
	Property: []Property{
		{Name: "BaseURL", Value: "http://luisfonseca.xyz/"},
		{Name: "MatchSchemaUrl", Value: fmt.Sprintf("http://0.0.0.0:%d/match-schema", ServerPort)},
		{Name: "FriendsServerAddress", Value: "luisfonseca.xyz"},
		{Name: "AccountServiceUrl", Value: "https://aox.luisfonseca.xyz/aomsvr/authenticate"},
		{Name: "MatchServiceUrl", Value: "http://localhost:8000/match1.aom.eso.com/WebServices/MatchService.asmx"},
		{Name: "GameListServiceUrl", Value: "http://aox.luisfonseca.xyz:8081/aomsvr/GameListService"},
		{Name: "ForwardingServerAddress", Value: "0.0.0.0"},
		{Name: "StatsServerAddress", Value: "aom.luisfonseca.xyz"},
		{Name: "AddressServer1", Value: "0.0.0.0"},
		{Name: "AddressServer2", Value: "0.0.0.0"},
		{Name: "RequiredBuildVersion", Value: "2"},
		{Name: "RequiredBuildEXEVersion", Value: "14"},
		{Name: "RequiredInternalVersion", Value: "1"},
		{Name: "ConfigRefreshInterval", Value: "12000"},
		{Name: "AddressRefreshInterval", Value: "5000"},
		{Name: "RefreshAddressServer", Value: "1"},
		{Name: "ESOOutageMessage", Value: ""},
		{Name: "MaxChatChannelPop", Value: "400"},
		{Name: "MOTD", Value: fmt.Sprintf("http://0.0.0.0:%d/motd?Language=US", ServerPort)},
		{Name: "ServerDateTime", Value: time.Now().Format("01/02/2006 15:04:05 PM")}, // TODO: Calculate this at every request
	},
	StringTable: StringTableConfig{
		URL: fmt.Sprintf("http://0.0.0.0:%d/string-table?Language=US", ServerPort), Version: "1",
	},
	ChatChannel: []ChatChannel{
		{Name: "General Chat 1", Address: "0.0.0.0", Port: "28805", InternalServiceName: "zA2XP_xx_x00", Population: "0", LastUpdate: "2.80"},
		{Name: "General Chat 2", Address: "0.0.0.0", Port: "28806", InternalServiceName: "zA2XP_xx_x01", Population: "0", LastUpdate: "40.94"},
		{Name: "General Chat 3", Address: "0.0.0.0", Port: "28807", InternalServiceName: "zA2XP_xx_x02", Population: "0", LastUpdate: "3.83"},
		{Name: "General Chat 4", Address: "0.0.0.0", Port: "28808", InternalServiceName: "zA2XP_xx_x03", Population: "0", LastUpdate: "19.14"},
		{Name: "Community Events", Address: "0.0.0.0", Port: "28809", InternalServiceName: "zA2XP_xx_x04", Population: "0", LastUpdate: "2.77"},
		{Name: "Beginners Only", Address: "0.0.0.0", Port: "28810", InternalServiceName: "zA2XP_xx_x05", Population: "0", LastUpdate: "36.60"},
		{Name: "Experts Only", Address: "0.0.0.0", Port: "28811", InternalServiceName: "zA2XP_xx_x06", Population: "0", LastUpdate: "5.95"},
		{Name: "Strategy", Address: "0.0.0.0", Port: "28812", InternalServiceName: "zA2XP_xx_x07", Population: "0", LastUpdate: "36.60"},
		{Name: "Lightning", Address: "0.0.0.0", Port: "28814", InternalServiceName: "zA2XP_xx_x09", Population: "0", LastUpdate: "18.86"},
		{Name: "Chinese Speakers", Address: "0.0.0.0", Port: "28815", InternalServiceName: "zA2XP_xx_x10", Population: "0", LastUpdate: "28.85"},
		{Name: "French Speakers", Address: "0.0.0.0", Port: "28816", InternalServiceName: "zA2XP_xx_x11", Population: "0", LastUpdate: "31.27"},
		{Name: "German Speakers", Address: "0.0.0.0", Port: "28817", InternalServiceName: "zA2XP_xx_x12", Population: "0", LastUpdate: "24.09"},
		{Name: "Italian Speakers", Address: "0.0.0.0", Port: "28818", InternalServiceName: "zA2XP_xx_x13", Population: "0", LastUpdate: "27.44"},
		{Name: "Japanese Speakers", Address: "0.0.0.0", Port: "28819", InternalServiceName: "zA2XP_xx_x14", Population: "0", LastUpdate: "25.17"},
		{Name: "Korean Speakers", Address: "0.0.0.0", Port: "28820", InternalServiceName: "zA2XP_xx_x15", Population: "0", LastUpdate: "16.27"},
		{Name: "Portuguese Speakers", Address: "0.0.0.0", Port: "28821", InternalServiceName: "zA2XP_xx_x16", Population: "0", LastUpdate: "25.17"},
		{Name: "Spanish Speakers", Address: "0.0.0.0", Port: "28822", InternalServiceName: "zA2XP_xx_x17", Population: "0", LastUpdate: "41.83"},
		{Name: "Taiwanese Speakers", Address: "0.0.0.0", Port: "28823", InternalServiceName: "zA2XP_xx_x18", Population: "0", LastUpdate: "25.24"},
		{Name: "Deathmatch", Address: "0.0.0.0", Port: "28813", InternalServiceName: "zAMXP_xx_x08", Population: "0", LastUpdate: "3.83"},
	},
	MessageServer: MessageServer{
		Name: "ZoneMsg (198292-FRONT3)", Address: "0.0.0.0", Port: "28801", InternalServiceName: "zMSGz000", Population: "1", LastUpdate: "0.63",
	},
	ChatChannelChatChannelConfig: []ChatChannelChatChannelConfig{
		{InternalServiceName: "zA2XP_xx_x00", IsLobby: "1", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x01", IsLobby: "1", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x02", IsLobby: "1", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x03", IsLobby: "1", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x04", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x05", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x06", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x07", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x08", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x09", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x10", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x11", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x12", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x13", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x14", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x15", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x16", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x17", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x18", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x19", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x20", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x21", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x22", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x23", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x24", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x25", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x26", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x27", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x28", IsLobby: "0", AppID: "0"},
		{InternalServiceName: "zA2XP_xx_x29", IsLobby: "0", AppID: "0"},
	},
}
