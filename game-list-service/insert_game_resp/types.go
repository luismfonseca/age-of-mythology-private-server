package insert_game_resp

type Envelope struct {
	XMLName string `xml:"soap:Envelope"`
	Soap    string `xml:"xmlns:soap,attr"`
	Xsi     string `xml:"xmlns:xsi,attr"`
	Xsd     string `xml:"xmlns:xsd,attr"`

	InsertGameResult string `xml:"soap:Body>InsertGameResponse>InsertGameResult"`
}
