package main

type MatchListResponse struct {
	CurrentTime int64    `json:"currentTime"`
	MatchIDs    []string `json:"matchIds"`
}

type MatchInfoResponse struct {
	MatchInfo MatchInfo `json:"matchInfo"`
	Players   []Player  `json:"players"`
	Teams     []Team    `json:"teams"`
	Coaches   []Coach   `json:"coaches"`
}

type MatchInfo struct {
	MatchId            string      `json:"matchId"`
	MapId              string      `json:"mapId"`
	GameVersion        string      `json:"gameVersion"`
	Region             string      `json:"region"`
	ProvisioningFlowId string      `json:"provisioningFlowId"`
	CustomGameName     string      `json:"customGameName"`
	QueueId            string      `json:"queueId"`
	GameMode           string      `json:"gameMode"`
	SeasonId           string      `json:"seasonId"`
	IsCompleted        bool        `json:"isCompleted"`
	IsRanked           bool        `json:"isRanked"`
	GameLengthMillis   int         `json:"gameLengthMillis"`
	GameStartMillis    int64       `json:"gameStartMillis"`
	PremierMatchInfo   interface{} `json:"premierMatchInfo"`
}

type Player struct {
	Puuid           string      `json:"puuid"`
	GameName        string      `json:"gameName"`
	TagLine         string      `json:"tagLine"`
	TeamId          string      `json:"teamId"`
	PartyId         string      `json:"partyId"`
	PlayerCard      string      `json:"playerCard"`
	PlayerTitle     string      `json:"playerTitle"`
	IsObserver      bool        `json:"isObserver"`
	AccountLevel    int         `json:"accountLevel"`
	CompetitiveTier int         `json:"competitiveTier"`
	Stats           PlayerStats `json:"stats"`
	CharacterId     interface{} `json:"characterId"`
}

type PlayerStats struct {
	Score        int                `json:"score"`
	Kills        int                `json:"kills"`
	Deaths       int                `json:"deaths"`
	Assists      int                `json:"assists"`
	RoundsPlayed int                `json:"roundsPlayed"`
	AbilityCasts PlayerAbilityCasts `json:"abilityCasts"`
}

type PlayerAbilityCasts struct {
	GrenadeCasts  int `json:"grenadeCasts"`
	Ability1Casts int `json:"ability1Casts"`
	Ability2Casts int `json:"ability2Casts"`
	UltimateCasts int `json:"ultimateCasts"`
}

type Team struct {
	TeamId       string `json:"teamId"`
	Won          bool   `json:"won"`
	RoundsPlayed int    `json:"roundsPlayed"`
	RoundsWon    int    `json:"roundsWon"`
	NumPoints    int    `json:"numPoints"`
}

type Coach struct {
	Puuid  string `json:"puuid"`
	TeamId string `json:"teamId"`
}

type AccountResponse struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}
