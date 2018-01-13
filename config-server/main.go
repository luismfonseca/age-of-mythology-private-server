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
	log.WithField("url", r.URL).Info("Received request for AoM config")

	res, _ := xml.MarshalIndent(BaseXMLConfiguration, "", "  ")
	w.Header().Set("Content-Type", "application/xml")

	fmt.Fprintf(w, "%s", res)
}

func stringTable(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.WithField("url", r.URL).Info("Received request for String Table")

	res, _ := xml.MarshalIndent(BaseXMLStringTable, "", "  ")
	w.Header().Set("Content-Type", "application/xml")

	fmt.Fprintf(w, "%s", res)
}

func main() {
	log.StandardLogger().Formatter = &log.TextFormatter{
		FullTimestamp: true,
	}
	log.WithField("port", ServerPort).Infoln("Starting configuration server")

	router := httprouter.New()
	router.GET("/aom", aomConfig)
	router.GET("/stringtable", stringTable)

	log.Warnln(http.ListenAndServe(fmt.Sprintf(":%d", ServerPort), router))
	log.WithField("port", ServerPort).Infoln("Shutting down server...")
}
