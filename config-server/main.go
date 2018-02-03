package main

import (
	"fmt"
	"net/http"

	"encoding/xml"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

const (
	ServerPort = 8080
)

func aomConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.WithField("url", r.RequestURI).Info("Received request for the AoM config")

	res, _ := xml.MarshalIndent(BaseXMLConfiguration, "", "  ")
	w.Header().Set("Content-Type", "application/xml")

	fmt.Fprintf(w, "%s", res)
}

func stringTable(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.WithField("url", r.RequestURI).Info("Received request for the String Table")

	res, _ := xml.MarshalIndent(BaseXMLStringTable, "", "  ")
	w.Header().Set("Content-Type", "application/xml")

	fmt.Fprintf(w, "%s", res)
}

func motd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.WithField("url", r.RequestURI).Info("Received request for the Message of the Day")

	fmt.Fprintf(w, "%s", "What is the meaning of life?")
}

func matchSchema(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.WithField("url", r.RequestURI).Info("Received request for the Match Schema")

	w.Header().Set("Content-Type", "application/xml")
	fmt.Fprintf(w, "%s", DefaultMatchSchema)
}

func accountServiceConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.WithField("url", r.RequestURI).Info("Received request for the Account Service Configuration")

	// TODO: parse request body to figure out if it's AoM or AoMX

	accountServiceConf, _ := GetAccountServiceConfig()
	res, _ := xml.MarshalIndent(accountServiceConf, "", "  ")
	w.Header().Set("Content-Type", "application/xml")

	fmt.Fprintf(w, "%s", res)
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
	log.WithField("port", ServerPort).Infoln("Starting configuration server")

	router := httprouter.New()
	router.GET("/aom", aomConfig)
	router.GET("/string-table", stringTable)
	router.GET("/motd", motd)
	router.GET("/match-schema", matchSchema)
	router.POST("/account-service-config", accountServiceConfig)
	router.NotFound = http.HandlerFunc(notFound)
	router.MethodNotAllowed = http.HandlerFunc(methodNotAllowed)

	log.Warnln(http.ListenAndServe(fmt.Sprintf(":%d", ServerPort), router))
	log.WithField("port", ServerPort).Infoln("Shutting down server...")
}
