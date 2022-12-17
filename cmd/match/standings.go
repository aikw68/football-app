package match

import (
	"strconv"

	"github.com/tidwall/gjson"
)

type StandingsList struct {
	Position    int
	TeamId      int
	TeamName    string
	PlayedGames int
	Won         int
	Draw        int
	Lost        int
	Points      int
}

type StandingsLists []StandingsList

// 順位表取得
func Standings(ch chan StandingsLists) {

	var rtnLists StandingsLists
	jsonBytes := ApiCall("standings")

	for i := 0; i < 20; i++ {

		arg := "standings.0.table." + strconv.Itoa(i)

		position := int(gjson.GetBytes(jsonBytes, arg+".position").Int())
		teamId := int(gjson.GetBytes(jsonBytes, arg+".team.id").Int())
		teamName := gjson.GetBytes(jsonBytes, arg+".team.name").String()
		playedGames := int(gjson.GetBytes(jsonBytes, arg+".playedGames").Int())
		won := int(gjson.GetBytes(jsonBytes, arg+".won").Int())
		draw := int(gjson.GetBytes(jsonBytes, arg+".draw").Int())
		lost := int(gjson.GetBytes(jsonBytes, arg+".lost").Int())
		points := int(gjson.GetBytes(jsonBytes, arg+".points").Int())

		rtnList := StandingsList{
			position,
			teamId,
			teamName,
			playedGames,
			won,
			draw,
			lost,
			points}

		rtnLists = append(rtnLists, rtnList)
	}
	ch <- rtnLists
}
