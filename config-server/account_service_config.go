package main

import "encoding/xml"

type AccountServiceGetConfigRequestEnvelope struct {
	SoapEnc string `xml:"soapenc,attr"`
	GameID  string `xml:"Body>GetConfig>gameID"`
	Soap    string `xml:"soap,attr"`
}

type AccountServiceGetConfigResponseEnvelope struct {
	Soap	string	`xml:"xmlns:soap,attr"`
	Xsi	string	`xml:"xmlns:xsi,attr"`
	Xsd	string	`xml:"xmlns:xsd,attr"`
	StatusCode int `xml:"soap:Body>GetConfigResponse>GetConfigResult>statusCode"`
	XmlBlob    string `xml:"soap:Body>GetConfigResponse>GetConfigResult>xmlBlob"`
	XMLName string `xml:"soap:Envelope"`
}

type AccountServiceConfig struct {
	PassportLoginUrl string `xml:"PassportLoginUrl"`
	HotmailLoginUrl  string `xml:"HotmailLoginUrl"`
	MSNLoginUrl      string `xml:"MSNLoginUrl"`
	PassportSiteID   int `xml:"PassportSiteID"`
	MasterTicketUrl  string `xml:"MasterTicketUrl"`
	WebServiceUrl    string `xml:"WebServiceUrl"`
	XMLName string `xml:"config"`
}

func GetAccountServiceConfig() (*AccountServiceGetConfigResponseEnvelope, error) {
	xmlBlob, err := xml.Marshal(AccountServiceConfig {
		PassportLoginUrl: "aox.luisfonseca.xyz/aomsvr/login.php",
		HotmailLoginUrl: "aox.luisfonseca.xyz/aomsvr/login.php",
		MSNLoginUrl: "aox.luisfonseca.xyz/aomsvr/login.php",
		PassportSiteID: 31071,
		MasterTicketUrl: "aox.luisfonseca.xyz/aomsvr",
		WebServiceUrl: "https://aox.luisfonseca.xyz/aomsvr/ZoneAccessService.php",
	})
	if err != nil {
		return nil, err
	}

	return &AccountServiceGetConfigResponseEnvelope{
		Soap: "http://schemas.xmlsoap.org/soap/envelope/",
		Xsi: "http://www.w3.org/2001/XMLSchema-instance",
		Xsd: "http://www.w3.org/2001/XMLSchema",
		StatusCode: 0,
		XmlBlob: string(xmlBlob),
	}, nil
}
