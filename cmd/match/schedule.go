package match

import (
	"fmt"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
)

type MatchList struct {
	Month           int
	Day             int
	WDays           string
	StartTime       string
	HomeTeamId      int
	AwayTeamId      int
	HomeTeamName    string
	AwayTeamName    string
	CurrentMatchday int
	StartDate       string
	EndDate         string
	Matchday        int
	HomeScore       int
	AwayScore       int
}

type GetSecretValueOutput struct {
	SecretString string
}

type MatchLists []MatchList

// 試合日程取得
func Schedule(ch chan MatchLists) {

	var rtnLists MatchLists

	jsonBytes := ApiCall("matches")
	games_num := int(gjson.GetBytes(jsonBytes, "resultSet.count").Int())
	wdays := []string{"日", "月", "火", "水", "木", "金", "土"}
	jst, _ := time.LoadLocation("Asia/Tokyo")

	for i := 1; i < games_num; i++ {

		arg := "matches." + strconv.Itoa(i)
		jdate := gjson.GetBytes(jsonBytes, arg+".utcDate").Time().In(jst)
		hour := fmt.Sprintf("%02d", jdate.Hour())
		minute := fmt.Sprintf("%02d", jdate.Minute())
		startTime := hour + ":" + minute

		homeTeamId := int(gjson.GetBytes(jsonBytes, arg+".homeTeam.id").Int())
		awayTeamId := int(gjson.GetBytes(jsonBytes, arg+".awayTeam.id").Int())
		homeTeamName := gjson.GetBytes(jsonBytes, arg+".homeTeam.name").String()
		awayTeamName := gjson.GetBytes(jsonBytes, arg+".awayTeam.name").String()
		currentMatchday := int(gjson.GetBytes(jsonBytes, arg+".season.currentMatchday").Int())
		startDate := gjson.GetBytes(jsonBytes, arg+".season.startDate").String()
		endDate := gjson.GetBytes(jsonBytes, arg+".season.endDate").String()
		matchday := int(gjson.GetBytes(jsonBytes, arg+".matchday").Int())
		homeScore := int(gjson.GetBytes(jsonBytes, arg+".score.fullTime.home").Int())
		awayScore := int(gjson.GetBytes(jsonBytes, arg+".score.fullTime.away").Int())

		rtnList := MatchList{
			int(jdate.Month()),
			int(jdate.Day()),
			wdays[jdate.Weekday()],
			startTime,
			homeTeamId,
			awayTeamId,
			homeTeamName,
			awayTeamName,
			currentMatchday,
			startDate[2:4],
			endDate[2:4],
			matchday,
			homeScore,
			awayScore}

		rtnLists = append(rtnLists, rtnList)
	}
	ch <- rtnLists
}
