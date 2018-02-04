package retrieve_games_resp

type Envelope struct {
	XMLName string `xml:"soap:Envelope"`
	Soap	string	`xml:"xmlns:soap,attr"`
	Xsi	string	`xml:"xmlns:xsi,attr"`
	Xsd	string	`xml:"xmlns:xsd,attr"`

	PublishedGameCount int `xml:"soap:Body>RetrieveGamesResponse>RetrieveGamesResult>PublishedGameCount"`
	Games	[]Game `xml:"soap:Body>RetrieveGamesResponse>RetrieveGamesResult>Games>Game"`
}

type Game struct {
	GameId	string	`xml:"GameId,attr"`
	GameName	string	`xml:"GameName,attr"`
	KVPs	[]KVP	`xml:"KVP"`
}

type KVP struct {
	K	string	`xml:"K,attr"`
	V	string	`xml:"V,attr"`
}
