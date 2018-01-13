package main

var DefaultMatchSchema = `
<matchschema>
	<conditionlist>
		<condition name="PlayerCount" type="int" bitlength="4" popularValue="4" ranking="1" />
		<condition name="NumTeams" type="int" bitlength="2" popularValue="2" ranking="1" />
		<condition name="MapSize" type="int" bitlength="2" popularValue="1" ranking="1" />
		<condition name="VictoryConditions" type="int" bitlength="2" popularValue="0" ranking="1" />
		<condition name="Handicap" type="int" bitlength="2" popularValue="0" ranking="2" />
		<condition name="Visibility" type="int" bitlength="1" popularValue="0" ranking="2" />
		<condition name="WorldResources" type="int" bitlength="2" popularValue="0" ranking="2" />
		<condition name="TeamSharedResources" type="int" bitlength="1" popularValue="0" ranking="2" />
		<condition name="TeamSharedPopulation" type="int" bitlength="1" popularValue="0" ranking="2" />
		<condition name="LockedTeams" type="int" bitlength="1" popularValue="0" ranking="3" />
		<condition name="AllowCheats" type="int" bitlength="1" popularValue="0" ranking="3" />
		<condition name="RecordGame" type="int" bitlength="1" popularValue="0" ranking="3" />
		<condition name="CoOpGame" type="int" bitlength="1" popularValue="0" ranking="3" />
		<condition name="RestrictPauses" type="int" bitlength="1" popularValue="0" ranking="4" />
		<condition name="PreferLanguage" type="int" bitlength="1" popularValue="0" ranking="4" />
		<condition name="LanguageRegion" type="string" popularValue="" ranking="4" />
		<condition name="MapType" type="string" popularValue="alfheim.xs" ranking="4" />
		<condition name="IsMapSet" type="int" bitlength="1" popularValue="0" ranking="5" />
		<condition name="CustomSetting" type="string" popularValue="" ranking="5" />
		<condition name="MatchType" type="int" bitlength="2" />
		<condition name="ServerPing" type="int" bitlength="2" />
		<condition name="MachineSpec" type="int" bitlength="2" />
		<!-- Continuous Match Conditions -->
		<condition name="PlayerRating" type="continuous" maximumdelta="100.0" />
	</conditionlist>
</matchschema>
`
