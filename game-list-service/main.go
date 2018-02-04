package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/luismfonseca/age-of-mythology-private-server/game-list-service/db"
	"github.com/luismfonseca/age-of-mythology-private-server/game-list-service/delete_game_req"
	"github.com/luismfonseca/age-of-mythology-private-server/game-list-service/insert_game_req"
	"github.com/luismfonseca/age-of-mythology-private-server/game-list-service/insert_game_resp"
	"github.com/luismfonseca/age-of-mythology-private-server/game-list-service/retrieve_games_req"
	"github.com/luismfonseca/age-of-mythology-private-server/game-list-service/retrieve_games_resp"
	"github.com/luismfonseca/age-of-mythology-private-server/game-list-service/update_game_req"
	"github.com/luismfonseca/age-of-mythology-private-server/game-list-service/update_game_resp"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

const (
	ServerPort = 8081
)

var gameDatabase db.Database

func gameListService(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.WithField("url", r.RequestURI).WithField("header", r.Header["Soapaction"]).Debugln("Got request")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithError(err).Warnln("Failed to read request body")
		return
	}

	switch r.Header["Soapaction"][0] {
	case `http://ensemblestudios.com/GameListService/RetrieveGames`:
		retrieveGamesReq := retrieve_games_req.Envelope{}
		err = xml.Unmarshal(body, &retrieveGamesReq)
		if err != nil {
			log.WithError(err).Warnln("Failed to unmarshal xml")
			return
		}

		gamesBytes, dbErr := gameDatabase.ListGames(retrieveGamesReq.Offset, retrieveGamesReq.Amount)
		if dbErr != nil {
			log.WithError(dbErr).Warnln("Failed to list the games")
			return
		}

		var games []retrieve_games_resp.Game
		for gameId, gameBytes := range gamesBytes {
			game := &retrieve_games_resp.Game{GameId: gameId}
			err := xml.Unmarshal(gameBytes, game)
			if err != nil {
				log.WithError(err).WithField("game", string(gameBytes)).Warnln("Failed to unmarshal xml")
			}

			games = append(games, *game)
		}

		gameCount := gameDatabase.GetNumberOfGames()

		retrieveGamesResp := retrieve_games_resp.Envelope{
			Soap:               "http://schemas.xmlsoap.org/soap/envelope/",
			Xsi:                "http://www.w3.org/2001/XMLSchema-instance",
			Xsd:                "http://www.w3.org/2001/XMLSchema",
			PublishedGameCount: gameCount, // fixme
			Games:              games,
		}
		log.Info("Listing games")

		responseBytes, _ := xml.MarshalIndent(retrieveGamesResp, "", "  ")
		w.Write(responseBytes)
	case `http://ensemblestudios.com/GameListService/InsertGame`:
		insertGameReq := insert_game_req.Envelope{}
		err = xml.Unmarshal(body, &insertGameReq)
		if err != nil {
			log.WithError(err).Warnln("Failed to unmarshal xml")
			return
		}

		gameId := uuid.New()
		insertGameResp := insert_game_resp.Envelope{
			Soap: "http://schemas.xmlsoap.org/soap/envelope/",
			Xsi:  "http://www.w3.org/2001/XMLSchema-instance",
			Xsd:  "http://www.w3.org/2001/XMLSchema",

			InsertGameResult: gameId,
		}

		game, _ := xml.MarshalIndent(insertGameReq.Game, "", "  ")
		log.WithField("game-id", gameId).Info("Creating game")

		dbErr := gameDatabase.InsertGame(gameId, game)
		if dbErr != nil {
			log.WithError(dbErr).Warnln("Failed to store the game")
			return
		}

		responseBytes, _ := xml.MarshalIndent(insertGameResp, "", "  ")
		w.Write(responseBytes)
	case `http://ensemblestudios.com/GameListService/UpdateGame`:
		updateGameReq := update_game_req.Envelope{}
		err = xml.Unmarshal(body, &updateGameReq)
		if err != nil {
			log.WithError(err).Warnln("Failed to unmarshal xml")
			return
		}

		gameId := updateGameReq.Game.GameId
		updateGameResp := update_game_resp.Envelope{
			Soap: "http://schemas.xmlsoap.org/soap/envelope/",
			Xsi:  "http://www.w3.org/2001/XMLSchema-instance",
			Xsd:  "http://www.w3.org/2001/XMLSchema",
		}

		game, _ := xml.MarshalIndent(updateGameReq.Game, "", "  ")
		log.WithField("game-id", gameId).Info("Updating game")

		dbErr := gameDatabase.InsertGame(gameId, game)
		if dbErr != nil {
			log.WithError(dbErr).Info("Failed to store the game")
			return
		}

		responseBytes, _ := xml.MarshalIndent(updateGameResp, "", "  ")
		w.Write(responseBytes)
	case `http://ensemblestudios.com/GameListService/DeleteGame`:
		deleteGameReq := delete_game_req.Envelope{}
		err = xml.Unmarshal(body, &deleteGameReq)
		if err != nil {
			log.WithError(err).Warnln("Failed to unmarshal xml")
			return
		}
		log.WithField("game-id", deleteGameReq.GameId).Info("Deleting game")

		dbErr := gameDatabase.DeleteGame(deleteGameReq.GameId)
		if dbErr != nil {
			log.WithError(dbErr).Warnln("Failed to store the game")
			return
		}

	default:
		log.WithField("soapaction", r.Header["Soapaction"][0]).Warn("Unknown soap action")
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	log.WithField("url", r.RequestURI).Warn("Received invalid request: returning 404 Not Found")

	http.NotFound(w, r)
}

func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
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
	var err error
	gameDatabase, err = db.New("hosted-games.leveldb")
	if err != nil {
		log.WithError(err).Error("Failed to open hosted games database")
		return
	}

	log.WithField("port", ServerPort).Infoln("Starting game list service")

	router := httprouter.New()
	router.POST("/aomsvr/GameListService", gameListService)
	router.NotFound = http.HandlerFunc(notFound)
	router.MethodNotAllowed = http.HandlerFunc(methodNotAllowed)

	log.Warnln(http.ListenAndServe(fmt.Sprintf(":%d", ServerPort), router))
	log.WithField("port", ServerPort).Infoln("Shutting down server...")
}
