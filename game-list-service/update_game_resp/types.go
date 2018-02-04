package update_game_resp

type Envelope struct {
	XMLName string `xml:"soap:Envelope"`
	Soap	string	`xml:"xmlns:soap,attr"`
	Xsi	string	`xml:"xmlns:xsi,attr"`
	Xsd	string	`xml:"xmlns:xsd,attr"`

	UpdateGameResponse string	`xml:"soap:Body>UpdateGameResponse"`
}
