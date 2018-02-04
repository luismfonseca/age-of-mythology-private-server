package delete_game_req
/*
<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
		<DeleteGame xmlns="http://ensemblestudios.com/GameListService">
			<GameId xmlns="http://ensemblestudios.com/GameListService">89efc9b4-5bc2-4291-9697-79fa36efc980</GameId>
		</DeleteGame>
	</soap:Body>
</soap:Envelope>
 */

type Envelope struct {
	GameId string	`xml:"Body>DeleteGame>GameId"`
}
