package update_game_req

/*
<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
		<UpdateGame xmlns="http://ensemblestudios.com/GameListService">
			<UpdatedGame GameName="asd - xx" GameId="6a6e62cb-c669-4499-8cec-b102d705bdac" xmlns="http://ensemblestudios.com/GameListService">
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="PlayerList" V="asd"/>
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="GameName" V="asd - xx"/>
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="NumGamePlayers" V="2"/>
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="ChatChannelName" V=""/>
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="TeamConfiguration" V="0"/>
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="MapSize" V="0"/>
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="GameType" V="0"/>
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="Handicap" V="0"/>
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="MapType" V="fastrandom.set"/>
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="Visibility" V="0"/>
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="WorldResources" V="1"/>
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="TeamSharedResources" V="0"/>
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="TeamSharedPopulation" V="0"/>
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="LockedTeams" V="1"/>
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="AllowCheats" V="0"/>
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="RecordGame" V="0"/>
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="CoOpGame" V="0"/>
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="RestrictPauses" V="0"/>
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="ServerPing" V="2325"/>
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="AverageRating" V="1600"/>
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="PreferLanguage" V="1"/>
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="LanguageRegion" V="0"/>
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="MachineSpec" V="0"/>
				<KVP xmlns="http://ensemblestudios.com/GameListService" K="IsMapSet" V="1"/>
			</UpdatedGame>
		</InsertGame>
	</soap:Body>
</soap:Envelope>
 */

type Envelope struct {
	Soap	string	`xml:"xmlns:soap,attr"`

	Game Game `xml:"Body>UpdateGame>UpdatedGame"`
}

type Game struct {
	GameName string	`xml:"GameName,attr"`
	GameId string	`xml:"GameId,attr"`
	KVPs	[]KVP	`xml:"KVP"`
}

type KVP struct {
	K	string	`xml:"K,attr"`
	V	string	`xml:"V,attr"`
}
