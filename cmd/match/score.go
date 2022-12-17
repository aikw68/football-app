package match

import (
	"strconv"

	"github.com/tidwall/gjson"
)

type ScoreList struct {
	Standings   int
	PlayerName  string
	Nationality string
	TeamId      int
	TeamName    string
	Goals       string
}

type ScoreLists []*ScoreList

// 得点ランキング取得
func Score(ch chan ScoreLists) {

	var rtnLists ScoreLists

	jsonBytes := ApiCall("scorers")
	games_num := int(gjson.GetBytes(jsonBytes, "count").Int())

	for i := 0; i < games_num; i++ {

		arg := "scorers." + strconv.Itoa(i)

		playerName := gjson.GetBytes(jsonBytes, arg+".player.name").String()
		nationality := gjson.GetBytes(jsonBytes, arg+".player.nationality").String()
		teamId := int(gjson.GetBytes(jsonBytes, arg+".team.id").Int())
		teamName := gjson.GetBytes(jsonBytes, arg+".team.name").String()
		goals := gjson.GetBytes(jsonBytes, arg+".goals").String()

		rtnList := ScoreList{
			i + 1,
			playerName,
			nationality,
			teamId,
			teamName,
			goals}

		rtnLists = append(rtnLists, &rtnList)
	}
	ch <- rtnLists
}
