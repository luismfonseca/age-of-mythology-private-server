package retrieve_games_req

type Envelope struct {
	Offset     int    `xml:"Body>RetrieveGames>Offset"`
	Amount     int    `xml:"Body>RetrieveGames>Amount"`
	SortColumn string `xml:"Body>RetrieveGames>SortColumn"`
	// TODO: Not implemented yet
	//FilterSettings	FilterSettings	`xml:"Body>RetrieveGames>filterSettings"`
	//PreferPlayers	PreferPlayers	`xml:"Body>RetrieveGames>PreferPlayers"`
	//ExcludePlayers	ExcludePlayers	`xml:"Body>RetrieveGames>ExcludePlayers"`
}
