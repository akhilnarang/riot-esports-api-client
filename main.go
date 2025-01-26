package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const BASE_URL = "https://esports.api.riotgames.com/val/match/v1/"
const RECENT_MATCHES_URL = BASE_URL + "recent-matches/by-queue/tournamentmode"
const MATCH_URL = BASE_URL + "matches/%s"

func main() {
	token := os.Getenv("TOKEN")
	if token == "" {
		panic("TOKEN environment variable not set")
	}

	token = "?api_key=" + token

	req, err := http.NewRequest("GET", RECENT_MATCHES_URL+token, nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func(body io.ReadCloser) {
		err = body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		panic(fmt.Errorf("request failed: %s", body))
	}

	// unmarshal into MatchListResponse
	var matchListResponse MatchListResponse
	if err = json.Unmarshal(body, &matchListResponse); err != nil {
		panic(err)
	}

	var choice = make(map[int]string)

	for i, matchID := range matchListResponse.MatchIDs {
		fmt.Printf("%d: %s\n", i+1, matchID)
		choice[i+1] = matchID
	}

	var input int
	fmt.Print("Enter the match number: ")
	_, err = fmt.Scan(&input)
	if err != nil {
		panic(err)
	}

	if input < 1 || input > len(matchListResponse.MatchIDs) {
		panic("invalid match number")
	}

	req, err = http.NewRequest("GET", fmt.Sprintf(MATCH_URL, choice[input])+token, nil)
	if err != nil {
		panic(err)
	}

	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func(body io.ReadCloser) {
		err = body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		panic(fmt.Errorf("request failed: %s", body))
	}

	var matchResponse MatchInfoResponse
	if err = json.Unmarshal(body, &matchResponse); err != nil {
		panic(err)
	}

	fmt.Printf("\nMatch: %s played on %s is completed: %t\n\n", matchResponse.MatchInfo.MatchId, matchResponse.MatchInfo.MapId, matchResponse.MatchInfo.IsCompleted)

	var teamMap = make(map[string][]Player)

	// players
	for _, player := range matchResponse.Players {
		teamMap[player.TeamId] = append(teamMap[player.TeamId], player)
	}

	var winningTeamID string

	for _, team := range matchResponse.Teams {
		if team.Won {
			winningTeamID = team.TeamId
			break
		}
	}

	// teams
	for team, players := range teamMap {
		var winner = ""
		if team == winningTeamID {
			winner = " (Winner 🎉)"
		}
		fmt.Println("Team: ", team, winner)
		fmt.Printf("\tSr.No.\tName\tScore\tKills\tDeaths\tAssists\tRounds\tFree Ability\tAbility 1\tAbility 2\tUltimate\n")
		for idx, player := range players {
			var obs = ""
			if player.IsObserver {
				obs = " (Observer)"
			}
			fmt.Printf("\t%d. %s%s\t%d\t%d\t%d\t%d\t%d\t%d\t%d\t%d\t%d\n", idx+1, player.GameName, obs, player.Stats.Score, player.Stats.Kills, player.Stats.Deaths, player.Stats.Assists, player.Stats.RoundsPlayed, player.Stats.AbilityCasts.GrenadeCasts, player.Stats.AbilityCasts.Ability1Casts, player.Stats.AbilityCasts.Ability2Casts, player.Stats.AbilityCasts.UltimateCasts)
		}
		fmt.Println()
		fmt.Println()
	}
}