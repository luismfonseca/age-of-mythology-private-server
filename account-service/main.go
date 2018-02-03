package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

const (
	ServerPort = 443
)

func passportLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.WithField("url", r.RequestURI).Warn("Got request")

	if len(r.Header["Authorization"]) != 1 {
		// return error
	}

	authorizationMap := make(map[string]string)
	for _, s := range strings.Split(r.Header["Authorization"][0], ",") {
		keyValue := strings.SplitN(s, "=", 2)
		authorizationMap[keyValue[0]] = keyValue[1]
	}
	log.WithField("Passport1.4 sign-in", authorizationMap["Passport1.4 sign-in"]).Warn("username")
	log.WithField("pwd", authorizationMap["pwd"]).Warn("password")


	w.Header().Set("Authentication-Info", "Passport1.4 da-status=success,from-PP="+authorizationMap["Passport1.4 sign-in"]+"                                                                                                                                                                                                          ,ru=http://espassport.eso.com?mkt=EN-US&lc=1033&id=31071")

	w.WriteHeader(http.StatusOK)
}

func zoneAccessService(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.WithField("url", r.RequestURI).WithField("header", r.Header["Soapaction"]).Warn("Got request")

	switch r.Header["Soapaction"][0] {
	case `"http://zone.msn.com/ActivateProductKey"`:
		w.Write([]byte(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
xmlns:xsd="http://www.w3.org/2001/XMLSchema">
  <soap:Body>
    <ActivateProductKeyResponse xmlns="http://zone.msn.com">
      <ActivateProductKeyResult>
        <statusCode>0</statusCode>
      </ActivateProductKeyResult>
    </ActivateProductKeyResponse>
  </soap:Body>
</soap:Envelope>
`))
	case `"http://zone.msn.com/GetCDKeyTicket"`:
		w.Write([]byte(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><GetCDKeyTicketResponse xmlns="http://zone.msn.com"><GetCDKeyTicketResult><statusCode>0</statusCode><xmlBlob>&lt;ticket&gt;&lt;duration&gt;300&lt;/duration&gt;&lt;sessionKey&gt;thisisacoolsession=&lt;/sessionKey&gt;&lt;serviceName&gt;CDKey&lt;/serviceName&gt;&lt;playerName&gt;&lt;/playerName&gt;&lt;playerID&gt;&lt;/playerID&gt;&lt;publicTokens&gt;&lt;/publicTokens&gt;&lt;iv&gt;IjZzGWahw01X5aozu+IcIQ==&lt;/iv&gt;&lt;productID&gt;DPD5/DsqjUAlujB4C1MA&lt;/productID&gt;&lt;encryptedText&gt;jPuo1cSisGD+I3UDxuV0S1cvv6e0ZFDgTbcmY0fMrOnDS/D4bvPej08vUguRhAldPKEKkYitYnTTujf6mxNwdERwW2Gy3Tt8a57kgeVR/EcDrvkT61P/nviKk3OG2nCwpOWI2OhdzTYQak6qIgt6fhlNTFWabVJkI4T4e2ROI0NHqoCxc2jwn4wAnJuS9AL2kro8ursMIoSV86c0NMdT5JOVqgvm9jvrmvsWqMW5Z66N7h+eGuFWVdXoH9v/tN8KJ677ell9/ybvMwPLRIyTTXJVm0RXmPFLFGJU0/ezSKDffvdap/Li/27Lx6ePS3eGv6i+n0N9VTAWg6wOUo2BsFl2Wk5PC0FYMZUYtnyKGiH2Wcz1WZQGAb9Ez8EyT71JN0DC1vm/SXpIY9IoltXJ3Jcoeg2GMjUHFozhdSvQ4n6sUuJk67L5VC8haK9Eq/czXRMg7j+p1gH/TgEqp+jKS+kRcimAzKQRPmbc0nWtVGAYFSFcNW+PsTO1lqUgfwjCFZ8/CG7I1D9XA7r4K73dalLiAoHjrmzhn3E5AvpT1pnh9IyENLbU1+sfVpRNgUr2hmX8N+ySiitzTHmBSX37s0AuzTgRMJQ0+OwGZdT+ePnbIPgIjd9vPOoZ0dVT6Dai4iWyEIHSj4zS8P3eQwCDXw3YjdW7zXah/AtlP2xkQUaYeUfdDN2eOdf4jt8UM3KWIuZ5pYl2DigUfuxazB7ZxHrvgz9BOmv426WkK7aEeKIRx6hjnQASaleed8jCAq1pkQSw04L30lpjpzXV5IWMdTtXew6ZM5kdsEQEZ7k2BTUufhN5D7LM352ZD7qFAzvvCRYOiS2jTUGtEqQoB3Sy/2vyZ/1KGpy/fcLSh8fkUUKyVm2EO9VmJw2D9tW6gUfJDoNWi8o3USHgAhxmcv/2AeQTiLuodDv40FuKsKnh630yRAa6CZfwr/BXXyN7qnmSbKWJHj+IocbIsO4i6M6JBFusNX8HQSo4qT02pVAQsMJJA5cOuBr5tCztwA3RaYDEzOxIL6t48IMRv2MGtIb5ZWOJHNYzRcFAhPIhxSHXWYCmFCyUSe9mdhhcj45PfXUAOAFgOAU2PeImeEWxTbObTUbwx6IfXpk7C2xVjnq+JCLgtmvICc41soWVt92Ec1Ad&lt;/encryptedText&gt;&lt;/ticket&gt;</xmlBlob></GetCDKeyTicketResult></GetCDKeyTicketResponse></soap:Body></soap:Envelope>`))
	case `"http://zone.msn.com/GetServiceTicket"`:
		w.Write([]byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><GetServiceTicketResponse xmlns="http://zone.msn.com"><GetServiceTicketResult><statusCode>0</statusCode><xmlBlob>&lt;ticket&gt;&lt;duration&gt;864000&lt;/duration&gt;&lt;sessionKey&gt;whatever=&lt;/sessionKey&gt;&lt;serviceName&gt;LegacyAuth&lt;/serviceName&gt;&lt;playerName&gt;asd&lt;/playerName&gt;&lt;playerID&gt;0&lt;/playerID&gt;&lt;publicTokens&gt;&lt;/publicTokens&gt;&lt;iv&gt;rXA/m+PRgvZ3lC9m4P7kWg==&lt;/iv&gt;&lt;productID&gt;DPD5/DsqjUAlujB4C1MA&lt;/productID&gt;&lt;encryptedText&gt;7epFsPMF2UAqmrs7Ow643cQsXyMntzTiHgI8hwx8LjTvtjW6Z4iY0lpl8/PGCgo6rIRfFyeIdP6/AipOu1BcHt/JoNhtWVxE5P/3fJTwfAZkc/tDZ3YX4fFHAcdsz6uuF8ulqorOFFSoCiygRnJqR557anL5T28dPjSguX/tPUVmn8lRgeYrx4zr7Wgl13wCxUm+mSqiVRThM8vj8GABeoCSLZQl7WA9GFd/kkswZPDUIU/c2a032OtNCoIQ+FZ8HnRRIpWZYmqIoF+L3F3U3pyW4brNJEk1qTwpuQz3Tde+RUgA0Ln+2PSwcA0JKUjiaIm8tlpowIcUynzztMPRo1VpC7Ls8dHBIJXltNGW6VNlrD2GFRfzs4ItXiYZrgigkhlM8TZx9JftRPOXdMzhPZYYc3lc3lDYFZ7MR99E56YcQj/ZNjTEO4FEDIUUid0Fkssug32ATFJI18VHdm4c0oP2iEvkj6HW20J0hUj2BRb/JAAptkA/hC5YZULeHFPoj0Rjc3w1CDuC3PA9YyfrUq+TIYpGzhr/fOTuevDlE37zIZrXiL5chn/eYNlh1nYLksWMDy6PYxaE7hBMe+7gwMIHZDcGF4gXQ5hKIrt4ziLmFkRY6jQYeH6hVgtk22gZNWZYGIT7hjikyHDG75q/LvgSXOYdWNpOecFcIxn+bVEyQh55shXhpubGocSw+r/QavZhDebZFeTipn5ZUCWdG+g01Pq+ard4v8MzhvLZmb/PEWmR10xqpl0qG6zesEwS9eQ0aFVQWxrLYVjo9DYCLXkimMEwA+kO1HldNf3lpe07n6NHpH/iArcJzCrDtjH7nIIudU9rGDEKYoN25AXRCLDGShy9xd2IxKA8/IP6w3bEnwVlzYhMnBK0Dx95BQVBZzcJvTQwdsn2CCR7NDox6+YGgzDTVFsGHzT2I1oMGt2BY/p/g5vOu3gBfNxK67oIhgGggp/U04Rmx+LvcNQUHLL+dxzuQBBlH0haRNwbBUtZY5qDEV1X9g96l6IJ0s7Okqc23gBSXKLpjDullU4t4Z0HTpOnr0lQMAzw8rKYB7P9Sa9pscZj20mzBD3SGUFcBxInc3VblSWvmmvAGMXQXSfsJF1A3mSj7MzLmUZRnYNHM9WTaqbsO28Rq5RWGjTtpLatFKU5M9PVuVCQbOPo/D0a5/nn0DyqOiigMmzGvKWMD/YRFjwTAdAiLCYCXGzW&lt;/encryptedText&gt;&lt;/ticket&gt;</xmlBlob></GetServiceTicketResult></GetServiceTicketResponse></soap:Body></soap:Envelope>`))
	case `"http://zone.msn.com/RegisterNickname"`:
		w.Write([]byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><RegisterNicknameResponse xmlns="http://zone.msn.com"><RegisterNicknameResult><statusCode>0</statusCode></RegisterNicknameResult></RegisterNicknameResponse></soap:Body></soap:Envelope>`))
	default:
		log.WithField("soapaction", r.Header["Soapaction"][0]).Warn("Unknown soap action")
	}
}

func authenticate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.WithField("url", r.RequestURI).WithField("header", r.Header["Soapaction"]).Warn("Got request")

	switch r.Header["Soapaction"][0] {
	case `http://zone.msn.com/AccountService/Authenticate`:
		w.Write([]byte(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><AuthenticateResponse xmlns="http://zone.msn.com/AccountService"><AccountID>17842503</AccountID><CanonicalAccountName>asd</CanonicalAccountName><AccountIDHex>1104147</AccountIDHex><Passkey>0000000000000000</Passkey><HasEmail>false</HasEmail></AuthenticateResponse></soap:Body></soap:Envelope>`))
	default:
		log.WithField("soapaction", r.Header["Soapaction"][0]).Warn("Unknown soap action")
	}
}

const padding = "============================================="

func masterTicket(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.WithField("url", r.RequestURI).Warn("Got request")

	if len(r.Header["Authorization"]) != 1 {
		// return error
	}

	authorizationMap := make(map[string]string)
	for _, s := range strings.Split(r.Header["Authorization"][0], ",") {
		keyValue := strings.SplitN(s, "=", 2)
		authorizationMap[keyValue[0]] = keyValue[1]
	}
	log.WithField("Passport1.4 sign-in", authorizationMap["Passport1.4 from-PP"]).Warn("username")
	log.WithField("pwd", authorizationMap).Warn("password")

	usernameEmail := strings.TrimSpace(authorizationMap[`Passport1.4 from-PP`]) // something@eso.com
	username := usernameEmail[:len(usernameEmail)-len("@eso.com")]
	sessionKey := []byte(padding)
	copy(sessionKey, []byte(username))
	responseBodyString :=
		`<XMLBlobResult xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://zone.msn.com">
			<statusCode>0</statusCode>
			<xmlBlob>&lt;ticket&gt;&lt;duration&gt;2880000&lt;/duration&gt;&lt;sessionKey&gt;`+string(sessionKey)+`&lt;/sessionKey&gt;&lt;serviceName&gt;Master&lt;/serviceName&gt;&lt;playerName&gt;&lt;/playerName&gt;&lt;playerID&gt;0&lt;/playerID&gt;&lt;publicTokens&gt;&lt;/publicTokens&gt;&lt;iv&gt;3OkGeRUnv4zhbiKYo0CfBg==&lt;/iv&gt;&lt;productID&gt;&lt;/productID&gt;&lt;encryptedText&gt;V4pavoD/HSbll8ecenKOQjoeZ6+WVumK9n+J27oBBoSz2YdSMJ1w6hiY/U3g4UhO/B3cLOibjcSHdLaB3UvhexOFoAp/qbUu6E0nhmLQMvwEhWdqZ7W7QaBQaaxZsQli0s6ECkEPR4EUkNHBsdr3iA5PGfHR8BP1YsrU7jbat1NZ9QjIhQjjIkJGOB4ufnxSnYfja5cWZyDdiwqQwMcGi/9ztKR8Zjw3okZwl2Jov02AO9SoQSdMlbd8l1apYSZ/Z00yQIq02/vWKWgRI9of8ZsF0+HKG7gtWokvIQR4zphlNAGl72D91ylwVc/QzDOubKmxH1/K1sp5lvU/MX0biLi4PQZkAtMJUh+Qh9tuyxRKraAYHO795R5qmjnq7nQ7va2E++ZdoNLsvrPxHqpgg6ndhboRJmRY2hn/i5zy2Bx9MQx9bwDBwsI9PhKqbd9N5YiO/R7vC7ciQZWHBL9xAgVMrOH46UJame9oWIYx4CCn2R/39PKLR1fcSiadYEtHiTO90j6eabDqTF+Tk9/pSMJSE9uxL6rjF+HgM2B2c/s3Dp0li+nO4KLi/W/gHirVd/55whRNtvFWrC/+6kBA0k1YFu5Po8wkumGmGekdj990XScM7R7bPdZC1sL54LYADp9L3udFpuGmwM98cvpG+AqNY+4lSJGY3gwhzRaSVEUfp9/NrXMWaDm5Y4G5VHDL30LO0ljtBEqwBeIeE2ADoSzuUxsgji+nXuSUk74w/qUwg3b/bgSyARIibtRGhms6y9zk7/iSUcITtUhXs7NVU2d4iUF6iZHSOXewAnDdyBrqTw4MFdd3+rnFtR7X4C8xlWP4XeJIoD9uPseUTerBDEGES5WbQIx2vqH/ktpkW7j2XWAO+LNTMRXFNo7J9PBK/vrPUdNv3nZ7WLT55yWTKllb5C/9dthp8dvwySmMggqor+TH62/tbMCaiPIrgbnT1MVr9UbZL41ixt1+WzTINSr4qMwFGKwO0iCy02FrVRv6y3Kk5KfiqSK9Ku7u3uT5ykO14QFyqtxJPi1VerdwVnvPvNsmG5f2W0dnULO+KlVZK6HfTCsn5WS/KVkRNsOw+1Lw90KhhPziWXjCioOeOw==&lt;/encryptedText&gt;&lt;/ticket&gt;</xmlBlob></XMLBlobResult>`


	w.Header().Set("Authentication-Info", "Passport1.4 da-status=success,from-PP='username=asd                                                                                                                                                                                                          ',ru=http://espassport.eso.com?mkt=EN-US&lc=1033&id=31071")
	w.Header().Set("Content-Type", "application/xml")

  log.Println(responseBodyString)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseBodyString))
}

func notFound(w http.ResponseWriter, r *http.Request){
	log.WithField("url", r.RequestURI).Warn("Received invalid request: returning 404 Not Found")

	http.NotFound(w, r)
}

func methodNotAllowed(w http.ResponseWriter, r *http.Request){
	log.WithField("url", r.RequestURI).Warn("Received invalid request: returning 405 Method Not Allowed")

	http.Error(w,
		http.StatusText(http.StatusMethodNotAllowed),
		http.StatusMethodNotAllowed,
	)
}


func main() {
	log.StandardLogger().Formatter = &log.TextFormatter{
		FullTimestamp: true,
	}
	log.WithField("port", ServerPort).Infoln("Starting account service")

	router := httprouter.New()
	router.NotFound = http.HandlerFunc(notFound)
	router.MethodNotAllowed = http.HandlerFunc(methodNotAllowed)

	router.GET("/aomsvr/login.php", passportLogin)
	router.POST("/aomsvr/MasterTickt.php", masterTicket)
	router.POST("/aomsvr/ZoneAccessService.php", zoneAccessService)
	router.POST("/aomsvr/authenticate", authenticate)

	log.Warnln(http.ListenAndServeTLS(fmt.Sprintf(":%d", ServerPort), "../certs/development.pem", "../certs/development.key", router))
	log.WithField("port", ServerPort).Infoln("Shutting down server...")
}
